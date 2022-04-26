package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	var s string
	var step = " "
	for i := 1; i < len(os.Args); i++ {
		s += step + os.Args[i]
	}
	//fmt.Println(s)
	//simpleServer()
	pointTest()

}

func forTest() {
	s, step := "", " " // 声明变量并赋值

	// 相当于Java里的 for (String arg : os.Args)
	// _ 是range产生的索引，不需要使用，用下划线代替
	for _, arg := range os.Args[1:] {
		s += step + arg
	}
	fmt.Println(s)
}

/**
输出字符出现行数 > 1的行
*/
func inputTest() {
	counts := make(map[string]int) // 定义一个Map
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		// 遍历每一行
		counts[input.Text()]++
	}
	// line 是key，n是value, counts相当于一个Map
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func simpleServer() {
	// 第一个参数为路径，第二个参数为处理函数
	http.HandleFunc("/", handler)
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func pointTest() {
	// 返回值是个指针
	var n = flag.Int("n", 1, "new")
	fmt.Printf("地址是%d", n)
	fmt.Printf("值是%d", *n)
}

/**
使用new来创建的变量类型是一个指针
*/
func newTest() {
	p := new(int)   // 默认值是0
	fmt.Println(*p) // 0
	*p = 2
	fmt.Println(*p) // 2
	var x = 1
	print(x)
}

/**
用var或者是 ：= 来声明变量
*/
func valueCreate() {
	var x = 1
	y := 1
	print(x)
	print(y)
}

/**
局部变量的逃逸
*/
var global *int // 声明一个int指针类型
func valueEscape() {
	var x int
	x = 1
	global = &x // 对x取地址，此时global就指向x的内存，所以x内存肯定是堆内存
}

/**
多重赋值
*/
/**
计算斐波那契数列
*/
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
