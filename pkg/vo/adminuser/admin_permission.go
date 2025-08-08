package adminuser

import "github.com/forroots/nomity-admin-api-v3/pkg/vo"

// AdminPermission は管理者の権限を表す列挙VO
type AdminPermission string

// AdminPermission として許容される値
const (
	// AdminPermissionCreateSuperAdmin はスーパー管理者の作成権限
	AdminPermissionCreateSuperAdmin AdminPermission = "CREATE_SUPER_ADMIN"
	// AdminPermissionCreateNormalAdmin は通常管理者の作成権限
	AdminPermissionCreateNormalAdmin AdminPermission = "CREATE_NORMAL_ADMIN"
)

// NewAdminPermission はAdminPermissionのコンストラクタ
func NewAdminPermission(value string) (AdminPermission, error) {
	permission := AdminPermission(value)
	switch permission {
	case AdminPermissionCreateSuperAdmin, AdminPermissionCreateNormalAdmin:
		return permission, nil
	default:
		return permission, vo.NewVOErrorf("invalid AdminPermission value '%s'", value)
	}
}

// String はAdminPermission値オブジェクトをstringに変換する
func (p AdminPermission) String() string {
	return string(p)
}

// Display はAdminPermission値オブジェクトを表示用の文字列に変換する
func (p AdminPermission) Display() string {
	switch p {
	case AdminPermissionCreateSuperAdmin:
		return "Create Super Admin"
	case AdminPermissionCreateNormalAdmin:
		return "Create Normal Admin"
	default:
		return "Unknown Value"
	}
}

// DisplayJP はAdminPermission値オブジェクトを日本語表示用の文字列に変換する
func (p AdminPermission) DisplayJP() string {
	switch p {
	case AdminPermissionCreateSuperAdmin:
		return "スーパー管理者の作成"
	case AdminPermissionCreateNormalAdmin:
		return "通常管理者の作成"
	default:
		return "不明な値"
	}
}
