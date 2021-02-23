.PHONY: build
build:
	@go build -v .

.PHONY: clean
clean:
	go clean -i .

.PHONY: dev
dev:
	@air

.PHONY: docker
docker:
	@docker build -t github.com/bakeruk/kubernetes-hello-world-api .

.PHONY: lint
lint:
	golint ./...

.PHONY: tools
tools:
	go install golang.org/x/lint/golint
	go install github.com/cosmtrek/air