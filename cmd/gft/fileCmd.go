package cmd

import (
	"fmt"
	//"../../pkg/auth"
	"github.com/spf13/cobra"
	"../../pkg/blockchain"
)

func init() {
	rootCmd.AddCommand(newFile)
	rootCmd.AddCommand(subscribe)
}

var newFile = &cobra.Command{
	Use: "new",
	Short: "Send a new file through GFT",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Sending a new file...")

		// 1) Open and Encrypt the file
		//hsh, aes, fEncrypted, rsaAes : = NewFileEncrypt()

		// 2) Transfer to File Server
		// Transfer EncryptedFile & Encrypted AES Key

		// 3) Create TX on Blockchain
		toAdd := "0xEBe15b10Dc453345f48303F5451C3c07557B8250"
		privKey := "0x0aa4df97fb37d17508ed208a80d4d6b784954dfb63363dc6128bd8dd3b155b34"
		data := "Hello World"
		rawTx := blockchain.CreateTx(toAdd, privKey, data)
		
		blockchain.SendTx(rawTx)
		// 4) Create timeline
	},
}

var subscribe = &cobra.Command{
	Use: "sub",
	Short: "Subscribe",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Sending a new file...")

		// 1) Open and Encrypt the file
		//hsh, aes, fEncrypted, rsaAes : = NewFileEncrypt()

		// 2) Transfer to File Server
		// Transfer EncryptedFile & Encrypted AES Key

		// 3) Create TX on Blockchain
		blockchain.Sub()
		// 4) Create timeline
	},
}