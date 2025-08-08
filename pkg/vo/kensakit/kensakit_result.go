package kensakit

import "github.com/forroots/nomity-admin-api-v3/pkg/vo"

// KensakitResult は検査結果を表す列挙VO
type KensakitResult int

// KensakitResult として許容される値
const (
	// KensakitResultNeedRetest は再検査が必要な状態を表す
	KensakitResultNeedRetest KensakitResult = 99
	// KensakitResultTypeA はA型の検査結果を表す
	KensakitResultTypeA KensakitResult = 1
	// KensakitResultTypeB はB型の検査結果を表す
	KensakitResultTypeB KensakitResult = 2
	// KensakitResultTypeC はC型の検査結果を表す
	KensakitResultTypeC KensakitResult = 3
	// KensakitResultTypeD はD型の検査結果を表す
	KensakitResultTypeD KensakitResult = 4
	// KensakitResultTypeE はE型の検査結果を表す
	KensakitResultTypeE KensakitResult = 5
)

// NewKensakitResult はKensakitResultのコンストラクタ
func NewKensakitResult(number int) (KensakitResult, error) {
	result := KensakitResult(number)
	switch result {
	case KensakitResultNeedRetest, KensakitResultTypeA, KensakitResultTypeB, KensakitResultTypeC, KensakitResultTypeD, KensakitResultTypeE:
		return result, nil
	default:
		return result, vo.NewVOErrorf("invalid KensakitResult value '%d'", number)
	}
}

// Int はKensakitResult値オブジェクトをintに変換する
func (r KensakitResult) Int() int {
	return int(r)
}

// Display はKensakitResult値オブジェクトを表示用の文字列に変換する
func (r KensakitResult) Display() string {
	switch r {
	case KensakitResultNeedRetest:
		return "Need Retest"
	case KensakitResultTypeA:
		return "A"
	case KensakitResultTypeB:
		return "B"
	case KensakitResultTypeC:
		return "C"
	case KensakitResultTypeD:
		return "D"
	case KensakitResultTypeE:
		return "E"
	default:
		return "Unknown Value"
	}
}

// DisplayJa はKensakitResult値オブジェクトを日本語の表示用の文字列に変換する
func (r KensakitResult) DisplayJa() string {
	switch r {
	case KensakitResultNeedRetest:
		return "再検査が必要"
	case KensakitResultTypeA:
		return "A"
	case KensakitResultTypeB:
		return "B"
	case KensakitResultTypeC:
		return "C"
	case KensakitResultTypeD:
		return "D"
	case KensakitResultTypeE:
		return "E"
	default:
		return "不明な値"
	}
}
