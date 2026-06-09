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

	b, err := os.ReadFile(filepath.Join(baseDir, ".\\res\\cover.jpg"))
	if err != nil {
		t.Error(err)
	}

	res, err := coverCreateor.CreateCover(b, "Lorem ipsum dolor sit amet, consectetur", "Lorem ipsum dolor sit amet, consectetur")

	if err != nil {
		t.Error(err)
	}

	if res == nil {
		t.Error("res is nil")
	} else {
		os.WriteFile(filepath.Join(baseDir, "cover-test.jpg"), res, 0644)
	}
}
