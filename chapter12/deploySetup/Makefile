PROJECT_NAME=bookServer
BINARY_NAME=app
GOCMD=go
GOBUILD=$(GOCMD) build
CGO_ENABLED=0
GOOS=linux
GOARCH=amd64

all:
	make build && make deploy

build:
	$(info Building the book server binary...)
	cd ${PROJECT_NAME} && ${GOBUILD} -o "$(BINARY_NAME)" -v

deploy:
	$(info Deploying...)
	docker-compose rm -f
	docker-compose up --build