# Oasis Starter dApp in Go

This is a skeleton for confidential Oasis dApps in Go.

## Prerequisites

To build the example MessageBox contract, you need `solc` and `abigen`.
If you're on Ubuntu, you can install both with:

```shell
make install-deps
```

To build the rest of the project, you need Go.

## Building

To build everything, simply run:

```shell
make
```

## Testing

To run the end-to-end test, you should first start the Sapphire Localnet
in Docker:

```shell
make run-localnet
```

After it has finished starting up, you can run the end-to-end test:

```shell
make test
```

This will deploy the MessageBox contract, store a message in it, retrieve it,
and check if the retrieved message matches the stored one.

## Running

Before deploying the contract or interacting with it, you should store your
account's hex-encoded private key in an environment variable (the `0x` prefix
is optional):

```shell
export PRIVATE_KEY=...
```

### Deploying the contract

You can deploy the contract on different networks by invoking:

```shell
./demo-starter deploy --network sapphire-localnet # Sapphire Localnet
./demo-starter deploy --network sapphire-testnet  # Sapphire Testnet
./demo-starter deploy --network sapphire          # Sapphire Mainnet
```

The deployed contract's address is printed to the standard output if the
deployment is successful.  You can store it in an environment variable,
as you will need it to interact with the contract.

### Interacting with the contract

The example MessageBox contract has two methods: `setMessage` and `message`.

To store a private message inside the deployed contract:

```shell
./demo-starter setMessage --network sapphire-localnet ${CONTRACT_ADDR} "Hello!"
```

To retrieve the private message from the deployed contract:

```shell
./demo-starter message --network sapphire-localnet ${CONTRACT_ADDR}
```

If you try to retrieve the message using a different account than the one
that was used to store it, the retrieval will fail.

## Debugging

For debugging purposes, you can also run the localnet in debug mode with:

```shell
make run-localnet-debug
```

Inside the Docker container, the Web3 gateway logs are located in
`/var/log/oasis-web3-gateway.log`, while the Oasis node logs are located
under the `/serverdir/node/net-runner/network/...` hierarchy (each node
has its subfolder and its log is in `node.log` inside that subfolder).

