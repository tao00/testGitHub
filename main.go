package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func studyone() {

	// fmt.Println("Hello, World!")
	//分配一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(2)

	//wg 用来等待程序完成
	//计数加2，表示要等待两个goroutine
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	//声明一个匿名函数 并创建一个goroutine
	go func() {
		//在函数退出时调用Done来通知main函数工作已经完成
		defer wg.Done()

		//显示字母表3次
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	go func() {
		//在函数退出时调用Done来通知main函数工作已经完成
		defer wg.Done()

		//显示字母表3次
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	//等待goroutine结束
	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}

var (
	counter int64
	wg      sync.WaitGroup
)

//这个实例程序展示如何使用atomic包来提供对值类型的安全访问
func studytwo() {
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()

	fmt.Println("Final counter:", counter)
}

func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		//安全地对counter加1
		atomic.AddInt64(&counter, 1)

		//当前 goroutine 从线程退出 并放回到队列
		runtime.Gosched()
	}
}

const (
	numberGoroutines = 4  //要使用的 goroutine的数量
	taskLoad         = 10 //要处理的工作的数量
)

//init初始化包，go语言运行时会在其他代码执行之前
//优先执行这个函数
func init() {
	//初始化随机种子
	rand.Seed(time.Now().Unix())
}

//这个示例展示如何使用有缓冲通道和固定数目的goroutine来处理一堆工作
func studythree() {
	//创建一个有缓冲通道来管理工作
	tasks := make(chan string, taskLoad)

	wg.Add(numberGoroutines)

	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	//增加一组要完成的工作
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task : %d", post)
	}

	//当所有工作都处理完时 关闭通道
	//以便所有goroutine退出
	close(tasks)

	//等待所有工作完成
	wg.Wait()
}

//worker 作为 goroutine启动来处理
//从有缓冲通道传入工作
func worker(tasks chan string, worker int) {
	//通知函数已经返回
	defer wg.Done()
	for {
		//等待分配工作
		task, ok := <-tasks
		if !ok {
			//这意味着通道已经空了，并且已经关闭
			fmt.Printf("worker: %d : shutting down\n", worker)
			return
		}
		//显示我们开始工作了
		fmt.Printf("worker: %d : started %s\n", worker, task)

		//随机等一段时间来模拟工作
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		//显示我们完成工作
		fmt.Printf("worker: %d : Completed %s\n", worker, task)
	}
}

func main() {
	fmt.Print("abc")
}
