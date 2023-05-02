package main

import (
	"flag"
	"fmt"
	"image/color"
	"image/jpeg"
	"os"

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

	flag.Parse()

	// Background and overlay
	bg, _ := gg.LoadJPG(*bgImg)
	image := gg.NewContext(*width, *height)
	image.DrawImage(bg, 0, 0)
	image.DrawRectangle(0, 0, float64(image.Width()), float64(image.Height()))
	image.SetColor(color.RGBA{0, 0, 0, 100})
	image.Fill()

	// Frame for title
	margin := 40.0
	x := margin
	y := margin
	w := float64(image.Width()) - (2.0 * margin)
	h := float64(image.Height()) - (2.0 * margin)
	image.DrawRectangle(x, y, w, h)
	image.SetColor(color.RGBA{0, 0, 0, 200})
	image.Fill()

	if err := image.LoadFontFace(*fontface, float64(*fontsize)); err != nil {
		fmt.Println(err)
	}

	textColor := color.RGBA{255, 255, 255, 255}
	textRightMargin := 100.0
	textTopMargin := 100.0
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
