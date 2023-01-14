package main

import (
	"fmt"

	"github.com/Shakkuuu/gacha-golang/gacha"
)

func main() {
	// section1()
	// section2()
	// section3()
	section4()
}

func section4() {
	p := gacha.NewPlayer(10, 100)

	n := inputN(p)
	results, summary := gacha.DrawN(p, n)

	fmt.Println(results)
	fmt.Println(summary)
}

func inputN(p *gacha.Player) int {
	max := p.DrawableNum()
	fmt.Printf("ガチャを引く回数を入力してください（最大:%d回）\n", max)

	var n int
	for {
		fmt.Print("ガチャを引く回数>")
		fmt.Scanln(&n)
		if 0 < n && n <= max {
			break
		}
		fmt.Printf("1以上%d以下の数を入力してください\n", max)
	}
	return n
}

// func section1() {
// 	// 乱数の種を設定する
// 	// 現在時刻をUNIX時間にしたものを種とする
// 	rand.Seed(time.Now().Unix())

// 	var n int
// 	fmt.Println("1: 単発 2:11連")

// LOOP:
// 	for {
// 		fmt.Print(">")
// 		var kind int
// 		fmt.Scanln(&kind)
// 		switch kind {
// 		case 1: // 単発ガチャ
// 			n = 1
// 			break LOOP
// 		case 2: // 11連ガチャ
// 			n = 11
// 			break LOOP
// 		default:
// 			fmt.Println("もう一度入力してください")
// 		}
// 	}

// 	for i := 1; i <= n; i++ {

// 		num := rand.Intn(100)

// 		fmt.Printf("%d回目 ", i)

// 		switch {
// 		case num < 80:
// 			fmt.Println("ノーマル")
// 		case num < 95:
// 			fmt.Println("R")
// 		case num < 99:
// 			fmt.Println("SR")
// 		default:
// 			fmt.Println("XR")
// 		}
// 	}
// }

// func section2() {
// 	slime := card{rarity: rarityN, name: "スライム"}
// 	fmt.Println(slime)

// 	dragon := card{rarity: raritySR, name: "ドラゴン"}
// 	fmt.Println(dragon)

// 	// 乱数の種を設定する
// 	// 現在時刻をUNIX時間にしたものを種とする
// 	rand.Seed(time.Now().Unix())

// 	var n int

// 	for {
// 		fmt.Print("何回引きますか?")
// 		fmt.Scanln(&n)

// 		if n > 0 {
// 			break
// 		}

// 		fmt.Println("もう一度入力してください")
// 	}

// 	result := map[string]int{}

// 	for i := 0; i < n; i++ {

// 		num := rand.Intn(100)

// 		// fmt.Printf("%d回目 ", i)

// 		switch {
// 		case num < 80:
// 			result["ノーマル"]++
// 		case num < 95:
// 			result["R"]++
// 		case num < 99:
// 			result["SR"]++
// 		default:
// 			result["XR"]++
// 		}
// 	}

// 	fmt.Println(result)
// }

// func section3() {
// 	// 乱数の種を設定する
// 	// 現在時刻をUNIX時間にしたものを種とする
// 	rand.Seed(time.Now().Unix())

// 	p := player{tickets: 10, coin: 100}

// 	n := inputN(&p)
// 	results, summary := drawN(&p, n)

// 	fmt.Println(results)
// 	fmt.Println(summary)
// }
