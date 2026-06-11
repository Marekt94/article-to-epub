package gonejackconverter

import (
	"os"
	"strings"
	"testing"
)

func TestHtmlToEpubConverter(t *testing.T) {
	const tempFileName = "..\\testdata\\test.html"

	b, err := os.ReadFile(tempFileName)
	if err != nil {
		t.Errorf(`No test file: %q`, err)
	}

	converter := NewGoneJackConverter()
	out, err := converter.ConvertHtmlToEpub(b, "Test title", "Test Author")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if out == nil {
		t.Errorf("No file stream!")
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
