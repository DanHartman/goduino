package main

import (
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
	"log"
	"os"
	"time"
)

func main() {
	// environment variable assigns port for app to listen on
	port := os.Getenv("DEV_PORT")
	if port == "" {
		log.Fatal("Env Var DEV_PORT must be set")
	}
	firmataAdaptor := firmata.NewAdaptor(port)
	led := gpio.NewLedDriver(firmataAdaptor, "13")

	work := func() {
		gobot.Every(1*time.Second, func() {
			led.Toggle()
		})
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{led},
		work,
	)

	robot.Start()
}
