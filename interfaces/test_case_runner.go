package interfaces

import (
	"github.com/Flyewzz/tester/models"
)

type ITestCaseRunner interface {
	Run(testCase *models.TestCase, app *Executable) *models.Message
}
