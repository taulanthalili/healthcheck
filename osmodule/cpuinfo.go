package osmodule

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const CpuFile string = "/proc/stat"

func getCPUUsage() (idle, total uint64) {
	contents, err := ioutil.ReadFile(CpuFile)
	if err != nil {
		return
	}

	///Only for test===========
	ttos := GetOs()
	if ttos != "linux" {
		contents, err = ioutil.ReadFile("/Users/taulant/Documents/projects/mgt-health-script/stat")
		if err != nil {
			return
		}
	}
	///=========================

	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if fields[0] == "cpu" {
			numFields := len(fields)
			for i := 1; i < numFields; i++ {
				val, err := strconv.ParseUint(fields[i], 10, 64)
				if err != nil {
					fmt.Println("Error: ", i, fields[i], err)
				}
				total += val // tally up all the numbers to get total ticks
				if i == 4 {  // idle is the 5th field in the cpu line
					idle = val
				}
			}
			return
		}
	}
	return
}

func GetCPUinfoTEST() {
	idle0, total0 := getCPUUsage()
	time.Sleep(3 * time.Second)
	idle1, total1 := getCPUUsage()

	idleTicks := float64(idle1 - idle0)
	totalTicks := float64(total1 - total0)
	cpuUsage := 100 * (totalTicks - idleTicks) / totalTicks

	fmt.Printf("CPU usage is %f%% [busy: %f, total: %f]\n", cpuUsage, totalTicks-idleTicks, totalTicks)
}

func GetCPUinfoOutput() bytes.Buffer {

	cmd := exec.Command("ps", "--no-headers", "-eo", "pcpu")
	ttos := GetOs()
	if ttos != "linux" {
		cmd = exec.Command("ps", "-Ao", "pcpu=")
	}

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return stderr
	}

	return out
}

func GetCPUinfo() {

	tmp := GetCPUinfoOutput()
	var t float32

	values := strings.Split(tmp.String(), "\n")

	for _, s := range values {
		j, _ := strconv.ParseFloat(strings.TrimSpace(s), 64)
		t = t + float32(j)
	}

	fmt.Println(t, "% CPU used")

}
