package htmltoepubconverter

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

type LogWriter struct {
}

type ErrLogWriter struct {
}

func (l *ErrLogWriter) Write(p []byte) (n int, err error) {
	message := string(p)
	if len(message) > 0 && message[len(message)-1] == '\n' {
		message = message[:len(message)-1]
	}
	logging.Global.Panicf(`[CONVERTER] %v`, message)
	return len(p), nil
}

func (l *LogWriter) Write(p []byte) (n int, err error) {
	message := string(p)
	if len(message) > 0 && message[len(message)-1] == '\n' {
		message = message[:len(message)-1]
	}
	logging.Global.Infof(`[CONVERTER] %v`, message)
	return len(p), nil
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

	converterExe := os.Getenv("EBOOK_CONVERTER_PATH")
	if converterExe == "" {
		converterExe = "ebook-convert.exe"
	}

	cmd := exec.Command(converterExe, filePathInt, filepath.Join(outputDir, outputFileName), cmdTitle, title, cmdAuthors, authors)
	logging.Global.Debugf(`cmd: %q`, cmd.Args)

	cmd.Stdout = &LogWriter{}
	cmd.Stderr = &ErrLogWriter{}

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
