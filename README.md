# http pprof

Minimal fork of [net/http/pprof](http://golang.org/pkg/net/http/pprof/) that
just removes the import side effect.

See http://godoc.org/github.com/lytics/httppprof for how to use the new API.

*Honestly, you should probably just never use the default HTTP mux for
non-trivial apps. Then you don't need this fork.*
