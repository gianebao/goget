package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var (
		err     error
		file    []byte
		changed bool
		pk      = &Package{}
	)

	defer func() {
		if r := recover(); r != nil {
			log.Fatal(r)
			os.Exit(1)
		}
	}()

	_, err = os.Stat(PkgName)

	if os.IsNotExist(err) {
		pk = NewPackage()
		changed = true
	} else {
		if file, err = ioutil.ReadFile(PkgName); err != nil {
			panic("Read file:" + err.Error())
		}

		if err = json.Unmarshal(file, pk); err != nil {
			panic("JSON decode:" + err.Error())
		}

		if changed, err = pk.Verify(); err != nil {
			panic(err.Error())
		}
	}

	pk.LoadAll()

	if changed {
		if file, err = json.MarshalIndent(pk, "", "\t"); err != nil {
			panic(err.Error())
		}
		ioutil.WriteFile(PkgName, file, 0644)
		log.Println("Configuration file `", PkgName, "` created")
	}
}
