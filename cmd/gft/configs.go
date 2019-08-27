package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"../../pkg/gftconfigs"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initProfile)
	rootCmd.AddCommand(getDeeID)
}

type initFile struct {
	DeeID   string `json:"deeID"`
	PubKey  string `json:"pubKey"`
	PrivKey string `json:"privKey"`
}

// This is to initialise the process
// record the essential information
// that allows the process to do anything
var initProfile = &cobra.Command{
	Use:   "init",
	Short: "Initiate user account",
	Run: func(cmd *cobra.Command, args []string) {

		initF := initFile{}

		// record the inputs from console
		initF.DeeID = lineReader("Your deeID")
		initF.PubKey = lineReader("Your Eth pubKey")
		initF.PrivKey = lineReader("Your Eth privKey")

		// create a JSON of the recorded strings
		jsonBytes, err := json.Marshal(&initF)
		if err != nil {
			log.Fatal(err)
		}

		// create a file out of this.
		// TODO: password protect this file
		ioutil.WriteFile("personal.gft", jsonBytes, 0644)

	},
}

var getDeeID = &cobra.Command{
	Use:   "get-deeID",
	Short: "Get the user account",
	Run: func(cmd *cobra.Command, args []string) {
		deeID := gftconfigs.GetMyDeeID()
		fmt.Println(deeID)
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
