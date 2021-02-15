package main

//go:generate qtc

import "github.com/julienschmidt/httprouter"

type redirconf struct {
	Source string
	Dest   string
}

func main() {
	router := httprouter.New()

	router.GET()
}
