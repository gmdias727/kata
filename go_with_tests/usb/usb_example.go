package example_usb

import "fmt"

func ListUSB() {
	fmt.Println(usb.Supported())
	fmt.Println()
}
