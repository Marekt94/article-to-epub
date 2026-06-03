package calibreconverter

import (
	"testing"
)

func TestHtmlToEpubConverter(t *testing.T) {
	const tempDir = "..\\temp"
	const tempFileName = "temp.epub"
	converter := &HtmlToEpubConverter{}
	out, err := converter.HtmlToEpubConverterInternal([]byte("<html><body><h1>Test</h1></body></html>"), ".\\temp", "temp.epub", "", "")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if out == nil {
		t.Errorf("No file stream!")
	}
}
