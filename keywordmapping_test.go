package adacore

import (
	"testing"
)

func TestLoadKeywordMappingList(t *testing.T) {
	lst, err := LoadKeywordMappingList("./unittest/keywordmapping.yaml")
	if err != nil {
		t.Fatalf("TestLoadKeywordMappingList LoadKeywordMappingList %v err is %v", "./unittest/keywordmapping.yaml", err)

		return
	}

	if len(lst.Keywords) != 2 {
		t.Fatalf("TestLoadKeywordMappingList LoadKeywordMappingList keywords.length is %v", len(lst.Keywords))

		return
	}

	if lst.Keywords[0].Keyword != "Ada" {
		t.Fatalf("TestLoadKeywordMappingList LoadKeywordMappingList keywords[0] is %v", lst.Keywords[0])

		return
	}

	if lst.Keywords[1].Keyword != "Zerro" {
		t.Fatalf("TestLoadKeywordMappingList LoadKeywordMappingList keywords[1] is %v", lst.Keywords[1])

		return
	}

	t.Logf("TestLoadKeywordMappingList OK")
}
