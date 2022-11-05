# middlechain

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
