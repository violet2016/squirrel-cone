package analyze

import "io/ioutil"
import "encoding/json"
import "log"

type Signature struct {
	Value []byte
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
