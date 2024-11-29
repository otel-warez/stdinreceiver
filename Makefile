
Makefile.Common:
	@wget -q https://raw.githubusercontent.com/otel-warez/build-tools/refs/heads/main/Makefile.Common

internal/tools/empty_test.go:
	@mkdir -p internal/tools
	@wget -q https://raw.githubusercontent.com/otel-warez/build-tools/refs/heads/main/tools/empty_test.go -O internal/tools/empty_test.go

internal/tools/go.mod:
	@mkdir -p internal/tools
	@wget -q https://raw.githubusercontent.com/otel-warez/build-tools/refs/heads/main/tools/go.mod -O internal/tools/go.mod

internal/tools/go.sum:
	@mkdir -p internal/tools
	@wget -q https://raw.githubusercontent.com/otel-warez/build-tools/refs/heads/main/tools/go.sum -O internal/tools/go.sum

internal/tools/tools.go:
	@mkdir -p internal/tools
	@wget -q https://raw.githubusercontent.com/otel-warez/build-tools/refs/heads/main/tools/tools.go -O internal/tools/tools.go

.PHONY: setup
setup: internal/tools/empty_test.go internal/tools/go.mod internal/tools/go.sum internal/tools/tools.go Makefile.Common