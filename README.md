# middlechain

[![CI](https://github.com/Pranc1ngPegasus/middlechain/actions/workflows/ci.yml/badge.svg)](https://github.com/Pranc1ngPegasus/middlechain/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/Pranc1ngPegasus/middlechain.svg)](https://pkg.go.dev/github.com/Pranc1ngPegasus/middlechain)

The simple HTTP middleware chainer for Golang.

## How to use

```go
mux := http.NewServeMux()

handler := middlechain.Chain(mux,
  someMiddleware,
)

// handle http request...

http.ListenAndServe(":8080", handler)
```
