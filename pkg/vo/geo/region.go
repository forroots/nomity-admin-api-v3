package geo

import "github.com/forroots/nomity-admin-api-v3/pkg/vo"

// Region は日本を8地域に区分した上位階層の列挙VO (1〜8)
type Region int

const (
	RegionUnknown  Region = 0
	RegionHokkaido Region = 1
	RegionTohoku   Region = 2
	RegionKanto    Region = 3
	RegionChubu    Region = 4
	RegionKinki    Region = 5
	RegionChugoku  Region = 6
	RegionShikoku  Region = 7
	RegionKyushu   Region = 8
)

// regionNamesEn と regionNamesJa は英語/日本語の地域名を配列などで管理 (index=region-1)
var regionNamesEn = []string{
	"Hokkaido",
	"Tohoku",
	"Kanto",
	"Chubu",
	"Kinki",
	"Chugoku",
	"Shikoku",
	"Kyushu",
}

var regionNamesJa = []string{
	"北海道",
	"東北",
	"関東",
	"中部",
	"近畿",
	"中国",
	"四国",
	"九州",
}

// NewRegion は Region のコンストラクタ。1〜8を許容
func NewRegion(code int) (Region, error) {
	if code < 1 || code > 8 {
		return RegionUnknown, vo.NewVOErrorf("invalid Region code '%d'", code)
	}
	return Region(code), nil
}

func (r Region) Int() int {
	return int(r)
}

func (r Region) Display() string {
	if r < 1 || r > 8 {
		return "Unknown Value"
	}
	return regionNamesEn[r-1]
}

func (r Region) DisplayJa() string {
	if r < 1 || r > 8 {
		return "不明な値"
	}
	return regionNamesJa[r-1]
}

// regionToPrefectures は各地域に属する都道府県一覧をマップで保持
// Prefecture は 1〜47 を想定
var regionToPrefectures = map[Region][]Prefecture{
	RegionHokkaido: {
		mustPrefecture(1), // 北海道
	},
	RegionTohoku: {
		mustPrefecture(2), // 青森
		mustPrefecture(3), // 岩手
		mustPrefecture(4), // 宮城
		mustPrefecture(5), // 秋田
		mustPrefecture(6), // 山形
		mustPrefecture(7), // 福島
	},
	RegionKanto: {
		mustPrefecture(8),  // 茨城
		mustPrefecture(9),  // 栃木
		mustPrefecture(10), // 群馬
		mustPrefecture(11), // 埼玉
		mustPrefecture(12), // 千葉
		mustPrefecture(13), // 東京
		mustPrefecture(14), // 神奈川
	},
	RegionChubu: {
		mustPrefecture(15), // 新潟
		mustPrefecture(16), // 富山
		mustPrefecture(17), // 石川
		mustPrefecture(18), // 福井
		mustPrefecture(19), // 山梨
		mustPrefecture(20), // 長野
		mustPrefecture(21), // 岐阜
		mustPrefecture(22), // 静岡
		mustPrefecture(23), // 愛知
	},
	RegionKinki: {
		mustPrefecture(24), // 三重
		mustPrefecture(25), // 滋賀
		mustPrefecture(26), // 京都
		mustPrefecture(27), // 大阪
		mustPrefecture(28), // 兵庫
		mustPrefecture(29), // 奈良
		mustPrefecture(30), // 和歌山
	},
	RegionChugoku: {
		mustPrefecture(31), // 鳥取
		mustPrefecture(32), // 島根
		mustPrefecture(33), // 岡山
		mustPrefecture(34), // 広島
		mustPrefecture(35), // 山口
	},
	RegionShikoku: {
		mustPrefecture(36), // 徳島
		mustPrefecture(37), // 香川
		mustPrefecture(38), // 愛媛
		mustPrefecture(39), // 高知
	},
	RegionKyushu: {
		mustPrefecture(40), // 福岡
		mustPrefecture(41), // 佐賀
		mustPrefecture(42), // 長崎
		mustPrefecture(43), // 熊本
		mustPrefecture(44), // 大分
		mustPrefecture(45), // 宮崎
		mustPrefecture(46), // 鹿児島
		mustPrefecture(47), // 沖縄
	},
}

// mustPrefecture はコンストラクタのエラー処理を省略するためのユーティリティ
func mustPrefecture(num int) Prefecture {
	p, err := NewPrefecture(num)
	if err != nil {
		panic(err) // データが壊れている想定はないのでpanicでOK
	}
	return p
}

// Prefectures は指定した Region に属する都道府県の一覧を返す (英語/日本語はPrefecture内メソッドで取得)
func (r Region) Prefectures() []Prefecture {
	if prefs, ok := regionToPrefectures[r]; ok {
		return prefs
	}
	return nil // or empty slice
}

// RegionOfPrefecture は都道府県から対応するRegionを返す（Lookup）
func RegionOfPrefecture(p Prefecture) Region {
	for region, prefs := range regionToPrefectures {
		for _, pf := range prefs {
			if pf == p {
				return region
			}
		}
	}
	return RegionUnknown
}
