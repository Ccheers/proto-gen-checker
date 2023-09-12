bin := protoc-gen-checker
outputDIR := ./bin
output := $(outputDIR)/$(bin)

VERSION := $(shell git describe --abbrev=0 --tags)
GOVERSION := $(shell go env GOVERSION)
BUILDTIME := $(shell date -u '+%Y-%m-%dT%H:%M:%SZ')
OSARCH := `go env GOOS`/`go env GOARCH`
GITCOMMIT := $(shell git rev-parse --short HEAD)
LDFLAGS :=  -X 'github.com/ccheers/proto-gen-checker/internal/meta.Version=$(VERSION)' \
	-X 'github.com/ccheers/proto-gen-checker/internal/meta.BuildTime=$(BUILDTIME)' \
	-X 'github.com/ccheers/proto-gen-checker/internal/meta.GitCommit=$(GITCOMMIT)' \
	-X 'github.com/ccheers/proto-gen-checker/internal/meta.GoVersion=$(GOVERSION)' \
	-X 'github.com/ccheers/proto-gen-checker/internal/meta.OSArch=$(OSARCH)'

.PHONY: gen_test
gen_test:
	protoc -I=. \
	       --go_out=paths=source_relative:. \
	       --checker_out=paths=source_relative:. \
	       ./test/v1/test.proto

.PHONY: install
install:
	cd ./cmd/protoc-gen-checker && go install

.PHONY: build
build:
	go build -ldflags "$(LDFLAGS)" -o $(output) ./cmd/protoc-gen-checker

.PHONY: package
package:
	CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o $(output) ./cmd/protoc-gen-checker
	tar -czf package/$(bin)-linux-x86.tar.gz $(output)
	CGO_ENABLE=0 GOOS=linux GOARCH=arm go build -ldflags "$(LDFLAGS)" -o $(output) ./cmd/protoc-gen-checker
	tar -czf package/$(bin)-linux-arm.tar.gz $(output)
	CGO_ENABLE=0 GOOS=darwin GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o $(output) ./cmd/protoc-gen-checker
	tar -czf package/$(bin)-darwin-x86.tar.gz $(output)
	CGO_ENABLE=0 GOOS=darwin GOARCH=arm64 go build -ldflags "$(LDFLAGS)" -o $(output) ./cmd/protoc-gen-checker
	tar -czf package/$(bin)-darwin-arm.tar.gz $(output)
	CGO_ENABLE=0 GOOS=windows GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o $(output).exe ./cmd/protoc-gen-checker
	tar -czf package/$(bin)-win.tar.gz $(output).exe
	rm -rf $(outputDIR)