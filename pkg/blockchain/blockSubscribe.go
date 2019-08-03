package blockchain

import (
    "context"
    "fmt"
    "log"
	//"github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/core/types"
)

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
			fmt.Println("New Header Coming In------------")
			fmt.Println(context.Background())
			fmt.Println("end of header heash----------------")
            fmt.Println(header.Hash().Hex()) // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
            block, err := client.BlockByHash(context.Background(), header.Hash())
            if err != nil {
                log.Fatal(err)
            }

			fmt.Println("Hel")
            fmt.Println(header.Hash().Hex()) // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
            fmt.Println(block.Hash().Hex())        // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
            fmt.Println(block.Number().Uint64())   // 3477413
            //fmt.Println(block.Time().Uint64())     // 1529525947
            fmt.Println(block.Nonce())             // 130524141876765836
            fmt.Println(len(block.Transactions())) // 7
        }
    }
}