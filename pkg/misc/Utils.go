package misc

import (
	"errors"
	urlLib "net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const DEFAULT_FILE_NAME = "Default_file_name"
const MAX_FILE_NAME_LENGTH = 32

func AdaptUrlToFileName(url string) (string, error) {
	if len(strings.TrimSpace(url)) == 0 {
		return DEFAULT_FILE_NAME, nil
	}
	u, err := urlLib.Parse(url)
	if err != nil {
		return DEFAULT_FILE_NAME, err
	}

	res := path.Base(u.Path)

	if len(res) > 32 {
		return res[:MAX_FILE_NAME_LENGTH], nil
	} else {
		return res, nil
	}
}

func GetAppDir() (string, error) {
	appDir := ".//"
	var err error
	mode := os.Getenv("MODE")
	if strings.ToUpper(mode) == "RELEASE" {
		appDir, err = os.Executable()
		if err != nil {
			return "", errors.New("cannot determine executable path: " + err.Error())
		}
	}

	return filepath.Dir(appDir), nil
}
