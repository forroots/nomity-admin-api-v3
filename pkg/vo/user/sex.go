package user

import "github.com/forroots/nomity-admin-api-v3/pkg/vo"

// Sex は性別を表す列挙VO
type Sex int

// Sex として許容される値
const (
	// SexUnanswered は未回答であることを表す
	SexUnanswered = 0
	// SexMale は男性であることを表す
	SexMale = 1
	// SexFemale は女性であることを表す
	SexFemale = 2
	// SexDonotWantToSay は回答したくないことを表す
	SexDonotWantToSay = 9
)

// NewSex はSexのコンストラクタ
func NewSex(number int) (Sex, error) {
	sex := Sex(number)
	switch sex {
	case SexUnanswered, SexMale, SexFemale, SexDonotWantToSay:
		return sex, nil
	default:
		return sex, vo.NewVOErrorf("invalid Sex value '%d'", number)
	}
}

// Int はSex値オブジェクトをintに変換する
func (s Sex) Int() int {
	return int(s)
}

// Display はSex値オブジェクトを表示用の文字列に変換する
func (s Sex) Display() string {
	switch s {
	case SexUnanswered:
		return "Unanswered"
	case SexMale:
		return "Male"
	case SexFemale:
		return "Female"
	case SexDonotWantToSay:
		return "Don't want to say"
	default:
		return "Unknown Value"
	}
}

// DisplayJa はSex値オブジェクトを日本語の表示用の文字列に変換する
func (s Sex) DisplayJa() string {
	switch s {
	case SexUnanswered:
		return "未回答"
	case SexMale:
		return "男性"
	case SexFemale:
		return "女性"
	case SexDonotWantToSay:
		return "回答したくない"
	default:
		return "不明な値"
	}
}
