package adacore

import (
	"sync"

	adacorebase "github.com/zhs007/adacore/base"
	"go.uber.org/zap"
)

// DownloadMgr - download manager
type DownloadMgr struct {
	workers []*DownloadWorker
	tasks   []*downloadTask
	mutex   sync.Mutex
}

// NewDownloadMgr - new DownloadMgr
func NewDownloadMgr(workernums int) *DownloadMgr {
	mgr := &DownloadMgr{}

	for i := 0; i < workernums; i++ {
		mgr.workers = append(mgr.workers, &DownloadWorker{
			WorkerIndex: i,
		})
	}

	return mgr
}

func (mgr *DownloadMgr) pushTask(url string, ondownload FuncOnDownloaded) {
	mgr.mutex.Lock()
	mgr.tasks = append(mgr.tasks, &downloadTask{
		url:        url,
		onDownload: ondownload,
	})
	mgr.mutex.Unlock()
}

func (mgr *DownloadMgr) pullTask() *downloadTask {
	mgr.mutex.Lock()
	defer mgr.mutex.Unlock()

	if len(mgr.tasks) > 0 {
		t := mgr.tasks[0]
		mgr.tasks = mgr.tasks[1:]
		return t
	}

	return nil
}

func (mgr *DownloadMgr) startWorker() {
	mgr.mutex.Lock()
	defer mgr.mutex.Unlock()

	for _, v := range mgr.workers {
		if !v.IsRunning {
			go func() {
				v.IsRunning = true

				for {
					t := mgr.pullTask()
					if t == nil {
						break
					}

					err := v.start(t.url)
					if t.onDownload != nil {
						t.onDownload(t.url, v.Buff, err)
					}
				}

				v.IsRunning = false
			}()

			return
		}
	}
}

// LogState - output state to log
func (mgr *DownloadMgr) LogState() {
	adacorebase.Info("DownloadMgr:LogState",
		zap.Int("workernums", len(mgr.workers)))

	for i := 0; i < len(mgr.workers); i++ {
		mgr.workers[i].LogState()
	}
}

// AddTask - add task
func (mgr *DownloadMgr) AddTask(url string, ondownloaded FuncOnDownloaded) {
	mgr.pushTask(url, ondownloaded)
	mgr.startWorker()
}
