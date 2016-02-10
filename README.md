# Forcing HTTPS with Go
An implementation of forcing https within the server code layer. Obviously there are several ways to achieve this
, but this is to demonstrate how easy it is to accomplish with go.

### Explanation
The objective was to keep it simple. All that's going on are two http methods that listen for incoming requests. It
wouldn't work to have them execute sequentially as the server would hang on the first method call and never reach the
second. The solution was to create a concurrent goroutine.

### Generating a Certificate and Private Key
Go comes with a certificate generator located in `$GOROOT/src/crypto/tls/generate_cert.go`. You just have to run it
with `go run <path-to>/generate_cert.go`. You can also use openssl which there are many tutorials on how to do so. Note
that even though these are self-signed, you will get warnings in all modern browsers. This is completely fine for
development, but you will eventually need to recieve from a Certificate Authority.
