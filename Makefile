LOCAL_BIN:=$(CURDIR)/bin

generate-api:
	make generate-sso-api
	make generate-s1-api
	make generate-s2-api
	make generate-s3-api
	make generate-s4-api
	make generate-s5-api
	make generate-s6-api
	make generate-s7-api
	make generate-support-api

generate-sso-api:
	mkdir -p sso_v1
	protoc --proto_path api \
	--go_out=./ --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=./ --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/sso_v1/sso.proto

generate-support-api:
	mkdir -p support_v1
	protoc --proto_path api \
	--go_out=./ --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=./ --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/support_v1/support.proto

generate-s1-api:
	mkdir -p service1
	protoc --proto_path api \
	--go_out=./ --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=./ --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/service1/service1.proto

generate-s2-api:
	mkdir -p service2
	protoc --proto_path api \
	--go_out=./ --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=./ --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/service2/service2.proto

generate-s3-api:
	mkdir -p service1
	protoc --proto_path api \
	--go_out=./ --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=./ --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/service3/service3.proto


generate-s4-api:
	mkdir -p service4
	protoc --proto_path api \
	--go_out=./ --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=./ --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/service4/service4.proto

generate-s5-api:
	mkdir -p service5
	protoc --proto_path api \
	--go_out=./ --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=./ --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/service5/service5.proto

generate-s6-api:
	mkdir -p service6
	protoc --proto_path api \
	--go_out=./ --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=./ --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/service6/service6.proto

generate-s7-api:
	mkdir -p service7
	protoc --proto_path api \
	--go_out=./ --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=./ --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/service7/service7.proto

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
