package checker

type Program interface {
	Run(input string) (string, error)
	Check() []*Verdict
}
