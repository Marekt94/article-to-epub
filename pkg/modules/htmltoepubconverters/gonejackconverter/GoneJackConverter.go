package gonejackconverter

import (
	"os"

	_ "embed"

	"github.com/Marekt94/go-kernel-mt/logging"
	"github.com/gonejack/html-to-epub/html2epub"
)

//go:embed res/cover.png
var defaultCover []byte

const logConverter = `[CONVERTER] %q`

type HtmlToEpubConverter struct {
}

func CreateRandomOutputFileName(ext string) (string, error) {
	pattern := `temp-*.`
	if ext != "" {
		if ext[0] == '.' {
			pattern += ext[1:]
		} else {
			pattern += ext
		}
	}
	f, err := os.CreateTemp("", pattern)
	if err != nil {
		return "", err
	} else {
		name := f.Name()
		defer os.Remove(f.Name())
		defer f.Close()
		return name, nil
	}
}

func (c *HtmlToEpubConverter) ConvertHtmlToEpub(htmlContent []byte, title string, authors string) ([]byte, error) {
	f, err := os.CreateTemp("", "art-*.html")
	if err != nil {
		return nil, err
	}
	tempFileName := f.Name()
	logging.Global.Debugf(logConverter, `tempFileName = `+tempFileName)
	defer os.Remove(tempFileName)

	_, err = f.Write(htmlContent)
	f.Close()
	if err != nil {
		return nil, err
	}

	output, err := CreateRandomOutputFileName(`.epub`)
	if err != nil {
		return nil, err
	}
	opt := html2epub.Options{
		Output:  output,
		HTML:    []string{tempFileName},
		Verbose: true,
	}
	cmd := html2epub.Cmd{
		Options:      opt,
		DefaultCover: defaultCover,
	}
	err = cmd.Run()
	var out []byte = nil
	if err == nil {
		out, err = os.ReadFile(output)
	}

	return out, err
}
