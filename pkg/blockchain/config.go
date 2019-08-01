package blockchain

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// Connecting to local machine
	// Have the address as a configurable element!!
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	_ = client // we'll use this in the upcoming sections
}
