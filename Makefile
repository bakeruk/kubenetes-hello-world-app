.PHONY: build
build:
	@go build -v .

.PHONY: clean
clean:
	go clean -i .

.PHONY: lint
lint:
	golint ./...

.PHONY: tools
tools:
	go install golang.org/x/lint/golint