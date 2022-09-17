package osmodule

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func DirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}

func CheckTopBigDir(s string) (string, int64) {

	tmpdir := s
	//tmpdir := "/Users"

	c1 := make(chan int64, 1)
	go func() {
		//time.Sleep(2 * time.Second)
		t, err := DirSize(tmpdir)
		if err == nil {
			c1 <- t / MB
		}
	}()

	select {
	case res := <-c1:
		return tmpdir, res
	case <-time.After(2 * time.Second):
		return tmpdir, -1
	}

}

func PrintTopBigDir(s string) {

	out := FindDirectory(s)
	dirout := strings.Split(out.String(), "\n")

	for _, dir := range dirout {
		if dir != "" {

			//fmt.Println("Main Dir:", dir)
			tmpdir, size := CheckTopBigDir(dir)

			if size > 999 || size == -1 {
				fmt.Println(tmpdir, size)
			}

		}

	}

}

//===============================================
