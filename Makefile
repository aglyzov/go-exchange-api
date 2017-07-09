
GO ?= go

.PHONY: help clean build test

help:                    # display this help
	@echo "Targets"
	@echo "-------"
	@grep -iE '^[a-z\-]+:' Makefile | sed -Ee 's/://; s/#/-/'

clean:                   # clean up working dir
	@echo "Cleaning"
	@echo "--------"
	rm -rf bin
	rm -rf pkg
	${GO} clean -v

build:                   # build the crypto-ticker binary
	@echo "Building"
	@echo "--------"
	mkdir -p bin
	${GO} build -o ./bin/crypto-ticker .

test:                    # test everything
	@echo "Testing"
	@echo "-------"
	${GO} test ./...
