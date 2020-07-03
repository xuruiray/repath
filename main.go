package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const (
	goPkgPath = "/pkg/mod/"
	goSrcPath = "/src/"
)

// 入参为包地址 ./ufs
func main() {
	// 加载环境变量
	gopath := os.Getenv("GOPATH")

	// 加载 local 参数
	if len(os.Args) != 2 {
		fmt.Println("arguments error")
		os.Exit(0)
	}

	packagePath := os.Args[1]
	packagePath = strings.Trim(packagePath, " ")
	srcPath := gopath + goPkgPath + packagePath
	destPath := gopath + goSrcPath + packagePath + string(os.PathSeparator)

	node := strings.Split(srcPath, string(os.PathSeparator))
	parentPath := strings.Join(node[:len(node)-1], string(os.PathSeparator))
	packageName := node[len(node)-1]
	dir, err := ioutil.ReadDir(parentPath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	find := false
	for _, fi := range dir {
		if strings.HasPrefix(fi.Name(), packageName) && fi.IsDir() {
			srcPath = parentPath + string(os.PathSeparator) + fi.Name() + string(os.PathSeparator)
			find = true
			break
		}
	}

	if !find {
		fmt.Println("package not exist")
		os.Exit(0)
	}

	copyDir(srcPath, destPath)
}

// copyDir 文件路径复制
// copyDir("d:/tmp", "d:/tmp2")
// copyDir("/home/tmp", "/home/tmp2")
func copyDir(src string, dest string) {
	src = formatPath(src)
	dest = formatPath(dest)

	cmd := exec.Command("mkdir", "-p", dest)
	fmt.Println(cmd)
	outPut, err := cmd.Output()
	if err != nil {
		fmt.Printf("%v failed %v\n", cmd, err)
		return
	}
	fmt.Println(string(outPut))

	cmd = exec.Command("cp", "-R", src, dest)
	fmt.Println(cmd)
	outPut, err = cmd.Output()
	if err != nil {
		fmt.Printf("%v failed %v\n", cmd, err)
		return
	}
	fmt.Println(string(outPut))
}

// formatPath 格式化路径
func formatPath(path string) string {
	switch runtime.GOOS {
	case "darwin", "linux":
		return strings.Replace(path, "\\", "/", -1)
	default:
		fmt.Println("only support linux,darwin, but os is " + runtime.GOOS)
		return path
	}
}
