package main

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	//指针
	var a int = 20
	fmt.Println("修改前=", a)
	add(&a)
	fmt.Println("修改后=", a)

	var arr = []int{1, 2, 3, 4, 5}
	multiply(&arr)
	fmt.Println("修改后=", arr)

	//Goroutine
	go printNum(true, "奇数")
	go printNum(false, "偶数")
	time.Sleep(time.Second)

	var list []func() string = []func() string{
		func() string {
			fmt.Println("开始执行任务1")
			time.Sleep(time.Second)
			return "任务1"
		},
		func() string {
			fmt.Println("开始执行任务2")
			time.Sleep(time.Second * 2)
			return "任务2"
		},
		func() string {
			fmt.Println("开始执行任务3")
			time.Sleep(time.Second * 3)
			return "任务3"
		},
	}
	taskSchedule(list)

	//面向对象
	r := Rectangle{
		width:  10,
		length: 20,
	}
	fmt.Println("面积:", r.Area())
	fmt.Println("周长:", r.Perimeter())

	c := &Circle{
		radius: 5,
	}
	fmt.Println("面积:", c.Area())
	fmt.Println("周长:", c.Perimeter())

	e := Employee{
		EmployeeId: 1,
		Person: Person{
			Name: "张三",
			Age:  18,
		},
	}
	e.printInfo()

	//channel
	ch := make(chan int, 10)
	go func() {
		for value := range ch {
			fmt.Print(value, ",")
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()
	time.Sleep(time.Second * 2)
	fmt.Println()

	//锁机制 使用  sync.Mutex
	var mu sync.Mutex
	count := 0
	var wg sync.WaitGroup
	for i := range 10 {
		wg.Add(1)
		go func(i int) {
			mu.Lock()
			defer wg.Done()
			for _ = range 1000 {
				count = count + 1
			}
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	fmt.Println("使用互斥锁计数:", count)

	//sync/atomic
	var count1 int64
	var wg1 sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg1.Add(1)
		go func() {
			defer wg1.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&count1, 1) // 原子增加
			}

		}()
	}

	wg1.Wait()
	fmt.Println("无锁计数:", count1)
}

// 指针
func add(a *int) {
	*a += 10
}

func multiply(arr *[]int) {

	for i := range *arr {
		(*arr)[i] *= 2
	}
}

// Goroutine
func printNum(isOdd bool, name string) {
	for i := range 10 {
		m := i % 2
		if isOdd && m != 0 {
			fmt.Println(name+"协程", i)
		}
		if !isOdd && m == 0 {
			fmt.Println(name+"协程", i)
		}
	}
}

func taskSchedule(taskList []func() string) {
	for i := range taskList {
		go func() {
			startTime := time.Now()
			taskName := taskList[i]()
			needTime := time.Now().Sub(startTime).Seconds()
			fmt.Println(taskName+"耗时：", needTime)
		}()
	}
	time.Sleep(time.Second * 10)
}

// 面向对象
type Shape interface {
	Area() float64
	Perimeter() float64
}
type Rectangle struct {
	width  float64
	length float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.length
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.length + 2*r.width
}

type Circle struct {
	radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *Circle) Perimeter() float64 {
	return math.Pi * c.radius * 2
}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeId int
}

func (e Employee) printInfo() {
	fmt.Println("Name:", e.Name, "Age:", e.Age, "EmployeeId:", e.EmployeeId)
}
