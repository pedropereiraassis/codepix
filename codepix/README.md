# Codepix platform
This is a learning project to run microservices with new techs for me like Golang, gRPC, Kafka, Nest.js and Next.js.

### Command to generate grpc/pb go files
```
protoc --go_out=application/grpc/pb --go_opt=paths=source_relative --go-grpc_out=application/grpc/pb --go-grpc_opt=paths=source_relative --proto_path=application/grpc/protofiles application/grpc/protofiles/*.proto
``` 