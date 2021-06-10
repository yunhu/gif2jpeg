package  gif2jpeg

import (
	"image"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"net/http"
	"os"
	"path"
)

/**
 * @Description:
 * @param sourceGif 源图地址
 * @param frame 取第几帧，传0或者过大过小都会取默认值
 * @return filepath 返回的生成的文件地址
 * @return err
 */

func TransGif2JpegFile(sourceGif string,frame int)  (filepath string,err error) {
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
	if frame <= 0 || frame > l {
		frame = l /2
	}

	p1 := image.NewPaletted(image.Rect(0, 0, all.Config.Width,all.Config.Height), palette.Plan9)
	draw.Draw(p1, p1.Bounds(), 	all.Image[frame-1], image.ZP, draw.Src) //添加图片

	op:= jpeg.Options{90}
	jpeg.Encode(jp,p1,&op)
	return filepath,err
}

func TransGif2JpegUrl(url string,frame int)  (filepath string,err error) {
	data, err := http.Get(url)
	if err !=nil{
		return "",err
	}
	fileName := path.Base(url)
	all, err := gif.DecodeAll(data.Body)
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
	if frame <= 0 || frame > l {
		frame = l /2
	}

	p1 := image.NewPaletted(image.Rect(0, 0, all.Config.Width,all.Config.Height), palette.Plan9)
	draw.Draw(p1, p1.Bounds(), 	all.Image[frame-1], image.ZP, draw.Src) //添加图片

	op:= jpeg.Options{90}
	jpeg.Encode(jp,p1,&op)
	return filepath,err
}
