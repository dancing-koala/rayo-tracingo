GORUN=go run
GOTEST=go test

run:
	$(GORUN) main.go vec3.go

test:
	$(GOTEST) -v ./...

bench:
	$(GOTEST) -v -bench=. ./...