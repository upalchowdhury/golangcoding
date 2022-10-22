.PHONY: protos

protos:
	protoc  -I protos/ protos/currency.proto --go_out=plugins=grpc:protos/
	
	
	
	# protoc protos/*.proto \
	#  --go_out=. \
	#    --go-grpc_out=. \
	#    --go-grpc_opt=paths=source_relative \
	#    --go_opt=paths=source_relative \
	#    --proto_path=.