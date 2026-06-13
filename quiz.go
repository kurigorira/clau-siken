package main

import (
	"bufio"
	"fmt"
	"strings"
)

// normalize は答え合わせ用に入力を正規化する。
// 空白除去・小文字化・全角英数記号を半角へ変換する。
func normalize(s string) string {
	s = strings.TrimSpace(s)
	var b strings.Builder
	for _, r := range s {
		switch {
		case r >= '０' && r <= '９': // 全角数字 → 半角
			b.WriteRune(r - '０' + '0')
		case r >= 'Ａ' && r <= 'Ｚ': // 全角大文字 → 半角
			b.WriteRune(r - 'Ａ' + 'A')
		case r >= 'ａ' && r <= 'ｚ': // 全角小文字 → 半角
			b.WriteRune(r - 'ａ' + 'a')
		case r == '　' || r == ' ': // 空白は無視
		case r == '＾': // 全角ハット
			b.WriteRune('^')
		case r == '×': // かけ算記号ゆれ
			b.WriteRune('*')
		default:
			b.WriteRune(r)
		}
	}
	return strings.ToLower(b.String())
}

// isCorrect は入力 in が問題 p の正解かどうか判定する。
func isCorrect(p Problem, in string) bool {
	n := normalize(in)
	if n == "" {
		return false
	}
	accept := p.Accept
	if len(accept) == 0 {
		accept = []string{p.Answer}
	}
	for _, a := range accept {
		if normalize(a) == n {
			return true
		}
	}
	return false
}

// RunQuiz は対象教科のクイズを対話形式で実行する。filter が空なら全教科。
// reader は標準入力などの読み取り元（メニューと共有して使う）。
func RunQuiz(subjects []Subject, filter string, reader *bufio.Reader) {
	grandTotal, grandCorrect := 0, 0

	for _, s := range subjects {
		if filter != "" && s.Name != filter {
			continue
		}
		fmt.Printf("\n========== %s（%s実施） ==========\n", s.Name, s.TestDate.Format("1/2"))
		fmt.Printf("範囲: %s\n", s.Scope)
		correct := 0
		for i, p := range s.Problems {
			fmt.Printf("\n[第%d問] %s\n", i+1, p.Q)
			for j, c := range p.Choices {
				fmt.Printf("   %d) %s\n", j+1, c)
			}
			fmt.Print("答え > ")
			line, _ := reader.ReadString('\n')
			line = strings.TrimRight(line, "\r\n")
			if isCorrect(p, line) {
				fmt.Println("  ⭕ 正解！")
				correct++
			} else {
				fmt.Printf("  ❌ 不正解。正解: %s\n", p.Answer)
			}
			if p.Explain != "" {
				fmt.Printf("  💡 %s\n", p.Explain)
			}
		}
		fmt.Printf("\n>> %s の結果: %d/%d 問正解\n", s.Name, correct, len(s.Problems))
		grandTotal += len(s.Problems)
		grandCorrect += correct
	}

	if grandTotal == 0 {
		fmt.Printf("教科「%s」が見つかりませんでした。\n", filter)
		return
	}
	fmt.Printf("\n==================================================\n")
	fmt.Printf(" 総合結果: %d/%d 問正解（正答率 %.0f%%）\n", grandCorrect, grandTotal, float64(grandCorrect)/float64(grandTotal)*100)
	switch {
	case grandCorrect == grandTotal:
		fmt.Println(" 🎉 全問正解！この調子で本番もいける！")
	case float64(grandCorrect)/float64(grandTotal) >= 0.8:
		fmt.Println(" 👍 よくできています。間違えた問題だけもう一度！")
	default:
		fmt.Println(" 📚 間違えた問題は解説を読んで、明日また挑戦しよう。")
	}
}
