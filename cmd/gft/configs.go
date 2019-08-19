package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initProfile)
}

type initFile struct {
	DeeID   string `json:"deeID"`
	PubKey  string `json:"pubKey"`
	PrivKey string `json:"privKey"`
}

var initProfile = &cobra.Command{
	Use:   "init",
	Short: "Initiate user account",
	Run: func(cmd *cobra.Command, args []string) {

		initF := initFile{}

		initF.DeeID = lineReader("Your deeID")
		initF.PubKey = lineReader("Your Eth pubKey")
		initF.PrivKey = lineReader("Your Eth privKey")

		jsonBytes, err := json.Marshal(&initF)
		if err != nil {
			log.Fatal(err)
		}

		ioutil.WriteFile("personal.gft", jsonBytes, 0644)

	},
}

func lineReader(text string) (output string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(text + ": ")
	output, _ = reader.ReadString('\n')
	output = strings.Replace(output, "\n", "", -1)
	fmt.Println(output)
	return
}
