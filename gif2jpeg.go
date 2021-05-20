package  gif2jpeg

import (
	"image"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"os"
	"path"
)



func TransGif2Jpeg(sourceGif string)  (filepath string,err error) {
	gifData, err := os.Open(sourceGif)
	if err != nil {
		return "",err
	}
	fileName := path.Base(sourceGif)
	all, err := gif.DecodeAll(gifData)
	if err != nil {
		return  "",err
	}
	filepath ="/tmp/"
	filepath += fileName +".jpeg"
	jp, err := os.Create(filepath)

	if err != nil {
		return  "",err
	}

	l := len(all.Image)
	if l%2 != 0 {
		l += 1
	}

	p1 := image.NewPaletted(image.Rect(0, 0, all.Config.Width,all.Config.Height), palette.Plan9)
	draw.Draw(p1, p1.Bounds(), 	all.Image[l/2], image.ZP, draw.Src) //添加图片

	op:= jpeg.Options{90}
	jpeg.Encode(jp,p1,&op)
	return filepath,err
}
