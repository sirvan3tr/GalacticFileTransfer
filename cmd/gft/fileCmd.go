package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"../../pkg/auth"
	"../../pkg/blockchain"
	"../../pkg/database"
	"../../pkg/gftconfigs"
	"../../pkg/serverhttp"
	"../../pkg/timeline"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(newFile)
	rootCmd.AddCommand(showSent)
	rootCmd.AddCommand(showReceived)
	rootCmd.AddCommand(subscribe)
	rootCmd.AddCommand(sendTransaction)
}

// ARGUMENTS:
// .. new [filename] [to:deeID]
var newFile = &cobra.Command{
	Use:   "new",
	Short: "Send a new file through GFT",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Sending a new file...")

		// 1) Open and Encrypt the file
		fileHash, aesRawKey, fEncrypted, rsaCipherString := auth.NewFileEncrypt(args[0])
		fmt.Println("Hash of the file:")
		fmt.Println(fileHash)

		fmt.Println("Raw AES key:")
		fmt.Println(aesRawKey)

		fmt.Println("Encrypted file:")
		fmt.Println(fEncrypted)

		fmt.Println("RSA Cipher")
		fmt.Println(rsaCipherString)

		// 2) Transfer to File Server
		// Transfer EncryptedFile & Encrypted AES Key
		serverhttp.SendPost([]byte(fEncrypted), []byte(rsaCipherString))

		// 3) Create TX on Blockchain
		toAdd := args[1]
		// TODO: runchecks on the toAdd to ensure it is a valid address

		//privKey := "c7803a01bd3f699467d8ae09138ce1d2f182e75a07040f6a62f7af90d049635e"
		privKey := gftconfigs.GetMyPrivKey()
		data := []string{"gft", fileHash}
		rawTx := blockchain.CreateTx(toAdd, privKey, strings.Join(data, ","))
		blockchain.SendTx(rawTx)
		//TODO: returns errors for SendTX

		// 4) Create timeline
		fmt.Println("Adding timeline...")
		node, err := timeline.AddFile(args[0], password)
		if err != nil {
			log.Fatal(err)
		}

		jsonBytes, err := json.Marshal(&node)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(jsonBytes))

		filename := args[0] + ".gft"
		ioutil.WriteFile(filename, jsonBytes, 0644)
		// 5:) Update file server with the tx-id and the timeline info
	},
}

var showSent = &cobra.Command{
	Use:   "show-sent",
	Short: "View sent files",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Showing sent files")
	},
}

var showReceived = &cobra.Command{
	Use:   "show-received",
	Short: "View received files",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Showing received files")

		// Shows all the files in the database
		// captured by subscribing to the blocks
		// TODO: have arg to show last 5 files
		// TODO: need to capture data from users messaging server too
		// -- because you could be missing a file if not subscribed
		database.ListAllFiles()

	},
}
var subscribe = &cobra.Command{
	Use:   "sub",
	Short: "Subscribe",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Listening to the blockchain for new transactions...")

		blockchain.Sub()

	},
}

var sendTransaction = &cobra.Command{
	Use:   "send",
	Short: "Send tx",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Sending transaction")

		toAdd := "0xf735d1bfb091ce9f50f797778633dfccb5d91310"
		privKey := "c7803a01bd3f699467d8ae09138ce1d2f182e75a07040f6a62f7af90d049635e"
		data := "Hello World"
		rawTx := blockchain.CreateTx(toAdd, privKey, data)
		blockchain.SendTx(rawTx)

		database.ListAllFiles()
	},
}
