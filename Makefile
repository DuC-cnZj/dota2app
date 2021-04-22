.PHONY: fmt
fmt:
	gofmt -w ./ && goimports -w ./