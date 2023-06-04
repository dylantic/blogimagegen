package main

import (
	"flag"
	"fmt"
	"image/color"
	"image/jpeg"
	"os"
	"strconv"
	"strings"

	"github.com/fogleman/gg"
)

func main() {
	title := flag.String("title", "Title here", "Set the title")
	bgImg := flag.String("bgimg", "./bg/neoncity.jpg", "Set a background image")
	height := flag.Int("height", 400, "Set the height")
	width := flag.Int("width", 1000, "Set the width")
	fontface := flag.String("font", "./fonts/OpenSans_SemiCondensed-SemiBold.ttf", "Set the fontface to use")
	fontsize := flag.Int("fontsize", 50, "Set the font size to use")
	output := flag.String("output", "output.png", "Set the filename for the output file")
	format := flag.String("format", "png", "Set the output format for the image (png/jpg)")
	quality := flag.Int("quality", 90, "Set quality option for format if it's supported (jpeg)")
	overlay := flag.String("overlay", "0, 0, 0, 0", "Set the overlay colour and opacity 'red, green, blue, opacity' (0-255)")
	frame := flag.String("frame", "0, 0, 0, 200", "Set the frame colour and opacity 'red, green, blue, opacity' (0-255)")

	flag.Parse()

	// Background and overlay
	bg, _ := gg.LoadJPG(*bgImg)
	image := gg.NewContext(*width, *height)
	image.DrawImage(bg, 0, 0)
	image.DrawRectangle(0, 0, float64(image.Width()), float64(image.Height()))
	image.SetColor(colorFromArray(*overlay))
	image.Fill()

	// Frame for title
	margin := 40.0
	x := margin
	y := margin
	w := float64(image.Width()) - (2.0 * margin)
	h := float64(image.Height()) - (2.0 * margin)
	image.DrawRectangle(x, y, w, h)
	image.SetColor(colorFromArray(*frame))
	image.Fill()

	if err := image.LoadFontFace(*fontface, float64(*fontsize)); err != nil {
		fmt.Println(err)
	}

	textColor := color.RGBA{255, 255, 255, 255}
	textRightMargin := float64(*width / 10.0)
	textTopMargin := float64(*width / 10.0)
	x = textRightMargin
	y = textTopMargin
	maxWidth := float64(image.Width()) - textRightMargin - textRightMargin
	image.SetColor(textColor)
	image.DrawStringWrapped(*title, x, y, 0, 0, maxWidth, 1.5, gg.AlignLeft)

	if *format == "png" {
		err := image.SavePNG(*output)
		if err != nil {
			fmt.Println(err)
		}
	} else if *format == "jpeg" || *format == "jpg" {
		img := image.Image()
		outfile, err := os.Create(*output)
		if err != nil {
			fmt.Println(err)
		}
		jpegopts := jpeg.Options{
			Quality: *quality,
		}
		jpeg.Encode(outfile, img, &jpegopts)
	} else {
		fmt.Println("Unsupported image format. Use one of: jpeg, png")
	}

}

func colorFromArray(rgbArray string) color.RGBA {
	rgbData := strings.Split(rgbArray, ",")
	var intArray []uint8
	if len(rgbData) < 4 {
		fmt.Println("Malformed RGBA string")
		os.Exit(1)
	}
	for _, v := range rgbData {
		intArray = append(intArray, uint8(parseUint(v, 10, 8)))
	}
	return color.RGBA{intArray[0], intArray[1], intArray[2], intArray[3]}

}

func parseUint(number string, base int, bits int) uint64 {
	value, _ := strconv.ParseUint(number, base, bits)
	return value
}
