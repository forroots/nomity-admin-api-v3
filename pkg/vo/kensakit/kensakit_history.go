package kensakit

import "github.com/forroots/nomity-admin-api-v3/pkg/vo"

// KensakitHistory は検査キットの履歴を表す列挙VO
type KensakitHistory string

// KensakitHistory として許容される値
const (
	// KensakitHistoryMissingNumber は検査キットの番号が欠番扱いにされたことを表す
	KensakitHistoryMissingNumber KensakitHistory = "MISSING_NUMBER"
	// KensakitHistoryNew は検査キットの新規発番されたことを表す
	KensakitHistoryNew KensakitHistory = "NEW"
	// KensakitHistoryInProduction は検査キットが製造中になったことを表す
	KensakitHistoryInProduction KensakitHistory = "IN_PRODUCTION"
	// KensakitHistoryDoneProduction は検査キットが製造完了になったことを表す
	KensakitHistoryDoneProduction KensakitHistory = "DONE_PRODCTION"
	// KensakitHistoryDoneAuthByUser は検査キットがユーザによって認証済みになったことを表す
	KensakitHistoryDoneAuthByUser KensakitHistory = "DONE_AUTH_BY_USER"
	// KensakitHistoryInTest は検査キットが検査中になったことを表す
	KensakitHistoryInTest KensakitHistory = "IN_TEST"
	// KensakitHistoryDoneTest は検査キットが検査完了になったことを表す
	KensakitHistoryDoneTest KensakitHistory = "DONE_TEST"
	// KensakitHistoryDoneSendTestResult は検査キットが検査結果通知済みになったことを表す
	KensakitHistoryDoneSendTestResult KensakitHistory = "DONE_SEND_TEST_RESULT"
	// KensakitHistoryDoneCheckTestResultByUserListAPI は検査キットがユーザが1度でも検査結果を確認済みになったことを表す（キット一覧取得API）
	KensakitHistoryDoneCheckTestResultByUserListAPI KensakitHistory = "DONE_CHECK_TEST_RESULT_BY_USER_LIST_API"
	// KensakitHistoryDoneCheckTestResultByAuthKeyAPI は検査キットがユーザが1度でも検査結果を確認済みになったことを表す（キット認証情報によるユーザ認証不要の結果確認API）
	KensakitHistoryDoneCheckTestResultByAuthKeyAPI KensakitHistory = "DONE_CHECK_TEST_RESULT_BY_AUTH_KEY_API"
)

// NewKensakitHistory はKensakitHistoryのコンストラクタ
func NewKensakitHistory(value string) (KensakitHistory, error) {
	history := KensakitHistory(value)
	switch history {
	case KensakitHistoryMissingNumber,
		KensakitHistoryNew,
		KensakitHistoryInProduction,
		KensakitHistoryDoneProduction,
		KensakitHistoryDoneAuthByUser,
		KensakitHistoryInTest,
		KensakitHistoryDoneTest,
		KensakitHistoryDoneSendTestResult,
		KensakitHistoryDoneCheckTestResultByUserListAPI,
		KensakitHistoryDoneCheckTestResultByAuthKeyAPI:
		return history, nil
	default:
		return history, vo.NewVOErrorf("invalid KensakitHistory value '%s'", value)
	}
}

// String はKensakitHistory値オブジェクトをstringに変換する
func (h KensakitHistory) String() string {
	return string(h)
}

// Display はKensakitHistory値オブジェクトを表示用の文字列に変換する
func (h KensakitHistory) Display() string {
	switch h {
	case KensakitHistoryMissingNumber:
		return "Missing Number"
	case KensakitHistoryNew:
		return "New"
	case KensakitHistoryInProduction:
		return "In Production"
	case KensakitHistoryDoneProduction:
		return "Done Production"
	case KensakitHistoryDoneAuthByUser:
		return "Done Auth By User"
	case KensakitHistoryInTest:
		return "In Test"
	case KensakitHistoryDoneTest:
		return "Done Test"
	case KensakitHistoryDoneSendTestResult:
		return "Done Send Test Result"
	case KensakitHistoryDoneCheckTestResultByUserListAPI:
		return "Done Check Test Result By User List API"
	case KensakitHistoryDoneCheckTestResultByAuthKeyAPI:
		return "Done Check Test Result By Auth Key API"
	default:
		return "Unknown Value"
	}
}

// DisplayJa はKensakitHistory値オブジェクトを日本語の表示用の文字列に変換する
func (h KensakitHistory) DisplayJa() string {
	switch h {
	case KensakitHistoryMissingNumber:
		return "検査キット番号欠番"
	case KensakitHistoryNew:
		return "新規発番"
	case KensakitHistoryInProduction:
		return "製造中"
	case KensakitHistoryDoneProduction:
		return "製造完了"
	case KensakitHistoryDoneAuthByUser:
		return "ユーザ認証済み"
	case KensakitHistoryInTest:
		return "検査中"
	case KensakitHistoryDoneTest:
		return "検査完了"
	case KensakitHistoryDoneSendTestResult:
		return "検査結果通知済み"
	case KensakitHistoryDoneCheckTestResultByUserListAPI:
		return "ユーザ確認済み（キット一覧取得API）"
	case KensakitHistoryDoneCheckTestResultByAuthKeyAPI:
		return "ユーザ確認済み（キット認証情報によるユーザ認証不要の結果確認API）"
	default:
		return "不明な値"
	}
}
