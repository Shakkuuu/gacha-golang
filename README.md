# gacha-golang

## 使い方

1. localサーバを起動  
``` go run main.go -coin=コインの枚数 チケットの枚数 ```  
``` 例) go run main.go -coin=200 20 ```
2. ブラウザで localhost:8080 にアクセス
3. 回数を入力して「ガチャを引く」
4. ガチャ結果に今引いたガチャの結果が表示される
5. 結果一覧にこれまで引いた結果とレア度ごとの個数が表示される
6. チケットやコインがなくなり、引ける回数が0になったら起動し直しましょう

## 注意事項

* 結果はsqliteで保存されています
* results.dbというファイルを削除することで、結果一覧がリセットされます
* 結果一覧は200件までしか表示されません

## 開発環境

* macbook air M1
* Visual Studio Code
* go: version go1.19.5 darwin/arm64
* 使用外部パッケージ  
```github.com/tenntenn/sqlite, github.com/Shakkuuu/gacha-golang/gacha```
