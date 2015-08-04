package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/pgrunde/back/server"
)

func main() {
	s, err := loadSettings("./settings.json")
	if err != nil {
		log.Fatalf("Load Settings Error: %s", err)
	}
	log.Fatal(server.New(s).ListenAndServe())
}

func loadSettings(file string) (s server.Settings, err error) {
	f, err := os.Open(file)
	if err != nil {
		return server.Settings{}, err
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return server.Settings{}, err
	}
	err = json.Unmarshal(b, &s)
	return
}
