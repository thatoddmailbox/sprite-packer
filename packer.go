package main

import (
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func Packer_Pack(o Options) {
	if o.Width * o.Height < len(o.Files) {
		log.Fatal("Spritesheet is too small for given number of files!")
	}

	out := image.NewRGBA(image.Rect(0, 0, o.Width*o.TileWidth, o.Height*o.TileHeight))

	x, y := 0, 0
	for _, fileInfo := range o.Files {
		file, err := os.Open(fileInfo.Path)
		if err != nil { panic(err) }
		defer file.Close()
		spriteImg, _, err := image.Decode(file)
		if err != nil { panic(err) }

		// very high-quality code here
		if fileInfo.Rotation == 90 {
			spriteImg = Utils_Rotate90(spriteImg)
		}
		if fileInfo.Rotation == 180 {
			spriteImg = Utils_Rotate90(spriteImg)
			spriteImg = Utils_Rotate90(spriteImg)
		}
		if fileInfo.Rotation == 270 {
			spriteImg = Utils_Rotate90(spriteImg)
			spriteImg = Utils_Rotate90(spriteImg)
			spriteImg = Utils_Rotate90(spriteImg)
		}

		draw.Draw(out, image.Rect(x*o.TileWidth, y*o.TileHeight, x*o.TileWidth + o.TileWidth, y*o.TileHeight + o.TileHeight), spriteImg, image.ZP, draw.Src)

		x += 1
		if x >= o.Width {
			x = 0
			y += 1
		}
	}

	outFile, err := os.OpenFile(o.OutputFile, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil { panic(err) }
	png.Encode(outFile, out)
}