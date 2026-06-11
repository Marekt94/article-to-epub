package gonejackconverter

import (
	"bytes"
	"errors"
	"image"
	"image/png"
	"os"
	"path/filepath"

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

	// _, thisFile, _, ok := runtime.Caller(0)
	// if !ok {
	// 	return nil, errors.New("runtime.Caller init error")
	// }
	thisFile, err := os.Executable()
	if err != nil {
		return nil, errors.New("cannot determine executable path: " + err.Error())
	}
	baseDir := filepath.Dir(thisFile)

	fontHeight := float64(ctx.Height()) / 20.0
	err = ctx.LoadFontFace(filepath.Join(baseDir, "res\\CrimsonPro-VariableFont_wght.ttf"), fontHeight)
	if err != nil {
		return nil, err
	}

	ax := float64(ctx.Width() / 2)
	marginY := float64(ctx.Height()) / 8.0
	marginX := 40.0

	ctx.SetHexColor("#000000")
	ctx.DrawStringWrapped(title, ax, marginY, 0.5, 0.5, float64(ctx.Width())-2*marginX, 1.5, gg.AlignCenter)
	ctx.DrawStringWrapped(author, ax, float64(ctx.Height())-marginY, 0.5, 0.5, float64(ctx.Width())-2*marginX, 1.5, gg.AlignCenter)

	var res bytes.Buffer
	imgRes := ctx.Image()
	err = png.Encode(&res, imgRes)
	if err != nil {
		return nil, err
	}

	return res.Bytes(), nil
}
