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

			log.Println(err)
			//log.Fatal(err)
		case header := <-headers:
			fmt.Println(header.Hash().Hex()) // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(header.Hash().Hex())     // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
			fmt.Println(block.Hash().Hex())      // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
			fmt.Println(block.Number().Uint64()) // 3477413
			//fmt.Println(block.Time().Uint64())     // 1529525947
			fmt.Println(block.Nonce()) // 130524141876765836
			fmt.Println("Txs in block: " + string(len(block.Transactions())))

			txs := block.Transactions()
			for _, element := range txs {
				fmt.Println("To: " + element.To().String()) //To() returns common.Address

				msg, err := element.AsMessage(types.NewEIP155Signer(element.ChainId()))
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println("From:")
				fmt.Println(msg.From().Hex())
				fmt.Println("Data: " + string(element.Data()))

				// Record files for us in the database
				database.RecordFile("e", string(element.Data()))
			}

		}
	}
}
