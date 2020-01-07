package adacore

import (
	"testing"
)

func TestDownloadMgr(t *testing.T) {
	tasksChan := make(chan int)
	mgr := NewDownloadMgr(2)

	mgr.AddTask("https://content.backcountry.com/images/items/large/TNF/TNF05JY/TNFBK.jpg",
		func(url string, buf []byte, err error) {
			if err != nil {
				t.Fatalf("TestDownloadMgr AddTask %v", err)
			}

			tasksChan <- 1
		})

	mgr.AddTask("https://content.backcountry.com/images/items/large/TNF/TNF05JY/VNWH_D2.jpg",
		func(url string, buf []byte, err error) {
			if err != nil {
				t.Fatalf("TestDownloadMgr AddTask %v", err)
			}

			tasksChan <- 1
		})

	mgr.AddTask("https://content.backcountry.com/images/items/small/TNF/TNF05JY/TNFBK_D2.jpg",
		func(url string, buf []byte, err error) {
			if err != nil {
				t.Fatalf("TestDownloadMgr AddTask %v", err)
			}

			tasksChan <- 1
		})

	mgr.AddTask("https://content.backcountry.com/images/items/900/TNF/TNF05JY/TNFBK_D3.jpg",
		func(url string, buf []byte, err error) {
			if err != nil {
				t.Fatalf("TestDownloadMgr AddTask %v", err)
			}

			tasksChan <- 1
		})

	mgr.AddTask("https://content.backcountry.com/images/items/small/TNF/TNF05JY/VNWH_D3.jpg",
		func(url string, buf []byte, err error) {
			if err != nil {
				t.Fatalf("TestDownloadMgr AddTask %v", err)
			}

			tasksChan <- 1
		})

	tasknums := 0
	for {
		r := <-tasksChan
		tasknums = tasknums + r
		if tasknums >= 5 {
			break
		}
	}

	t.Logf("TestDownloadMgr OK")
}
