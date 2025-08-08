package kensakit

import "github.com/forroots/nomity-admin-api-v3/pkg/vo"

type KensakitReceiptImportStatus string

const (
	KensakitReceiptImportStatusPending KensakitReceiptImportStatus = "pending"
	KensakitReceiptImportStatusDone    KensakitReceiptImportStatus = "done"
	KensakitReceiptImportStatusError   KensakitReceiptImportStatus = "error"
)

func NewKensakitReceiptImportStatus(value string) (KensakitReceiptImportStatus, error) {
	status := KensakitReceiptImportStatus(value)
	switch status {
	case KensakitReceiptImportStatusPending, KensakitReceiptImportStatusDone, KensakitReceiptImportStatusError:
		return status, nil
	default:
		return status, vo.NewVOErrorf("invalid KensakitReceiptImportStatus value '%s'", value)
	}
}

func (s KensakitReceiptImportStatus) String() string {
	return string(s)
}
