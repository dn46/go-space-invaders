build:
	@go build -o bin/goRayLib

run: build
	@./bin/goRayLib

test:
	@go test -v ./...