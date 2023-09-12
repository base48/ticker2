package main

import (
	"encoding/json"
	resty "github.com/go-resty/resty/v2"
	echo "github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
	"path"
	"strconv"
	"strings"
)

func GetText(c echo.Context) error {
	t := c.QueryParam("text")
	ch <- t
	return c.String(http.StatusOK, "text: "+t)
}

func GetPrice() int {
	client := resty.New()

	resp, err := client.R().
		EnableTrace().
		Get("https://api.gemini.com/v1/pubticker/btcusd")
//		Get("https://api.bitfinex.com/v1/pubticker/btcusd")
	if err == nil {
		type Pars struct {
			Price string `json:"last"`
		}
		pars := Pars{}
		json.Unmarshal(resp.Body(), &pars)
		btc := strings.Split(pars.Price, ".")
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
			Temp float32 `json:"temperature"`
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
			Press float32 `json:"pressure"`
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
		Get("http://localhost:10000/")
	if err == nil {
		type Pars struct {
			Ill int `json:"illuminance"`
		}
		pars := Pars{}
		json.Unmarshal(resp.Body(), &pars)
		return pars.Ill
	}
	return 0
}

func TextToRGB(fg uint32, bg uint32, text string, data []uint32) {
	for i := 0; i < len(text) && i < panels; i++ {
		CharToRGB(fg, bg, text[i], (i*panel)+first, data)
	}
}

func CharToRGB(fg uint32, bg uint32, c byte, offset int, data []uint32) {
	len_x := 5
	len_y := 8
	font_path := "font"
	if c == 46 {
		c = 63
	}
	letter_file := path.Join(font_path, string(c))
	content, _ := ioutil.ReadFile(letter_file)

	matrix := make([][]byte, 0)
	for _, line := range strings.Split(string(content), "\n") {
		if line == "" {
			continue
		}
		matrix = append(matrix, []byte(line))
	}
	if len(matrix) == 0 {
		return
	}

	for i := 0; i < len_x; i++ {
		for j := 0; j < len_y; j++ {
			j_it := j
			if i%2 == 1 {
				j_it = len_y - j - 1
			}
			data_pos := offset + i*len_y + j_it
			color := fg
			if matrix[j][i] == byte('0') {
				color = bg
			} else {
				print()
			}
			data[data_pos] = color
		}
	}
}
