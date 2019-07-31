protoc -I adarenderpb/ adarenderpb/adarender.proto --go_out=plugins=grpc:adarenderpb
protoc -I proto/ proto/adacore.proto --go_out=plugins=grpc:proto