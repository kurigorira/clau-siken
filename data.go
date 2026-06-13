package main

import "time"

// Problem は1問の問題を表す。
//   - Choices が空ならば記述式（自由入力）、要素があれば選択式。
//   - Accept は正解として認める答えの一覧（表記ゆれ対応）。空の場合は Answer のみ。
type Problem struct {
	Q       string   // 問題文
	Choices []string // 選択肢（記述式なら空）
	Answer  string   // 模範解答（表示用）
	Accept  []string // 正解と判定する入力（正規化して比較）
	Explain string   // 解説
}

// Subject は1教科分のテスト範囲・提出物・問題集をまとめる。
type Subject struct {
	Name     string    // 教科名
	TestDate time.Time // テスト実施日
	Scope    string    // テスト範囲
	Submit   string    // 提出物・締切（なければ空）
	Advice   string    // 先生からのアドバイス要約
	Problems []Problem
}

// date はその年の月日から time.Time を作るヘルパー（テスト範囲は令和8年＝2026年）。
func date(month, day int) time.Time {
	return time.Date(2026, time.Month(month), day, 0, 0, 0, 0, time.Local)
}

// Subjects はテスト日順に並べた全教科データ。
// 問題はテスト範囲（教科書・ワークの単元）に沿った基礎〜標準レベルを中心に作成。
var Subjects = []Subject{
	{
		Name:     "社会",
		TestDate: date(6, 29),
		Scope:    "【歴史】教科書 P168〜239 / ワーク P58②〜P67,68① 【時事】4〜6月のニュース",
		Submit:   "ワーク提出：6/29(月)",
		Advice:   "重要語句＋歴史の流れ・地図・資料も確認。単元テストを見直す。時事は新聞/ニュースで補強。",
		Problems: []Problem{
			{Q: "第一次世界大戦のきっかけとなった、1914年の暗殺事件を何という？",
				Answer: "サラエボ事件", Accept: []string{"サラエボ事件", "サラエヴォ事件"},
				Explain: "オーストリア皇位継承者夫妻がサラエボで暗殺され、第一次世界大戦が始まった。"},
			{Q: "1919年に結ばれた、第一次世界大戦の講和条約は？",
				Answer: "ベルサイユ条約", Accept: []string{"ベルサイユ条約", "ヴェルサイユ条約"},
				Explain: "ドイツに巨額の賠償金と軍備制限を課した。"},
			{Q: "1929年にアメリカから始まり世界中に広がった不況を何という？",
				Answer: "世界恐慌", Accept: []string{"世界恐慌", "世界大恐慌"},
				Explain: "ニューヨークの株価大暴落をきっかけに各国へ波及した。"},
			{Q: "日本が満州事変への国際的批判を受け、1933年に脱退した国際機関は？",
				Answer: "国際連盟", Accept: []string{"国際連盟", "連盟"},
				Explain: "リットン調査団の報告を不服として脱退し、国際的に孤立していった。"},
			{Q: "1945年8月6日に世界で初めて原子爆弾が投下された都市は？",
				Choices: []string{"東京", "広島", "長崎", "大阪"},
				Answer:  "広島", Accept: []string{"広島", "2"},
				Explain: "8月6日広島、8月9日長崎に原子爆弾が投下された。"},
			{Q: "日本が無条件降伏を受け入れた、1945年の連合国の宣言は？",
				Answer: "ポツダム宣言", Accept: []string{"ポツダム宣言"},
				Explain: "1945年8月14日に受諾、15日に終戦を国民に伝えた（玉音放送）。"},
			{Q: "日本国憲法の三大原則のうち、戦争放棄を定めた条文は第何条？",
				Choices: []string{"第1条", "第9条", "第25条", "第99条"},
				Answer:  "第9条", Accept: []string{"第9条", "9条", "9", "2"},
				Explain: "第9条は戦争放棄・戦力不保持・交戦権の否認を定める。"},
			{Q: "1956年の日ソ共同宣言の後、日本が加盟を認められた国際機関は？",
				Answer: "国際連合", Accept: []string{"国際連合", "国連"},
				Explain: "ソ連の支持が得られ、1956年に国際連合へ加盟した。"},
		},
	},
	{
		Name:     "数学",
		TestDate: date(7, 3),
		Scope:    "教科書 P12〜P56 / ワーク P4〜P43 / 単元テスト「多項式」（展開・因数分解・根号を含む数）",
		Submit:   "",
		Advice:   "ワーク・教科書を繰り返し解き直す。展開/因数分解/根号の計算＋証明など応用にも挑戦。",
		Problems: []Problem{
			{Q: "(x+3)(x-5) を展開せよ",
				Answer: "x^2-2x-15", Accept: []string{"x^2-2x-15", "x²-2x-15"},
				Explain: "x²+(3-5)x+(3×-5)=x²-2x-15。"},
			{Q: "(2x+1)^2 を展開せよ",
				Answer: "4x^2+4x+1", Accept: []string{"4x^2+4x+1", "4x²+4x+1"},
				Explain: "(a+b)²=a²+2ab+b²。(2x)²+2·2x·1+1²=4x²+4x+1。"},
			{Q: "x^2+7x+12 を因数分解せよ",
				Answer: "(x+3)(x+4)", Accept: []string{"(x+3)(x+4)", "(x+4)(x+3)"},
				Explain: "かけて12・たして7 → 3と4。"},
			{Q: "x^2-9 を因数分解せよ",
				Answer: "(x+3)(x-3)", Accept: []string{"(x+3)(x-3)", "(x-3)(x+3)"},
				Explain: "a²-b²=(a+b)(a-b) の公式。"},
			{Q: "√12 を a√b の形に簡単にせよ",
				Answer: "2√3", Accept: []string{"2√3", "2ルート3"},
				Explain: "√12=√(4×3)=2√3。"},
			{Q: "√50 - √18 を計算せよ",
				Answer: "2√2", Accept: []string{"2√2", "2ルート2"},
				Explain: "5√2-3√2=2√2。"},
			{Q: "(√3+1)(√3-1) を計算せよ",
				Answer: "2", Accept: []string{"2"},
				Explain: "(a+b)(a-b)=a²-b² → 3-1=2。"},
			{Q: "6/√2 の分母を有理化せよ",
				Answer: "3√2", Accept: []string{"3√2", "3ルート2"},
				Explain: "6/√2 = 6√2/2 = 3√2。"},
		},
	},
	{
		Name:     "国語",
		TestDate: date(7, 6),
		Scope:    "『握手』『俳句の世界』『和語・漢語・外来語』ほか / 範囲の漢字 / これまでの小テスト",
		Submit:   "ワーク P6〜47",
		Advice:   "教科書を読み直し、授業・家庭学習プリントを復習。漢字の読み書きを確実に。",
		Problems: []Problem{
			{Q: "小説『握手』の作者は誰？",
				Answer: "井上ひさし", Accept: []string{"井上ひさし"},
				Explain: "ルロイ修道士との交流を描いた井上ひさしの作品。"},
			{Q: "俳句で、季節を表すために詠み込む言葉を何という？",
				Answer: "季語", Accept: []string{"季語"},
				Explain: "1つの俳句に季語は原則1つ（季重なりに注意）。"},
			{Q: "俳句の基本の音数（リズム）を数字で答えよ（例:5-7-5）",
				Answer: "5-7-5", Accept: []string{"5-7-5", "五七五", "575"},
				Explain: "五・七・五の十七音が基本。字余り・字足らずもある。"},
			{Q: "次のうち「外来語」はどれ？",
				Choices: []string{"山(やま)", "登山(とざん)", "ピアノ", "手紙"},
				Answer:  "ピアノ", Accept: []string{"ピアノ", "3"},
				Explain: "和語=訓読み中心、漢語=音読み中心、外来語=主に欧米から入った語（カタカナ）。"},
			{Q: "「山(やま)」のように訓読みで読む、日本固有の言葉の種類は？",
				Answer: "和語", Accept: []string{"和語", "やまとことば"},
				Explain: "和語は訓読み、漢語は音読みが基本。"},
			{Q: "漢字の読みを答えよ：「貢献」",
				Answer: "こうけん", Accept: []string{"こうけん"},
				Explain: "「貢献（こうけん）」＝役に立つように力をつくすこと。"},
			{Q: "漢字の読みを答えよ：「拍手」",
				Answer: "はくしゅ", Accept: []string{"はくしゅ"},
				Explain: "「拍手（はくしゅ）」。"},
			{Q: "ひらがなを漢字に直せ：「ちんもく」",
				Answer: "沈黙", Accept: []string{"沈黙"},
				Explain: "「沈黙（ちんもく）」＝だまっていること。"},
		},
	},
	{
		Name:     "英語",
		TestDate: date(7, 10),
		Scope:    "教科書 P8〜P37 / ワーク P4〜P37（P35を除く）",
		Submit:   "ワーク提出：7/10(金)（テスト終了後に回収）",
		Advice:   "ワークを何回も解き直す。対話文（AさんBさんの文）を確認する。",
		Problems: []Problem{
			{Q: "現在完了「すでに〜した」：I have ( ) finished my homework.",
				Choices: []string{"already", "yet", "ago", "since"},
				Answer:  "already", Accept: []string{"already", "1"},
				Explain: "肯定文の「すでに」は already。yet は疑問・否定で使う。"},
			{Q: "経験：Have you ever ( ) Kyoto?",
				Choices: []string{"visit", "visits", "visited", "visiting"},
				Answer:  "visited", Accept: []string{"visited", "3"},
				Explain: "現在完了は have/has + 過去分詞。visit-visited-visited。"},
			{Q: "継続：She ( ) lived here for ten years.",
				Choices: []string{"have", "has", "is", "had"},
				Answer:  "has", Accept: []string{"has", "2"},
				Explain: "主語 She（3人称単数）なので has。"},
			{Q: "関係代名詞：This is the book ( ) I bought yesterday.",
				Choices: []string{"who", "which", "what", "whose"},
				Answer:  "which", Accept: []string{"which", "2"},
				Explain: "先行詞が物（the book）で目的格なので which（that も可）。"},
			{Q: "受け身：This temple ( ) built 500 years ago.",
				Choices: []string{"is", "was", "has", "did"},
				Answer:  "was", Accept: []string{"was", "2"},
				Explain: "過去の受け身は was/were + 過去分詞。"},
			{Q: "現在完了の継続で「〜の間」を表す語は？（for / since のうち期間に使う方）",
				Answer: "for", Accept: []string{"for"},
				Explain: "for＋期間（for ten years）、since＋起点（since 2010）。"},
			{Q: "日本語にせよ：I have just arrived at the station.",
				Answer: "私はちょうど駅に着いたところだ", Accept: []string{"ちょうど駅に着いたところ", "ちょうど駅についたところ", "私はちょうど駅に着いたところだ"},
				Explain: "完了用法 have just + 過去分詞＝「ちょうど〜したところ」。"},
		},
	},
	{
		Name:     "理科",
		TestDate: date(7, 13),
		Scope:    "教科書 P6〜83（単元：運動とエネルギー）/ ワーク P2〜33",
		Submit:   "ワーク提出：7/1(水)（丸付け・訂正まで／教科連絡係が回収）",
		Advice:   "教科書・ノートを見直す。作図や計算問題は必ず出る。ワークを何回も解き直す。",
		Problems: []Problem{
			{Q: "一直線上を一定の速さで進む運動を何という？",
				Answer: "等速直線運動", Accept: []string{"等速直線運動"},
				Explain: "力がはたらかない（つり合う）とき、物体は等速直線運動を続ける。"},
			{Q: "物体がそのままの運動を続けようとする性質に関する法則を何という？",
				Answer: "慣性の法則", Accept: []string{"慣性の法則", "慣性"},
				Explain: "外から力がはたらかなければ、静止は静止、運動は等速直線運動を続ける。"},
			{Q: "仕事の単位は何？（記号で）",
				Choices: []string{"N", "J", "W", "Pa"},
				Answer:  "J", Accept: []string{"j", "ジュール", "2"},
				Explain: "仕事〔J〕＝力〔N〕×力の向きに動いた距離〔m〕。"},
			{Q: "100Nの物体を2m持ち上げたときの仕事は何J？",
				Answer: "200J", Accept: []string{"200", "200j"},
				Explain: "仕事＝100N×2m＝200J。"},
			{Q: "一定時間あたりにする仕事の量（仕事率）の単位は？（記号で）",
				Choices: []string{"N", "J", "W", "m/s"},
				Answer:  "W", Accept: []string{"w", "ワット", "3"},
				Explain: "仕事率〔W〕＝仕事〔J〕÷時間〔s〕。"},
			{Q: "高い位置にある物体がもつエネルギーを何という？",
				Answer: "位置エネルギー", Accept: []string{"位置エネルギー"},
				Explain: "高さが高いほど、質量が大きいほど大きい。"},
			{Q: "位置エネルギーと運動エネルギーの和を何という？",
				Answer: "力学的エネルギー", Accept: []string{"力学的エネルギー"},
				Explain: "摩擦などがなければ一定に保たれる（力学的エネルギー保存の法則）。"},
			{Q: "120kmを2時間で進んだときの平均の速さは時速何km？",
				Answer: "60km/h", Accept: []string{"60", "60km/h", "時速60km", "60km"},
				Explain: "速さ＝距離÷時間＝120÷2＝60km/h。"},
		},
	},
}
