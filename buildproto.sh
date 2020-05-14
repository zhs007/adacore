export PATH="$PATH:$(go env GOPATH)/bin"
protoc --proto_path=proto/ --go_out=plugins=grpc:adacorepb --go_opt=paths=source_relative proto/*.proto
protoc --proto_path=adarenderproto/ --go_out=plugins=grpc:adarenderpb --go_opt=paths=source_relative adarenderproto/*.proto