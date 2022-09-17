package osmodule

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
)

type DiskStatus struct {
	All  uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
}

//Human reading values for Disk size
const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

func GetOsArch() string {
	arch := runtime.GOARCH
	return arch
}

func GetOs() string {
	os := runtime.GOOS
	return os
}

//TEST
func TestOsModule(s string) string {
	CheckIfExists("/home")
	GettheDirectorywithHighnumFiles()

	CheckTopProcessMemory()
	CheckTopProcessCPU()
	CheckTopProcessSwap()
	CheckServerLoadIO()

	return string("Hello: " + s)
}

//Locate and return the Directory with many files
func GettheDirectorywithHighnumFiles() string {

	//empty
	return string("/FullPathDirectoryName")
}

//Check If File or Directory exists.
func CheckIfExists(s string) bool {
	if _, err := os.Stat(s); os.IsNotExist(err) {
		// path/to/whatever does not exist
		return false
	}
	return true
}

//===============================================
//Check if Swap & Memory/RAM is more than 95% used by checking the values
func CheckifServerisOutofMemory() bool {
	meminfo := &MemInfo{}
	err := meminfo.Update()
	if err != nil {
		panic(err)
	}

	fmt.Println(meminfo.Used(), "% Memory used") // Total - Used
	return true
}

//Find the process that uses the most of RAM
func CheckTopProcessMemory() {

	//BASH
	//ps aux --sort=-%mem | head -4
	//ps --no-headers -eo pmem,rss,vsize,time,pid,args | sort -k 1 -n -r | head -1
	CheckifServerisOutofMemory()
	cmd_string := "ps --no-headers -eo pmem,rss,vsize,time,pid,args | sort -k 1 -n -r | head -1"

	ttos := GetOs()
	if ttos != "linux" {
		cmd_string = "ps -Ao pmem,rss,vsize,time,pid,args -r | sort -k 1 -n -r | head -1"
	}

	//Print the process that uses the Memory
	fmt.Println(CheckTopProcess(cmd_string))
}

//===============================================

//===============================================
//Check if Swap is more than 95% used by checking the values
func CheckifServerisOutofSwap() bool {
	meminfo := &MemInfo{}
	err := meminfo.Update()
	if err != nil {
		panic(err)
	}

	fmt.Println(meminfo.Swap(), "% Swap used") // Swap % Used Memory
	return true
}

//Find the process that uses the most of Swap
func CheckTopProcessSwap() {
	CheckifServerisOutofSwap()

	cmd_string_swap := "grep VmSwap /proc/[0-9]*/status | sort -k2 -n | tail -1 | cut -d\"/\" -f3"
	ttos := GetOs()
	if ttos != "linux" {
		//cmd_string = "ps -o pmem,rss,vsize,time,pid,args -p " + strconv.FormatUint(swapinfo.GetPID(), 10) + " | sort -k 1 -n -r | head -1"

		cmd_string_swap = "grep -r VmSwap /Users/taulant/Documents/projects/mgt-health-script/[0-9]*/status | sort -k2 -n | tail -1 | cut -d\"/\" -f7"
	}

	read_line := strings.TrimSuffix(CheckTopProcess(cmd_string_swap), "\n")
	cmd_string := "ps --no-headers -o pid,time,args -p " + read_line

	if ttos != "linux" {
		cmd_string = "ps -o pid,time,args -p " + read_line + " | sort -k 1 -n -r | head -1"
	}

	fmt.Println("Swap:\n", CheckTopProcess(cmd_string))
}

//===============================================

//===============================================
//Check if CPU is more than 95% used
func CheckifServerHighCPU() bool {

	//
	//ps --no-headers -eo %cpu

	return true
}

//Find the process that uses the most of CPU
func CheckTopProcessCPU() {
	//BASH examples
	//ps aux --sort=-%cpu | head -4
	//ps -eo pcpu,time,pid,args | sort -k 1 -n -r | head -1

	//ps -eo pcpu,time,pid,args --sort -pcpu
	//ps --no-headers -eo pcpu,time,pid,args --sort -pcpu | head -1
	GetCPUinfo()
	CheckifServerHighCPU()
	cmd_string := "ps --no-headers -eo pcpu,time,pid,args --sort -pcpu | head -1"

	ttos := GetOs()
	if ttos != "linux" {
		cmd_string = "ps -Ao pcpu,pid,args -r | sort -k 1 -n -r | head -1"
	}

	//Print the process that use much CPU:
	fmt.Println(CheckTopProcess(cmd_string))

}

//Run bash command
func CheckTopProcess(s string) string {

	cmd := exec.Command("bash", "-c", s)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
	}

	return out.String()
}

//===============================================

func CheckServerLoadIO() string {
	return string("Server Load information")
}

//===============================================
// disk usage for a given partition
func DiskUsage(path string) (disk DiskStatus) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return
	}
	disk.All = fs.Blocks * uint64(fs.Bsize)
	disk.Free = fs.Bfree * uint64(fs.Bsize)
	disk.Used = disk.All - disk.Free
	return
}

// Check disk free space percentage
func CheckPartitionFreePercent(s string) int32 {
	disk := DiskUsage(s)
	diskall := float64(disk.All) / float64(GB)
	diskfree := float64(disk.Free) / float64(GB)

	dfpercent := float64(diskfree / diskall)
	if math.IsNaN(dfpercent) {
		return -1
	}

	p := int32(dfpercent * 100)

	if p <= 10 {
		p = -1
	}

	return p

}

func FindDirectory(name string) bytes.Buffer {
	cmd := exec.Command("/usr/bin/find", "/home",
		"-maxdepth", "1",
		"-type", "d", "!", "-name", ".")

	if name == "/" {
		cmd = exec.Command("/usr/bin/find", name,
			"-not", "-path", "/home",
			"-not", "-path", "/home/*",
			"-not", "-path", "/dev",
			"-not", "-path", "/dev/*",
			"-not", "-path", "/proc",
			"-not", "-path", "/proc/*",
			"-not", "-path", "/boot",
			"-not", "-path", "/boot/*",
			"-not", "-path", "/sys",
			"-not", "-path", "/sys/*",
			"-maxdepth", "7",
			"-type", "d", "!", "-name", ".")

	} else {
		cmd = exec.Command("/usr/bin/find", name,
			"-not", "-path", "/home/mysql",
			"-not", "-path", "/home/mysql/*",
			"-not", "-path", "/home/.swap",
			"-maxdepth", "7",
			"-type", "d", "!", "-name", ".")

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

func PrintFilesForDirectory(s string) {

	out := FindDirectory(s)
	dirout := strings.Split(out.String(), "\n")
	var max int64
	var tmp int64
	var tmpsize int64
	var tmpfile string
	var tmppath string
	for _, dir := range dirout {
		if dir != "" {

			//fmt.Println("Main Dir:", dir)

			f, err := os.Open(dir)
			if err != nil {
				log.Fatal(err)
			}
			files, err := f.Readdir(-1)
			f.Close()
			if err != nil {
				log.Fatal(err)
			}

			for _, file := range files {
				if !file.IsDir() && file.Name() != ".swap" {
					t := file.Size()
					max = t
					if max > tmp {
						tmppath = dir
						tmpfile = file.Name()
						tmpsize = file.Size()
						tmp = max
					} else {
						max = tmp
					}

				}
			}

		}

	}
	fmt.Println(tmppath, tmpfile, tmpsize)

}

//===============================================
