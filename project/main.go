package main

import (
	"fmt"

	"github.com/karalabe/usb"
)

func main() {
	// list := usb.DeviceInfo()
	fmt.Println(usb.Enumerate(0, 0))
	fmt.Println(usb.Supported())
	usb.

}
