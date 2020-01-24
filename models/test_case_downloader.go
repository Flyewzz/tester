package models

type TestCaseDownloader struct {
}

func (tcd *TestCaseDownloader) Download(dirPath string) *TestCase {
	return NewTestCase()
}
