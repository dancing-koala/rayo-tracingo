GORUN=go run
GOTEST=go test

run:
	$(GORUN) main.go

test:
	$(GOTEST) -v ./...

bench:
	$(GOTEST) -v -bench=. ./...