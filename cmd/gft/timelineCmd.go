package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"../../pkg/timeline"
	"io/ioutil"
	"log"
)

func init() {
	rootCmd.AddCommand(addFile)
	rootCmd.AddCommand(modifyFile)
}


var addFile = &cobra.Command{
	Use:   "add",
	Short: "Add File to GFT",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Adding File...")
		node, err := timeline.AddFile(args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}

		jsonBytes, err := json.Marshal(&node)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(jsonBytes))

		filename := args[1] + ".gft"
		ioutil.WriteFile(filename, jsonBytes, 0644)

	},
}

var modifyFile = &cobra.Command{
	Use:   "modify",
	Short: "Modify file in GFT",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Modifying File...")

		gftFile, err := ioutil.ReadFile(args[1])
		if err != nil {
			log.Fatal(err)
		}

		oldFileNode := timeline.Node{}
		err =json.Unmarshal([]byte(gftFile), &oldFileNode)
		if err != nil {
			log.Fatal(err)
		}

		newFileNode, err := timeline.ModifyFile(args[0], oldFileNode)
		if err != nil {
			log.Fatal(err)
		}

		jsonBytes, err := json.Marshal(&newFileNode)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(jsonBytes))
	},
}