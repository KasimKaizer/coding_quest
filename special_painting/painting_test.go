package painting

import (
	"testing"
)

var TestCases = []struct {
	Description  string
	InputImgPath string
	ProcessColor Color
	Expected     string
}{
	{
		"Base Case Test",
		"base_case_img.png",
		Red,
		"world",
	},
	{
		"Real Case Test",
		"real_test_img.png",
		Red,
		"cake",
	},
}

func TestDecodeImage(t *testing.T) {
	for _, tt := range TestCases {
		t.Run(tt.Description, func(t *testing.T) {
			got, err := DecodeImage(tt.InputImgPath, tt.ProcessColor)
			if err != nil {
				t.Fatal(err)
			}
			if got != tt.Expected {
				t.Fatalf("expected: %s, got: %s", tt.Expected, got)
			}
		})
	}
}

/*
// Created to process the the base case image, as the one provided on the website was too big for
// aurate testing.
// so I took the hex data shown on the website and converted that into a 15X15 image.


func createBaseImg(imgDataPath string) error {
	f, err := os.Open(imgDataPath)
	if err != nil {
		return err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	size := 15
	topLeft := image.Point{X: 0, Y: 0}
	bottomRight := image.Point{X: size, Y: size}
	img := image.NewRGBA(image.Rectangle{topLeft, bottomRight})
	y := 0
	for scanner.Scan() {
		splitData := strings.Fields(scanner.Text())
		for idx, pix := range splitData {
			img.Set(idx, y, hexColor(pix))
		}
		y++
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	nf, err := os.Create("base_case_img.png")
	if err != nil {
		return err
	}
	return png.Encode(nf, img)
}

func hexColor(hex string) color.RGBA {
	values, _ := strconv.ParseUint(strings.Repeat(hex, 3), 16, 32)
	return color.RGBA{R: uint8(values >> 16), G: uint8((values >> 8) & 0xFF), B: uint8(values & 0xFF), A: 255}
}
*/
