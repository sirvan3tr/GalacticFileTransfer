package cmd

import (
	"fmt"
	//"../../pkg/auth"
	"../../pkg/blockchain"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(newFile)
	rootCmd.AddCommand(showSent)
	rootCmd.AddCommand(subscribe)
}

var newFile = &cobra.Command{
	Use:   "new",
	Short: "Send a new file through GFT",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Sending a new file...")

		// 1) Open and Encrypt the file
		//hsh, aes, fEncrypted, rsaAes := NewFileEncrypt()

		// 2) Transfer to File Server
		// Transfer EncryptedFile & Encrypted AES Key

		// 3) Create TX on Blockchain
		toAdd := "0x16978b95a180bf35a40f0fafa68e73d87aab4232"
		privKey := "2c2952291448595ffe14276e8fc914644988625c1f441d6f7afd7cba1edd18ab"
		data := "Hello World"
		rawTx := blockchain.CreateTx(toAdd, privKey, data)
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

var subscribe = &cobra.Command{
	Use:   "sub",
	Short: "Subscribe",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Listening to the blockchain for new transactions...")

		blockchain.Sub()

	},
}
