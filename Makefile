
.EXPORT_ALL_VARIABLES:
GO111MODULE := on
SRC_DIRS := slice utils

all: generate copy-generated lint fmt sec test

dirs:
	@echo $(SRC_DIRS)

get:
	@go get -u .

generate:
	@go run -v ./codegen

copy-generated:
	@for dir in $(SRC_DIRS); do \
		cp -r ./codegen/$$dir/generated/*.go $$dir; \
	done

test:
	@for dir in $(SRC_DIRS); do \
		(echo $$dir && cd $$dir && go test . -coverpkg=./... -covermode=atomic -coverprofile=./coverage.txt -v -count 1;) \
	done

bench:
	@for dir in $(SRC_DIRS); do \
		(echo $$dir && cd $$dir && go test -bench=.;) \
	done

tidy:
	@go mod tidy -v

fmt:
	@for dir in $(SRC_DIRS); do \
		gofmt -s -w ./$$dir/*.go; \
	done

lint:
	@for dir in $(SRC_DIRS); do \
		golint $$dir/...; \
	done
	golint codegen/...

# https://github.com/securego/gosec
sec:
	@gosec -exclude-dir codegen ./...

# https://github.com/fzipp/gocyclo
cyclo-top-10:
	@gocyclo -top 10 -ignore "codegen" .

go-mark-doc:
	@for dir in $(SRC_DIRS); do \
		gomarkdoc ./$$dir > $$dir.md; \
	done

.PHONY: get generate copy-generated test build lint fmt sec cyclo-top-10 tidy go-mark-doc