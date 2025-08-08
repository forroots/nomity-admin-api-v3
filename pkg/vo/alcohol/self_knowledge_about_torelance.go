package alcohol

import "github.com/forroots/nomity-admin-api-v3/pkg/vo"

// お酒に対する強さについての自覚を表す列挙VO
type SelfKnowledgeAboutAlcoholTorelance int

// SelfKnowledgeAboutAlcoholTorelance として許容される値
const (
	// SelfKnowledgeAboutAlcoholTorelanceUnanswered は未回答であることを表す
	SelfKnowledgeAboutAlcoholTorelanceUnanswered SelfKnowledgeAboutAlcoholTorelance = 0
	// SelfKnowledgeAboutAlcoholTorelanceDontKnow はわからないことを表す
	SelfKnowledgeAboutAlcoholTorelanceDontKnow SelfKnowledgeAboutAlcoholTorelance = 1
	// SelfKnowledgeAboutAlcoholTorelanceVeryWeak はとても弱いことを表す
	SelfKnowledgeAboutAlcoholTorelanceVeryWeak SelfKnowledgeAboutAlcoholTorelance = 2
	// SelfKnowledgeAboutAlcoholTorelanceWeak は少し弱いことを表す
	SelfKnowledgeAboutAlcoholTorelanceWeak SelfKnowledgeAboutAlcoholTorelance = 3
	// SelfKnowledgeAboutAlcoholTorelanceDontKnow はどちらともいえないこと（普通である）を表す
	SelfKnowledgeAboutAlcoholTorelanceNormal SelfKnowledgeAboutAlcoholTorelance = 4
	// SelfKnowledgeAboutAlcoholTorelanceNormal はまあまあ強いことを表す
	SelfKnowledgeAboutAlcoholTorelanceStrong SelfKnowledgeAboutAlcoholTorelance = 5
	// SelfKnowledgeAboutAlcoholTorelanceVeryStrong はとても強いことを表す
	SelfKnowledgeAboutAlcoholTorelanceVeryStrong SelfKnowledgeAboutAlcoholTorelance = 6
	// SelfKnowledgeAboutAlcoholTorelanceDontWantToSay は回答したくないことを表す
	SelfKnowledgeAboutAlcoholTorelanceDontWantToSay SelfKnowledgeAboutAlcoholTorelance = 99
)

// NewSelfKnowledgeAboutAlcoholTorelance はSelfKnowledgeAboutAlcoholTorelanceのコンストラクタ
func NewSelfKnowledgeAboutAlcoholTorelance(number int) (SelfKnowledgeAboutAlcoholTorelance, error) {
	torelance := SelfKnowledgeAboutAlcoholTorelance(number)
	switch torelance {
	case SelfKnowledgeAboutAlcoholTorelanceUnanswered,
		SelfKnowledgeAboutAlcoholTorelanceDontKnow,
		SelfKnowledgeAboutAlcoholTorelanceVeryWeak,
		SelfKnowledgeAboutAlcoholTorelanceWeak,
		SelfKnowledgeAboutAlcoholTorelanceNormal,
		SelfKnowledgeAboutAlcoholTorelanceStrong,
		SelfKnowledgeAboutAlcoholTorelanceVeryStrong,
		SelfKnowledgeAboutAlcoholTorelanceDontWantToSay:
		return torelance, nil
	default:
		return torelance, vo.NewVOErrorf("invalid SelfKnowledgeAboutAlcoholTorelance value '%d'", number)
	}
}

// Int はSelfKnowledgeAboutAlcoholTorelance値オブジェクトをintに変換する
func (t SelfKnowledgeAboutAlcoholTorelance) Int() int {
	return int(t)
}

// Display はSelfKnowledgeAboutAlcoholTorelance値オブジェクトを表示用の文字列に変換する
func (t SelfKnowledgeAboutAlcoholTorelance) Display() string {
	switch t {
	case SelfKnowledgeAboutAlcoholTorelanceUnanswered:
		return "Unanswered"
	case SelfKnowledgeAboutAlcoholTorelanceDontKnow:
		return "Don't know"
	case SelfKnowledgeAboutAlcoholTorelanceVeryWeak:
		return "Very weak"
	case SelfKnowledgeAboutAlcoholTorelanceWeak:
		return "Weak"
	case SelfKnowledgeAboutAlcoholTorelanceNormal:
		return "Normal"
	case SelfKnowledgeAboutAlcoholTorelanceStrong:
		return "Strong"
	case SelfKnowledgeAboutAlcoholTorelanceVeryStrong:
		return "Very strong"
	case SelfKnowledgeAboutAlcoholTorelanceDontWantToSay:
		return "Don't want to say"
	default:
		return "Unknown Value"
	}
}

// DisplayJa はSelfKnowledgeAboutAlcoholTorelance値オブジェクトを日本語の表示用の文字列に変換する
func (t SelfKnowledgeAboutAlcoholTorelance) DisplayJa() string {
	switch t {
	case SelfKnowledgeAboutAlcoholTorelanceUnanswered:
		return "未回答"
	case SelfKnowledgeAboutAlcoholTorelanceDontKnow:
		return "わからない"
	case SelfKnowledgeAboutAlcoholTorelanceVeryWeak:
		return "とても弱い"
	case SelfKnowledgeAboutAlcoholTorelanceWeak:
		return "弱い"
	case SelfKnowledgeAboutAlcoholTorelanceNormal:
		return "普通"
	case SelfKnowledgeAboutAlcoholTorelanceStrong:
		return "強い"
	case SelfKnowledgeAboutAlcoholTorelanceVeryStrong:
		return "とても強い"
	case SelfKnowledgeAboutAlcoholTorelanceDontWantToSay:
		return "回答したくない"
	default:
		return "不明な値"
	}
}
