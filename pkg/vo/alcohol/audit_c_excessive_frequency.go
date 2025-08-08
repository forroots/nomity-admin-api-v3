package alcohol

import "github.com/forroots/nomity-admin-api-v3/pkg/vo"

// AuditCFrequencyOfExcessiveDrinking は一度に６ドリンク以上飲酒することがどのくらいの頻度であるかを表す列挙VO（AUDIT-C）
type AuditCFrequencyOfExcessiveDrinking int

// AuditCFrequencyOfExcessiveDrinking として許容される値
const (
	// AuditCFrequencyOfExcessiveDrinkingUnanswered は未回答
	AuditCFrequencyOfExcessiveDrinkingUnanswered AuditCFrequencyOfExcessiveDrinking = 0
	// AuditCFrequencyOfExcessiveDrinkingNever はない
	AuditCFrequencyOfExcessiveDrinkingNever AuditCFrequencyOfExcessiveDrinking = 1
	// AuditCFrequencyOfExcessiveDrinkingMonthlyOrLess は月に一度未満
	AuditCFrequencyOfExcessiveDrinkingMonthlyOrLess AuditCFrequencyOfExcessiveDrinking = 2
	// AuditCFrequencyOfExcessiveDrinkingMonthly は月に一度
	AuditCFrequencyOfExcessiveDrinkingMonthly AuditCFrequencyOfExcessiveDrinking = 3
	// AuditCFrequencyOfExcessiveDrinkingWeekly は週に一度
	AuditCFrequencyOfExcessiveDrinkingWeekly AuditCFrequencyOfExcessiveDrinking = 4
	// AuditCFrequencyOfExcessiveDrinkingDailyOrAlmostDaily はほぼ毎日
	AuditCFrequencyOfExcessiveDrinkingDailyOrAlmostDaily AuditCFrequencyOfExcessiveDrinking = 5
	// AuditCFrequencyOfExcessiveDrinkingDonotWantToSay は回答したくない
	AuditCFrequencyOfExcessiveDrinkingDonotWantToSay AuditCFrequencyOfExcessiveDrinking = 99
)

// NewAuditCFrequencyOfExcessiveDrinking はAuditCFrequencyOfExcessiveDrinkingのコンストラクタ
func NewAuditCFrequencyOfExcessiveDrinking(number int) (AuditCFrequencyOfExcessiveDrinking, error) {
	frequency := AuditCFrequencyOfExcessiveDrinking(number)
	switch frequency {
	case AuditCFrequencyOfExcessiveDrinkingUnanswered,
		AuditCFrequencyOfExcessiveDrinkingNever,
		AuditCFrequencyOfExcessiveDrinkingMonthlyOrLess,
		AuditCFrequencyOfExcessiveDrinkingMonthly,
		AuditCFrequencyOfExcessiveDrinkingWeekly,
		AuditCFrequencyOfExcessiveDrinkingDailyOrAlmostDaily,
		AuditCFrequencyOfExcessiveDrinkingDonotWantToSay:
		return frequency, nil
	default:
		return frequency, vo.NewVOErrorf("invalid AuditCFrequencyOfExcessiveDrinking value '%d'", number)
	}
}

// Int はAuditCFrequencyOfExcessiveDrinking値オブジェクトをintに変換する
func (f AuditCFrequencyOfExcessiveDrinking) Int() int {
	return int(f)
}

// Display はAuditCFrequencyOfExcessiveDrinking値オブジェクトを表示用の文字列に変換する
func (f AuditCFrequencyOfExcessiveDrinking) Display() string {
	switch f {
	case AuditCFrequencyOfExcessiveDrinkingUnanswered:
		return "Unanswered"
	case AuditCFrequencyOfExcessiveDrinkingNever:
		return "Never"
	case AuditCFrequencyOfExcessiveDrinkingMonthlyOrLess:
		return "Monthly or less"
	case AuditCFrequencyOfExcessiveDrinkingMonthly:
		return "Monthly"
	case AuditCFrequencyOfExcessiveDrinkingWeekly:
		return "Weekly"
	case AuditCFrequencyOfExcessiveDrinkingDailyOrAlmostDaily:
		return "Daily or almost daily"
	default:
		return "Unknown Value"
	}
}

// DisplayJa はAuditCFrequencyOfExcessiveDrinking値オブジェクトを日本語の表示用の文字列に変換する
func (f AuditCFrequencyOfExcessiveDrinking) DisplayJa() string {
	switch f {
	case AuditCFrequencyOfExcessiveDrinkingUnanswered:
		return "未回答"
	case AuditCFrequencyOfExcessiveDrinkingNever:
		return "ない"
	case AuditCFrequencyOfExcessiveDrinkingMonthlyOrLess:
		return "月に一度未満"
	case AuditCFrequencyOfExcessiveDrinkingMonthly:
		return "月に一度"
	case AuditCFrequencyOfExcessiveDrinkingWeekly:
		return "週に一度"
	case AuditCFrequencyOfExcessiveDrinkingDailyOrAlmostDaily:
		return "ほぼ毎日"
	default:
		return "不明な値"
	}
}
