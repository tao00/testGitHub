package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"strings"
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
	// fmt.Println("abc")

	// t := Teacher{}
	// t.ShowA()
	// t.ShowB()
	// a := 1
	// b := 2
	// defer calc("1", a, calc("10", a, b))
	// a = 0
	// defer calc("2", a, calc("20", a, b))
	// b = 1

	// input := "askjavsaishdhafnkasocfhiasdohs"
	// fmt.Printf("%s\n", input)
	// countMinchar(input)

	// one := Node{4, nil, nil}
	// two := Node{5, nil, nil}
	// three := Node{6, nil, nil}
	// four := Node{7, nil, nil}

	// lastOne := Node{2, &one, &two}

	// lastTwo := Node{2, &three, &four}

	// aim := Node{20, &lastOne, &lastTwo}

	// res := aim.calcuAllValue()
	// fmt.Printf("所有叶子数是： %d\n", res)

}

func mission() {
	//创建一个有缓冲的通道
	aisle := make(chan int, 10)
	// aisle <- 1

	wg.Add(5)

	go produceNum("大佬", aisle)
	go takeNum("甲", aisle)
	go takeNum("乙", aisle)
	go takeNum("丙", aisle)
	go takeNum("丁", aisle)

	// for i := 0; i <= 10; i++ {
	// 	num := rand.Intn(100)
	// 	aisle <- num
	// }
	//关闭通道
	// close(aisle)

	//等待所有工作完成
	wg.Wait()
}

func produceNum(mom string, aisle chan int) {
	defer wg.Done()

	for i := 0; i <= 200; i++ {
		num := rand.Intn(100)
		aisle <- num

		fmt.Printf("%s 产出了数字 %d\n", mom, num)

		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
	}
	fmt.Printf("producer: %s : 工作完成\n", mom)
	//关闭通道
	close(aisle)
}

func takeNum(son string, aisle chan int) {
	defer wg.Done()
	for {
		num, ok := <-aisle
		if !ok {
			//这意味着通道已经空了，并且已经关闭
			fmt.Printf("worker: %s : 工作完成\n", son)
			return
		}
		fmt.Printf("%s 取到的随机数 %d\n", son, num)

		sleep := rand.Int63n(500)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
	}
}

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func inputContent() string {
	fmt.Println("请输入一个字符串：")
	var input string
	fmt.Scanf("%s", &input)
	return input
}

func countMinchar(input string) {
	for i := 'a'; i <= 'z'; i++ {
		x := string(i)
		count := strings.Count(input, x)
		if count != 0 {
			fmt.Printf("字母 %c 出现的次数: %d\n", i, count)
		}
	}
}

func sqrt(x int) int {
	i := 1
	for {
		num := i * i
		if x < num {
			i--
			return i
		}
		i++
	}
}

// type Node struct {
// 	value int   //叶子或节点对应的数值
// 	left  *Node //左子节点
// 	right *Node //右子节点
// }

// func (n *Node) calcuAllValue() int {
// 	allNum := 0
// 	//叶子节点
// 	if n.left == nil && n.right == nil {
// 		allNum = n.value
// 	}
// 	if n.left != nil {
// 		leftNum := n.left.calcuAllValue()
// 		allNum += leftNum
// 	}
// 	if n.right != nil {
// 		rightNum := n.right.calcuAllValue()
// 		allNum += rightNum
// 	}
// 	return allNum
// }

type DepositMSG struct {
	Name        string
	Addr        string
	Depositname string
	Amount      int
}

func testMarshal() []byte {
	depositmsg := DepositMSG{
		Name:        "root",
		Addr:        "0x006",
		Depositname: "test",
		Amount:      100,
	}

	data, err := json.Marshal(depositmsg)
	fmt.Println(data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func testUnmarshal(data []byte) {
	var depositmsg DepositMSG
	err := json.Unmarshal(data, &depositmsg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(depositmsg)
}

func testRead() []byte {
	fp, err := os.OpenFile("./data.json", os.O_RDONLY, 0755)
	defer fp.Close()
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 100)
	n, err := fp.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(data[:n]))
	return data[:n]
}

func testWrite(data []byte) {
	fp, err := os.OpenFile("data.json", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	_, err = fp.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}

// type User struct {
// 	Name   string
// 	Amount int64
// }

// func testMarshal() []byte {
// 	user := User{
// 		Name:   "root",
// 		Amount: 18,
// 	}
// 	data, err := json.Marshal(user)
// 	fmt.Println(data)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return data
// }

// func testUnmarshal(data []byte) {
// 	var user User
// 	err := json.Unmarshal(data, &user)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(user)
// }

type TxMessage struct {
	Tx           Transtion
	Time         int
	Currentnodes []Node
	Historynodes []Node
}

type Transtion struct {
	Txid       string
	Trans_type TranstionType
	Parentnode Node
	Status     Status
}

type Node struct {
	Nodeid int
}

type Status int

const (
	Pending Status = iota
	Fail
	Success
)

type TranstionType int

const (
	Transfer TranstionType = iota
	NewAccount
)
