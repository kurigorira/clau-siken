// saiya-study は才弥さんの期末テスト対策用CLIツール。
//
// 使い方:
//
//	go run . schedule        学習スケジュールを表示
//	go run . quiz            全教科の問題集に挑戦
//	go run . quiz 数学        教科を指定して問題集に挑戦
//	go run . scope           各教科のテスト範囲・提出物を表示
//	go run .                 ヘルプを表示
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	args := os.Args[1:]
	cmd := "help"
	if len(args) > 0 {
		cmd = args[0]
	}

	switch cmd {
	case "schedule", "s":
		// 今日(またはテスト範囲の起点)から計画を生成。
		start := time.Now()
		if start.Before(date(6, 13)) {
			start = date(6, 13)
		}
		PrintSchedule(GenerateSchedule(Subjects, start))
	case "quiz", "q":
		filter := ""
		if len(args) > 1 {
			filter = args[1]
		}
		RunQuiz(Subjects, filter)
	case "scope":
		printScope()
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
  go run . schedule     学習スケジュールを表示
  go run . scope        各教科のテスト範囲・提出物を表示
  go run . quiz         全教科の問題集に挑戦
  go run . quiz 数学     教科を指定して挑戦（社会/数学/国語/英語/理科）

点数アップの作戦は docs/点数アップ施策.md を見てね。`)
}
