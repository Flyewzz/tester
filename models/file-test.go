package models

type FileTest struct {
	Expected string
}

func NewFileTest(exp string) *FileTest {
	return &FileTest{
		Expected: exp,
	}
}

func (ft *FileTest) Run() {

}
