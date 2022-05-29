.PHONY: build test clean

GO=go

build:
	@$(GO) build .

test:
	@$(GO) test -v .

clean:
	@rm co
