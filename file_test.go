package goexamples

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	//"io/ioutil"
	"os"
	"testing"
	"strings"
)

const (
	exampleFile = "/tmp/goexample.txt"
)

func TestFile(t *testing.T) {
	fmt.Println("testing TestFile")
	// Clean up log file just in case.
	err := os.RemoveAll(exampleFile)
	assert.NoError(t, err)

	//logContent, err := ioutil.ReadFile(exampleFile)
	//assert.NoError(t, err)
	//lines := strings.Split(strings.TrimSpace(string(logContent)), "\n")

	lines := strings.Split(strings.TrimSpace(string("test1\ntest2")), "\n")
	expectedContent := []string {
		"test1",
		"test2"}

	for i, line := range lines {
		assert.Contains(t, line, expectedContent[i])
	}

	// clean up file.
	err = os.RemoveAll(exampleFile)
	assert.NoError(t, err)
}
