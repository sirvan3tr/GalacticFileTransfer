package gftconfigs

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type initFile struct {
	DeeID   string `json:"deeID"`
	PubKey  string `json:"pubKey"`
	PrivKey string `json:"privKey"`
}

func InitGFT() {
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
}

func GetMyPrivKey() (privKey string) {
	initF := GetProfile()
	return initF.PrivKey
}

func GetMyDeeID() (deeID string) {
	initF := GetProfile()
	return initF.DeeID
}

func GetMyPubKey() (pubKey string) {
	initF := GetProfile()
	return initF.PubKey
}

func GetProfile() (profile initFile) {
	jsonBytes, err := ioutil.ReadFile("personal.gft")
	if err != nil {
		log.Fatal(err)
	}

	initF := initFile{}

	err = json.Unmarshal(jsonBytes, &initF)
	if err != nil {
		log.Fatalln("error:", err)
	}

	return initF
}
func lineReader(text string) (output string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(text + ": ")
	output, _ = reader.ReadString('\n')
	output = strings.Replace(output, "\n", "", -1)
	fmt.Println(output)
	return
}
