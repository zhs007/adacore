package chatbotada

import (
	"io/ioutil"
	"os"
	"testing"
	"text/template"

	"gopkg.in/yaml.v2"
)

type yamlObj struct {
	Description string
	Other       string
}

type testDebugExcel struct {
	DebugExcel yamlObj
}

func Test_debugexcel(t *testing.T) {
	fi, err := os.Open("../lang/chatbot.en.yaml")
	if err != nil {
		t.Fatalf("Test_debugexcel Open err %v", err)

		return
	}

	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		t.Fatalf("Test_debugexcel ReadAll err %v", err)

		return
	}

	de := &testDebugExcel{}

	err = yaml.Unmarshal(fd, de)
	if err != nil {
		t.Fatalf("Test_debugexcel Unmarshal err %v", err)

		return
	}

	lstct := []ExcelColumnType{
		ColumnPrimaryKey,
		ColumnInfo,
	}
	lstctn := []string{
		"column0",
		"column1",
	}
	var lstallcts []DebugExcelColumnType

	for i, v := range lstctn {
		lstallcts = append(lstallcts, DebugExcelColumnType{
			Name: v,
			Type: ExcelColumnType2String(lstct[i]),
		})
	}

	mParams := make(map[string]interface{})

	mParams["Columns"] = lstallcts

	tmp1, err := template.New("tmpl").Parse(de.DebugExcel.Other)
	if err != nil {
		t.Fatalf("Test_debugexcel Parse err %v", err)

		return
	}

	err = tmp1.Execute(os.Stdout, mParams)
	if err != nil {
		t.Fatalf("Test_debugexcel Execute err %v", err)

		return
	}

	t.Logf("Test_debugexcel OK")
}
