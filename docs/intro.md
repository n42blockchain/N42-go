# N42 DOCS
Documentation provided for N42 users and developers.

[![Telegram Chat][tg-badge]][tg-url]

N42 (N42 Chain) is a blockchain full node implementation characterized by being user-friendly, highly modular, and fast and efficient.

## What is this about?

N42 is a node implementation compatible with all node protocols that support N42 Chain.

It was originally built and promoted by N42, licensed under Apache and MIT licenses.

As a complete N42 Chain node, N42 allows users to connect to the N42 Chain network and interact with the N42 Chain blockchain.

This includes sending and receiving transactions, querying logs and traces, as well as accessing and interacting with smart contracts.

Creating a successful N42 Chain node requires a high-quality implementation that is both secure and efficient, and easy to use on consumer hardware. It also requires building a strong community of contributors to help support and improve the software.

## What are the goals of N42?

**1. Modularity**

Every component of N42 is built as a library: well-tested, heavily documented, and benchmarked. We envision developers importing the node's packages, mixing and matching, and innovating on top of them.

Examples of such usage include, but are not limited to, launching standalone P2P networks, talking directly to a node's database, or "unbundling" the node into the components you need.

To achieve this, we are licensing N42 under the Apache/MIT permissive license.

**2. Performance**

N42 aims to be fast, so we used Golang and parallel virtual machine sync node architecture.

We also used tested and optimized N42 Chain libraries.

**3. Free for anyone to use any way they want**

N42 is free open-source software, built by the community for the community.

By licensing the software under the Apache/MIT license, we want developers to use it without being bound by business licenses, or having to think about the implications of GPL-like licenses.

**4. Client Diversity**

The N42 Chain protocol becomes more antifragile when no node implementation dominates. This ensures that if there's a software bug, the network does not confirm a wrong block. By building a new client, we hope to contribute to N42 Chain's antifragility.

**5. Used by a wide demographic**

We aim to solve for node operators who care about fN42 historical queries, but also for hobbyists who cannot operate on large hardware.

We also want to support teams and individuals who want both sync from genesis and via "fast sync".

We envision that N42 will be flexible enough for the trade-offs each team faces.

## Who is this for?

N42 is a new N42 Chain full node allowing users to sync and interact with the entire blockchain, including its historical state if in archive mode.

- Full node: It can be used as a full node, storing and processing the entire blockchain, validating blocks and transactions, and participating in the consensus process.

- Archive node: It can also be used as an archive node, storing the entire history of the blockchain, which is useful for applications that need access to historical data. As a data engineer/analyst, or as a data indexer, you'll want to use Archive mode. For all other use cases where historical access is not needed, you can use Full mode.

## Is this secure?

N42 implements the specification of N42 Chain as defined in the repository. To ensure the node is built securely, we run the following tests:

1. Virtual machine state tests are run on every Pull Request.
2. We regularly re-sync multiple nodes from scratch.
3. We operate multiple nodes at the tip of N42 Chain mainnet and various testnets.
4. We extensively unit test, fuzz test, and document all our code, while also restricting PRs with aggressive lint rules.
5. We also plan to audit/fuzz the virtual machine & parts of the codebase. Please reach out if you're interested in collaborating on securing this codebase.

We intend to also audit/fuzz the EVM & parts of the codebase. Please reach out if you're interested in collaborating on securing this codebase.

## Sections

Here are some useful sections to jump to:

- Install N42 by following the [guide](./installation/installation.md).
- Sync your node on any [official network](./run/run-a-node.md).
- View [statistics and metrics](./run/observability.md) about your node.
- Query the [JSON-RPC](./jsonrpc/intro.md) using Foundry's `cast` or `curl`.
- Set up your [development environment and contribute](./developers/contribute.md)!

> 📖 **About this book**
>
> The book is continuously rendered [here](https://github.com/WeAreN42/N42/docs)!
> You can contribute to this book on [GitHub][gh-book].

[tg-badge]: https://img.shields.io/endpoint?color=neon&logo=telegram&label=chat&url=https%3A%2F%2Ftg.sumanjay.workers.dev%2Fparadigm%5Freth
[tg-url]: https://t.me/N42
[gh-book]: https://github.com/WeAreN42/N42/docs
