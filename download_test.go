package adacore

import (
	"testing"
)

func TestDownloadFile(t *testing.T) {
	err := DownloadFile("./unittest/downloadfile001.jpg", "https://content.backcountry.com/images/items/large/TNF/TNF05JY/TNFBK.jpg")
	if err != nil {
		t.Fatalf("DownloadFile DownloadFile %v", err)
	}

	t.Logf("DownloadFile OK")
}

func TestDownloadBuff(t *testing.T) {
	buf, err := DownloadBuff("https://content.backcountry.com/images/items/large/TNF/TNF05JY/TNFBK.jpg")
	if err != nil {
		t.Fatalf("TestDownloadBuff DownloadBuff %v", err)
	}

	if len(buf) != 15395 {
		t.Fatalf("TestDownloadBuff DownloadBuff invalid buffer length %v", len(buf))
	}

	t.Logf("TestDownloadBuff OK")
}
