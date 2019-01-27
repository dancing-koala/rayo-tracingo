GORUN=go run
GOTEST=go test
BIN=rayo-tracingo

run:
	$(GORUN) *.go

test:
	$(GOTEST) -v ./...

bench:
	$(GOTEST) -v -bench=. ./...

build:
	go build -o $(BIN)

profiling: run
	pprof -http=:5000 ./$(BIN).prof