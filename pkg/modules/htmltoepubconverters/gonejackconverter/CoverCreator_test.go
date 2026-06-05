package gonejackconverter

import (
	"article-to-epub/pkg/modules"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestGenerateCover(t *testing.T) {
	t.Helper()

	var coverCreateor modules.CoverCreatorIntf = &CoverCreator{}
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller failed")
	}
	baseDir := filepath.Dir(thisFile)

	b, err := os.ReadFile(filepath.Join(baseDir, ".\\res\\cover.png"))
	if err != nil {
		t.Error(err)
	}

	res, err := coverCreateor.CreateCover(b, "Test Title example", "Name Surname-Surname")

	if err != nil {
		t.Error(err)
	}

	if res == nil {
		t.Error("res is nil")
	} else {
		os.WriteFile(filepath.Join(baseDir, "cover-test.png"), res, 0644)
	}
}
