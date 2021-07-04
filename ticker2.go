package main

//import "fmt"
import "time"
import ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"

func main(){
	opt := ws2811.DefaultOptions
	opt.Channels[0].Brightness = 100
	opt.Channels[0].LedCount = 8
	opt.Channels[0].GpioPin = 10

	dev, _ := ws2811.MakeWS2811(&opt)

	dev.Init()
	defer dev.Fini()

	dev.Leds(0)[0] = 0xff0000
	dev.Leds(0)[1] = 0x00ff00
	dev.Leds(0)[2] = 0x0000ff
	dev.Leds(0)[3] = 0xff00ff
	dev.Leds(0)[4] = 0xffff00
	dev.Leds(0)[5] = 0x00ffff
	dev.Leds(0)[6] = 0xffffff
	dev.Leds(0)[7] = 0xf0f0f0
	dev.Render()

	time.Sleep(10000 * time.Millisecond)

//	fmt.Println("Last: ", GetPrice())
}
