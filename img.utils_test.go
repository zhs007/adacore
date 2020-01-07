package adacore

import (
	"testing"

	"github.com/disintegration/imaging"
)

func TestResizeImage(t *testing.T) {
	err := ResizeImage("./unittest/sample001_180.jpg", 180, 0, "./unittest/sample001.jpg", imaging.NearestNeighbor)
	if err != nil {
		t.Fatalf("TestResizeImage ResizeImage %v", err)

		return
	}

	err = ResizeImage("./unittest/sample001_240x240.jpg", 240, 240, "./unittest/sample001.jpg", imaging.NearestNeighbor)
	if err != nil {
		t.Fatalf("TestResizeImage ResizeImage %v", err)

		return
	}

	t.Logf("TestResizeImage OK")
}

func TestFitImage(t *testing.T) {
	err := FitImage("./unittest/sample001_800x600.jpg", 800, 600, "./unittest/sample001.jpg", imaging.NearestNeighbor)
	if err != nil {
		t.Fatalf("TestFitImage FitImage %v", err)

		return
	}

	t.Logf("TestFitImage OK")
}

func TestFillImage(t *testing.T) {
	err := FillImage("./unittest/sample001_800x600c.jpg", 800, 600, "./unittest/sample001.jpg", imaging.Center, imaging.NearestNeighbor)
	if err != nil {
		t.Fatalf("TestFillImage FillImage %v", err)

		return
	}

	err = FillImage("./unittest/sample001_800x600t.jpg", 800, 600, "./unittest/sample001.jpg", imaging.Top, imaging.NearestNeighbor)
	if err != nil {
		t.Fatalf("TestFillImage FillImage %v", err)

		return
	}

	t.Logf("TestFillImage OK")
}
