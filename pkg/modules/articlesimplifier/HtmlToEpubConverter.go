package articlesimplifier

import (
	"os"
	"os/exec"
)

const tempHtmlDir = ".\\temp"
const tempHtmlPath = tempHtmlDir + "\\temp.html"
const outputEpubPath = tempHtmlDir + "\\temp.epub"

type HtmlToEpubConverter struct {
}

func (c *HtmlToEpubConverter) ConvertHtmlToEpub(html []byte) ([]byte, error) {
	if err := os.Mkdir(tempHtmlDir, 0755); err != nil && !os.IsExist(err) {
		return nil, err
	}
	if err := os.WriteFile(tempHtmlPath, html, 0644); err != nil {
		return nil, err
	}

	cmd := exec.Command("ebook-convert.exe", tempHtmlPath, outputEpubPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return nil, err
	}

	var out []byte
	var err error
	if out, err = os.ReadFile(outputEpubPath); err != nil {
		return nil, err
	}

	return out, nil
}
