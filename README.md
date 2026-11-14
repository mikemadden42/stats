# stats

Display operating system stats

![example workflow](https://github.com/mikemadden42/stats/actions/workflows/go.yml/badge.svg)

## Install

```sh
go install github.com/mikemadden42/stats@latest
```

Or run it from a clone with `go run .`.

## Example output

```
arch: arm64
boot time: 3 days ago
hostname: example.local
ips:
	127.0.0.1/8
	::1/128
	10.0.0.19/24
kernel: 27.2.0
os: macOS 27.2.1
timezone: CST
timezone offset: -21600
id: 00000000-0000-0000-0000-000000000000
uptime: 72h12m49s

available mem: 9.7 GiB
free mem: 496 MiB
total mem: 24 GiB
used mem: 14 GiB

go arch: arm64
go cores: 8
go os: darwin
go version: go1.27.4
```

Sections the OS doesn't report (such as MAC addresses or virtual memory on
macOS) are left out.

> **Note:** the `id:` line is the machine's unique ID, and the output also
> includes the hostname and IP addresses. Redact them before sharing output
> publicly.
