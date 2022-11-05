.PHONY: format
format:
	$(call format)

define format
	@go run mvdan.cc/gofumpt -l -w .
	@go run golang.org/x/tools/cmd/goimports -w .
	@go mod tidy
endef
