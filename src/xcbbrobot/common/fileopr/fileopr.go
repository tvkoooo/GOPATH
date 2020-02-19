package fileopr

import (
	"fmt"
	"os"
)

//创建文件夹
func CreateDir(dir *string) {
	//打开/新建今日文件夹
	exist, err := PathExists(dir)
	if err != nil {
		fmt.Printf("get dir error![%v]\n", err)
		os.Exit(1)
	}
	if exist {
		fmt.Printf("has dir![%v]\n", *dir)
	} else {
		// 创建路径文件夹
		err := os.Mkdir(*dir, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed![%v]\n", err)
			os.Exit(1)
		}
	}
}

// 判断文件夹是否存在
func PathExists(path *string) (bool, error) {
	_, err := os.Stat(*path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
