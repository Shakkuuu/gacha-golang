package main

import (
	// TODO: fmtパッケージをインポートする

	"fmt"
	"math/rand"
	"time"
)

func main() {

	// 乱数の種を設定する
	// 現在時刻をUNIX時間にしたものを種とする
	rand.Seed(time.Now().Unix())

	var n int
	fmt.Println("1: 単発 2:11連")

LOOP:
	for {
		fmt.Print(">")
		var kind int
		fmt.Scanln(&kind)
		switch kind {
		case 1: // 単発ガチャ
			n = 1
			break LOOP
		case 2: // 11連ガチャ
			n = 11
			break LOOP
		default:
			fmt.Println("もう一度入力してください")
		}
	}

	for i := 1; i <= n; i++ {

		num := rand.Intn(100)

		fmt.Printf("%d回目 ", i)

		switch {
		case num < 80:
			fmt.Println("ノーマル")
		case num < 95:
			fmt.Println("R")
		case num < 99:
			fmt.Println("SR")
		default:
			fmt.Println("XR")
		}
	}
}
