package painting

import (
	"image"
	"image/draw"
	"os"
	"regexp"
	"strings"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

type Color int

const (
	Red Color = iota + 1
	Green
	Blue
)

func DecodeImage(imgPath string, color Color) (string, error) {
	imgFile, err := os.Open(imgPath)
	if err != nil {
		return "", err
	}
	defer imgFile.Close()
	img, _, err := image.Decode(imgFile)
	if err != nil {
		return "", err
	}
	data := constructData(img, color)
	var message strings.Builder
	curBit := 0
	bitPos := 0
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			curBit <<= 1
			if data[i][j]&1 == 1 {
				curBit += 1
			}
			bitPos++
			if bitPos != 8 {
				continue
			}
			char := byte(curBit)
			if char == '\x00' {
				return getLastWord(message.String()), nil
			}
			message.WriteByte(char)
			bitPos, curBit = 0, 0
		}
	}
	return message.String(), nil
}

func constructData(img image.Image, color Color) [][]uint8 {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	rgba := image.NewRGBA(bounds)
	draw.Draw(rgba, bounds, img, bounds.Min, draw.Src)

	var out [][]uint8

	for y := 0; y < height; y++ {
		var temp []uint8
		for x := 0; x < width; x++ {
			idx := ((y * width) + x) * 4
			val := rgba.Pix[idx]
			if color == Green {
				val = rgba.Pix[idx+1]
			}
			if color == Blue {
				val = rgba.Pix[idx+2]
			}
			temp = append(temp, val)
		}
		out = append(out, temp)
	}

	return out
}

func getLastWord(text string) string {
	out := text[strings.LastIndex(text, " ")+1:]
	return regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(out, "")
}
