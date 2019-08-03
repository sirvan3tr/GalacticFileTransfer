package blockchain

import (
    "context"
    "crypto/ecdsa"
    "encoding/hex"
    "fmt"
    "log"
    "math/big"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
)

// toAdd: To address in Hex
// privKey: Private key in Hex
func CreateTx(toAdd string, privKey string, data string) (rawTxHex string){
	// Connect to the Blockchain
    client := connect()

    privateKey, err := crypto.HexToECDSA(privKey[2:])
    if err != nil {
        log.Fatal(err)
	}
    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
    }

    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        log.Fatal(err)
    }

    value := big.NewInt(1000000000000000000) // in wei
    gasLimit := uint64(28000) // in units
    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    toAddress := common.HexToAddress(toAdd)
    tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, []byte(data))

    chainID, err := client.NetworkID(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
    if err != nil {
        log.Fatal(err)
    }

    ts := types.Transactions{signedTx}
    rawTxBytes := ts.GetRlp(0)
    rawTxHex = hex.EncodeToString(rawTxBytes)
	log.Println("Raw TX has been created.")
	return
}

func SendTx(rawTx string) {
	// Connect to the Blockchain
	client := connect()

    rawTxBytes, err := hex.DecodeString(rawTx)

    tx := new(types.Transaction)
    rlp.DecodeBytes(rawTxBytes, &tx)

    err = client.SendTransaction(context.Background(), tx)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0xc429e5f128387d224ba8bed6885e86525e14bfdc2eb24b5e9c3351a1176fd81f
}