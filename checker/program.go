package checker

import "context"

type Program interface {
	Run(ctx context.Context, input string) (string, error)
	Check() []*Verdict
}
