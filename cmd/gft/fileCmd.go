package cmd

import (
	"fmt"
	"../../pkg/auth"
)

func init() {
	rootCmd.AddCommand(newFile)
	rootCmd.AddCommand(modifyFile)
}

var newFile = &cobra.Command{
	use: "new",
	short: "Send a new file through GFT",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Sending a new file...")

		// 1) Open and Encrypt the file
		hsh, aes, fEncrypted, rsaAes : = NewFileEncrypt()

		// 2) Transfer to File Server
		// Transfer EncryptedFile & Encrypted AES Key

		// 3) Create TX on Blockchain

		// 4) Create timeline
	}
}