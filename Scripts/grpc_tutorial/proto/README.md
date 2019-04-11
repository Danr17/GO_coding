How to use protobuf + grpc

1. Install:
- https://developers.google.com/protocol-buffers/docs/downloads   
- go get google.golang.org/grpc    
- go get -u google.golang.org/protobuf/protoc-gen-go   

2. Run:
protoc.exe --proto_path=proto --proto_path=C:\Users\xxx\protoc-3.7.1-win64\include --go_out=plugins=grpc:proto service.proto
