package filedir

import (
	"fmt"
	"go.k6.io/k6/js/modules"
	"os"
)

func init() {
	modules.Register("k6/x/filedir", new(FILEDIR))
}

type FILEDIR struct{}

func (*FILEDIR) HasDir(path string) bool {
	_, _err := os.Stat(path)
	if _err == nil {
		return true
	}
	if os.IsNotExist(_err) {
		return false
	}
	return false
}

func (*FILEDIR) CreateDir(path string) {
	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
		fmt.Printf("创建目录异常 -> %v\n", err)
	} else {
		fmt.Println("创建成功!")
	}
}
