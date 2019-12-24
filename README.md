# About escpos #

This is a simple [Go][1] package that provides [ESC-POS][2] library functions
to help with sending control codes to a ESC-POS capable printer such as an
Epson TM-T82 or similar.

These printers are often used in retail environments in conjunction with a
point-of-sale (POS) system.


## Installation ##

Install the package via the following:

    go get -u github.com/epsimatic/escpos

## Example epos-server ##

An example EPOS server implementation is available in [original repository](https://github.com/CloudInn/escpos/tree/master/cmd/epos-server).
That example server is more or less compatible with [Epson TM-Intelligent][4] printers and print server implementations.

Please note that it requires `gokogiri`, `libxml` and `cgo` to build

## Usage ##

The escpos package can be used similarly to the following:

```go
package main

import (
    "bufio"
	"log"
    "os"

    "github.com/epsimatic/escpos"
)

func main() {
	f, err := os.OpenFile("/dev/usb/lp0", os.O_WRONLY, 0644)
	println("Port opened")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("Printer close failed: %v", err)
		}
	}()

	w := bufio.NewWriter(f)
	defer func() {
		if err := w.Flush(); err != nil {
			log.Fatalf("Print buffer flush failed: %v", err)
		}
	}()

	p, err := escpos.NewPrinter(w)
	if err != nil {
		log.Fatalf("Print init failed: %v", err)
	}
	defer p.End()
    p.Init()

	println("Printer opened")

    p.SetSmooth(1)
    p.SetFontSize(2, 3)
    p.SetFont("A")
    p.WriteLn("test ")
    p.SetFont("B")
    p.WriteLn("test2 ")
    p.SetFont("C")
    p.WriteLn("test3 ")
    p.Formfeed()

    p.SetFont("B")
    p.SetFontSize(1, 1)

    p.SetEmphasize(1)
    p.WriteLn("halle")
    p.Formfeed()

    p.SetUnderline(1)
    p.SetFontSize(4, 4)
    p.WriteLn("halle")

    p.SetReverse(1)
    p.SetFontSize(2, 4)
    p.WriteLn("halle")
    p.Formfeed()

    p.SetFont("C")
    p.SetFontSize(8, 8)
    p.WriteLn("halle")
    p.FormfeedN(5)

    p.Cut()
}
```

## NOTE
The Imported font inside the code is a system font called DejaVuSansMono-Bold.ttfsoyou shoukd make sure it exists in the system and it'splaced in the "/usr/share/fonts/truetype/dejavu/"


## TODO
- Fix barcode support

## Credits
- Repo forked from [kenshaw](https://github.com/kenshaw/escpos) escpos
- Repo forked from [CloudInn](https://github.com/CloudInn/escpos/tree/master/cmd/epos-server) escpos
- Some of the work in this repo is based on [python-escpos](https://github.com/python-escpos/python-escpos) and [escpos-php](https://github.com/mike42/escpos-php) packages


[1]: http://www.golang.org/project
[2]: https://en.wikipedia.org/wiki/ESC/P
[4]: https://c4b.epson-biz.com
