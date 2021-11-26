protoc --go_out=plugins=grpc,paths=source_relative:gen/v1 auth.proto    

protoc --grpc-gateway_out=paths=source_relative,grpc_api_configuration=auth.yaml:gen/v1 auth.proto