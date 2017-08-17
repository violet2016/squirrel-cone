package analyze

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Signature struct {
	Value string `json:"Value"`
	// SigScanAddress int    `json:"SigScanAddress"`
	// ASMSignature   bool   `json:"ASMSignature"`
	// Key            string `json:"Key"`
}

func LoadSigsFile(filename string) []Signature {
	sigs := []Signature{}
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(content, &sigs)
	return sigs
}
