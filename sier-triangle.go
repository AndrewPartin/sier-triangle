package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"math/rand"
	"os"
	"strconv"

	"github.com/cheggaaa/pb/v3" // progress bar
)

func errorMsg() {
	fmt.Println("Useage:")
	fmt.Println("sier-triangle {scale} {iterations}")
	os.Exit(1)
}

func main() {

	if len(os.Args) != 3 {
		errorMsg()
	}

	// create vars
	scale, err := strconv.Atoi(os.Args[1])
	if err != nil {
		errorMsg()
	}
	itera, err := strconv.Atoi(os.Args[2])
	if err != nil {
		errorMsg()
	}
	x, y := .5*float64(scale), .5*float64(scale)
	var n int

	// create progress bar
	bar := pb.StartNew(itera)

	// create image
	img := image.NewRGBA(image.Rect(0, 0, scale, scale))

	// set image background to all black
	draw.Draw(img, img.Bounds(), &image.Uniform{color.RGBA{0, 0, 0, 255}}, image.ZP, draw.Src)

	// main loop; runs through iterations
	for i := 0; i < itera; i++ {

		// increment progress bar
		bar.Increment()

		// set pixel at x, y green
		img.SetRGBA(int(math.Round(x)), int(math.Round(y)), color.RGBA{0, 255, 0, 255})

		// algorithm wizardry
        n = rand.Intn(3)
		x = .5 * (x + .5*float64(n*scale))
		y = .5 * y
		if n == 1 {
			y += .5 * float64(scale)
		}

	}

	// create out.png file
	out, err := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// encode image to file
	png.Encode(out, img)
}
