package handler

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"strings"
)

//ReadFile read from a file and put into a slice, separate by /n
func ReadFile(path string) ([]string, error) {
	fmt.Println("Start to read : " + path)

	readFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var readSlice []string
	if runtime.GOOS == "windows" {
		readSlice = strings.Split(string(readFile), "\r\n")
	} else {
		readSlice = strings.Split(string(readFile), "\n")
	}

	fmt.Println("Read " + path + " complete")
	fmt.Println()

	return readSlice, nil
}
