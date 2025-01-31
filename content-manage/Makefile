GOHOSTOS := $(shell go env GOHOSTOS)
GOPATH := $(shell go env GOPATH)
VERSION := $(shell git describe --tags --always)

ifeq ($(GOHOSTOS), windows)
	# 在 Windows 系统上设置 Git_Bash 路径并转换分隔符
	Git_Bash := $(subst \,/,$(subst cmd\,bin\bash.exe,$(dir $(shell where git))))

	# 使用 Git Bash 的 find 命令查找 proto 文件
	INTERNAL_PROTO_FILES := $(shell $(Git_Bash) -c "find internal -name '*.proto'")
	API_PROTO_FILES := $(shell $(Git_Bash) -c "find api -name '*.proto'")
else
	# 非 Windows 系统直接使用 find 命令
	INTERNAL_PROTO_FILES := $(shell find internal -name '*.proto')
	API_PROTO_FILES := $(shell find api -name '*.proto')
endif

.PHONY: init
# 初始化环境
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/google/wire/cmd/wire@latest

.PHONY: config
# 生成 internal proto
config:
	protoc --proto_path=./internal \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./internal \
	       $(INTERNAL_PROTO_FILES)

.PHONY: api
# 生成 api proto
api:
	protoc --proto_path=./api \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./api \
 	       --go-http_out=paths=source_relative:./api \
 	       --go-grpc_out=paths=source_relative:./api \
	       --openapi_out=fq_schema_naming=true,default_response=false:. \
	       $(API_PROTO_FILES)

.PHONY: build
# 编译
build:
	go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

.PHONY: generate
# 生成
generate:
	go mod tidy
	go get github.com/google/wire/cmd/wire@latest
	go generate ./...

.PHONY: all
# 生成所有文件
all: api config generate

# 帮助信息
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.PHONY: print
# 打印 Git_Bash 变量（用于调试）
print:
	$(info Git_Bash=$(Git_Bash))

.DEFAULT_GOAL := help
