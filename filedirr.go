package filedirr

import (
	"bufio"
	"fmt"
	"go.k6.io/k6/js/modules"
	"os"
	"path/filepath"
)

func init() {
	modules.Register("k6/x/filedirr", new(FileDir))
}

type FileDir struct{}

func (*FileDir) HasDir(path string) bool {
	_, _err := os.Stat(path)
	if _err == nil {
		return true
	}
	if os.IsNotExist(_err) {
		return false
	}
	return false
}

func (*FileDir) CreateDir(path string) {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		fmt.Printf("创建目录异常 -> %v\n", err)
	} else {
		fmt.Println("创建成功!")
	}
}

func (*FileDir) FileDirs(path string) {
	var files []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println(file)
	}
}

func (*FileDir) ReadTxtFile(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}
