package interfaces

import (
	"github.com/Flyewzz/tester/models"
)

type ITestCaseDownloader interface {
	Download(dirPath string) *models.TestCase
}
