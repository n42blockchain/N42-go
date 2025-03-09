# N42 Blockchain

[![Go](https://img.shields.io/badge/go-1.19%2B-blue.svg)](https://golang.org)
[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/n42world/n42/ci.yml?branch=main)](https://github.com/n42world/n42/actions)
[![GitHub License](https://img.shields.io/github/license/n42world/n42)](https://github.com/n42world/n42/blob/main/LICENSE)
[![GitHub Issues](https://img.shields.io/github/issues/n42world/n42)](https://github.com/n42world/n42/issues)
[![GitHub Pull Requests](https://img.shields.io/github/issues-pr/n42world/n42)](https://github.com/n42world/n42/pulls)
[![GitHub Stars](https://img.shields.io/github/stars/n42world/n42)](https://github.com/n42world/n42/stargazers)
[![GitHub Forks](https://img.shields.io/github/forks/n42world/n42)](https://github.com/n42world/n42/network/members)

## Introduction

N42 introduces a secure, efficient, and globally interconnected digital ecosystem, empowering developers to build applications with maximum autonomy and interoperability. Designed as a high-performance public blockchain, N42 utilizes Go for its exceptional concurrency, scalability, and ease of deployment, ensuring a robust and highly efficient environment.

With its modular and sharded architecture, N42 provides advanced transaction throughput and data processing capabilities, critical for developing a globally connected digital infrastructure. Its permissionless design facilitates seamless integration and efficient data exchange across diverse application environments, laying the foundation for the next generation of decentralized internet services.

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
git clone https://github.com/N42world/n42.git
cd n42
make n42
./build/bin/n42
```

### Windows

Windows users may run n42 in 3 possible ways:

* Build executable binaries natively for Windows using [Chocolatey package manager](https://chocolatey.org/)
* Use Docker :  see [docker-compose.yml](./docker-compose.yml)
* Use WSL (Windows Subsystem for Linux) **strictly on version 2**. Under this option you can build n42 just as you would on a regular Linux distribution. You can point your data also to any of the mounted Windows partitions (eg. `/mnt/c/[...]`, `/mnt/d/[...]` etc) but in such case be advised performance is impacted: this is due to the fact those mount points use `DrvFS` which is a [network file system](#blocks-execution-is-slow-on-cloud-network-drives) and, additionally, MDBX locks the db for exclusive access which implies only one process at a time can access data.  This has consequences on the running of `rpcdaemon` which has to be configured as [Remote DB](#for-remote-db) Even if it runs on the same computer, limitations may still apply. However, if your data is stored on the native Linux filesystem, these limitations do not apply. **Please be aware that the default WSL2 environment operates with a distinct IP address, separate from the Windows host's network interface. This distinction is crucial when setting up NAT for port 30303 on your router.  **

### Docker container
Docker allows for building and running n42 via containers. This alleviates the need for installing build dependencies onto the host OS.
see [docker-compose.yml](./docker-compose.yml) [dockerfile](./Dockerfile).
For convenience we provide the following commands:
```sh
make images # build docker images than contain executable n42 binaries
make up # alias for docker-compose up -d && docker-compose logs -f 
make down # alias for docker-compose down && clean docker data
make start #  alias for docker-compose start && docker-compose logs -f 
make stop # alias for docker-compose stop
```

## Executables

The N42 project comes with one wrappers/executables found in the `cmd`
directory.

|    Command    | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| :-----------: | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
|  **`N42`**   | Our main N42 CLI client.  It can be used by other processes as a gateway into the N42 network via JSON RPC endpoints exposed on top of HTTP transports. `N42 --help`  for command line options.          |


## n42 ports

| Port  | Protocol |         Purpose          |       Expose       |
|:-----:|:--------:|:------------------------:|:------------------:|
| 61015 |   UDP    | The port used by discv5. |       Public       |
| 61016 |   TCP    | The port used by libp2p. |       Public       |
| 20012 |   TCP    |      Json rpc/HTTP       |       Public       |
| 20013 |   TCP    |    Json rpc/Websocket    |       Public       |
| 20014 |   TCP    | Json rpc/HTTP/Websocket  | JWT Authentication |
| 4000  |   TCP    |   BlockChain Explorer    |       Public       |
| 6060  |   TCP    |         Metrics          |      Private       | 
| 6060  |   TCP    |          Pprof           |      Private       | 

## License
The N42 library is licensed under the
[GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html).

