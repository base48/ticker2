package main

import "fmt"
import "time"
import "strings"
import "strconv"
import ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"

const(
	first= 5
	brig = 255
	poli = 8
	pole = (8*5)
	ledc = (pole*poli)+first // ((8*5 zhora dolu) * 5 poli) + 5 na zacatku
)

func CharToRGB(fg uint32, bg uint32, c byte, i int, data []uint32){
	switch c{
		case '1':
			data[i+0] = bg; data[i+15] = bg; data[i+16] = fg; data[i+31] = bg; data[i+32] = bg
			data[i+1] = bg; data[i+14] = fg; data[i+17] = fg; data[i+30] = bg; data[i+33] = bg
			data[i+2] = bg; data[i+13] = bg; data[i+18] = fg; data[i+29] = bg; data[i+34] = bg
			data[i+3] = bg; data[i+12] = bg; data[i+19] = fg; data[i+28] = bg; data[i+35] = bg
			data[i+4] = bg; data[i+11] = bg; data[i+20] = fg; data[i+27] = bg; data[i+36] = bg
			data[i+5] = bg; data[i+10] = bg; data[i+21] = fg; data[i+26] = bg; data[i+37] = bg
			data[i+6] = bg; data[i+9]  = bg; data[i+22] = fg; data[i+25] = bg; data[i+38] = bg
			data[i+7] = bg; data[i+8]  = fg; data[i+23] = fg; data[i+24] = fg; data[i+39] = bg
		case '2':
			data[i+0] = bg; data[i+15] = fg; data[i+16] = fg; data[i+31] = fg; data[i+32] = bg
			data[i+1] = fg; data[i+14] = bg; data[i+17] = bg; data[i+30] = bg; data[i+33] = fg
			data[i+2] = bg; data[i+13] = bg; data[i+18] = bg; data[i+29] = bg; data[i+34] = fg
			data[i+3] = bg; data[i+12] = bg; data[i+19] = bg; data[i+28] = bg; data[i+35] = fg
			data[i+4] = bg; data[i+11] = bg; data[i+20] = bg; data[i+27] = fg; data[i+36] = bg
			data[i+5] = bg; data[i+10] = bg; data[i+21] = fg; data[i+26] = bg; data[i+37] = bg
			data[i+6] = bg; data[i+9]  = fg; data[i+22] = bg; data[i+25] = bg; data[i+38] = bg
			data[i+7] = fg; data[i+8]  = fg; data[i+23] = fg; data[i+24] = fg; data[i+39] = fg
		case '3':
			data[i+0] = fg; data[i+15] = fg; data[i+16] = fg; data[i+31] = fg; data[i+32] = fg
			data[i+1] = bg; data[i+14] = bg; data[i+17] = bg; data[i+30] = bg; data[i+33] = fg
			data[i+2] = bg; data[i+13] = bg; data[i+18] = bg; data[i+29] = fg; data[i+34] = bg
			data[i+3] = bg; data[i+12] = bg; data[i+19] = fg; data[i+28] = bg; data[i+35] = bg
			data[i+4] = bg; data[i+11] = bg; data[i+20] = bg; data[i+27] = fg; data[i+36] = bg
			data[i+5] = bg; data[i+10] = bg; data[i+21] = bg; data[i+26] = bg; data[i+37] = fg
			data[i+6] = fg; data[i+9]  = bg; data[i+22] = bg; data[i+25] = bg; data[i+38] = fg
			data[i+7] = bg; data[i+8]  = fg; data[i+23] = fg; data[i+24] = fg; data[i+39] = bg
		case '4':
			data[i+0] = fg; data[i+15] = bg; data[i+16] = bg; data[i+31] = bg; data[i+32] = bg
			data[i+1] = fg; data[i+14] = bg; data[i+17] = bg; data[i+30] = bg; data[i+33] = bg
			data[i+2] = fg; data[i+13] = bg; data[i+18] = fg; data[i+29] = bg; data[i+34] = bg
			data[i+3] = fg; data[i+12] = bg; data[i+19] = fg; data[i+28] = bg; data[i+35] = bg
			data[i+4] = fg; data[i+11] = bg; data[i+20] = fg; data[i+27] = bg; data[i+36] = bg
			data[i+5] = fg; data[i+10] = fg; data[i+21] = fg; data[i+26] = fg; data[i+37] = fg
			data[i+6] = bg; data[i+9]  = bg; data[i+22] = fg; data[i+25] = bg; data[i+38] = bg
			data[i+7] = bg; data[i+8]  = bg; data[i+23] = fg; data[i+24] = bg; data[i+39] = bg
		case '5':
			data[i+0] = fg; data[i+15] = fg; data[i+16] = fg; data[i+31] = fg; data[i+32] = fg
			data[i+1] = fg; data[i+14] = bg; data[i+17] = bg; data[i+30] = bg; data[i+33] = bg
			data[i+2] = fg; data[i+13] = bg; data[i+18] = bg; data[i+29] = bg; data[i+34] = bg
			data[i+3] = fg; data[i+12] = fg; data[i+19] = fg; data[i+28] = fg; data[i+35] = bg
			data[i+4] = bg; data[i+11] = bg; data[i+20] = bg; data[i+27] = bg; data[i+36] = fg
			data[i+5] = bg; data[i+10] = bg; data[i+21] = bg; data[i+26] = bg; data[i+37] = fg
			data[i+6] = fg; data[i+9]  = bg; data[i+22] = bg; data[i+25] = bg; data[i+38] = fg
			data[i+7] = bg; data[i+8]  = fg; data[i+23] = fg; data[i+24] = fg; data[i+39] = bg
		case '6':
			data[i+0] = bg; data[i+15] = fg; data[i+16] = fg; data[i+31] = fg; data[i+32] = bg
			data[i+1] = fg; data[i+14] = bg; data[i+17] = bg; data[i+30] = bg; data[i+33] = fg
			data[i+2] = fg; data[i+13] = bg; data[i+18] = bg; data[i+29] = bg; data[i+34] = bg
			data[i+3] = fg; data[i+12] = fg; data[i+19] = fg; data[i+28] = fg; data[i+35] = bg
			data[i+4] = fg; data[i+11] = bg; data[i+20] = bg; data[i+27] = bg; data[i+36] = fg
			data[i+5] = fg; data[i+10] = bg; data[i+21] = bg; data[i+26] = bg; data[i+37] = fg
			data[i+6] = fg; data[i+9]  = bg; data[i+22] = bg; data[i+25] = bg; data[i+38] = fg
			data[i+7] = bg; data[i+8]  = fg; data[i+23] = fg; data[i+24] = fg; data[i+39] = bg
		case '7':
			data[i+0] = fg; data[i+15] = fg; data[i+16] = fg; data[i+31] = fg; data[i+32] = fg
			data[i+1] = bg; data[i+14] = bg; data[i+17] = bg; data[i+30] = bg; data[i+33] = fg
			data[i+2] = bg; data[i+13] = bg; data[i+18] = bg; data[i+29] = fg; data[i+34] = bg
			data[i+3] = bg; data[i+12] = bg; data[i+19] = fg; data[i+28] = bg; data[i+35] = bg
			data[i+4] = bg; data[i+11] = bg; data[i+20] = fg; data[i+27] = bg; data[i+36] = bg
			data[i+5] = bg; data[i+10] = fg; data[i+21] = bg; data[i+26] = bg; data[i+37] = bg
			data[i+6] = bg; data[i+9]  = fg; data[i+22] = bg; data[i+25] = bg; data[i+38] = bg
			data[i+7] = bg; data[i+8]  = fg; data[i+23] = bg; data[i+24] = bg; data[i+39] = bg
		case '8':
			data[i+0] = bg; data[i+15] = fg; data[i+16] = fg; data[i+31] = fg; data[i+32] = bg
			data[i+1] = fg; data[i+14] = bg; data[i+17] = bg; data[i+30] = bg; data[i+33] = fg
			data[i+2] = fg; data[i+13] = bg; data[i+18] = bg; data[i+29] = bg; data[i+34] = fg
			data[i+3] = bg; data[i+12] = fg; data[i+19] = fg; data[i+28] = fg; data[i+35] = bg
			data[i+4] = fg; data[i+11] = bg; data[i+20] = bg; data[i+27] = bg; data[i+36] = fg
			data[i+5] = fg; data[i+10] = bg; data[i+21] = bg; data[i+26] = bg; data[i+37] = fg
			data[i+6] = fg; data[i+9]  = bg; data[i+22] = bg; data[i+25] = bg; data[i+38] = fg
			data[i+7] = bg; data[i+8]  = fg; data[i+23] = fg; data[i+24] = fg; data[i+39] = bg
		case '9':
			data[i+0] = bg; data[i+15] = fg; data[i+16] = fg; data[i+31] = fg; data[i+32] = bg
			data[i+1] = fg; data[i+14] = bg; data[i+17] = bg; data[i+30] = bg; data[i+33] = fg
			data[i+2] = fg; data[i+13] = bg; data[i+18] = bg; data[i+29] = bg; data[i+34] = fg
			data[i+3] = fg; data[i+12] = bg; data[i+19] = bg; data[i+28] = bg; data[i+35] = fg
			data[i+4] = bg; data[i+11] = fg; data[i+20] = fg; data[i+27] = fg; data[i+36] = fg
			data[i+5] = bg; data[i+10] = bg; data[i+21] = bg; data[i+26] = bg; data[i+37] = fg
			data[i+6] = fg; data[i+9]  = bg; data[i+22] = bg; data[i+25] = bg; data[i+38] = fg
			data[i+7] = bg; data[i+8]  = fg; data[i+23] = fg; data[i+24] = fg; data[i+39] = bg
		case '0':
			data[i+0] = bg; data[i+15] = fg; data[i+16] = fg; data[i+31] = fg; data[i+32] = bg
			data[i+1] = fg; data[i+14] = bg; data[i+17] = bg; data[i+30] = bg; data[i+33] = fg
			data[i+2] = fg; data[i+13] = bg; data[i+18] = bg; data[i+29] = bg; data[i+34] = fg
			data[i+3] = fg; data[i+12] = bg; data[i+19] = bg; data[i+28] = fg; data[i+35] = fg
			data[i+4] = fg; data[i+11] = bg; data[i+20] = fg; data[i+27] = bg; data[i+36] = fg
			data[i+5] = fg; data[i+10] = fg; data[i+21] = bg; data[i+26] = bg; data[i+37] = fg
			data[i+6] = fg; data[i+9]  = bg; data[i+22] = bg; data[i+25] = bg; data[i+38] = fg
			data[i+7] = bg; data[i+8]  = fg; data[i+23] = fg; data[i+24] = fg; data[i+39] = bg
		default :
			data[i+0] = bg; data[i+15] = bg; data[i+16] = bg; data[i+31] = bg; data[i+32] = bg
			data[i+1] = bg; data[i+14] = bg; data[i+17] = bg; data[i+30] = bg; data[i+33] = bg
			data[i+2] = bg; data[i+13] = bg; data[i+18] = bg; data[i+29] = bg; data[i+34] = bg
			data[i+3] = bg; data[i+12] = bg; data[i+19] = bg; data[i+28] = bg; data[i+35] = bg
			data[i+4] = bg; data[i+11] = bg; data[i+20] = bg; data[i+27] = bg; data[i+36] = bg
			data[i+5] = bg; data[i+10] = bg; data[i+21] = bg; data[i+26] = bg; data[i+37] = bg
			data[i+6] = bg; data[i+9]  = bg; data[i+22] = bg; data[i+25] = bg; data[i+38] = bg
			data[i+7] = bg; data[i+8]  = bg; data[i+23] = bg; data[i+24] = bg; data[i+39] = bg
	}
}

func TextToRGB(fg uint32, bg uint32, text string, data []uint32){
	for i := 0; i < len(text) && i < poli; i++ {
		CharToRGB(fg, bg, text[i], (i*pole)+first, data)
	}
}

func main(){
	opt := ws2811.DefaultOptions
	opt.Channels[0].Brightness = brig
	opt.Channels[0].LedCount = ledc
	opt.Channels[0].GpioPin = 18

	dev, _ := ws2811.MakeWS2811(&opt)

	dev.Init()
	defer dev.Fini()
	for {

		btc := strings.Split(GetPrice(),".")
		btci, _ := strconv.Atoi(btc[0])
		format := fmt.Sprintf("%8d", btci)
	//	fmt.Println(format)
		TextToRGB(0x00ff00, 0x000000, format, dev.Leds(0))
		dev.Render()
		time.Sleep(10000 * time.Millisecond)
	}
}
