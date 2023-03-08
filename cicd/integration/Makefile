# A=$(shell echo hi)

ARCH=$(shell go env GOARCH)
MARCH=$(shell go env GOOS)-$(shell go env GOARCH)

.PHONY: test
test: bb


.PHONY: bb
bb:
	@echo "hihi"
	@mkdir -p hi
	@echo $(ARCH)
	@echo $(MARCH)


GOOS=darwin GOARCH=arm64 go build