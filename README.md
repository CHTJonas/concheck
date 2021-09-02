# Concheck

Here lies the code for a teeny tiny go application to test your internet connectivity. It tries to ping a few endpoints, resolve a few DNS queries and send a few HTTP requests. If these are successful then it exits cleanly, otherwise it quits with a non-zero exit code.

## Usage

Just run the `concheck` binary - that's it! I designed it to be run as part of a shell pipeline in a cronjob thereby avoiding cronspam when my (ADSL) internet connection dies e.g.:

```* * * * * /usr/local/bin/concheck && curl -sS https://api.service.com/webhook```

On Linux you will need to run the binary with `CAP_NET_RAW` privileges in order to send ICMP packets (running as root is strongly discouraged). You can do this by running `sudo setcap cap_net_raw=+ep /usr/local/bin/concheck` in a terminal.

The (non-zero) exit code determines the type of connectivity failure:

| Exit Code | What's not working?      |
| --------- | ------------------------:|
| 1         | HTTP connectivity broken |
| 2         | ICMP connectivity broken |
| 3         | DNS resolution broken    |

There are two command line flags to explicitly force the connectivity checks to use either IPv4 or IPv6. These are `-4` and `-6` respectively. In most situations however, you will not need to use either. We're clever enough to figure out when you're running with an IPv4- or IPv6-only connection and squelch out the "network unreachable" errors in those situations.

## Installation

Pre-built binaries for a variety of operating systems and architectures are available to download from [GitHub Releases](https://github.com/CHTJonas/concheck/releases). If you wish to compile from source then you will need a suitable [Go toolchain installed](https://golang.org/doc/install). After that just clone the project using Git and run Make! Cross-compilation is easy in Go so by default we build for all targets and place the resulting executables in `./bin`:

```bash
git clone https://github.com/CHTJonas/concheck.git
cd concheck
make clean && make all
```

## Copyright

concheck is licensed under the [BSD 2-Clause License](https://opensource.org/licenses/BSD-2-Clause).

Copyright (c) 2021 Charlie Jonas.
