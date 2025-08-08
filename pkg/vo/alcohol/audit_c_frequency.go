package alcohol

import "github.com/forroots/nomity-admin-api-v3/pkg/vo"

// AuditCFrequency は最近1か月の平均的な飲酒の頻度を表す列挙VO（AUDIT-C）
type AuditCFrequency int

// AuditCFrequency として許容される値
const (
	// AuditCFrequencyUnanswered は未回答
	AuditCFrequencyUnanswered AuditCFrequency = 0
	// AuditCFrequencyNever は飲まない
	AuditCFrequencyNever AuditCFrequency = 1
	// AuditCFrequencyMonthlyOrLess は月に一度以下
	AuditCFrequencyMonthlyOrLess AuditCFrequency = 2
	// AuditCFrequency2To4TimesAMonth は月に2〜4度
	AuditCFrequency2To4TimesAMonth AuditCFrequency = 3
	// AuditCFrequency2To3TimesAWeek は週に2〜3度
	AuditCFrequency2To3TimesAWeek AuditCFrequency = 4
	// AuditCFrequency4OrMoreTimesAWeek は週に4度以上
	AuditCFrequency4OrMoreTimesAWeek AuditCFrequency = 5
	// AuditCFrequencyDonotWantToSay は回答したくない
	AuditCFrequencyDonotWantToSay AuditCFrequency = 99
)

// NewAuditCFrequency はAuditCFrequencyのコンストラクタ
func NewAuditCFrequency(number int) (AuditCFrequency, error) {
	frequency := AuditCFrequency(number)
	switch frequency {
	case AuditCFrequencyUnanswered,
		AuditCFrequencyNever,
		AuditCFrequencyMonthlyOrLess,
		AuditCFrequency2To4TimesAMonth,
		AuditCFrequency2To3TimesAWeek,
		AuditCFrequency4OrMoreTimesAWeek,
		AuditCFrequencyDonotWantToSay:
		return frequency, nil
	default:
		return frequency, vo.NewVOErrorf("invalid AuditCFrequency value '%d'", number)
	}
}

// Int はAuditCFrequency値オブジェクトをintに変換する
func (f AuditCFrequency) Int() int {
	return int(f)
}

// Display はAuditCFrequency値オブジェクトを表示用の文字列に変換する
func (f AuditCFrequency) Display() string {
	switch f {
	case AuditCFrequencyUnanswered:
		return "Unanswered"
	case AuditCFrequencyNever:
		return "Never"
	case AuditCFrequencyMonthlyOrLess:
		return "Monthly or less"
	case AuditCFrequency2To4TimesAMonth:
		return "2 to 4 times a month"
	case AuditCFrequency2To3TimesAWeek:
		return "2 to 3 times a week"
	case AuditCFrequency4OrMoreTimesAWeek:
		return "4 or more times a week"
	default:
		return "Unknown Value"
	}
}

// DisplayJa はAuditCFrequency値オブジェクトを日本語の表示用の文字列に変換する
func (f AuditCFrequency) DisplayJa() string {
	switch f {
	case AuditCFrequencyUnanswered:
		return "未回答"
	case AuditCFrequencyNever:
		return "飲まない"
	case AuditCFrequencyMonthlyOrLess:
		return "月に一度以下"
	case AuditCFrequency2To4TimesAMonth:
		return "月に2〜4度"
	case AuditCFrequency2To3TimesAWeek:
		return "週に2〜3度"
	case AuditCFrequency4OrMoreTimesAWeek:
		return "週に4度以上"
	default:
		return "不明な値"
	}
}
