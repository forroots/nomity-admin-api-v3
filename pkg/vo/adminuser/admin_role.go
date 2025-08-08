package adminuser

import (
	"encoding/json"

	"github.com/forroots/nomity-admin-api-v3/pkg/vo"
)

// AdminRole は管理者の権限を表す列挙VO
type AdminRole string

// AdminRole として許容される値
const (
	// AdminRoleSuperAdmin はスーパー管理者権限
	AdminRoleSuperAdmin AdminRole = "SuperAdmin"
	// AdminRoleSupport は一般管的なサポート業務用の権限
	AdminRoleSupport AdminRole = "Support"
)

// NewAdminRole はAdminRoleのコンストラクタ
func NewAdminRole(value string) (AdminRole, error) {
	role := AdminRole(value)
	switch role {
	case AdminRoleSuperAdmin, AdminRoleSupport:
		return role, nil
	default:
		return role, vo.NewVOErrorf("invalid AdminRole value '%s'", value)
	}
}

// String はAdminRole値オブジェクトをstringに変換する
func (r AdminRole) String() string {
	return string(r)
}

func (r AdminRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

// Display はAdminRole値オブジェクトを表示用の文字列に変換する
func (r AdminRole) Display() string {
	switch r {
	case AdminRoleSuperAdmin:
		return "Super"
	case AdminRoleSupport:
		return "Normal"
	default:
		return "Unknown Value"
	}
}

// DisplayJP はAdminRole値オブジェクトを日本語表示用の文字列に変換する
func (r AdminRole) DisplayJP() string {
	switch r {
	case AdminRoleSuperAdmin:
		return "スーパー管理者"
	case AdminRoleSupport:
		return "一般管理者"
	default:
		return "不明な値"
	}
}
