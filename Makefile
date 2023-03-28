BUILDPATH=$(CURDIR)
API_NAME=http_tool

build: 
	@echo "Building binary file ..."
	@go build -mod=vendor -ldflags '-s -w' -o $(BUILDPATH)/build/${API_NAME} cmd/main.go
	@echo "Binary file created at build/${API_NAME}"

test: 
	@echo "Running tests..."
	@go test ./... --coverprofile coverfile_out >> /dev/null
	@go tool cover -func coverfile_out
