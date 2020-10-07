package checker

import (
	"context"

	"github.com/Flyewzz/tester/models"
)

type Program interface {
	Run(ctx context.Context, input string) (models.Result, error)
	Check() []*Verdict
}
