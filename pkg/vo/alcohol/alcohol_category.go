package alcohol

import "github.com/forroots/nomity-admin-api-v3/pkg/vo"

// AlcoholCategory はアルコールのカテゴリーを表す列挙VO
type AlcoholCategory string

// AlcoholCategory として許容される値
const (
	// AlcoholCategoryBeer はビール類
	AlcoholCategoryBeer AlcoholCategory = "BEER"
	// AlcoholCategoryWine はワイン・スパークリングワイン
	AlcoholCategoryWine AlcoholCategory = "WINE"
	// AlcoholCategorySour はサワー・チューハイ類
	AlcoholCategorySour AlcoholCategory = "SOUR"
	// AlcoholCategoryShochu は焼酎
	AlcoholCategoryShochu AlcoholCategory = "SHOCHU"
	// AlcoholCategorySake は日本酒
	AlcoholCategorySake AlcoholCategory = "SAKE"
	// AlcoholCategoryPlumWine は梅酒
	AlcoholCategoryPlumWine AlcoholCategory = "PLUM_WINE"
	// AlcoholCategorySpirits はスピリッツ類
	AlcoholCategorySpirits AlcoholCategory = "SPIRITS"
	// AlcoholCategoryCocktail はカクテル・リキュール類
	AlcoholCategoryCocktail AlcoholCategory = "COCKTAIL"
	// AlcoholCategoryWhiskey はウィスキー・ブランデー
	AlcoholCategoryWhiskey AlcoholCategory = "WHISKEY"
	// AlcoholCategoryOthers はその他
	AlcoholCategoryOthers AlcoholCategory = "OTHERS"
	// AlcoholCategoryNotDrink は飲まない
	AlcoholCategoryNotDrink AlcoholCategory = "NOT_DRINK"
)

// NewAlcoholCategory はAlcoholCategoryのコンストラクタ
func NewAlcoholCategory(value string) (AlcoholCategory, error) {
	category := AlcoholCategory(value)
	switch category {
	case AlcoholCategoryBeer,
		AlcoholCategoryWine,
		AlcoholCategorySour,
		AlcoholCategoryShochu,
		AlcoholCategorySake,
		AlcoholCategoryPlumWine,
		AlcoholCategorySpirits,
		AlcoholCategoryCocktail,
		AlcoholCategoryWhiskey,
		AlcoholCategoryOthers,
		AlcoholCategoryNotDrink:
		return category, nil
	default:
		return category, vo.NewVOErrorf("invalid AlcoholCategory value '%s'", value)
	}
}

// String はAlcoholCategory値オブジェクトをstringに変換する
func (c AlcoholCategory) String() string {
	return string(c)
}

// Display はAlcoholCategory値オブジェクトを表示用の文字列に変換する
func (c AlcoholCategory) Display() string {
	switch c {
	case AlcoholCategoryBeer:
		return "Beer"
	case AlcoholCategoryWine:
		return "Wine"
	case AlcoholCategorySour:
		return "Sour"
	case AlcoholCategoryShochu:
		return "Shochu"
	case AlcoholCategorySake:
		return "Sake"
	case AlcoholCategoryPlumWine:
		return "Plum Wine"
	case AlcoholCategorySpirits:
		return "Spirits"
	case AlcoholCategoryCocktail:
		return "Cocktail"
	case AlcoholCategoryWhiskey:
		return "Whiskey"
	case AlcoholCategoryOthers:
		return "Others"
	case AlcoholCategoryNotDrink:
		return "Not Drink"
	default:
		return "Unknown Value"
	}
}

// DisplayJa はAlcoholCategory値オブジェクトを日本語表示用の文字列に変換する
func (c AlcoholCategory) DisplayJa() string {
	switch c {
	case AlcoholCategoryBeer:
		return "ビール"
	case AlcoholCategoryWine:
		return "ワイン・スパークリングワイン"
	case AlcoholCategorySour:
		return "サワー・チューハイ"
	case AlcoholCategoryShochu:
		return "焼酎"
	case AlcoholCategorySake:
		return "日本酒"
	case AlcoholCategoryPlumWine:
		return "梅酒"
	case AlcoholCategorySpirits:
		return "スピリッツ"
	case AlcoholCategoryCocktail:
		return "カクテル・リキュール"
	case AlcoholCategoryWhiskey:
		return "ウィスキー・ブランデー"
	case AlcoholCategoryOthers:
		return "その他"
	case AlcoholCategoryNotDrink:
		return "飲まない"
	default:
		return "不明な値"
	}
}
