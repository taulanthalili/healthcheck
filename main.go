package main

import (
	"fmt"
	"github.com/taulanthalili/healthcheck/osmodule"
)

var partitions = make([]string, 3)

func main() {

	//TEST OSMODULE
	//fmt.Println(osmodule.GetOsArch())
	//fmt.Println(osmodule.TestOsModule("OS"))

	ostmp := osmodule.GetOs()

	//Check Partitions we use
	partitions[0] = "/"
	partitions[1] = "/home"
	partitions[2] = "/data"

	for i := 0; i < len(partitions); i++ {
		t := osmodule.CheckPartitionFreePercent(partitions[i])

		if t > 0 {
			fmt.Println("Partition ", partitions[i], "free space in %:", t)

			if ostmp == "linux" {
				//osmodule.PrintTopBigDir(partitions[i])
				osmodule.PrintFilesForDirectory(partitions[i])
			}
		}
	}

	osmodule.CheckTopProcessSwap()
	//osmodule.GetCPUinfoTEST()
	osmodule.CheckTopProcessCPU()
	osmodule.CheckTopProcessMemory()

	// osmodule.PrintTopBigDir("/Users/taulant/Documents")
	//osmodule.PrintTopBigDir("/private/var")

	//TEST NGINXMOD
	// fmt.Println(nginxmod.TestNginxmod("Nginx"))

	//TEST Service
	// fmt.Println(services.TestServices("Services"))

	//TEST Magento
	// fmt.Println(magento.TestMagento("Magento"))

}
