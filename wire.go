//go:build wireinject
// +build wireinject

package main

import (
	"fmt"
	"github.com/google/wire"
	"time"
)

type GlobalBrandImage struct {
	URL string
}

type GlobalBrandImageFuture struct {
	Provider[GlobalBrandImage]
}

func GetGlobalBrandImage() GlobalBrandImageFuture {
	return GlobalBrandImageFuture{
		NewProvider(func() GlobalBrandImage {
			time.Sleep(2 * time.Second)
			fmt.Println("hello from GetImage")
			return GlobalBrandImage{URL: "KFC image URL"}
		})}
}

type GlobalBrandName string
type GlobalBrandNameFuture struct {
	Provider[GlobalBrandName]
}

func GetGlobalBrandName() GlobalBrandNameFuture {
	return GlobalBrandNameFuture{
		NewProvider(func() GlobalBrandName {
			time.Sleep(2 * time.Second)
			fmt.Println("hello from GetGlobalBrandName")
			return GlobalBrandName("KFC")
		})}
}

type GlobalBrand struct {
	Name GlobalBrandName
	Img  GlobalBrandImage
}
type GlobalBrandFuture struct {
	Provider[GlobalBrand]
}

func GetGlobalBrand(name GlobalBrandNameFuture, img GlobalBrandImageFuture) GlobalBrandFuture {
	return GlobalBrandFuture{
		NewProvider(func() GlobalBrand {
			time.Sleep(1 * time.Second)
			fmt.Println("hello from GetGlobalBrand")
			return GlobalBrand{
				Name: name.Value(),
				Img:  img.Value(),
			}
		})}
}

type HomeScreen struct {
	Title GlobalBrandName
	Brand GlobalBrand
}

func NewHomeScreen(gb GlobalBrandFuture, name GlobalBrandNameFuture) HomeScreen {
	return HomeScreen{
		Title: name.Value(),
		Brand: gb.Value(),
	}
}

type MenuScreen struct {
	Icon  GlobalBrandImage
	Brand GlobalBrand
}

func NewMenuScreen(gb GlobalBrandFuture, img GlobalBrandImageFuture) MenuScreen {
	return MenuScreen{
		Icon:  img.Value(),
		Brand: gb.Value(),
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
