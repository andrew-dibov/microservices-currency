CMD_NAME := microservices-currency

BIN_NAME := $(CMD_NAME)-bin
IMG_NAME := $(CMD_NAME):latest

API_KEY := $(CMD_NAME)-key

.DEFAULT_GOAL := help
.PHONY: get_deps get_protoc gen_protos build_bin run_bin run_app build_container run_container stop_container clean help

# ---

get_deps:
	@ go mod tidy

get_protoc:
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

gen_protos:
	@protoc --go_out=. --go-grpc_out=. proto/currency/currency.proto
	@protoc --go_out=. --go-grpc_out=. proto/conversion/conversion.proto

# ---

build_bin: get_deps get_protoc gen_protos
	@CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o bin/$(BIN_NAME) cmd/$(CMD_NAME)/main.go

run_bin: build_bin
	@./bin/$(BIN_NAME)

# ---

run_app: get_deps get_protoc gen_protos
	@go run cmd/$(CMD_NAME)/main.go

# ---

build_container:
	@docker image build -t $(IMG_NAME) . > /dev/null 2>&1

run_container: build_container
	@docker container run -d --rm --name $(CMD_NAME) $(IMG_NAME)

stop_container:
	@docker container stop $(CMD_NAME) > /dev/null 2>&1

# ---

clean:
	@rm -rf bin

# ---

help:
	@echo "get_deps        установить зависимости"
	@echo "get_protoc      установить protoc"
	@echo "gen_protos      сгенерировать .proto"
	@echo "build_bin       собрать бинарник"
	@echo "run_bin         запустить бинарник"
	@echo "run_app         запустить приложение"
	@echo "build_container собрать образ"
	@echo "run_container   запустить контейнер"
	@echo "stop_container  остановить контейнер"
	@echo "clean           удалить сертификаты и бинарник"
	@echo "help            показать справку"