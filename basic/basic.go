package main

import (
	"fmt"
	"time"
)

/**
 * 函数定义
 */
func calculate(a int, b int) (int, error) {
	// 计算参数相加的结果
	var sum = a + b
	// 修改参数的值
	b = 14
	// 返回函数调用结果
	return sum, nil
}

func main() {
	//////////////////////////////////// 变量及其定义 ////////////////////////////////////
	// 通过 var 关键字定义变量，该变量的类型为 int
	var v1 int
	// 给变量赋值为 1
	v1 = 1
	// 通过 := 语法，定义变量并赋值 2，变量的类型会在编译时通过值推断(即通过值 2 推断其类型为 int)。
	v2 := 2
	// 定义指针变量
	var v3 *int
	fmt.Printf("v1:(%v,%T), v2:(%v,%T), v3:(%v,%T)\n", v1, v1, v2, v2, v3, v3)

	// 定义变量后默认初始化
	var (
		i1 int
		i2 string
		i3 bool
		i4 *int
		i5 []int
		i6 [1]int
	)
	fmt.Printf("init value-> i1:(%v,%T), i2:(%v,%T), i3:(%v,%T), i4:(%v,%T), i5:(%v,%T), i6:(%v,%T)\n", i1, i1, i2, i2, i3, i3, i4, i4, i5, i5, i6, i6)

	//////////////////////////////////// 常量及其定义 ////////////////////////////////////
	// 定义常量并赋值
	const c1 int = 3
	// 定义2个常量并赋值
	const (
		c2           = 12.1
		c3 rune      = 3
		c4 complex64 = 12 + 3i
		c5 string    = "constValue"
	)
	fmt.Printf("c1:(%v,%T), c2:(%v,%T), c3:(%v,%T), c4:(%v,%T), c5:(%v,%T)\n", c1, c1, c2, c2, c3, c3, c4, c4, c5, c5)

	//////////////////////////////////// 运算符 ////////////////////////////////////
	var o1 = 1*4 + 2 - 10/2  // 数学运算，结果 = 1
	var o2 = "str" + "Value" // 字符串相加，结果 = strValue
	var o3 = 10 % 3          // 取余，结果 = 1
	var o4 = 1 << 2          // 移位，结果 = 4
	var o5 = 3 > 2           // 比较，结果 = true
	fmt.Printf("o1:(%v,%T), o2:(%v,%T), o3:(%v,%T), o4:(%v,%T), o5:(%v,%T)\n", o1, o1, o2, o2, o3, o3, o4, o4, o5, o5)

	//////////////////////////////////// 函数调用 ////////////////////////////////////
	p1 := 4
	f1, _ := calculate(3, p1)
	fmt.Printf("func call: (%v, %T), (%v, %T)\n", f1, f1, p1, p1)

	//////////////////////////////////// 常见语句 statements ////////////////////////////////////
	ifStatements()
	forStatements()
	switchStatements()

	//////////////////////////////////// 数据类型 data types ////////////////////////////////////
	dataTypes1()
	dataTypes2()

	//////////////////////////////////// defer语句 ////////////////////////////////////
	deferStatement()

	//////////////////////////////////// go语句 ////////////////////////////////////
	goStatement()

	//////////////////////////////////// select语句 ////////////////////////////////////
	selectStatement()

	//////////////////////////////////// break语句 ////////////////////////////////////
	breakStatement()

	//////////////////////////////////// break语句 ////////////////////////////////////
	continueStatement()

	//////////////////////////////////// break语句 ////////////////////////////////////
	gotoStatement()
}
func gotoStatement() {
Err:
	fmt.Println("occour error")

	for i := 0; i < 10; i++ {
		if i > 2 {
			goto Success
		}
		if i > 5 {
			goto Err
		}
		fmt.Println("for i = ", i)
	}

	fmt.Println("after for loop")

Success:
	fmt.Println("goto success")

}

func continueStatement() {
OuterLoop:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i > 2 || j > 2 {
				continue OuterLoop
			}
			fmt.Printf("continue for: %v, %v\n", i, j)
		}
	}
}

func breakStatement() {
OuterLoop:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i > 2 || j > 2 {
				break OuterLoop
			}
			fmt.Printf("break for: %v, %v\n", i, j)
		}

	}
}

func selectStatement() {
	c1 := make(chan int, 10)
	c2 := make(chan int, 10)
	go func(c1 chan int, c2 chan int) {
		time.Sleep(3 * time.Second)
		c1 <- 10
		c2 <- 20
	}(c1, c2)

	arr := []int{1, 2, 3, 4, 5, 6, 7}
	// 循环7次，等待从channel中读取数据
	for i, _ := range arr {
		select {
		case v := <-c1:
			fmt.Println("recevie data from c1: ", i, v)
			c2 <- 1
		case v := <-c2:
			c1 <- 2
			fmt.Println("recevie data from c2: ", i, v)
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("recevie default sleep")
		}
	}

	// c3 为nil，如果select语句块没有default的话，则会一直阻塞导致程序异常。
	var c3 chan int
	select {
	case i := <-c3:
		fmt.Println("receive data from c3:", i)
	default:
		fmt.Println("receive c3 default")
	}
}

func goStatement() {
	i := 1
	go func(a int) {
		time.Sleep(time.Second)
		fmt.Println("goroutine: ", a)
	}(i)
	i = 2
	fmt.Println("func run: ", i)
	time.Sleep(2 * time.Second)
}

func deferStatement() {
	i := 1
	defer fmt.Printf("defer 1, i value = %v\n", i)
	i = 2
	defer fmt.Printf("defer 2, i value = %v\n", i)
	i = 3
	fmt.Printf("i value is %v\n", i)
}

/**
 * 结构体定义
 */
type Dog struct {
	Name string
}

/**
 * 结构体定义
 */
type Duck struct {
	Name string
}

// 自定义接口，定义函数集合。如果某个结构体定义了该接口集合的所有函数，则可以认为其是该接口的实现
type Animal interface {
	Greet(msg string) error
	Step(distance int) error
}

// 实现 Step(distance int) error 函数
func (d Dog) Step(distance int) error {
	fmt.Printf("%v step %v\n", d.Name, distance)
	return nil
}

// 实现 Greet(message string) error 函数
func (d Dog) Greet(message string) error {
	fmt.Printf("%v greet %v !\n", d.Name, message)
	return nil
}

func dataTypes2() {
	// 自定义结构体类型及指针，其中tag `json:"name,omitempty"` 表示在进行json序列化/反序列时使用的json字段名称为name，omitempty表示是否忽略空值
	type Person struct {
		Name string `json:"name,omitempty"`
		Age  int    `json:"age,omitempty"`
	}
	var person Person = Person{Name: "Arloz", Age: 18}
	var personPtr *Person = &person
	fmt.Printf("structValues: person(%v,%T), personPtr(%v,%T)\n", person, person, personPtr, personPtr)

	// 定义接口类型的变量，可以赋值为实现了该接口的struct类型对象
	var an Animal = Dog{Name: "MyDog"}
	an.Greet("hello")
	an.Step(10)
	switch an.(type) {
	case Animal: // 匹配Animal类型，输出：case animal:
		fmt.Printf("case animal: %v\n", an)
	case Dog:
		fmt.Printf("case dog: %v\n", an)
	default:
		fmt.Printf("case default: %v\n", an)
	}

	// 接口类型转换，如果类型匹配，则ok为true，否则为false
	dog, ok := an.(Dog)
	fmt.Printf("interface type convert: (%v,%T), ok = %v\n", dog, dog, ok)
}

func dataTypes1() {
	// 基础类型
	var (
		v1 bool   = true
		v2 int64  = 32
		v3 string = "stringValue"
		v4 [4]int = [4]int{1, 2, 3, 4}
		v5 []int  = []int{1, 2, 3, 4, 5}
	)
	fmt.Printf("dataValues: v1(%v,%T), v2(%v,%T), v3(%v,%T), v4(%v,%T), v5(%v,%T)\n", v1, v1, v2, v2, v3, v3, v4, v4, v5, v5)

	// 定义key为string类型， value为int类型的map映射
	var m map[string]int = make(map[string]int, 0)
	m["key1"] = 1
	m["key2"] = 2
	m["key3"] = 3
	fmt.Printf("dataValues map: %v, %T\n", m, m)

	// 定义一个channel通过，内部元素的数据类型为int，通道的缓冲区大小为10
	var ch chan int = make(chan int, 10)
	// 往channle中添加一个元素
	ch <- 1
	// 从channle中读取一个元素
	c := <-ch
	fmt.Printf("dataValues channel: %v\n", c)
}

func switchStatements() {
	// 支持简单初始化语句执行: i := 2
	switch i := 2; i {
	default:
		fmt.Printf("default: %v\n", i)
	case 1, 2:
		fmt.Printf("case 1, 2: %v\n", i)
	case 3, 4:
		fmt.Printf("case 3, 4: %v\n", i)
	}

	// 无初始化语句
	x := "abc"
	switch x {
	case "abc":
		fmt.Printf("case abc: %v\n", x)
	default:
		fmt.Printf("case default: %v\n", x)
	}

	// 省略表达式语句，默认为 switch true {}
	switch {
	case true:
		fmt.Printf("case true: %v\n", x)
	case false:
		fmt.Printf("case false: %v\n", x)
	default:
		fmt.Printf("case default: %v\n", x)
	}

	// 基于类型的switch匹配, 变量y必须是 interface 类型
	var y interface{}
	y = 2
	switch y.(type) {
	case int:
		fmt.Printf("case int: %T\n", y)
	case string:
		fmt.Printf("case string: %T\n", y)
	default:
		fmt.Printf("case default:%T\n", y)
	}
}

func forStatements() {
	// 基础的for循环语句
	for i := 0; i < 5; i++ {
		fmt.Printf("for statements: i = %v\n", i)
	}

	// 只有一个条件表达式的for语句，类似于其他语言（JAVA）中的while语句
	x := 0
	for x < 5 {
		x += 1
	}
	fmt.Printf("for while x:%v\n", x)

	// 基于range的for循环语句
	arr := [...]int{1, 2, 3, 4, 5, 6}
	for idx, v := range arr {
		fmt.Printf("range array:%v = %v\n", idx, v)
	}

	sli := []int{1, 2, 3, 4, 5, 6}
	for idx, v := range sli {
		fmt.Printf("range slice:%v = %v\n", idx, v)
	}

	str := "rangeValue"
	for idx, v := range str {
		fmt.Printf("range string:%v = %v\n", idx, v)
	}

	m := make(map[string]int)
	m["key1"] = 1
	m["key2"] = 2
	m["key3"] = 3
	for k, v := range m {
		fmt.Printf("range map:%v = %v\n", k, v)
	}

	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	// 关闭通道，如果通道未关闭，则for range遍历时会一直等待不会结束
	close(ch)
	for c := range ch {
		fmt.Printf("range channel:%v\n", c)
	}

}

func ifStatements() {
	x := 5
	if x < 10 {
		fmt.Printf("simple if statements, %v\n", x)
	}

	if i := 5; i < 3 {
		fmt.Printf("if statements, %v\n", i)
	} else if i = 3; i > 4 {
		fmt.Printf("if else statements, %v\n", i)
	} else {
		fmt.Printf("else statements, %v\n", i)
	}
}
