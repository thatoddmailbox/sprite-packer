# sprite-packer
This program packs multiple, separate image files into a single spritesheet. It currently assumes that every tile in the spritesheet has the same dimension, and only supports PNG files.

## Usage
First, create a list of the images you want to include in your spritesheet. This list should be a text file where each line corresponds to a file name. You may include comments with the semicolon (`;`) character. For example:

```
; tiles
tile1.png
tile2.png

; other
other1.png
yet_another_file.png
again_a_file.png
```

Then, run the `sprite-packer` program, like so: `./sprite-packer -fileList list.txt -tw 32 -th 32 -w 2 -h 3 -output output.png`. This command will read the file list from `list.txt`, create a spritesheet sized 2 tiles by 3 tiles, where tiles have a size of 32 pixels by 32 pixels, and save that spritesheet to `output.png`.