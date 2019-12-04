package main

import (
	"fmt"

	"./cw"
)

var keyword = []string{"蔡英文", "韓國瑜"}

func main() {
	cwSystem := cw.System{}

	PTTArticles := cwSystem.LoadPTT("./data/ptt.json")
	fmt.Printf("%+v\n", PTTArticles.Articles[0])

	FBArticles := cwSystem.LoadFB("./data/fb.json")
	fmt.Printf("%+v\n", FBArticles.Articles[0])

	cwSystem.CountCyberWarriors()
	fmt.Println(cwSystem)
}
