// 这个示例程序展示如何声明,并使用方法
package main

import (
	"fmt"
)

// user 在程序里定义一个用户类型
type user struct {
	name  string
	email string
}

// notify 使用值接收者实现了一个方法
func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n",
		u.name,
		u.email)
}

// change Email 使用指针接收者实现了一个方法
func (u *user) changeEmail(email string) {
	u.email = email
}

// main 是应用程序的入口
func main() {
	//Go 语言里有两种类型的接收者：值接收者和指针接收者
	//notify 方法的接收者被声明为user类型的值。如果使用值接收者声明方法，调用时会使用这个值的一个副本来执行
	// user 类型的值可以用来调用
	// 使用值接收者声明的方法
	bill := user{"Bill", "bill@email.com"}
	bill.notify()
	//使用这个指针变量来调用notify方法。为了支持这种方法调用，Go语言调整了指针的值，来符合方法接收者的定义,
	//go进行后边的动作 修改成(*lisa).notify()，注意，其实notify还是操作的是一份副本
	// 指向user类型值的指针也可以用来调用
	// 使用值接收者声明的方法
	lisa := &user{"Lisa", "lisa@email.com"}
	lisa.notify()
	// user类型的值可以用来调用
	// 使用指针接收者声明的方法
	//Go语言再一次对值做了调整，使之符合函数的接收者，进行调用，(&bill).change Email ("bill@newdomain.com")
	bill.changeEmail("bill@newdomain.com")
	bill.notify()
	//指向user类型值的指针可以用来调用
	// 使用指针接收者声明的方法
	//一旦change Email 调用返回，这个调用对值做的修改也会反映在 lisa指针所指向的值上。
	// 这是因为 change Email 方法使用了指针接收者。
	// 总结一下，值接收者使用值的副本来调用方法，而指针接受者使用实际值来调用方法。
	lisa.changeEmail("lisa@newdomain.com")
	lisa.notify()
}
