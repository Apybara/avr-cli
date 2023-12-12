# Aleo Validator Registry CLI

## Prerequisites for avr.aleo owner
- Install SnarkOS (this requires the installation of SnarkOS and its dependencies)
- This CLI requires the `snarkos` binary to be on the environment path.
- Deploy the avr.aleo program https://github.com/Apybara/avr-program
- Make sure that the owner of the avr.aleo program has enough funds to pay for the transaction fees. 

## Prerequisites for validator
- Install SnarkOS (this requires the installation of SnarkOS and its dependencies)
- This CLI requires the `snarkos` binary to be on the environment path.
- Make sure that the validator has enough funds to pay for the transaction fees.

## Installation
- `make build`

## Running as a avr.aleo owner
### Deploy the avr.aleo program
go to https://github.com/Apybara/avr-program and deploy the aleo program

### Register new set of validators
**Note:** This command can only be called by the owner of the avr.aleo Program.

```bash
./avr-cli register-validator --name="aleo node" --description="aleo validator node" --website-url="https://aleo.org" --logo-url="https://aleo.org" --validator=aleo1rt3vjrusjvd6wje97efl3ra78k0d6f4c3zn8avuym0qwkl4njv9shhmfsk --private-key=<owner of the avr.aleo private key>`
```

## Running as a validator
### Register validator information
**Note:** This command can only be called by the calling validator.

```bash
./avr-cli register-my-validator --name="aleo node" --description="aleo validator node" --website-url="https://aleo.org" --logo-url="https://aleo.org" --validator=aleo1rt3vjrusjvd6wje97efl3ra78k0d6f4c3zn8avuym0qwkl4njv9shhmfsk --private-key=<validator private key>`
```

### Utilities
#### Generate a field value for an input
```bash
./avr-cli input-field --value=126207244316550804821666916field
Value: hello world
```

#### Convert the field input to a string value
```bash
./avr-cli output-field --value="hello world"
Value: 126207244316550804821666916field
```

## Author
apybara Engineering Team