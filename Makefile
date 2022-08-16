GODOC=localhost:4242
IGNORED_FOLDER=.ignore
COVERAGE_FILE=coverage.txt
MODULE := $(shell go list -m)

#all:	@ run all main rules test, lint and build (assume the dependencies and tools are installed, else voir make `install`, `tool` and `citool`) 
all: test lint build

.PHONY: help
#help:	@ List available rules on this project
help: 
	@grep -E '[a-zA-Z\.\-]+:.*?@ .*$$' $(MAKEFILE_LIST)| sort | tr -d '#'  | awk 'BEGIN {FS = ":.*?@ "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: install
#install:	@ Install dependencies
install:
	@go mod download

.PHONY: tool
#tool:	@ Install tooling
tool:
	go install golang.org/x/tools/...@latest
	go install github.com/vektra/mockery/v2/...@latest
	@echo "mockery: " `mockery --version`

.PHONY: citool
#citool:	@ ci tooling installation
citool:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.48.0
	@echo "golangci-lint: " `golangci-lint version`

.PHONY: build
#build:	@ Build packages or binaries
build:
	CGO_ENABLED=1 go build -tags static -ldflags "-s -w" ./...

.PHONY: test-bench
#test-bench:	@ run benchmarks tests
test-bench:
	CGO_ENABLED=1 go test -count=1 -run "^$$" -benchmem -benchtime 1000x -bench "^(Benchmark).*" -v ./...

.PHONY: test-race
#test-race:	@ Run race detection tests
test-race:
	CGO_ENABLED=1 go build -race -a -tags static -ldflags "-s -w" ./...

.PHONY: test-unit
#test-unit:	@ Run units tests
test-unit: --private-create-ignored-folder
	go test -v -count=1 -race -coverprofile=${IGNORED_FOLDER}/${COVERAGE_FILE} -covermode=atomic ./... | sed ''/PASS/s//`printf "\033[32mPASS\033[0m"`/'' | sed ''/FAIL/s//`printf "\033[35mFAIL\033[0m"`/'' | sed ''/RUN/s//`printf "\033[36mRUN\033[0m"`/''
	@go tool cover -func ${IGNORED_FOLDER}/${COVERAGE_FILE} | grep total:

.PHONY: test
#test:	@ Run the tests suite (bench, race and units)
test: test-bench test-race test-unit
	@echo "tests executed !"

.PHONY: generate
#generate:	@ Update the generated resources
generate:
	@GOMODLOCATION=$$PWD go generate ./...

.PHONY: lint
#lint:	@ Run linter validation
lint:
	golangci-lint run --skip-dirs mocks

.PHONY: godoc
#godoc:	@ Run a server to render godoc server
godoc:
	$(eval pid := ${shell nohup godoc -http=${GODOC} >> /dev/null & echo $$! ; })
	@echo "server started:"
	@echo "\tDoc location: http://${GODOC}/pkg/${MODULE}"
	@echo "\texecute the following command to turn off server: kill $(pid)"

.PHONY: clean
#clean:	@ cleanup the ignored and vendor folders
clean:
	@rm -rf ${IGNORED_FOLDER}

.PHONY: fclean
#fclean:	@ invoke clean and remove binaries
fclean: clean

# Privates function

--private-create-ignored-folder:
	@if [ ! -d ${IGNORED_FOLDER} ]; then \
		mkdir -p ${IGNORED_FOLDER}; \
	fi
