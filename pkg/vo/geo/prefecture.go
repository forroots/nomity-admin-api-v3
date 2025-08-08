package geo

import "github.com/forroots/nomity-admin-api-v3/pkg/vo"

// Prefecture は都道府県を表す列挙VO
type Prefecture int

// 都道府県の英語表記をインデックス(0-based)で管理 (p-1) で参照
var prefectureNamesEn = []string{
	"Hokkaido",
	"Aomori",
	"Iwate",
	"Miyagi",
	"Akita",
	"Yamagata",
	"Fukushima",
	"Ibaraki",
	"Tochigi",
	"Gunma",
	"Saitama",
	"Chiba",
	"Tokyo",
	"Kanagawa",
	"Niigata",
	"Toyama",
	"Ishikawa",
	"Fukui",
	"Yamanashi",
	"Nagano",
	"Gifu",
	"Shizuoka",
	"Aichi",
	"Mie",
	"Shiga",
	"Kyoto",
	"Osaka",
	"Hyogo",
	"Nara",
	"Wakayama",
	"Tottori",
	"Shimane",
	"Okayama",
	"Hiroshima",
	"Yamaguchi",
	"Tokushima",
	"Kagawa",
	"Ehime",
	"Kochi",
	"Fukuoka",
	"Saga",
	"Nagasaki",
	"Kumamoto",
	"Oita",
	"Miyazaki",
	"Kagoshima",
	"Okinawa",
}

// 都道府県の日本語表記をインデックス(0-based)で管理 (p-1) で参照
var prefectureNamesJa = []string{
	"北海道",
	"青森",
	"岩手",
	"宮城",
	"秋田",
	"山形",
	"福島",
	"茨城",
	"栃木",
	"群馬",
	"埼玉",
	"千葉",
	"東京",
	"神奈川",
	"新潟",
	"富山",
	"石川",
	"福井",
	"山梨",
	"長野",
	"岐阜",
	"静岡",
	"愛知",
	"三重",
	"滋賀",
	"京都",
	"大阪",
	"兵庫",
	"奈良",
	"和歌山",
	"鳥取",
	"島根",
	"岡山",
	"広島",
	"山口",
	"徳島",
	"香川",
	"愛媛",
	"高知",
	"福岡",
	"佐賀",
	"長崎",
	"熊本",
	"大分",
	"宮崎",
	"鹿児島",
	"沖縄",
}

// NewPrefecture はPrefectureのコンストラクタ (1〜47が有効)
func NewPrefecture(num int) (Prefecture, error) {
	if num < 1 || num > 47 {
		return 0, vo.NewVOErrorf("invalid Prefecture value '%d'", num)
	}
	return Prefecture(num), nil
}

// Int はPrefecture値オブジェクトをintに変換する
func (p Prefecture) Int() int {
	return int(p)
}

// Display はPrefecture値オブジェクトを英語表記で返す
func (p Prefecture) Display() string {
	if p < 1 || p > 47 {
		return "Unknown Value"
	}
	return prefectureNamesEn[p-1]
}

// DisplayJa はPrefecture値オブジェクトを日本語表記で返す
func (p Prefecture) DisplayJa() string {
	if p < 1 || p > 47 {
		return "不明な値"
	}
	return prefectureNamesJa[p-1]
}

func GetAllPrefectureNamesEn() []string {
	return prefectureNamesEn
}

func GetAllPrefectureNamesJa() []string {
	return prefectureNamesJa
}
