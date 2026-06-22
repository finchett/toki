
linter:
	docker run -t --rm -v $$(pwd):/app -w /app \
	-v $$(go env GOCACHE):/.cache/go-build -e GOCACHE=/.cache/go-build \
	-v $$(go env GOMODCACHE):/.cache/mod -e GOMODCACHE=/.cache/mod \
	-v ~/.cache/golangci-lint:/.cache/golangci-lint -e GOLANGCI_LINT_CACHE=/.cache/golangci-lint \
	-e CGO_CFLAGS="-D_LARGEFILE64_SOURCE" \
	golangci/golangci-lint:v2.12.2-alpine golangci-lint run --fix --config .golangci.yaml --timeout 5m --concurrency 4

test:
	docker run -t --rm -v $$(pwd):/app -w /app \
	-v $$(go env GOCACHE):/.cache/go-build -e GOCACHE=/.cache/go-build \
	-v $$(go env GOMODCACHE):/.cache/mod -e GOMODCACHE=/.cache/mod \
	--entrypoint "" golang:1.26.3 sh -c "go test -v -count=1 -p 4 -coverprofile=coverage.out ./... && go tool cover -func=coverage.out && go tool cover -html=coverage.out -o coverage.html"

build:
	docker run -t --rm -v $$(pwd):/app -w /app \
	-v $$(go env GOCACHE):/.cache/go-build -e GOCACHE=/.cache/go-build \
	-v $$(go env GOMODCACHE):/.cache/mod -e GOMODCACHE=/.cache/mod \
	--entrypoint "" golang:1.26.3 sh -c "go build -o tock ./cmd/tock"

# Refresh test data by running the script that generates it.
# By default, it refreshes data for the last 1 day,
refresh-test-data:
	python3 scripts/refresh_test_data.py --days $(or $(DAYS),1)

# ── Desktop app (Wails) ──────────────────────────────────────────────────
# These targets run on the host (Wails can't cross-compile macOS in Docker).
# Install the CLI first:  go install github.com/wailsapp/wails/v2/cmd/wails@latest

WAILS ?= wails
DESKTOP_DIR := cmd/tock-desktop

# Build a .app for the host architecture (fastest).
desktop-build:
	cd $(DESKTOP_DIR) && $(WAILS) build -clean
	@rm -rf $(DESKTOP_DIR)/build/bin/Toki.app
	@mv $(DESKTOP_DIR)/build/bin/tock-desktop.app $(DESKTOP_DIR)/build/bin/Toki.app
	@echo "Built $(DESKTOP_DIR)/build/bin/Toki.app"

# Build a universal (arm64 + amd64) .app suitable for distribution.
desktop-build-universal:
	cd $(DESKTOP_DIR) && $(WAILS) build -clean -platform darwin/universal
	@rm -rf $(DESKTOP_DIR)/build/bin/Toki.app
	@mv $(DESKTOP_DIR)/build/bin/tock-desktop.app $(DESKTOP_DIR)/build/bin/Toki.app
	@echo "Built $(DESKTOP_DIR)/build/bin/Toki.app"

# Build and open the resulting .app.
desktop-run: desktop-build
	open $(DESKTOP_DIR)/build/bin/Toki.app

# Live-reload dev server with Go bindings exposed at http://localhost:34115.
desktop-dev:
	cd $(DESKTOP_DIR) && $(WAILS) dev

# Check that the Wails CLI and its prerequisites are installed.
desktop-doctor:
	cd $(DESKTOP_DIR) && $(WAILS) doctor

# Regenerate THIRD_PARTY_NOTICES.txt for bundled Go and npm deps.
# Requires go-licenses, license-checker-rseidelsohn, and jq on PATH.
notices:
	./scripts/gen-notices.sh

.PHONY: desktop-build desktop-build-universal desktop-run desktop-dev desktop-doctor notices
