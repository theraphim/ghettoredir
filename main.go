package main

//go:generate qtc

import (
	"bytes"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"time"

	"github.com/coreos/go-systemd/v22/activation"
	"github.com/coreos/go-systemd/v22/daemon"
	"github.com/julienschmidt/httprouter"
	"github.com/kelseyhightower/envconfig"

	log "github.com/sirupsen/logrus"
)

type redirconf struct {
	ListenHTTP []string
	Host       string
	DropPrefix string `default:"/ghetto"`
	VCS        string `default:"git"`
}

func (s *redirconf) serve(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pn := ps.ByName("path")
	if pn == "" {
		http.Error(w, "foo", http.StatusBadRequest)
		return
	}

	// 2-component paths
	pn = strings.TrimPrefix(pn, "/")
	splits := strings.Split(pn, "/")
	if len(splits) < 2 {
		http.Error(w, "foo", http.StatusBadRequest)
		return
	}
	base := splits[0] + "/" + splits[1]

	buf := meta(s.Host+"/"+base, s.VCS, "https://"+s.Host+"/"+base+".git", pn)
	http.ServeContent(w, r, "", time.Time{}, bytes.NewReader([]byte(buf)))
}

func main() {
	var conf redirconf

	if err := envconfig.Process("ghettoredir", &conf); err != nil {
		log.Fatal(err)
	}
	router := httprouter.New()

	router.GET(conf.DropPrefix+"/*path", conf.serve)

	listeners, err := activation.Listeners()
	if err != nil {
		log.Fatal(err)
	}
	var wg sync.WaitGroup
	for _, v := range listeners {
		if v == nil {
			continue
		}
		v := v
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := http.Serve(v, router); err != nil {
				log.Error(err)
			}
		}()
	}
	for _, v := range conf.ListenHTTP {
		v := v
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := http.ListenAndServe(v, router); err != nil {
				log.Error(err)
			}
		}()
	}

	daemon.SdNotify(false, daemon.SdNotifyReady)

	dchan := make(chan struct{}, 1)
	go func() {
		defer close(dchan)
		wg.Wait()
	}()
	control := make(chan os.Signal, 1)
	signal.Notify(control, os.Interrupt)
	select {
	case sig := <-control:
		log.Infof("Signal caught: %v", sig)
	case <-dchan:
		log.Info("All listeners gone")
	}
	daemon.SdNotify(false, daemon.SdNotifyStopping)
}
