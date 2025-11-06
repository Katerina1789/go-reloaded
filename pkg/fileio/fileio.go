package fileio

import "io/ioutil"

func ReadFile(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func WriteFile(filename, content string) error {
	return ioutil.WriteFile(filename, []byte(content), 0644)
}