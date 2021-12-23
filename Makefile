BINARY_NAME = k7z

.PHONY: build
build:
	go build -v -o $(BINARY_NAME) ./cmd/k7z

.PHONY: run
run: build
	./$(BINARY_NAME) 

.PHONY: test
test: 
	$(TEST_COMMAND) -cover -parallel 5 -failfast  ./... 

.PHONY: tidy
tidy:
	go mod tidy

# auto restart bot (using fiber CLI <3)
.PHONY: dev
dev:
	fiber dev -t ./cmd/k7z/
