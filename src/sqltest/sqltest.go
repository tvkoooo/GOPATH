package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func main() {
	//取出参数查看
	getMainNum := 0
	for i, args := range os.Args {
		fmt.Printf("args[%d]=%s\n", i, args)
		getMainNum++
	}

	//打开数据库
	//DSN数据源字符串：用户名:密码@协议(地址:端口)/数据库?参数=参数值
	//db, err := sql.Open("mysql", "tvkoooo:dk2012@tcp(127.0.0.1:3306)/db_mm_cherry?charset=utf8")
	db, err := sql.Open("mysql", os.Args[1])
	if err != nil {
		fmt.Println(err)
	}

	//关闭数据库，db会被多个goroutine共享，可以不调用
	defer db.Close()

	//sql 命令数量
	fmt.Printf("\n\nsql 命令数量%d\n",getMainNum-2)

	//解析所有sql命令
	for jjj:=0;jjj<getMainNum-2 ;jjj++  {
		//查询数据，取所有字段 返回sql.Rows结果集
		//rows, err := db.Query("select `id`,`name` from db_mm_cherry.t_user_basic")
		errFlag := false
		rows2, err := db.Query(os.Args[2 + jjj])
		if err != nil {
			fmt.Printf("\nsql命令%d\tSQL错误:\n",jjj+1)
			fmt.Println(err)//输出错误
			errFlag = true
		}
		//如果有数据
		if !errFlag {
			//返回所有列
			cols, _ := rows2.Columns()
			//这里表示一行所有列的值，用[]byte表示
			vals := make([][]byte, len(cols))
			//这里表示一行填充数据
			scans := make([]interface{}, len(cols))
			//这里scans引用vals，把数据填充到[]byte里
			for k, _ := range vals {
				scans[k] = &vals[k]
			}

			i := 0
			result := make(map[int]map[string]string)
			for rows2.Next() {
				//填充数据
				rows2.Scan(scans...)
				//每行数据
				row := make(map[string]string)
				//把vals中的数据复制到row中
				for k, v := range vals {
					key := cols[k]
					//这里把[]byte数据转成string
					row[key] = string(v)
				}
				//放入结果集
				result[i] = row
				i++
			}
			fmt.Printf("\nsql命令%d\t结果集:\n",jjj+1)
			fmt.Println(result)
			rows2.Close()
		}
	}
}
