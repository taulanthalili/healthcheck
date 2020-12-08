package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

const NginxLogPath string = "../access.log"

func ContainsI(a string, b string) bool {
	return strings.Contains(
		strings.ToLower(a),
		strings.ToLower(b),
	)
}

// Exists reports whether the named file or directory exists.
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func checkIfContainString(name string, pathFile string) bool {

	// read the whole file at once
	b, err := ioutil.ReadFile(pathFile)
	if err != nil {
		panic(err)
	}
	s := string(b)
	// //check whether s contains substring text
	return ContainsI(s, name)

}

func main() {

	var text string
	text = "select"

	if Exists(NginxLogPath) {
		fmt.Println("Nginx log file exist")
		//check whether s contains substring text
		fmt.Println(checkIfContainString(text, NginxLogPath))

	} else {
		fmt.Println("nginx log does not exists")
	}

	checkService("varnish")

}
