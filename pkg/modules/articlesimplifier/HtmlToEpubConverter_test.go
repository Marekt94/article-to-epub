package articlesimplifier

import (
	"testing"
)

func TestHtmlToEpubConverter(t *testing.T) {
	converter := &HtmlToEpubConverter{}
	_, err := converter.ConvertHtmlToEpub([]byte("<html><body><h1>Test</h1></body></html>"))
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
