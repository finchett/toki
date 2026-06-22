# Toki desktop

Wails (Go + WebKit) shell wrapping the Toki time tracker, with a React + TypeScript frontend.

## Prerequisites

- macOS 11+ (Big Sur or newer)
- Go (matching the version in `go.mod`)
- Node.js 18+
- Wails CLI:

  ```sh
  go install github.com/wailsapp/wails/v2/cmd/wails@latest
  ```

  Ensure `$(go env GOPATH)/bin` is on your `PATH`. Run `wails doctor` (or `make desktop-doctor` from the repo root) to verify the toolchain.

## Build

From the repo root:

```sh
make desktop-build              # host architecture, ~7s incremental
make desktop-build-universal    # arm64 + amd64 fat binary
make desktop-run                # build, then `open Toki.app`
```

The packaged app lands at `cmd/tock-desktop/build/bin/Toki.app` (rename happens after Wails packaging because Wails derives the bundle directory from the project name).

## Develop

```sh
make desktop-dev
```

Runs `wails dev` with Vite hot reload. Go methods are reachable from a browser at <http://localhost:34115> for direct devtools poking.

## Menu bar

Toki lives in the macOS menu bar while it's running. The status item shows `●`
plus the elapsed time of the running activity (or `○` when nothing is tracked).
Closing the window leaves Toki in the menu bar — click **Show Toki** there to
bring the window back, or **Quit Toki** to fully exit.

## Configuration

App identity (bundle id, version, copyright) lives in `wails.json` under the `info` block, with platform-specific overrides in `build/darwin/Info.plist` (and `Info.dev.plist` for `wails dev`).
