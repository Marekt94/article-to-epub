package gonejackconverter

import (
	"article-to-epub/pkg/misc"
	"article-to-epub/pkg/modules"
	"os"
	"path/filepath"

	_ "embed"

	"github.com/Marekt94/go-kernel-mt/logging"
	"github.com/gonejack/html-to-epub/html2epub"
)

//go:embed res/cover.jpg
var defaultCover []byte

const LOG_CONVERTER = `[CONVERTER] %q`
const TEMP_PATTERN = `art-*`

func NewGoneJackConverter() *HtmlToEpubConverter {
	return &HtmlToEpubConverter{Covercreator: &CoverCreator{}}
}

type HtmlToEpubConverter struct {
	Covercreator modules.CoverCreatorIntf
}

func CreateRandomOutputFileName(dir string, ext string) (string, error) {
	pattern := `temp-*.`
	if ext != "" {
		if ext[0] == '.' {
			pattern += ext[1:]
		} else {
			pattern += ext
		}
	}
	f, err := os.CreateTemp(dir, pattern)
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
	appDir, err := misc.GetAppDir()
	if err != nil {
		return nil, err
	}
	tempDir, err := os.MkdirTemp(appDir, TEMP_PATTERN)
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(tempDir)
	logging.Global.Debugf(LOG_CONVERTER, `tempDir = `+tempDir)

	tempFilePrefix := filepath.Join(tempDir, filepath.Base(tempDir))
	htmlFileName := tempFilePrefix + ".html"
	f, err := os.Create(htmlFileName)
	if err != nil {
		return nil, err
	}
	logging.Global.Debugf(LOG_CONVERTER, `tempFileName = `+htmlFileName)
	defer os.Remove(htmlFileName)

	_, err = f.Write(htmlContent)
	f.Close()
	if err != nil {
		return nil, err
	}

	output := tempFilePrefix + ".epub"
	logging.Global.Debugf(LOG_CONVERTER, `outputFileName = `+output)
	if err != nil {
		return nil, err
	}

	cover, err := c.Covercreator.CreateCover(defaultCover, title, authors)
	if err != nil {
		return nil, err
	}

	opt := html2epub.Options{
		Output:    output,
		HTML:      []string{htmlFileName},
		Verbose:   true,
		ImagesDir: tempDir,
	}
	cmd := html2epub.Cmd{
		Options:      opt,
		DefaultCover: cover,
	}
	err = cmd.Run()
	var out []byte = nil
	if err == nil {
		out, err = os.ReadFile(output)
	}

	return out, err
}
