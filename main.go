package main

import (
	"fmt"
	"github.com/stephennancekivell/go-future/future"
	"time"
)

func main() {
	//one call
	home := BuildHomeScreen()
	fmt.Println("home title", home.Title)
	fmt.Println("home brand name", home.Brand.Name)
	fmt.Println("home brand image url", home.Brand.Img.URL)

	//next call
	fmt.Println()

	menu := BuildMenuScreen()
	fmt.Println("menu icon url", menu.Icon.URL)
	fmt.Println("menu brand name", menu.Brand.Name)
	fmt.Println("menu brand image url", menu.Brand.Img.URL)
}

func main2() {
	f := future.New(func() string {
		fmt.Println("before sleep")
		time.Sleep(2 * time.Second)
		fmt.Println("after sleep")
		return "value"
	})

	fmt.Println("before get")
	value := f.Get()
	fmt.Println("after get")
	fmt.Println(value)

	value2 := f.Get()
	fmt.Println(value2)
}
