# gRPC

## what is gRPC?
An RPC library and framework

gRPC is a modern, open source, high-performance remote procedure call
(RPC) framework that can run anywhere.
gRPC enables client and server applications to communicate transparently,
and simplifies the building of connected systems.

## quickStart
let's try [helloworld](https://www.grpc.io/docs/quickstart/go/) first!

```
cd grpc/helloworld
# gen pb.go
sh genCode.sh
# run client and server
go run greeter_client/main.go world &
go run greeter_server/main.go
```
it's pretty easy case.
we only need define our service with request and response in `xxx.proto` file.
then we can generate the code we need.

for generate env we need install
- protoc(for compiler)
- protoc-gen-go (for go code generate)

for our own bussiness logic, just need implement the method we had defined,
like `SayHello`.

## advance
### grpc type
the method `SayHello` is a kind of `Unary RPCs`,
there are other kind is supported([see in Service definition](https://www.grpc.io/docs/guides/concepts/)).
- Server streaming RPCs
- Client streaming RPCs
- Bidirectional streaming RPCs

### interceptor
there are four kind of interceptor.
[see in go-grpc](https://github.com/grpc/grpc-go/blob/347a6b4db3/examples/features/interceptor/README.md)

- Client Side
    - Unary Interceptor

        An implementation of a unary interceptor can usually be divided into three parts: pre-processing, invoking RPC method, and post-processing.
    - Stream Interceptor

        An implementation of a stream interceptor usually include pre-processing, and stream operation interception.
- Server Side (refer to client side)
    - Unary Interceptor
    - Stream Interceptor

About `metadata` in example which usually used for authentication. [see](https://davidsbond.github.io/2019/06/14/creating-grpc-interceptors-in-go.html)
### http grpc status mapping

[see](https://github.com/grpc/grpc/blob/master/doc/http-grpc-status-mapping.md)

### middleware
[go-grpc-middleware](https://github.com/grpc-ecosystem/go-grpc-middleware)
[see jwk with grpc](jwk/quickStart.md#grpc-with-jwt)

## request tools
### [grpcurl](https://github.com/fullstorydev/grpcurl)
```
brew install grpcurl
```
- list service

`grpcurl -import-path grpc/helloworld/ -proto helloworld.proto list`
- list method

`grpcurl -import-path grpc/helloworld/ -proto helloworld.proto list helloworld.Greeter`
- describe method

`grpcurl -import-path grpc/helloworld/ -proto helloworld.proto describe helloworld.Greeter.SayHello`
- request

```
go run grpc/helloworld/greeter_server/main.go
grpcurl -plaintext -import-path grpc/helloworld/ -proto helloworld.proto -d '{"name":"world"}' localhost:50051 helloworld.Greeter/SayHello
```
`-plaintext` for Use plain-text HTTP/2 when connecting to server (no TLS).
without it u will get error like

**Failed to dial target host "localhost:50051": tls: first record does not look like a TLS handshake**

if u want use endpoint directly, check [enable-server-reflection](https://github.com/grpc/grpc-go/blob/master/Documentation/server-reflection-tutorial.md#enable-server-reflection)

```
grpcurl -plaintext -d '{"name":"world"}' localhost:50051 helloworld.Greeter/SayHello
```

### [bloomRPC](https://github.com/uw-labs/bloomrpc)
The missing GUI Client for GRPC services.
```
brew cask install bloomrpc
```