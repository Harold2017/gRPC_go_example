#

## try gRPC with golang

gRPC-go with a better ecosystem than python

> go-grpc-middleware: https://github.com/grpc-ecosystem/go-grpc-middleware

Those middlewares/interceptors enable many useful functions like logging in a easier implementation way compared with python gRPC interceptors (what i did server logging by wrapper...)

> go mod instead of GOPATH: https://golang.org/doc/go1.12#modules

It just works like pipenv with go.mod like Pipfile and go.sum like Pipfile.lock

goroutine vs coroutine:
> async/await provides cooperative concurrency. Multiple coroutines run in the same thread, but only one is active at any time. A coroutine suspends/swaps execution at explicitly defined points.
> 
> go provides pooled parallelism. Multiple subroutines run across multiple threads, with several being active at any time. A subroutine suspends/swaps execution at implicitly defined points.

goroutine with channel:
> https://medium.com/rungo/anatomy-of-channels-in-go-concurrency-in-go-1ec336086adb
> 
> https://codeburst.io/diving-deep-into-the-golang-channels-549fd4ed21a8
