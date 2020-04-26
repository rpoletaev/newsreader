VERSION := $(shell git describe --tags --abbrev=0)
GIT_SHA := $(shell git rev-parse --short HEAD)
DATE := $(shell date +%s)
MODULE := github.com/rpoletaev/newsreader
GO = GO111MODULE=on CGO_ENABLED=0 go
GO_FLAGS := -mod=vendor -tags production -installsuffix cgo


LDFLAGS += -X $(MODULE)/pkg.Timestamp=$(DATE)
LDFLAGS += -X $(MODULE)/pkg.Version=$(VERSION)
LDFLAGS += -X $(MODULE)/pkg.GitSHA=$(GIT_SHA)


build: $(GO) build $(GO_FLAGS) -o bin/newsreader -ldflags "$(LDFLAGS)" ./cmd

gen:
	mockgen -source=./internal/backend.go -destination=./mock/backend.go -package=mock