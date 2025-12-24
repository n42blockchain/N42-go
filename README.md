# N42 Blockchain

[![Go](https://img.shields.io/badge/go-1.19%2B-blue.svg)](https://golang.org)
[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/n42blockchain/n42/ci.yml?branch=main)](https://github.com/n42blockchain/n42/actions)
[![GitHub License](https://img.shields.io/github/license/n42blockchain/n42)](https://github.com/n42blockchain/n42/blob/main/LICENSE)
[![GitHub Issues](https://img.shields.io/github/issues/n42blockchain/n42)](https://github.com/n42blockchain/n42/issues)
[![GitHub Pull Requests](https://img.shields.io/github/issues-pr/n42blockchain/n42)](https://github.com/n42blockchain/n42/pulls)
[![GitHub Stars](https://img.shields.io/github/stars/n42blockchain/n42)](https://github.com/n42blockchain/n42/stargazers)
[![GitHub Forks](https://img.shields.io/github/forks/n42blockchain/n42)](https://github.com/n42blockchain/n42/network/members)

## Introduction

N42 establishes a secure, efficient, and globally interconnected digital ecosystem, giving developers high autonomy and seamless interoperability for building applications. As a high-performance public blockchain, N42 is built using Go, leveraging its superior concurrency, scalability, and ease of deployment to deliver a robust, high-performance infrastructure.

Featuring a modular, sharded architecture, N42 delivers high transaction throughput and advanced data processing—key elements for building globally connected digital infrastructure. Its permissionless framework enables effortless integration and efficient data exchange across a wide range of applications, paving the way for the next generation of decentralized internet services.

**Disclaimer:** This software is currently a tech preview. We will do our best to keep it stable and avoid breaking changes, but we make no guarantees.

## System Requirements

- **Storage**: ≥ 200 GB (SSD or NVMe recommended; HDD not recommended)
- **Memory**: ≥ 16 GB RAM
- **CPU**: 64-bit architecture
- **Go Version**: [≥ 1.19](https://golang.org/doc/install)

## Building from Source

### Linux and macOS

To build N42 from source, you must have the latest version of Go installed.

- Installation instructions: [Go installation page](https://golang.org/doc/install)
- Download Go: [Go download page](https://golang.org/dl/)

Clone the repository and compile:

```sh
git clone https://github.com/n42blockchain/n42.git
cd n42
make n42
./build/bin/n42
```

### Windows

Windows users may run N42 in three ways:

- **Native binaries**: Build using [Chocolatey](https://chocolatey.org/)
- **Docker**: See [docker-compose.yml](./docker-compose.yml)
- **WSL2 (Windows Subsystem for Linux)**:
    - Install and build as on Linux
    - Ensure storage is on Linux filesystem for best performance

### Docker Container

Docker allows easy building and running without installing dependencies on the host OS.

See: [docker-compose.yml](./docker-compose.yml), [Dockerfile](./Dockerfile)

Convenient Docker commands:

```sh
make images # Build Docker images containing N42 binaries
make up     # docker-compose up -d && docker-compose logs -f
make down   # docker-compose down && clean docker data
make start  # docker-compose start && docker-compose logs -f
make stop   # docker-compose stop
```

## Executables

N42 includes one main executable located in the `cmd` directory:

| Command | Description |
|---------|-------------|
| **`n42`** | Main CLI client, provides JSON RPC endpoints over HTTP transports. Use `n42 --help` for options. |

## Network Ports

| Port  | Protocol | Purpose                      | Exposure            |
|-------|----------|------------------------------|---------------------|
| 61015 | UDP      | Discovery v5                 | Public              |
| 61016 | TCP      | libp2p Communication         | Public              |
| 20012 | TCP      | JSON RPC over HTTP           | Public              |
| 20013 | TCP      | JSON RPC over Websocket      | Public              |
| 20014 | TCP      | Secure JSON RPC (JWT Auth)   | Authenticated       |
| 4000  | TCP      | Blockchain Explorer          | Public              |
| 6060  | TCP      | Metrics & Profiling (pprof)  | Private             |

## License

N42 is licensed under the [GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html).

