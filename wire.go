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

type GlobalBrandImageProvider struct {
	*Provider[GlobalBrandImage]
}

func GetGlobalBrandImage() GlobalBrandImageProvider {
	return GlobalBrandImageProvider{
		NewProvider(func() GlobalBrandImage {
			time.Sleep(2 * time.Second)
			fmt.Println("hello from GetImage")
			return GlobalBrandImage{URL: "KFC image URL"}
		})}
}

type GlobalBrandName string
type GlobalBrandNameProvider struct {
	*Provider[GlobalBrandName]
}

func GetGlobalBrandName() GlobalBrandNameProvider {
	return GlobalBrandNameProvider{
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
type GlobalBrandProvider struct {
	*Provider[GlobalBrand]
}

func GetGlobalBrand(name GlobalBrandNameProvider, img GlobalBrandImageProvider) GlobalBrandProvider {
	return GlobalBrandProvider{
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

func NewHomeScreen(gb GlobalBrandProvider, name GlobalBrandNameProvider) HomeScreen {
	return HomeScreen{
		Title: name.Value(),
		Brand: gb.Value(),
	}
}

type MenuScreen struct {
	Icon  GlobalBrandImage
	Brand GlobalBrand
}

func NewMenuScreen(gb GlobalBrandProvider, img GlobalBrandImageProvider) MenuScreen {
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
