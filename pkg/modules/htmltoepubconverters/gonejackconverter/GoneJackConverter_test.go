package gonejackconverter

import (
	"os"
	"strings"
	"testing"
)

func TestHtmlToEpubConverter(t *testing.T) {
	const tempDir = "..\\testdata"
	const tempFileName = "..\\testdata\\test.html"
	const outp = "..\\testdata\\temp.epub"

	b, err := os.ReadFile(tempFileName)
	if err != nil {
		t.Errorf(`No test file: %q`, err)
	}

	converter := &HtmlToEpubConverter{}
	out, err := converter.ConvertHtmlToEpub(b, outp, "Test Author")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if out == nil {
		t.Errorf("No file stream!")
	} else {
		os.Remove(outp)
	}
}

func TestCreateRandomOutputFileName(t *testing.T) {
	defExt := `.epub`
	out, err := CreateRandomOutputFileName(`.epub`)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !strings.HasSuffix(out, defExt) {
		t.Errorf(`Expected suffix %s in %s`, defExt, out)
	}
	_, err = os.ReadFile(out)
	if err == nil {
		t.Errorf(`Temp file exists but it should not`)
	}

	out, err = CreateRandomOutputFileName(`epub`)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !strings.HasSuffix(out, defExt) {
		t.Errorf(`Expected suffix %s in %s`, defExt, out)
	}
	_, err = os.ReadFile(out)
	if err == nil {
		t.Errorf(`Temp file exists but it should not`)
	}
}
