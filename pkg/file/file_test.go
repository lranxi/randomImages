package file

import (
	"fmt"
	"path"
	"testing"
)

func TestIsExists(t *testing.T) {
	path := "/Users/wenlys/code/images/api/pkg/file/file.go"
	fi, ok := IsExists(path)
	if ok {
		t.Log(fi.Name())
	}

}

func TestRename(t *testing.T) {
	fmt.Println(path.Join("abs", "1"))
}
