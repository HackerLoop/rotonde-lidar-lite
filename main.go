package main

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/hybridgroup/gobot/platforms/i2c"
	"github.com/hybridgroup/gobot/platforms/raspi"
)

func main() {
	host := raspi.NewRaspiAdaptor("host")

	dev := i2c.NewLIDARLiteDriver(host, "dev")

	dev.Start()

	go func() {
		if dist, err := dev.Distance(); err != nil {
			log.Error(err)
		} else {
			log.Info(dist)
		}
		time.Sleep(time.Second * 1)
	}()

	select {}
}
