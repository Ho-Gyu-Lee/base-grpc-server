# base-grpc-server

```bash
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

```bash
  protoc --go_out=. --go-grpc_out=. ./api/*.proto
```

```bash
  protoc -I=. --csharp_out=. --grpc_out=. --plugin=protoc-gen-grpc=grpc_csharp_plugin.exe ./api/common.proto  
  protoc -I=. --csharp_out=. --grpc_out=. --plugin=protoc-gen-grpc=grpc_csharp_plugin.exe ./api/slave.proto  
  //protoc -I=. --csharp_out=. --grpc_out=. --plugin=protoc-gen-grpc="C:\Users\HoGyu\Documents\protoc\bin\grpc_csharp_plugin.exe" ./api/common.proto
  //protoc -I=. --csharp_out=. --grpc_out=. --plugin=protoc-gen-grpc="C:\Users\HoGyu\Documents\protoc\bin\grpc_csharp_plugin.exe" ./api/slave.proto
```

## 참고 자료

 - [Golang gRPC Quick start](https://grpc.io/docs/languages/go/quickstart/)
 - [gRPC Package](https://packages.grpc.io/)