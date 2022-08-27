# grpc-gateway-usage

## Running the servers

```sh
go run server/server.go
go run proxy/proxy.go
```

## Testing

```sh
curl -X POST localhost:8081/v1/example/echo
```
