package adacore

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// KeywordMapping - keyword mapping
type KeywordMapping struct {
	// Keyword - keyword
	Keyword string
	// URL - url
	URL string
}

// KeywordMappingList - KeywordMapping list
type KeywordMappingList struct {
	Keywords []*KeywordMapping
}

// LoadKeywordMappingList - load keyword mapping file
func LoadKeywordMappingList(fn string) (*KeywordMappingList, error) {
	fi, err := os.Open(fn)
	if err != nil {
		return nil, err
	}

	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		return nil, err
	}

	lst := &KeywordMappingList{}

	err = yaml.Unmarshal(fd, lst)
	if err != nil {
		return nil, err
	}

	return lst, nil
}
