package blockchain

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func connect() (client *ethclient.Client){
	// Connecting to local machine
	// Have the address as a configurable element!!
	
	//client, err := ethclient.Dial("wss://ropsten.infura.io/ws")
	client, err := ethclient.Dial("ws://localhost:9545")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	_ = client

	return
}
