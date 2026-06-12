package misc

import (
	"errors"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

func AdaptUrlToFileName(url string) string {
	re := regexp.MustCompile(`[^\pL]+`)
	url = strings.ReplaceAll(url, `\`, "")
	return re.ReplaceAllString(url, "_")
}

func GetAppDir() (string, error) {
	appDir := ""
	var err error
	mode := os.Getenv("MODE")
	if strings.ToUpper(mode) == "RELEASE" {
		appDir, err = os.Executable()
		if err != nil {
			return "", errors.New("cannot determine executable path: " + err.Error())
		}
	} else {
		var ok bool
		_, appDir, _, ok = runtime.Caller(0)
		if !ok {
			return "", errors.New("cannot determine executable path: runtime.Caller init error")
		}
	}

	return filepath.Dir(appDir), nil
}
