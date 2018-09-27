package fileopr

import (
	"os"
	"fmt"
)

//创建文件夹
func CreateDir(dir *string)()  {
	//打开/新建今日文件夹
	exist, err := PathExists(dir)
	if err != nil {
		fmt.Printf("CreateDir::get dir error![%v]\n", err)
		os.Exit(1)
	}
	if exist {
		//logfile.SystemLogPrintf("info","Folder already exists![%v]\n", *dir)
	} else{
		// 创建路径文件夹
		err := os.Mkdir(*dir, os.ModePerm)
		if err != nil {
			fmt.Printf("CreateDir::mkdir failed![%v]\n", err)
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

func CheckFileSize(path *string,size int64) (int64) {
	fileInfo, err := os.Stat(*path)
	if err == nil {
		//文件大小
		fileSize:= fileInfo.Size()
		if fileSize>size {
			//如果超过最大值，给-1
			return -1
		}else {
			return fileSize
		}
	}else {
		//如果文件不存在，返回0值
		return 0
	}

}

func RenameFile(path *string,newName *string)(bool)  {
	//today := time.Now().Format("20060102")
	err :=os.Rename(*path,*newName)
	if err == nil {
		fmt.Println("file:",*path,"rename OK!")
		return true
	}else {
		fmt.Println("file:",*path,"rename false!")
		return false
	}

}