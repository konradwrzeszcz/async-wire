//go:build wireinject
// +build wireinject

package main

import (
	"fmt"
	"github.com/google/wire"
	"github.com/stephennancekivell/go-future/future"
	"time"
)

type MyFuture[T any] struct {
	F future.Future[T]
}

type GlobalBrandImage struct {
	URL string
}

func GetGlobalBrandImage() MyFuture[GlobalBrandImage] {
	f := future.New(func() GlobalBrandImage {
		time.Sleep(2 * time.Second)
		fmt.Println("hello from GetImage")
		return GlobalBrandImage{URL: "KFC image URL"}
	})

	return MyFuture[GlobalBrandImage]{F: f}
}

type GlobalBrandName string

func GetGlobalBrandName() MyFuture[GlobalBrandName] {
	f := future.New(func() GlobalBrandName {
		time.Sleep(2 * time.Second)
		fmt.Println("hello from GetGlobalBrandName")
		return GlobalBrandName("KFC")
	})
	return MyFuture[GlobalBrandName]{F: f}
}

type GlobalBrand struct {
	Name GlobalBrandName
	Img  GlobalBrandImage
}

func GetGlobalBrand(name MyFuture[GlobalBrandName], img MyFuture[GlobalBrandImage]) MyFuture[GlobalBrand] {
	f := future.New(func() GlobalBrand {
		time.Sleep(1 * time.Second)
		fmt.Println("hello from GetGlobalBrand")
		return GlobalBrand{
			Name: name.F.Get(),
			Img:  img.F.Get(),
		}
	})
	return MyFuture[GlobalBrand]{F: f}
}

type HomeScreen struct {
	Title GlobalBrandName
	Brand GlobalBrand
}

func NewHomeScreen(gb MyFuture[GlobalBrand], name MyFuture[GlobalBrandName]) HomeScreen {
	return HomeScreen{
		Title: name.F.Get(),
		Brand: gb.F.Get(),
	}
}

type MenuScreen struct {
	Icon  GlobalBrandImage
	Brand GlobalBrand
}

func NewMenuScreen(gb MyFuture[GlobalBrand], img MyFuture[GlobalBrandImage]) MenuScreen {
	return MenuScreen{
		Icon:  img.F.Get(),
		Brand: gb.F.Get(),
	}
}

func BuildHomeScreen() HomeScreen {
	wire.Build(GetGlobalBrandName, GetGlobalBrandImage, GetGlobalBrand, NewHomeScreen)
	return HomeScreen{}
}

func BuildMenuScreen() MenuScreen {
	wire.Build(GetGlobalBrandName, GetGlobalBrandImage, GetGlobalBrand, NewMenuScreen)
	return MenuScreen{}
}
