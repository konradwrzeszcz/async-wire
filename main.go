package main

import (
	"fmt"
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
