package handler

import (
	"fmt"
	"io/ioutil"
	"strings"
)

//ReadFile read from a file and put into a slice, separate by /n
func ReadFile(path string) ([]string, error) {
	fmt.Println("Start to read : " + path)
	readFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	readSlice := strings.Split(string(readFile), "\n")
	fmt.Println("Read " + path + " complete")
	fmt.Println()
	return readSlice, nil
}
