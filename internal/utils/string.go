package utils

import "fmt"

func FormatPhoneNumber(phone string) string {
	if len(phone) == 10 {
		return fmt.Sprintf("%s-%s-%s", phone[:3], phone[3:6], phone[6:])
	} else if len(phone) == 11 {
		return fmt.Sprintf("%s-%s-%s", phone[:3], phone[3:7], phone[7:])
	}
	return phone // 変換できない場合はそのまま返す
}
