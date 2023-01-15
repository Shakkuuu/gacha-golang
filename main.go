package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/Shakkuuu/gacha-golang/gacha"
)

var (
	flagCoin    int
	flagResults string
	flagSummary string
)

var (
	regexpResults = regexp.MustCompile(`^results.*\.txt$`)
	regexpSummary = regexp.MustCompile(`^summary.*\.txt$`)
)

func init() {
	flag.IntVar(&flagCoin, "coin", 0, "コインの初期枚数")
	flag.StringVar(&flagResults, "results", "results.txt", "結果ファイルの名前")
	flag.StringVar(&flagSummary, "summary", "summary.txt", "集計ファイルの名前")
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// section1()
	// section2()
	// section3()
	// section4()
}

func run() error {
	flag.Parse()

	if !regexpResults.MatchString(flagResults) {
		return fmt.Errorf("結果ファイル名が不正(%s)", flagResults)
	}

	if !regexpSummary.MatchString(flagSummary) {
		return fmt.Errorf("集計ファイル名が不正(%s)", flagSummary)
	}

	tickets, err := initialTickets()
	if err != nil {
		return err
	}

	p := gacha.NewPlayer(tickets, flagCoin)
	play := gacha.NewPlay(p)

	n := inputN(p)
	for play.Draw() {
		if n <= 0 {
			break
		}
		fmt.Println(play.Result())
		n--
	}

	if err != nil {
		return fmt.Errorf("ガチャを%d回引く:%w", n, err)
	}

	if err := saveResults(play.Results()); err != nil {
		return err
	}

	if err := saveSummary(play.Summary()); err != nil {
		return err
	}

	return nil
}

func initialTickets() (int, error) {
	if flag.NArg() == 0 {
		return 0, errors.New("ガチャチケットの枚数を入力してください")
	}

	num, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		return 0, fmt.Errorf("ガチャチケット数のパース(%q):%w", flag.Arg(0), err)
	}

	return num, nil
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

func saveResults(results []*gacha.Card) (rerr error) {
	f, err := os.Create(flagResults)
	if err != nil {
		return fmt.Errorf("%sの作成:%w", flagResults, err)
	}

	defer func() {
		if err := f.Close(); err != nil && rerr == nil {
			rerr = fmt.Errorf("%sのクローズ:%w", flagResults, err)
		}
	}()

	for _, result := range results {
		fmt.Fprintln(f, result)
	}

	return nil
}

func saveSummary(summary map[gacha.Rarity]int) (rerr error) {
	f, err := os.Create(flagSummary)
	if err != nil {
		return fmt.Errorf("%sの作成:%w", flagSummary, err)
	}

	defer func() {
		if err := f.Close(); err != nil && rerr == nil {
			rerr = fmt.Errorf("%sのクローズ:%w", flagSummary, err)
		}
	}()

	for rarity, count := range summary {
		fmt.Fprintf(f, "%s %d\n", rarity.String(), count)
	}

	return nil
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

// func section4() {
// 	p := gacha.NewPlayer(10, 100)

// 	n := inputN(p)
// 	results, summary := gacha.DrawN(p, n)

// 	fmt.Println(results)
// 	fmt.Println(summary)
// }
