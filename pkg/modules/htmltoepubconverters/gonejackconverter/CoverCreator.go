package gonejackconverter

import (
	"bytes"
	"errors"
	"image"
	"image/png"
	"path/filepath"
	"runtime"

	"github.com/fogleman/gg"
)

type CoverCreator struct {
}

func (c *CoverCreator) CreateCover(canvas []byte, title string, author string) ([]byte, error) {
	img, _, err := image.Decode(bytes.NewReader(canvas))
	if err != nil {
		return nil, err
	}
	ctx := gg.NewContextForImage(img)

	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		return nil, errors.New("runtime.Caller init error")
	}
	baseDir := filepath.Dir(thisFile)

	err = ctx.LoadFontFace(filepath.Join(baseDir, "res\\arial.ttf"), 120)
	if err != nil {
		return nil, err
	}

	ax := float64(ctx.Width() / 2)
	margin := 250.0

	ctx.SetHexColor("#000000")
	ctx.DrawStringAnchored(title, ax, margin, 0.5, 0.5)
	ctx.DrawStringAnchored(author, ax, float64(ctx.Height())-margin, 0.5, 0.5)

	var res bytes.Buffer
	imgRes := ctx.Image()
	err = png.Encode(&res, imgRes)
	if err != nil {
		return nil, err
	}

	return res.Bytes(), nil
}
