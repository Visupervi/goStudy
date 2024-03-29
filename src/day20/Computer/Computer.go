package Computer

import (
	_interface "day20/interface"
)

type Computer struct {
}

func (c *Computer) Working(usb _interface.IUsb) {
	usb.Stop()
	usb.Start()
}
