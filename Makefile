APP_NAME        = "bsuir-schedule"
DOCKER_TAG      ?= "v0.0.1"
GOPATH          := $(shell go env GOPATH)
GRC             := $(shell which grc)

test:
	@$(GRC) go test -v ./...

build:
	@go build -ldflags "-X main.Version=$(DOCKER_TAG)" -o ./bin/$(APP_NAME)

build_windows_x86:
	@GOOS=windows GOARCH=386 go build -ldflags "-X main.Version=$(DOCKER_TAG)" -o ./bin/$(APP_NAME)_x86.exe

build_windows_x64:
	@GOOS=windows GOARCH=amd64 go build -ldflags "-X main.Version=$(DOCKER_TAG)" -o ./bin/$(APP_NAME)_x64.exe

install:
	@go build -ldflags "-X main.Version=$(DOCKER_TAG)" -o ${GOPATH}/bin/$(APP_NAME)

run:
	@./bin/$(APP_NAME)

format:
	@go fmt ./...

docker-build:
	@docker build . -t $(APP_NAME):$(DOCKER_TAG)

docker-run:
	@docker run --rm -it --name $(APP_NAME) $(APP_NAME):$(DOCKER_TAG) 

# release: docker-build
# 	@docker tag $(DOCKER_REGISTRY)/$(APP_NAME):$(DOCKER_TAG) $(DOCKER_REGISTRY)/$(APP_NAME):latest
# 	@docker push $(DOCKER_REGISTRY)/$(APP_NAME):$(DOCKER_TAG)
# 	@docker push $(DOCKER_REGISTRY)/$(APP_NAME):latest