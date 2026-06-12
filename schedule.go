package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

// Task は1日の中の1つの学習タスク。
type Task struct {
	Subject string
	Minutes int
	Detail  string
}

// DayPlan は1日分の学習計画。
type DayPlan struct {
	Date      time.Time
	Tasks     []Task
	Reminders []string // ワーク提出などの注意
	TestToday string   // その日がテスト本番ならば教科名
}

var weekdayJP = []string{"日", "月", "火", "水", "木", "金", "土"}

// phaseDetail はテストまでの残り日数に応じた学習内容を返す。
func phaseDetail(s Subject, daysUntil int) string {
	switch {
	case daysUntil <= 2:
		return "総仕上げ：暗記の最終確認＋単元テスト/ワークの間違い直し"
	case daysUntil <= 6:
		return "問題演習：ワークを2周目、間違えた問題を中心に解き直し"
	default:
		return "基礎固め：教科書を読み直し＋ワークを1周（範囲: " + shorten(s.Scope) + "）"
	}
}

func shorten(s string) string {
	if i := strings.Index(s, " /"); i > 0 {
		return s[:i]
	}
	r := []rune(s)
	if len(r) > 24 {
		return string(r[:24]) + "…"
	}
	return s
}

// reminderFor はその日に出すべき提出物リマインドを返す。
// 提出日の3日前から当日まで通知する。
func reminderFor(day time.Time) []string {
	var out []string
	type due struct {
		date time.Time
		text string
	}
	dues := []due{
		{date(7, 1), "理科ワーク提出（7/1 水・丸付け＆訂正まで）"},
		{date(6, 29), "社会ワーク提出（6/29 月）"},
		{date(7, 10), "英語ワーク提出（7/10 金・テスト後回収）"},
	}
	for _, d := range dues {
		diff := int(d.date.Sub(day).Hours() / 24)
		if diff >= 0 && diff <= 3 {
			if diff == 0 {
				out = append(out, "★本日提出: "+d.text)
			} else {
				out = append(out, fmt.Sprintf("提出まであと%d日: %s", diff, d.text))
			}
		}
	}
	return out
}

// GenerateSchedule は start から最終テスト日までの日別計画を作る。
// 平日は60分、土日は120分を基本に、直近のテスト教科へ重点配分する。
func GenerateSchedule(subjects []Subject, start time.Time) []DayPlan {
	subs := append([]Subject(nil), subjects...)
	sort.Slice(subs, func(i, j int) bool { return subs[i].TestDate.Before(subs[j].TestDate) })

	last := subs[len(subs)-1].TestDate
	var plans []DayPlan

	for d := start; !d.After(last); d = d.AddDate(0, 0, 1) {
		plan := DayPlan{Date: d, Reminders: reminderFor(d)}

		// 本日テストがあるか
		for _, s := range subs {
			if sameDay(s.TestDate, d) {
				plan.TestToday = s.Name
			}
		}

		// まだ受けていない教科（テスト日が今日以降）を抽出
		var upcoming []Subject
		for _, s := range subs {
			if !s.TestDate.Before(d) {
				upcoming = append(upcoming, s)
			}
		}
		if len(upcoming) == 0 {
			plans = append(plans, plan)
			continue
		}

		budget := 60
		if wd := d.Weekday(); wd == time.Saturday || wd == time.Sunday {
			budget = 120
		}

		next := upcoming[0]
		daysUntil := int(next.TestDate.Sub(d).Hours() / 24)

		// テスト前日・当日朝は直近教科に全集中
		if daysUntil <= 2 {
			plan.Tasks = append(plan.Tasks, Task{next.Name, budget, phaseDetail(next, daysUntil)})
		} else if len(upcoming) >= 2 {
			// 直近教科に多め、2番目の教科に残りを配分（先取り）
			primary := budget * 2 / 3
			second := upcoming[1]
			plan.Tasks = append(plan.Tasks,
				Task{next.Name, primary, phaseDetail(next, daysUntil)},
				Task{second.Name, budget - primary, phaseDetail(second, int(second.TestDate.Sub(d).Hours()/24))},
			)
		} else {
			plan.Tasks = append(plan.Tasks, Task{next.Name, budget, phaseDetail(next, daysUntil)})
		}

		plans = append(plans, plan)
	}
	return plans
}

func sameDay(a, b time.Time) bool {
	return a.Year() == b.Year() && a.YearDay() == b.YearDay()
}

// PrintSchedule は計画を見やすく出力する。
func PrintSchedule(plans []DayPlan) {
	fmt.Println("==================================================")
	fmt.Println(" 才弥さん Ⅰ学期期末テスト 学習スケジュール")
	fmt.Println("==================================================")
	totalMin := 0
	for _, p := range plans {
		wd := weekdayJP[int(p.Date.Weekday())]
		fmt.Printf("\n【%s(%s)】\n", p.Date.Format("1/2"), wd)
		if p.TestToday != "" {
			fmt.Printf("  ★★ テスト本番：%s ★★\n", p.TestToday)
		}
		for _, t := range p.Tasks {
			totalMin += t.Minutes
			fmt.Printf("  ・%s %d分 … %s\n", t.Subject, t.Minutes, t.Detail)
		}
		for _, r := range p.Reminders {
			fmt.Printf("  ⚠ %s\n", r)
		}
	}
	fmt.Printf("\n--------------------------------------------------\n")
	fmt.Printf(" 合計学習時間の目安：約%.1f時間（1日30分〜2時間ペース）\n", float64(totalMin)/60)
	fmt.Printf(" ヒント: `go run . quiz` で問題集に挑戦、`go run . quiz 数学` で教科指定。\n")
}
