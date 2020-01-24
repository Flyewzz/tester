package models

type IOFileManager struct {
}

func (iofm *IOFileManager) Read(path string) (string, error) {
	return "", nil
}

func (iofm *IOFileManager) Write(text string, path string) error {
	return nil
}
