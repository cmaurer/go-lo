
.EXPORT_ALL_VARIABLES:

SRC_DIRS := slice utils

all: generate copy-generated lint fmt test

dirs:
	@echo $(SRC_DIRS)

generate:
	go run -v ./codegen

copy-generated:
	@for dir in $(SRC_DIRS); do \
		cp -r ./codegen/$$dir/generated/*.go $$dir; \
	done

test:
	@for dir in $(SRC_DIRS); do \
		(echo $$dir && cd $$dir && go test . -coverpkg=./... -covermode=atomic -coverprofile=./coverage.txt -v -count 1;) \
	done

fmt:
	@for dir in $(SRC_DIRS); do \
		gofmt -s -w ./$$dir/*.go; \
	done


lint:
	@for dir in $(SRC_DIRS); do \
		golint $$dir/...; \
	done
	golint codegen/...


.PHONY: generate copy-generated test build lint fmt