package connection

import (
	"io"
	"net"
	"os"

	"github.com/epsimatic/escpos"
)

//NewConnection creats a connection with a usb printer or a network printer and
//returns an object to use escops package functions with
func NewConnection(connectionType string, connectionHost string) (*escpos.Printer, error) {
	var f io.Writer
	var err error

	if connectionType == "usb" {
		f, err = os.OpenFile(connectionHost, os.O_WRONLY, 0)
	} else if connectionType == "network" {
		f, err = net.Dial("tcp", connectionHost)
	}
	if err != nil {
		return nil, err
	}
	printerObj, err := escpos.NewPrinter(f)
	if err != nil {
		return nil, err
	}
	return printerObj, nil

}
