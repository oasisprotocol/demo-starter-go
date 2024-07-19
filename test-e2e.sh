#!/bin/sh

set -eu

# Private key of the first test account.
export PRIVATE_KEY="ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"

# Network to use.
NETWORK="sapphire-localnet"

# Message to store in the box.
MSG="Hello, world!"

# Deploy the contract, its address is returned on stdout.
ADDR=`./demo-starter deploy --network ${NETWORK}`
echo "Contract address is: ${ADDR}"

# Store the message inside the deployed contract.
./demo-starter setMessage --network ${NETWORK} "${ADDR}" "${MSG}"

# Retrieve the message from the deployed contract.
MSG_GOT=`./demo-starter message --network ${NETWORK} "${ADDR}"`

# Check if the retrieved message matches the stored message.
if [ "x${MSG}" = "x${MSG_GOT}" ]; then
	echo "Test passed!"
	echo "Stored \"${MSG}\", got \"${MSG_GOT}\"."
	exit 0
else
	echo "Test failed!"
	echo "Expected \"${MSG}\", got \"${MSG_GOT}\"."
	exit 1
fi

