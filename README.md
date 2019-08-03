# GalacticFileTransfer
Data transfer system using blockchain


Using the following structure:
https://github.com/golang-standards/project-layout

## Args
Arguments passed to the app to carry out various functions.

1. new
2. sub
3. ...

## Other
### Bugs and errors
1. There seems to be a bug with ganache-cli that won't allow us to subscribe to new blocks
2. Similar error on the truffle dev environment - this time we get an error if the transaction hex values begin with a zero - e.g. if the 's' value in the tx sig begins with zero (e.g. 0x0e2221d639e2db5c65b61e67f49ae2c5f5bac88ec52334a3c0369019f39b3618) then we get an error of: 'json: cannot unmarshal hex number with leading zero digits into Go struct field rpcBlock.transactions of type *hexutil.Big'