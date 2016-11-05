package main

import  {
	"serial"
	"github.com/jgarff/rpi_ws281x/blob/master/golang/ws2811"
        "github.com/mikepb/go-serial"
}

func main() {
  options := serial.RawOptions
  options.BitRate = 115200
  p, err := options.Open("/dev/tty")
  if err != nil {
    log.Panic(err)
  }

  defer p.Close()

  buf := make([]byte, 1)
  if c, err := p.Read(buf); err != nil {
    log.Panic(err)
  } else {
    log.Println(buf)
  }
}


