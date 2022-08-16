APP_NAME        = "bsuir-schedule"
DOCKER_TAG      ?= "v0.0.2"
GOPATH          := $(shell go env GOPATH)
GRC             := $(shell which grc)

vet:
	@$(GRC) go vet ./...

test: format vet
	@$(GRC) go test -v -cover -coverprofile=.coverprofile ./...

build:
	@go build -ldflags "-X main.Version=$(DOCKER_TAG)" -o ./bin/$(APP_NAME)

install:
	@go build -ldflags "-X main.Version=$(DOCKER_TAG)" -o ${GOPATH}/bin/$(APP_NAME)
	@echo "Installed $(APP_NAME) to ${GOPATH}/bin/$(APP_NAME)"

run:
	@./bin/$(APP_NAME)

format:
	@go fmt ./...

docker-build:
	@docker build . -t $(APP_NAME):$(DOCKER_TAG)

docker-run:
	@docker run --rm -it --name $(APP_NAME) $(APP_NAME):$(DOCKER_TAG) 