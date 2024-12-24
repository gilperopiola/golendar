package main

import (
	"log"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	robotoFont font.Face
)

func init() {
	fontBytes, err := os.ReadFile("./fonts/Roboto-Bold.ttf")
	if err != nil {
		log.Fatal(err)
	}

	ttfFont, err := opentype.Parse(fontBytes)
	if err != nil {
		log.Fatal(err)
	}

	robotoFont, err = opentype.NewFace(ttfFont, &opentype.FaceOptions{
		Size:    32,
		DPI:     72,
		Hinting: font.HintingNone,
	})
	if err != nil {
		log.Fatal(err)
	}
}
