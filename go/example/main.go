package main

import "fmt"

var version string // ビルド時にldflagsフラグ経由でバージョンを埋め込むための変数

func main() {
	if version == "" {
		version = "dev" // バージョン未設定時に空文字を表示してしまう不具合を修正
	}
	fmt.Printf("Hello from Example %s!\n", version)
}
