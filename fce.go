package main

import (
		"strings"
		"strconv"
		"encoding/json"
		"github.com/go-resty/resty/v2"
		"net/http"
		"github.com/labstack/echo/v4"
)

func GetText(c echo.Context) error {
	t := c.QueryParam("text")
	ch <- t
	return c.String(http.StatusOK, "text: " + t)
}

func GetPrice() int {
	client := resty.New()

	resp, err := client.R().
		EnableTrace().
		Get("https://api.bitfinex.com/v1/pubticker/btcusd")
	if err == nil {
		type Pars struct {
			Price	string `json:"last_price"`
		}
		pars := Pars{}
		json.Unmarshal(resp.Body(), &pars)
		btc := strings.Split(pars.Price,".")
		btci, _ := strconv.Atoi(btc[0])
		return btci
	}
	return 0
}

func GetTemp() float32 {
	client := resty.New()

	resp, err := client.R().
		EnableTrace().
		Get("http://172.17.0.133:10000/")
	if err == nil {
		type Pars struct {
			Temp	float32 `json:"temperature"`
		}
		pars := Pars{}
		json.Unmarshal(resp.Body(), &pars)
		return pars.Temp
	}
	return 0
}

func GetPress() float32 {
	client := resty.New()

	resp, err := client.R().
		EnableTrace().
		Get("http://172.17.0.133:10000/")
	if err == nil {
		type Pars struct {
			Press	float32 `json:"pressure"`
		}
		pars := Pars{}
		json.Unmarshal(resp.Body(), &pars)
		return pars.Press
	}
	return 0
}

func GetIll() int {
	client := resty.New()

	resp, err := client.R().
		EnableTrace().
		Get("http://172.17.0.123:10000/")
	if err == nil {
		type Pars struct {
			Ill		int `json:"illuminance"`
		}
		pars := Pars{}
		json.Unmarshal(resp.Body(), &pars)
		return pars.Ill
	}
	return 0
}

func TextToRGB(fg uint32, bg uint32, text string, data []uint32){
    for i := 0; i < len(text) && i < panels; i++ {
        CharToRGB(fg, bg, text[i], (i*panel)+first, data)
    }
}

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
		case '.':
			data[i+0] = bg; data[i+15] = bg; data[i+16] = bg; data[i+31] = bg; data[i+32] = bg
			data[i+1] = bg; data[i+14] = bg; data[i+17] = bg; data[i+30] = bg; data[i+33] = bg
			data[i+2] = bg; data[i+13] = bg; data[i+18] = bg; data[i+29] = bg; data[i+34] = bg
			data[i+3] = bg; data[i+12] = bg; data[i+19] = bg; data[i+28] = bg; data[i+35] = bg
			data[i+4] = bg; data[i+11] = bg; data[i+20] = bg; data[i+27] = bg; data[i+36] = bg
			data[i+5] = bg; data[i+10] = bg; data[i+21] = bg; data[i+26] = bg; data[i+37] = bg
			data[i+6] = bg; data[i+9]  = bg; data[i+22] = bg; data[i+25] = bg; data[i+38] = bg
			data[i+7] = bg; data[i+8]  = bg; data[i+23] = fg; data[i+24] = bg; data[i+39] = bg
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
