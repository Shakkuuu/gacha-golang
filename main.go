package main

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"

	"github.com/Shakkuuu/gacha-golang/gacha"
	"github.com/tenntenn/sqlite"
)

type TmpResults struct {
	DB      []*gacha.Card
	One     []string
	Msg     string
	Tickets int
	Coins   int
	Kaisu   int
	Rari    []string
}

var tmpl = template.Must(template.New("index").Parse(`<!DOCTYPE html>
<html>
	<head><title>ガチャ</title></head>
	<body>
		<p>{{.Msg}}</p>
		<p>チケット数:{{.Tickets}} コイン数:{{.Coins}} 引ける回数:{{.Kaisu}}
		<form action="/draw">
			<label for="num">回数</input>
			<input type="number" name="num" min="1" value="1">
			<input type="submit" value="ガチャを引く">
		</form>
		<h2>ガチャ結果</h2>
		<ol>{{range $o := .One}}
		<li>{{$o}}</li>
		{{end}}</ol>
		<h2>結果一覧</h2>
		<p>{{.Rari}}</p>
		<ol>{{range $d := .DB}}
		<li>{{$d}}</li>
		{{end}}</ol>
	</body>
</html>`))

var (
	flagCoin int
)

func init() {
	flag.IntVar(&flagCoin, "coin", 0, "コインの初期枚数")
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	flag.Parse()

	tickets, err := initialTickets()
	if err != nil {
		return err
	}

	db, err := sql.Open(sqlite.DriverName, "results.db")
	if err != nil {
		return fmt.Errorf("データベースのOpen:%w", err)
	}

	if err := createTable(db); err != nil {
		return err
	}

	// チケット数とコイン数の設定
	p := gacha.NewPlayer(tickets, flagCoin)

	play := gacha.NewPlay(p)

	// var numnum int
	var onere []string
	// var rere []string
	var msg string

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		results, err := getResults(db, 200)

		ti, co := p.Maisu()
		kai := p.DrawableNum()
		fmt.Printf("チケット:%d コイン:%d 引ける回数:%d \n", ti, co, kai)

		reamap := map[gacha.Rarity]int{}
		for _, reav := range results {
			reamap[reav.Rarity]++
		}
		var rea []string
		for rarity, count := range reamap {
			countStr := strconv.Itoa(count)
			rea = append(rea, rarity.String()+":"+countStr)
		}

		// if len(onere) > 0 {
		// 	lenlen := len(onere)
		// 	rere = onere[lenlen-numnum:]
		// 	fmt.Println(rere)
		// }

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		rr := TmpResults{
			DB:      results,
			One:     onere,
			Msg:     msg,
			Tickets: ti,
			Coins:   co,
			Kaisu:   kai,
			Rari:    rea,
		}

		if err := tmpl.Execute(w, rr); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		msg = ""
		onere = nil
	})

	http.HandleFunc("/draw", func(w http.ResponseWriter, r *http.Request) {
		num, err := strconv.Atoi(r.FormValue("num"))
		kai := p.DrawableNum()
		if kai == 0 {
			msg = "チケットあるいはコインがありません"
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		if num > kai {
			// fmt.Println("引ける回数を超えてます")
			msg = "引ける回数を超えてます"
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		// numnum = num
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		for i := 0; i < num; i++ {
			if !play.Draw() {
				if err := saveResult(db, play.Result()); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				onere = append(onere, play.Result().String())

				break
			}

			if err := saveResult(db, play.Result()); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			onere = append(onere, play.Result().String())
		}

		if err := play.Err(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusFound)
	})

	return http.ListenAndServe(":8080", nil)
}

func createTable(db *sql.DB) error {
	const sqlStr = `CREATE TABLE IF NOT EXISTS results(
		id        INTEGER PRIMARY KEY,
		rarity	  TEXT NOT NULL,
		name      TEXT NOT NULL
	);`

	_, err := db.Exec(sqlStr)
	if err != nil {
		return fmt.Errorf("テーブル作成:%w", err)
	}

	return nil
}

func saveResult(db *sql.DB, card *gacha.Card) error {
	const sqlStr = `INSERT INTO results(rarity, name) VALUES (?,?);`

	_, err := db.Exec(sqlStr, card.Rarity.String(), card.Name)
	if err != nil {
		return err
	}
	return nil
}

func getResults(db *sql.DB, limit int) ([]*gacha.Card, error) {
	const sqlStr = `SELECT rarity, name FROM results LIMIT ?`
	rows, err := db.Query(sqlStr, limit)
	if err != nil {
		return nil, fmt.Errorf("%qの実行:%w", sqlStr, err)
	}
	defer rows.Close()

	var results []*gacha.Card
	for rows.Next() {
		var card gacha.Card
		err := rows.Scan(&card.Rarity, &card.Name)
		if err != nil {
			return nil, fmt.Errorf("Scan:%w", err)
		}
		results = append(results, &card)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("結果の取得:%w", err)
	}

	return results, nil
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
