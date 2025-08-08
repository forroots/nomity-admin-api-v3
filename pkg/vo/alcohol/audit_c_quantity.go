package alcohol

import "github.com/forroots/nomity-admin-api-v3/pkg/vo"

// AuditCQuantity は飲酒するときには通常どのくらいの量を飲むかを表す列挙VO（AUDIT-C）
type AuditCQuantity int

// AuditCQuantity として許容される値
const (
	// AuditCQuantityUnanswered は未回答
	AuditCQuantityUnanswered AuditCQuantity = 0
	// AuditCQuantity1To2 は1~2ドリンク
	AuditCQuantity1To2 AuditCQuantity = 1
	// AuditCQuantity3To4 は3~4ドリンク
	AuditCQuantity3To4 AuditCQuantity = 2
	// AuditCQuantity5To6 は5~6ドリンク
	AuditCQuantity5To6 AuditCQuantity = 3
	// AuditCQuantity7To9 は7~9ドリンク
	AuditCQuantity7To9 AuditCQuantity = 4
	// AuditCQuantity10OrMore は10ドリンク以上
	AuditCQuantity10OrMore AuditCQuantity = 5
	// AuditCQuantityDonotWantToSay は回答したくない
	AuditCQuantityDonotWantToSay AuditCQuantity = 99
)

// NewAuditCQuantity はAuditCQuantityのコンストラクタ
func NewAuditCQuantity(number int) (AuditCQuantity, error) {
	quantity := AuditCQuantity(number)
	switch quantity {
	case AuditCQuantityUnanswered,
		AuditCQuantity1To2,
		AuditCQuantity3To4,
		AuditCQuantity5To6,
		AuditCQuantity7To9,
		AuditCQuantity10OrMore,
		AuditCQuantityDonotWantToSay:
		return quantity, nil
	default:
		return quantity, vo.NewVOErrorf("invalid AuditCQuantity value '%d'", number)
	}
}

// Int はAuditCQuantity値オブジェクトをintに変換する
func (q AuditCQuantity) Int() int {
	return int(q)
}

// Display はAuditCQuantity値オブジェクトを表示用の文字列に変換する
func (q AuditCQuantity) Display() string {
	switch q {
	case AuditCQuantityUnanswered:
		return "Unanswered"
	case AuditCQuantity1To2:
		return "1~2 drinks"
	case AuditCQuantity3To4:
		return "3~4 drinks"
	case AuditCQuantity5To6:
		return "5~6 drinks"
	case AuditCQuantity7To9:
		return "7~9 drinks"
	case AuditCQuantity10OrMore:
		return "10 or more drinks"
	case AuditCQuantityDonotWantToSay:
		return "Don't want to say"
	default:
		return "Unknown Value"
	}
}

// DisplayJa はAuditCQuantity値オブジェクトを日本語表示用の文字列に変換する
func (q AuditCQuantity) DisplayJa() string {
	switch q {
	case AuditCQuantityUnanswered:
		return "未回答"
	case AuditCQuantity1To2:
		return "1~2ドリンク"
	case AuditCQuantity3To4:
		return "3~4ドリンク"
	case AuditCQuantity5To6:
		return "5~6ドリンク"
	case AuditCQuantity7To9:
		return "7~9ドリンク"
	case AuditCQuantity10OrMore:
		return "10ドリンク以上"
	case AuditCQuantityDonotWantToSay:
		return "回答したくない"
	default:
		return "不明な値"
	}
}
