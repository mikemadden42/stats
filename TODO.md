# TODO

## Bugs and misleading output

- [ ] **Memory sizes are in base-1000 units.** `humanize.Bytes` (`main.go:39-45`) uses powers of 1000, so a 24 GiB machine shows `total mem: 26 GB`. Switch to `humanize.IBytes`.
- [ ] **Virtual memory lines show `0 B` on macOS.** go-sysinfo doesn't fill in `VirtualFree/Total/Used` there. Print them only when `VirtualTotal > 0`.
- [ ] **Uptime prints to the microsecond** (e.g. `72h6m11.983006s`). Use `hostInfo.Uptime().Truncate(time.Second)`.
- [ ] **MAC and IP lists repeat entries.** macOS hides real MACs, so `02:00:00:00:00:00` appears many times, and some IPs are duplicated. Remove duplicates and skip all-zero/placeholder MACs.
- [ ] **Extra space in list items.** `fmt.Println("\t", ip)` adds a space after the tab. Use `fmt.Printf("\t%s\n", ip)`.
- [ ] **Errors panic.** `checkErr` panics with a stack trace. Use `log.Fatal(err)` (or stderr + `os.Exit(1)`), and make a failing `host.Memory()` non-fatal so the Go info still prints.

## Lint config (`.golangci.yml`) is invalid for golangci-lint v2

`golangci-lint config verify` rejects it, so settings are ignored. `golangci-lint migrate` can convert it.

- [x] `version: 2` → `version: "2"`
- [x] `linters.disable-all: true` → `linters.default: none`
- [x] `linters-settings:` → `linters.settings:`; move goimports to `formatters.settings.goimports.local-prefixes` (as a list)
- [x] `issues.exclude-dirs` → `linters.exclusions.paths`
- [x] Add a package comment to `main.go` (revive `package-comments`)

## GoReleaser config (`.goreleaser.yaml`) is out of date

`goreleaser check` fails.

- [x] Add `version: 2` at the top
- [x] `snapshot.name_template` → `snapshot.version_template`
- [x] Remove the no-op `go generate ./...` hook and the boilerplate comments
- [x] Add `dist/` to `.gitignore`

## CI (`.github/workflows/go.yml`)

- [x] Add a `golangci/golangci-lint-action` step
- [x] Target `oldstable` and `stable` in the build matrix
- [ ] Add a `go run .` smoke-test step

## Dependencies

- [ ] Update deps (`go get -u ./... && go mod tidy`): `go-sysinfo` v1.15.5, `go-humanize` v1.1.0, `golang.org/x/sys` v0.25.0 → v0.50.0 (indirect)

## Housekeeping

- [ ] Delete `update_go.log` or add it to `.gitignore`
- [ ] Expand the README with sample output and install/run instructions
- [ ] Note that the output includes the machine's unique ID (`id:`), so be careful pasting it publicly
