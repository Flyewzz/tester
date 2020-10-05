package checker

import (
	"io/ioutil"
	"strings"
)

func GetTestsNames(path string) ([]string, error) {
	tests, err := ioutil.ReadDir(path)
	if err != nil {
		return []string{}, nil
	}
	var dirNames []string
	for _, t := range tests {

		if t.IsDir() {
			dirNames = append(dirNames, t.Name())
		}
	}

	return dirNames, nil
}

func GetTest(path string) (*Test, error) {
	testFolder, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	var test Test
	subPaths := strings.Split(path, "/")
	test.Name = subPaths[len(subPaths)-1]
	if path[len(path)-1] != '/' {
		test.Name = subPaths[len(subPaths)-1]
		path += "/"
	}
	for _, t := range testFolder {
		if !t.IsDir() {
			filePath := path + t.Name()
			file, err := ioutil.ReadFile(filePath)
			if err != nil {
				return nil, err
			}
			switch t.Name() {
			case "input.txt":
				test.Input = string(file)
			case "output.txt":
				test.Output = string(file)
			}
		}
	}
	return &test, nil
}
