GO_ENTRYPOINT	=	.\src

PROTO_SRC	=	sidelogs-proto\common.proto \
				sidelogs-proto\logs.proto \

PROTO_DEST	=	.\gen

run:
	go run $(GO_ENTRYPOINT)

proto:
	protoc --go_out=$(PROTO_DEST) --go_opt=paths=source_relative \
    --go-grpc_out=$(PROTO_DEST) --go-grpc_opt=paths=source_relative \
    $(PROTO_SRC)
