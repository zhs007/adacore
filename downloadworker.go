package adacore

import (
	"time"

	adacorebase "github.com/zhs007/adacore/base"
	"go.uber.org/zap"
)

// DownloadWorker - download worker
type DownloadWorker struct {
	WorkerIndex     int
	CurURL          string
	StartTime       int64
	LastTime        int64
	Buff            []byte
	IsRunning       bool
	TotalTaskNums   int
	TotalBuffNums   int64
	TotalFailedNums int
	TotalTime       int64
}

// start - start download
func (dw *DownloadWorker) start(url string) error {
	// dw.IsRunning = true
	dw.TotalTaskNums++

	// defer func() {
	// 	dw.IsRunning = false
	// }()

	dw.CurURL = url
	dw.StartTime = time.Now().UnixNano()

	buf, err := DownloadBuff(url)
	if err != nil {
		dw.TotalFailedNums++

		return err
	}

	et := time.Now().UnixNano()

	dw.Buff = buf

	dw.TotalBuffNums += int64(len(buf))
	dw.TotalTime += et - dw.StartTime

	return nil
}

// LogState - output state to log
func (dw *DownloadWorker) LogState() {
	adacorebase.Info("DownloadWorker:LogState",
		zap.Int("workerindex", dw.WorkerIndex),
		zap.Int("tasknums", dw.TotalTaskNums),
		zap.Int("failednums", dw.TotalFailedNums),
		zap.Int64("totalbytes", dw.TotalBuffNums),
		zap.Int64("totaltime", dw.TotalTime))
}
