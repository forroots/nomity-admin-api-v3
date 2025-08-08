package kensakit

import "github.com/forroots/nomity-admin-api-v3/pkg/vo"

// KensakitReceiptImportItemStatus は t_kensakit_receipt_import_items.status の列挙VO
type KensakitReceiptImportItemStatus string

const (
	KensakitReceiptImportItemStatusPending KensakitReceiptImportItemStatus = "pending"
	KensakitReceiptImportItemStatusDone    KensakitReceiptImportItemStatus = "done"
	KensakitReceiptImportItemStatusError   KensakitReceiptImportItemStatus = "error"
)

// NewKensakitReceiptImportItemStatus は KensakitReceiptImportItemStatus のバリデーション付きコンストラクタ
func NewKensakitReceiptImportItemStatus(value string) (KensakitReceiptImportItemStatus, error) {
	status := KensakitReceiptImportItemStatus(value)
	switch status {
	case KensakitReceiptImportItemStatusPending,
		KensakitReceiptImportItemStatusDone,
		KensakitReceiptImportItemStatusError:
		return status, nil
	default:
		return status, vo.NewVOErrorf("invalid KensakitReceiptImportItemStatus value '%s'", value)
	}
}

// String は KensakitReceiptImportItemStatus を文字列として返す
func (s KensakitReceiptImportItemStatus) String() string {
	return string(s)
}
