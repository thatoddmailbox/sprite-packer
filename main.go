package main

import (
	"flag"
	"log"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

type Options struct {
	FileListFile string
	Files []FileEntry
	OutputFile string
	TileWidth int
	TileHeight int
	Width int
	Height int
}

type FileEntry struct {
	Path string
	Rotation int
}

func main() {
	log.Println("sprite-packer")

	options := Options{}

	flag.StringVar(&options.FileListFile, "fileList", "", "List of files to pack into spritesheet")
	flag.StringVar(&options.OutputFile, "output", "", "Where to output the spritesheet (PNG format)")
	flag.IntVar(&options.TileWidth, "tw", 0, "The width, in pixels, of one tile")
	flag.IntVar(&options.TileHeight, "th", 0, "The height, in pixels, of one tile")
	flag.IntVar(&options.Width, "w", 0, "The width, in tiles, of the spritesheet")
	flag.IntVar(&options.Height, "h", 0, "The height, in tiles, of the spritesheet")

	flag.Parse()

	if options.FileListFile == "" {
		log.Fatal("Please specify the file list")
	}
	if options.OutputFile == "" {
		log.Fatal("Please specify the output file")
	}
	if options.TileWidth == 0 || options.TileHeight == 0 {
		log.Fatal("Please specify the tile dimensions")
	}
	if options.Width == 0 || options.Height == 0 {
		log.Fatal("Please specify the spritesheet dimensions")
	}

	// parse file list
	// kinda crappy way of doing it but too bad
	fileListContents, err := ioutil.ReadFile(options.FileListFile)
	if err != nil {
		panic(err)
	}

	relBasePath := filepath.Dir(options.FileListFile)
	fullBasePath, err := filepath.Abs(relBasePath)
	if err != nil {
		panic(err)
	}

	fileListItems := strings.Split(string(fileListContents), "\n")
	for _, relativePath := range fileListItems {
		inputLine := strings.TrimSpace(relativePath)
		if inputLine == "" || strings.HasPrefix(inputLine, ";") {
			continue
		}

		parts := strings.Split(inputLine, ";")
		inputPath := parts[0]
		rotation := 0

		if len(parts) > 1 {
			rotation, err = strconv.Atoi(parts[1])
			if err != nil { panic(err) }
		}

		options.Files = append(options.Files, FileEntry{filepath.Join(fullBasePath, inputPath), rotation})
	}

	Packer_Pack(options)
}