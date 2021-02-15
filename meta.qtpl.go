// Code generated by qtc from "meta.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line meta.qtpl:1
package main

//line meta.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line meta.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line meta.qtpl:3
func streammeta(qw422016 *qt422016.Writer, importRoot, vcs, vcsRoot, suffix string) {
//line meta.qtpl:3
	qw422016.N().S(`
<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
<meta name="go-import" content="`)
//line meta.qtpl:8
	qw422016.E().S(importRoot)
//line meta.qtpl:8
	qw422016.N().S(` `)
//line meta.qtpl:8
	qw422016.E().S(vcs)
//line meta.qtpl:8
	qw422016.N().S(` `)
//line meta.qtpl:8
	qw422016.E().S(vcsRoot)
//line meta.qtpl:8
	qw422016.N().S(`">
<meta http-equiv="refresh" content="0; url=https://godoc.org/`)
//line meta.qtpl:9
	qw422016.E().S(importRoot)
//line meta.qtpl:9
	qw422016.E().S(suffix)
//line meta.qtpl:9
	qw422016.N().S(`">
</head>
<body>
Redirecting to docs at <a href="https://godoc.org/`)
//line meta.qtpl:12
	qw422016.E().S(importRoot)
//line meta.qtpl:12
	qw422016.E().S(suffix)
//line meta.qtpl:12
	qw422016.N().S(`">godoc.org/`)
//line meta.qtpl:12
	qw422016.E().S(importRoot)
//line meta.qtpl:12
	qw422016.E().S(suffix)
//line meta.qtpl:12
	qw422016.N().S(`</a>...
</body>
</html>
`)
//line meta.qtpl:15
}

//line meta.qtpl:15
func writemeta(qq422016 qtio422016.Writer, importRoot, vcs, vcsRoot, suffix string) {
//line meta.qtpl:15
	qw422016 := qt422016.AcquireWriter(qq422016)
//line meta.qtpl:15
	streammeta(qw422016, importRoot, vcs, vcsRoot, suffix)
//line meta.qtpl:15
	qt422016.ReleaseWriter(qw422016)
//line meta.qtpl:15
}

//line meta.qtpl:15
func meta(importRoot, vcs, vcsRoot, suffix string) string {
//line meta.qtpl:15
	qb422016 := qt422016.AcquireByteBuffer()
//line meta.qtpl:15
	writemeta(qb422016, importRoot, vcs, vcsRoot, suffix)
//line meta.qtpl:15
	qs422016 := string(qb422016.B)
//line meta.qtpl:15
	qt422016.ReleaseByteBuffer(qb422016)
//line meta.qtpl:15
	return qs422016
//line meta.qtpl:15
}