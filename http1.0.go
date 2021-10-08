// go build hoge.goで簡単にネイティブコードができる。
// コンパイラ言語は初めて使ったのですが、想像以上に簡単にコンパイルできます！

package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// local:8000でウェブサーバーを立てている想定。
	// respはhttp.Responseオブジェクト
	resp, _ := http.Get("http://localhost:8000")

	// deferで関数を抜けた後に通信を閉じる
	defer resp.Body.Close()

	// Responseオブジェクトをバイト配列に変換
	body, _ := ioutil.ReadAll(resp.Body)

	log.Println(resp.Header)
	log.Println(string(body))
}
