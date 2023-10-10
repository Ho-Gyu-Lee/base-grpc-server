# base-grpc-server

```bash
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

```bash
  protoc --go_out=. --go-grpc_out=. ./api/*.proto
```

```bash
  protoc -I=. --cpp_out=. --grpc_out=. --plugin=protoc-gen-grpc=grpc_cpp_plugin.exe ./api/greeter.proto
  //protoc -I=. --cpp_out=. --grpc_out=. --plugin=protoc-gen-grpc="C:\Users\HoGyu\Documents\Protobuf\grpc_cpp_plugin.exe" ./api/greeter.proto
```

```bash
  protoc -I=. --csharp_out=. --grpc_out=. --plugin=protoc-gen-grpc=grpc_csharp_plugin.exe ./api/greeter.proto
  //protoc -I=. --csharp_out=. --grpc_out=. --plugin=protoc-gen-grpc="C:\Users\HoGyu\Documents\Protobuf\grpc_csharp_plugin.exe" ./api/greeter.proto
```

## 참고 자료

 - [Golang gRPC Quick start](https://grpc.io/docs/languages/go/quickstart/)
 - [gRPC Package](https://packages.grpc.io/)
 - gRPC Package Version : 2022-09-27T01:37:21-07:00