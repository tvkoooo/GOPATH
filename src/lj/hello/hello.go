/*use the test
see this laji*/
package main

import (
	"fmt"
	"sync"
	_"time"
)
//this the main
func main(){
	var j int = 5
	var waitGroup sync.WaitGroup //我们应该充分利用waitGroup 来控制go 的执行生命周期
	waitGroup.Add(2)//我们会启动两个go，也可以每次启动go前，用add(1),在go里面去done来减一

	//注意：使用go来分布式执行，你不确定是谁先结束这个go，如果不对这个go做生命周期维护，
	// 很可能出现多个go执行异步，多个go执行的先后顺序和你想的相差甚远，甚至出现主程序结束，其中某些go还未执行完毕，go的逻辑丢失
	go func(pwaitGroup *sync.WaitGroup,j int) {
		fmt.Println(" 1 j=",j)
		//各协程间通信，go建议是用通道channel来通信，主程序和go也是不同的协程
		//需要特别注意，这里的waitGroup，必须是外部控制逻辑的的指针，需要操作的是外部的控制。
		//如果外部传递进来的是拷贝值，go里面的只会拷贝一份副本，因此做done不会对外部控制器减一
		//测试情况下传递进来是外部的&waitGroup（就是他的指针）发现，waitGroup应该用(*pwaitGroup)来实现，但是(pwaitGroup)也没有异常？
		//如果是同函数内，其实也可以不用通过指针传递进来，可以直接使用函数的全局变量处理，因为go提供是原子操作（add/done）都是非嵌入式的（锁安全的）
		//但是由于go和主程序是分开的协程，如果是对同一个数据做处理会出现竞态，因此该数据是危险的，需要加锁
		(*pwaitGroup).Done()//实现控制器减一

	}(&waitGroup,j)

	j = 6
	go func(waitGroup *sync.WaitGroup,j int) {
		fmt.Println(" 2 j=",j)
		(*waitGroup).Done()
	}(&waitGroup,j)
	waitGroup.Wait()//等待所有go函数执行结束，退出



	var k int = 5
	a := func() func() {
		var i int = 10 //初始化只会执行一次
		//里面函数只会执行一次
		func(){
			fmt.Println(" 3 i=",i,"k = ",k)
		}()
		i -=2 //表达式只会执行一次
		k -=2 //表达式只会执行一次
		//表达式只会执行一次
		fmt.Println(" 3 i=",i,"k = ",k)
		//return是函数逻辑结束，退回执行
		return func() {
			//返回逻辑执行期间，改变数据
			var ooo int =110//a()每次都会初始化，因此ooo每次都一样，这个会保持不变
			k +=3//a()每次都会原来基础上加3，注意，由于k是外部变量，因此k会被外部修改
			i +=3//a()每次都会原来基础上加3
			fmt.Println(" 4 i=",i,"k = ",k,"ooo = ",ooo)
		}
	}() //末尾的括号表明匿名函数被调用，并将返回的函数指针赋给变量a
	//末尾的括号和func()是相对应的，程序执行时，会在出现a()前，优先执行return函数之前所有表达式，类似初始化了该函数所有功能，并运行结束
	//到return后（或者理解为执行完毕，结束return 进行函数输出，但是本次不执行return里面的内容）跳出执行到末尾的(),把指针交给a，
	// 在遇到第一个调用a()，开始执行return内部函数（闭环数据不释放），执行完毕后跳转到末尾的(),把指针交给a，
	//同理，如果还有下个a()，继续执行return内部函数（闭环数据不释放），执行完毕后末尾的(),把指针交给a，直到最后所有结束退出

	//末尾的括号和func()是相对应的，而func()是和return是对应的，可以理解为，下次出现a后面调用的(),就是执行的return这个逻辑
	a()
	k *= 2
	a()

	bbb := func(){
		var kkkkkk int =11111
		kkkkkk +=3
		fmt.Println(" 555555",kkkkkk)
	}
	//本次匿名函数和上面的区别在于没有后面的(),因此每次只有在调用bbb()才执行所有里面的功能
	bbb()
	bbb()

	var str string
	fmt.Scanln(&str)
	fmt.Printf("INPUT :%s\n", str)
}