package articlesimplifier

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/Marekt94/go-kernel-mt/logging"
)

const tempHtmlDir = ".\\temp"
const tempHtmlFileName = "temp.html"
const outputEpubFileName = "temp.epub"

const cmdTitle = "--title"
const cmdAuthors = "--authors"

type HtmlToEpubConverter struct {
}

func (c *HtmlToEpubConverter) HtmlToEpubConverterInternal(html []byte, outputDir string, outputFileName string, title string, authors string) ([]byte, error) {
	if err := os.Mkdir(outputDir, 0755); err != nil && !os.IsExist(err) {
		return nil, err
	}
	filePathInt := filepath.Join(outputDir, tempHtmlFileName)
	if err := os.WriteFile(filePathInt, html, 0644); err != nil {
		return nil, err
	}
	defer os.Remove(filePathInt)

	cmd := exec.Command("ebook-convert.exe", filePathInt, filepath.Join(outputDir, outputFileName), cmdTitle, title, cmdAuthors, authors)
	logging.Global.Debugf(`cmd: %q`, cmd.Args)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return nil, err
	}

	var out []byte
	var err error
	filePathInt = filepath.Join(outputDir, outputFileName)
	if out, err = os.ReadFile(filePathInt); err != nil {
		return nil, err
	}
	defer os.Remove(filePathInt)

	return out, nil
}

func (c *HtmlToEpubConverter) ConvertHtmlToEpub(html []byte, title string, authors string) ([]byte, error) {
	epub, err := c.HtmlToEpubConverterInternal(html, tempHtmlDir, outputEpubFileName, title, authors)
	return epub, err
}
