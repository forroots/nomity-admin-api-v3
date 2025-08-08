package kensakit

import "github.com/forroots/nomity-admin-api-v3/pkg/vo"

// KensakitStatusForUser はユーザーが見る検査キットのステータスを表す列挙VO
type KensakitStatusForUser int

// KensakitStatusForUser として許容される値
const (
	// KensakitStatusForUserNotReceived は検査キットが未受領の状態を表す
	KensakitStatusForUserNotReceived KensakitStatusForUser = 0
	// KensakitStatusForUserTesting は検査キットが検査中の状態を表す
	KensakitStatusForUserTesting KensakitStatusForUser = 1
	// KensakitStatusForUserDoneTesting は検査キットが検査完了の状態を表す
	KensakitStatusForUserDoneTesting KensakitStatusForUser = 2
)

// ErrInvalidKensakitStatusForUser はKensakitStatusForUser値オブジェクトが不正な値であることを示すエラー

// NewKensakitStatusForUser はKensakitStatusForUserのコンストラクタ
func NewKensakitStatusForUser(number int) (KensakitStatusForUser, error) {
	status := KensakitStatusForUser(number)
	switch status {
	case KensakitStatusForUserNotReceived, KensakitStatusForUserTesting, KensakitStatusForUserDoneTesting:
		return status, nil
	default:
		return status, vo.NewVOErrorf("invalid KensakitStatusForUser value '%d'", number)
	}
}

// Int はKensakitStatusForUser値オブジェクトをintに変換する
func (s KensakitStatusForUser) Int() int {
	return int(s)
}

// Display はKensakitStatusForUser値オブジェクトを表示用の文字列に変換する
func (s KensakitStatusForUser) Display() string {
	switch s {
	case KensakitStatusForUserNotReceived:
		return "Not Received"
	case KensakitStatusForUserTesting:
		return "Testing"
	case KensakitStatusForUserDoneTesting:
		return "Done Testing"
	default:
		return "Unknown Value"
	}
}

// DisplayJa はKensakitStatusForUser値オブジェクトを日本語の表示用の文字列に変換する
func (s KensakitStatusForUser) DisplayJa() string {
	switch s {
	case KensakitStatusForUserNotReceived:
		return "未受領"
	case KensakitStatusForUserTesting:
		return "検査中"
	case KensakitStatusForUserDoneTesting:
		return "検査完了"
	default:
		return "不明な値"
	}
}

// ステータスが「検査中」以降であるかどうかを判定する（true => 少なくとも受領登録済み）
func (s KensakitStatusForUser) IsTestingOrLater() bool {
	return s >= KensakitStatusForUserTesting
}

// ステータスが「検査完了」以降であるかどうかを判定する（true => 検査完了）
func (s KensakitStatusForUser) IsDoneTestingOrLater() bool {
	return s >= KensakitStatusForUserDoneTesting
}
