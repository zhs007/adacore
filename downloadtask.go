package adacore

// FuncOnDownloaded - func (url string, buf []byte, err error)
type FuncOnDownloaded func(url string, buf []byte, err error)

// downloadTask - download task
type downloadTask struct {
	url        string
	onDownload FuncOnDownloaded
}
