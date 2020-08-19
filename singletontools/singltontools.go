package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	pathFlag := flag.String("path", "E:\\vip-singleton\\fork\\singleton\\test", "scan base path")
	topLine := flag.Int("topLine", 6, "matching lines content")
	oldYear := flag.String("oldYear", "2019", "matching copyright old year")
	newYear := flag.String("newYear", "2020", "repalce the copyright old year to current year")
	fileSuffix := flag.String("suffix", "java,go", "matching file suffix")
	flag.Parse()
	println(*pathFlag, *topLine, *oldYear, *newYear, *fileSuffix)
	fileSuffixArr := strings.Split(*fileSuffix, ",")
	updateCopyRigFile(*pathFlag, *topLine, *oldYear, *newYear, fileSuffixArr)
}

func listDir(dirPth string, fileSuffixArr []string) (files []string, err error) {
	//fmt.Println(dirPth)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	pthSep := string(os.PathSeparator)
	//    suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			tempFiles, _ := listDir(dirPth+pthSep+fi.Name(), fileSuffixArr)
			for _, tempFile := range tempFiles {
				files = append(files, tempFile)
			}
		} else {
			for _, sufStr := range fileSuffixArr {
				if strings.HasSuffix(fi.Name(), strings.TrimSpace(sufStr)) {
					files = append(files, dirPth+pthSep+fi.Name())
					break
				}
			}
		}
	}
	return files, nil
}

func updateCopyRigFile(path string, topLine int, oldYear string, newYear string, fileSuffixArr []string) {
	fmt.Println("begin read the Dir", path)
	fileNames, err := listDir(path, fileSuffixArr)
	if err != nil {
		panic(err)
	}
	for _, fileName := range fileNames {
		file, err1 := os.Open(fileName)
		if err1 != nil {
			panic(err1)
		}
		rd := bufio.NewReader(file)
		var buffer bytes.Buffer
		for i := 0; i < topLine; i++ {
			line, err2 := rd.ReadString('\n')
			if err2 != nil || io.EOF == err2 {
				break
			} else {
				buffer.WriteString(line)
			}
		}
		file.Close()
		bufferStr := buffer.String()
		matchLine := strings.Replace(bufferStr, oldYear, newYear, 1)
		if len(bufferStr) == len(matchLine) {
			file1, err3 := os.OpenFile(fileName, os.O_WRONLY, 0644)
			if err3 != nil {
				panic(err3)
			}
			defer file1.Close()
			file1.WriteString(matchLine)
			fmt.Println(file1.Name())
		}
	}
}
