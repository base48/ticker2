package main

import (
		"fmt"
		"time"
		ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
)

const (
	first= 5 // first 5 leds before panels 
	ti = 5 // time interval in seconds
	ilt = 0 // illumination threshold
	brig = 255
	panels = 8
	panel = (8*5)
	ledc = (panels*panel)+first // ((8*5 from top to bottom) * 5 panels) + leds from start 
)

func main(){
	opt := ws2811.DefaultOptions
	opt.Channels[0].Brightness = brig
	opt.Channels[0].LedCount = ledc
	opt.Channels[0].GpioPin = 18

	dev, _ := ws2811.MakeWS2811(&opt)

	lt := time.Now()
	format := "0"
	stav := 0
	var fg, bg uint32

	dev.Init()
	defer dev.Fini()
	for {
		t := time.Now()
		if t.After(lt.Add(ti * time.Second)) {
			il := GetIll()
			if stav == 0 {
				format = fmt.Sprintf("%8d", GetPrice())
				fg, bg = 0xffff00, 0x000000
				if il <= ilt { fg, bg = 0x706000, 0x000000 }
				stav = 1
			} else if stav == 1 {
				format = fmt.Sprintf("%8.1f", GetTemp())
				fg, bg = 0xff0040, 0x000000
				if il <= ilt { fg, bg = 0x600010, 0x000000 }
				stav = 2
			} else {
				format = fmt.Sprintf("%8.1f", GetPress()/100)
				fg, bg = 0x00ff00, 0x000000
				if il <= ilt { fg, bg = 0x006000, 0x000000 }
				stav = 0
			}
			fmt.Println(format)
			lt = t
		}

		TextToRGB(fg, bg, format, dev.Leds(0))
		dev.Render()
		time.Sleep(10 * time.Millisecond)
	}
}
