package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "os"
    "os/exec"
)


const NGINX_LOG_PATH string = "../access.log"


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


func check_Service(name string) {
  cmd := exec.Command("systemctl", "status", name)
  out, err := cmd.CombinedOutput()
  if err != nil {
    if exitErr, ok := err.(*exec.ExitError); ok {
      fmt.Printf("systemctl finished with non-zero: %v\n", exitErr)
    } else {
      fmt.Printf("failed to run systemctl: %v", err)
      os.Exit(1)
    }
  }
  fmt.Printf("Status is: %s\n", string(out))
}



func main() {

    var text string
    text = "select"

    if Exists(NGINX_LOG_PATH) {
       fmt.Println("Nginx log file exist")
       //check whether s contains substring text
       fmt.Println(checkIfContainString(text, NGINX_LOG_PATH))

    } else {
       fmt.Println("nginx log does not exists")
    }

    check_Service("varnish")


}