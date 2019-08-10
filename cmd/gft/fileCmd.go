package cmd

import (
	"../../pkg/auth"
	"../../pkg/blockchain"
	"../../pkg/serverhttp"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

func init() {
	rootCmd.AddCommand(newFile)
	rootCmd.AddCommand(showSent)
	rootCmd.AddCommand(showReceived)
	rootCmd.AddCommand(subscribe)
	rootCmd.AddCommand(sendTransaction)
}

var newFile = &cobra.Command{
	Use:   "new",
	Short: "Send a new file through GFT",
	Run: func(cmd *cobra.Command, args []string) {
		// .. new [filename] [to:deeID]
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

		serverhttp.SendPost()
		// 2) Transfer to File Server
		// Transfer EncryptedFile & Encrypted AES Key

		// 3) Create TX on Blockchain
		toAdd := "0x16978b95a180bf35a40f0fafa68e73d87aab4232"
		privKey := "2c2952291448595ffe14276e8fc914644988625c1f441d6f7afd7cba1edd18ab"
		data := []string{"gft", fileHash}
		rawTx := blockchain.CreateTx(toAdd, privKey, strings.Join(data, ","))
		blockchain.SendTx(rawTx)
		// 4) Create timeline
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
		fmt.Println("Showing sent files")
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
		privKey := "2c2952291448595ffe14276e8fc914644988625c1f441d6f7afd7cba1edd18ab"
		data := "Hello World"
		rawTx := blockchain.CreateTx(toAdd, privKey, data)
		blockchain.SendTx(rawTx)

	},
}