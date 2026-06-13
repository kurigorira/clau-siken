// saiya-study は才弥さんの期末テスト対策用CLIツール。
//
// 引数なしで起動すると、番号を選んで進める対話メニューになる。
//
//	go run .                 対話メニュー（おすすめ）
//	go run . schedule        学習スケジュールを表示
//	go run . quiz [教科]      問題集に挑戦（教科を省くと全教科）
//	go run . scope           各教科のテスト範囲・提出物を表示
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	args := os.Args[1:]

	// 引数なし → 対話メニュー
	if len(args) == 0 {
		RunMenu(reader)
		return
	}

	switch args[0] {
	case "schedule", "s":
		PrintSchedule(GenerateSchedule(Subjects, studyStart()))
	case "quiz", "q":
		filter := ""
		if len(args) > 1 {
			filter = args[1]
		}
		RunQuiz(Subjects, filter, reader)
	case "scope":
		printScope()
	case "menu", "m":
		RunMenu(reader)
	case "serve", "web", "ui":
		Serve()
	default:
		printHelp()
	}
}

func printScope() {
	fmt.Println("========== Ⅰ学期期末テスト 範囲・提出物 ==========")
	for _, s := range Subjects {
		fmt.Printf("\n■ %s（%s 実施）\n", s.Name, s.TestDate.Format("1/2"))
		fmt.Printf("  範囲: %s\n", s.Scope)
		if s.Submit != "" {
			fmt.Printf("  提出: %s\n", s.Submit)
		}
		fmt.Printf("  助言: %s\n", s.Advice)
	}
}

func printHelp() {
	fmt.Println(`才弥さん 期末テスト対策ツール (saiya-study)

使い方:
  go run .              対話メニュー（番号を選ぶだけ・おすすめ）
  go run . serve        Web版（アイコンで操作する画面）をブラウザで開く
  go run . schedule     学習スケジュールを表示
  go run . scope        各教科のテスト範囲・提出物を表示
  go run . quiz [教科]   問題集に挑戦（社会/数学/国語/英語/理科。省略で全教科）

点数アップの作戦は docs/点数アップ施策.md を見てね。`)
}
