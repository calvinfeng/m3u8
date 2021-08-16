package downloader

type Downloader interface {
	Start(concurrency int, mergedFilename string) error
}
