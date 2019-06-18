package path

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func GetProjectPath(level int) (string, error) {
	var path string
	_, filename, _, ok := runtime.Caller(0)

	path = filename
	if ! ok {
		return "", fmt.Errorf("%v about '%v'","No caller information", filename)
	}

	for i := 0; i < level; i++ {
		path = filepath.Dir(path)
	}
	return path, nil
}