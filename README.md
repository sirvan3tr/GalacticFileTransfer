# GalacticFileTransfer
Data transfer system using blockchain


Using the following structure:
https://github.com/golang-standards/project-layout
### Requirements
In order to run this app you will require an Ethereum test node and a file server.

1. Currently for development and testing purposes we are using truffle which is running the deeID application. deeID truffle app can be found here: https://github.com/sirvan3tr/deeID.Web

2. File server: The architecture assumes that everyone has their own file server. In the future you should be able to connect together direclty - p2p.
Use this simple file server: https://github.com/sirvan3tr/Gft.FileServer

## Assumptions
- Each node has their own particular file-server (though in the future we can make it a p2p as well)
## Args
Arguments passed to the app to carry out various functions.

``./gft new [file] [to: deeID address]``

Subscribe to the blockchain in order to record files that are being sent to me.
``./gft sub``
**TO DO: Have the user's messaging server handle this.

``./gft view contacts``

``./gft view files``

``./gft get-deeID``

## To do:
1. Password protect or encrypt the file that has the private keys and deeID. Can use the mobile deeID app to do this for us.


## Other
### Bugs and errors
1. There seems to be a bug with ganache-cli that won't allow us to subscribe to new blocks
2. Similar error on the truffle dev environment - this time we get an error if the transaction hex values begin with a zero - e.g. if the 's' value in the tx sig begins with zero (e.g. 0x0e2221d639e2db5c65b61e67f49ae2c5f5bac88ec52334a3c0369019f39b3618) then we get an error of: 'json: cannot unmarshal hex number with leading zero digits into Go struct field rpcBlock.transactions of type *hexutil.Big'