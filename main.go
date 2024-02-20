package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math/rand"
	"os"
	"time"

	"github.com/andrewmthomas87/wfc/wfc"
)

func main() {
	imgs, err := decodeImages()
	if err != nil {
		panic(err)
	}

	compatMatrix := buildCompatMatrix(arrows)
	compatFn := func(s1, s2 wfc.State, d wfc.Direction) bool {
		return compatMatrix[s1][s2][d]
	}

	rand := rand.New(rand.NewSource(time.Now().Unix()))
	w := wfc.NewWave(rand, len(arrows), weights, 48, 48, compatFn)

	img := generateImage(imgs, w)
	err = writeImage(0, img)
	if err != nil {
		panic(err)
	}

	i := 0
	for !(w.IsContradictory() || w.IsCollapsed()) {
		w.Iterate()

		img := generateImage(imgs, w)
		err = writeImage(i+1, img)
		if err != nil {
			panic(err)
		}

		i++
	}

	img = generateImage(imgs, w)
	err = writeImage(0, img)
	if err != nil {
		panic(err)
	}

	if w.IsContradictory() {
		os.Exit(1)
	}

	fmt.Printf("%d iterations\n", i)
}

func decodeImages() ([]image.Image, error) {
	imgs := make([]image.Image, len(arrows))
	for i := 0; i < len(arrows); i++ {
		f, err := os.Open(fmt.Sprintf("arrows/%d.png", i))
		if err != nil {
			return nil, err
		}
		defer f.Close()

		img, _, err := image.Decode(f)
		if err != nil {
			return nil, err
		}
		imgs[i] = img
	}

	return imgs, nil
}

func generateImage(imgs []image.Image, w *wfc.Wave) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, w.W*16, w.H*16))
	draw.Draw(
		img,
		img.Bounds(),
		&image.Uniform{color.RGBA{255, 255, 255, 255}},
		image.Point{},
		draw.Src,
	)

	for y, col := range w.Grid {
		for x, c := range col {
			if c.IsContradictory() {
				draw.Draw(
					img,
					image.Rect(16*x, 16*y, 16*(x+1), 16*(y+1)),
					image.NewUniform(color.RGBA{255, 0, 0, 255}),
					image.Point{},
					draw.Src,
				)
				continue
			}

			a := 255 / len(c.States)
			mask := image.NewUniform(color.Alpha{A: uint8(a)})

			for _, s := range c.States {
				src := imgs[s]
				draw.DrawMask(
					img,
					image.Rect(16*x, 16*y, 16*(x+1), 16*(y+1)),
					src,
					src.Bounds().Min,
					mask,
					image.Point{},
					draw.Over,
				)
			}
		}
	}

	return img
}

func writeImage(i int, img image.Image) error {
	f, err := os.Create(fmt.Sprintf("output/%06d.png", i))
	if err != nil {
		return err
	}
	defer f.Close()

	return png.Encode(f, img)
}
