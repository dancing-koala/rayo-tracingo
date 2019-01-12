GORUN=go run
GOTEST=go test

run:
	$(GORUN) *.go

test:
	$(GOTEST) -v ./...

bench:
	$(GOTEST) -v -bench=. ./...