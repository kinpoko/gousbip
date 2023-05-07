package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strings"

	"github.com/ktr0731/go-fuzzyfinder"
)

func main() {
	var attach bool
	var detach bool
	flag.BoolVar(&attach, "a", false, "attach device")
	flag.BoolVar(&detach, "d", false, "detach device")
	flag.Parse()

	if attach == detach {
		fmt.Println("Please specify -a or -d.")
		return
	}

	cmd := exec.Command("powershell.exe", "-ExecutionPolicy", "Bypass", "usbipd", "wsl", "list")
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	lines := strings.Split(string(stdout), "\n")
	devices := lines[1:]

	idx, err := fuzzyfinder.Find(devices, func(i int) string {
		return devices[i]
	},
		fuzzyfinder.WithHeader("Please select a device:"),
	)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	deviceid := strings.Fields(devices[idx])[0]

	if attach {
		cmd = exec.Command("powershell.exe", "-ExecutionPolicy", "Bypass", "usbipd", "wsl", "attach", "--busid", deviceid)
		stdout, err = cmd.Output()
		if err != nil {
			fmt.Println(err.Error())
			fmt.Printf("Could not attach %s\n", deviceid)
			return
		}
		fmt.Printf("Device %s has been attached.\n", deviceid)
		return
	}

	if detach {
		cmd = exec.Command("powershell.exe", "-ExecutionPolicy", "Bypass", "usbipd", "wsl", "detach", "--busid", deviceid)
		stdout, err = cmd.Output()
		if err != nil {
			fmt.Println(err.Error())
			fmt.Printf("Could not detach %s\n", deviceid)
			return
		}
		fmt.Printf("Device %s has been detached.\n", deviceid)
		return
	}
}
