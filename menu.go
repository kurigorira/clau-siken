package main

import (
	"bufio"
	"fmt"
	"strings"
	"time"
)

// RunMenu は番号を選んで進める対話メニューを実行する。
// 各操作のあとメニューに戻り、0 で終了する。
func RunMenu(reader *bufio.Reader) {
	for {
		today := studyStart()
		printMenuHeader(today)
		fmt.Println("  1) 今日やること（スケジュール）")
		fmt.Println("  2) 問題集に挑戦")
		fmt.Println("  3) テスト範囲・提出物を見る")
		fmt.Println("  4) 点数アップのヒント")
		fmt.Println("  0) おわる")
		switch ask(reader, "えらんでね > ") {
		case "1":
			scheduleMenu(reader, today)
		case "2":
			quizMenu(reader)
		case "3":
			printScope()
			pause(reader)
		case "4":
			printTips()
			pause(reader)
		case "0", "q", "quit", "exit":
			fmt.Println("おつかれさま！また明日 👋")
			return
		default:
			fmt.Println("（0〜4 の番号を入れてね）")
		}
	}
}

// printMenuHeader はタイトルと「次のテストまであと何日」を表示する。
func printMenuHeader(today time.Time) {
	wd := weekdayJP[int(today.Weekday())]
	fmt.Println("\n==================================================")
	fmt.Println(" 才弥さん 期末テスト対策ツール")
	if name, days, ok := NextTest(Subjects, today); ok {
		switch {
		case days == 0:
			fmt.Printf(" 今日 %s(%s) ／ 本番：%s！\n", today.Format("1/2"), wd, name)
		default:
			fmt.Printf(" 今日 %s(%s) ／ 次のテスト：%s（あと%d日）\n", today.Format("1/2"), wd, name, days)
		}
	} else {
		fmt.Printf(" 今日 %s(%s) ／ テストはすべて終了！\n", today.Format("1/2"), wd)
	}
	fmt.Println("==================================================")
}

// scheduleMenu は今日の予定を見せ、希望すれば全日程も表示する。
func scheduleMenu(reader *bufio.Reader, today time.Time) {
	plans := GenerateSchedule(Subjects, studyStart())
	fmt.Println()
	PrintToday(plans, today)
	if yes(ask(reader, "\n全日程も見る？ (y/N) > ")) {
		PrintSchedule(plans)
	}
	pause(reader)
}

// quizMenu は教科を選んで問題集を実行する。
func quizMenu(reader *bufio.Reader) {
	fmt.Println("\n── どの教科の問題に挑戦する？ ──")
	for i, s := range Subjects {
		if _, days, ok := NextTest([]Subject{s}, studyStart()); ok {
			fmt.Printf("  %d) %s（%s・あと%d日）\n", i+1, s.Name, s.TestDate.Format("1/2"), days)
		} else {
			fmt.Printf("  %d) %s（%s・終了）\n", i+1, s.Name, s.TestDate.Format("1/2"))
		}
	}
	fmt.Printf("  %d) 全教科とおして挑戦\n", len(Subjects)+1)
	fmt.Println("  0) もどる")

	choice := ask(reader, "えらんでね > ")
	switch choice {
	case "0", "":
		return
	case fmt.Sprintf("%d", len(Subjects)+1), "全部", "全教科":
		RunQuiz(Subjects, "", reader)
	default:
		idx := atoiSafe(choice)
		if idx >= 1 && idx <= len(Subjects) {
			RunQuiz(Subjects, Subjects[idx-1].Name, reader)
		} else if s := findSubject(choice); s != "" {
			RunQuiz(Subjects, s, reader)
		} else {
			fmt.Println("（番号で選んでね）")
			return
		}
	}
	pause(reader)
}

// printTips は点数アップのヒントを要約表示する（詳細は docs を案内）。
func printTips() {
	fmt.Println("\n========== 点数アップのヒント ==========")
	tips := []string{
		"① ワークは2周＋間違い直し。これが一番点になる。",
		"② 提出物は締切前に終わらせる（社会6/29・理科7/1・英語7/10）。",
		"③ 間違えた問題はノートに集めて、前日に見返す。",
		"④ 25分集中＋5分休憩。スマホは別の部屋に。",
		"⑤ 数学・理科は途中式と単位を必ず書く（ミス防止）。",
		"⑥ 英語は対話文を音読して暗唱、不規則動詞は毎日5個。",
		"⑦ 国語は漢字で確実に得点（読み書き両方）。",
	}
	for _, t := range tips {
		fmt.Println("  " + t)
	}
	fmt.Println("\n  くわしくは docs/点数アップ施策.md を見てね。")
}

// ---- 入力まわりの小さなヘルパー ----

func ask(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}

func yes(s string) bool {
	s = strings.ToLower(strings.TrimSpace(s))
	return s == "y" || s == "yes" || s == "はい"
}

func pause(reader *bufio.Reader) {
	ask(reader, "\n[Enterでメニューに戻る] ")
}

func atoiSafe(s string) int {
	n := 0
	for _, r := range strings.TrimSpace(s) {
		if r < '0' || r > '9' {
			return -1
		}
		n = n*10 + int(r-'0')
	}
	if s == "" {
		return -1
	}
	return n
}

func findSubject(s string) string {
	for _, sub := range Subjects {
		if sub.Name == strings.TrimSpace(s) {
			return sub.Name
		}
	}
	return ""
}

// studyStart は計画の起点（今日。ただしテスト範囲の起点6/13より前なら6/13）。
func studyStart() time.Time {
	now := time.Now()
	if now.Before(date(6, 13)) {
		return date(6, 13)
	}
	return now
}
