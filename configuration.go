package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Montage struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
}

func (montage Montage) String() string {
	return fmt.Sprintf("[%02d] %s [%s]", montage.Id, montage.Name, montage.Path)
}

type Configuration struct {
	Autor    string    `json:"author"`
	Title    string    `json:"title"`
	Montages []Montage `json:"montages"`
	Files    []string  `json:"files"`
	Double   bool      `json:"double"` // when double compilation is required
}

func (configuration *Configuration) load() {
	var err error
	var file *os.File
	if file, err = os.Open(jsonPath); err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var content []byte
	if content, err = ioutil.ReadAll(file); err != nil {
		log.Fatal(err)
	}
	if err = json.Unmarshal(content, &configuration); err != nil {
		log.Fatalf("failed to parse configuration: %s\n ⮞ %v", jsonPath, err)
	}
	checkConfiguration(configuration)
}

func checkConfiguration(configuration *Configuration) {
	for _, file := range configuration.Files {
		path := filepath.Join(basedir, file)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			log.Fatalf("Error in configuration file: %q\nFile listed not found on disk: %q\n", jsonPath, path)
		}
	}
	for _, montage := range configuration.Montages {
		path := filepath.Join(basedir, montage.Path)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			log.Fatalf("Error in configuration file: %q\nMontage listed not found on disk: %+v\nFile not found: %q\n", jsonPath, montage, path)
		}
	}
}
