package blockchain

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"../database"

	//"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/core/types"
)

func prettyprint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.Bytes(), err
}

func Sub() {
	client := connect()

	/*
		client, err := ethclient.Dial("wss://ropsten.infura.io/ws")
		if err != nil {
			log.Fatal(err)
		}
	*/

	headers := make(chan *types.Header)

	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(sub)
	for {
		log.Println(headers)
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			txs := block.Transactions()
			fmt.Println("Txs in block: " + string(len(txs)))
			for _, element := range txs {
				fmt.Println("To: " + element.To().String()) //To() returns common.Address

				msg, err := element.AsMessage(types.NewEIP155Signer(element.ChainId()))
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println("From: " + msg.From().Hex())
				fmt.Println("Data: " + string(element.Data()))

				// Record files for us in the database
				// tx_address, tx_data, file_type
				database.RecordFile(string(element.Hash().Hex()), string(element.Data()), "received")
			}

		}
	}
}
