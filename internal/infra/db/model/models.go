// internal/infra/db/model/models.go
package model

import (
	"encoding/json"
	"time"

	"github.com/uptrace/bun"
)

/* ========== KV storages ========== */

type FiberStorage struct {
	bun.BaseModel `bun:"table:fiber_storage"`

	K string `bun:"k,pk,notnull,default:''"`
	V []byte `bun:"v,notnull"`
	E int64  `bun:"e,notnull,default:0"`
}

type FiberStorageAdmin struct {
	bun.BaseModel `bun:"table:fiber_storage_admin"`

	K string `bun:"k,pk,notnull,default:''"`
	V []byte `bun:"v,notnull"`
	E int64  `bun:"e,notnull,default:0"`
}

/* ========== Admin users & logs ========== */

type TAdminUser struct {
	bun.BaseModel `bun:"table:t_admin_users"`

	ID               int64      `bun:"id,pk"`
	CreatedAt        *time.Time `bun:"created_at"`
	UpdatedAt        *time.Time `bun:"updated_at"`
	DeletedAt        *time.Time `bun:"deleted_at,soft_delete"`
	Name             string     `bun:"name,notnull"`
	Email            string     `bun:"email,notnull"`
	Role             string     `bun:"role,notnull"`
	Active           bool       `bun:"active,notnull"`
	EmailConfirmedAt *time.Time `bun:"email_confirmed_at"`

	// Relations
	ActionLogs         []TAdminActionLog            `bun:"rel:has-many,join:id=t_admin_user_id"`
	DeviceSessions     []WAdminDeviceSession        `bun:"rel:has-many,join:id=t_admin_user_id"`
	ReceiptImports     []TKensakitReceiptImport     `bun:"-"` // FK未定義（created_by_admin_user_id）はマッピングしない
	ReceiptImportItems []TKensakitReceiptImportItem `bun:"-"` // 同上
}

type TAdminActionLog struct {
	bun.BaseModel `bun:"table:t_admin_action_logs"`

	ID           int64            `bun:"id,pk"`
	CreatedAt    time.Time        `bun:"created_at,notnull"`
	TAdminUserID int64            `bun:"t_admin_user_id,notnull"`
	ActionType   string           `bun:"action_type,notnull"`
	Message      *string          `bun:"message"`
	Details      *json.RawMessage `bun:"details,type:jsonb"`
	RequestID    string           `bun:"request_id,notnull"`
	IPAddress    string           `bun:"ip_address,notnull"`

	// Relations
	AdminUser *TAdminUser `bun:"rel:belongs-to,join:t_admin_user_id=id"`
}

type WAdminDeviceSession struct {
	bun.BaseModel `bun:"table:w_admin_device_sessions"`

	SessionID    string     `bun:"session_id,pk"`
	CreatedAt    *time.Time `bun:"created_at"`
	TAdminUserID *int64     `bun:"t_admin_user_id"`
	Email        *string    `bun:"email"`
	OTP          *string    `bun:"otp"`
	OTPExpiresAt *time.Time `bun:"otp_expires_at"`

	// Relations
	AdminUser *TAdminUser `bun:"rel:belongs-to,join:t_admin_user_id=id"`
}

/* ========== Articles & tags ========== */

type TArticle struct {
	bun.BaseModel `bun:"table:t_articles"`

	ID           int64      `bun:"id,pk"`
	CreatedAt    *time.Time `bun:"created_at"`
	UpdatedAt    *time.Time `bun:"updated_at"`
	DeletedAt    *time.Time `bun:"deleted_at,soft_delete"`
	Title        string     `bun:"title,notnull"`
	ThumbnailURL string     `bun:"thumbnail_url,notnull"`
	ContentURL   string     `bun:"content_url,notnull"`
	ReleaseDate  time.Time  `bun:"release_date,notnull"`

	// Relations
	Tags []TArticleTag `bun:"rel:has-many,join:id=t_article_id"`
}

type TArticleTag struct {
	bun.BaseModel `bun:"table:t_article_tags"`

	ID           int64      `bun:"id,pk"`
	CreatedAt    *time.Time `bun:"created_at"`
	UpdatedAt    *time.Time `bun:"updated_at"`
	DeletedAt    *time.Time `bun:"deleted_at,soft_delete"`
	TArticleID   int64      `bun:"t_article_id,notnull"`
	Tag          string     `bun:"tag,notnull"`
	DisplayOrder int64      `bun:"display_order,notnull"`

	// Relations
	Article *TArticle `bun:"rel:belongs-to,join:t_article_id=id"`
}

/* ========== Users domain ========== */

type TUser struct {
	bun.BaseModel `bun:"table:t_users"`

	ID           int64      `bun:"id,pk"`
	CreatedAt    *time.Time `bun:"created_at"`
	UpdatedAt    *time.Time `bun:"updated_at"`
	DeletedAt    *time.Time `bun:"deleted_at,soft_delete"`
	ParentUserID *int64     `bun:"parent_user_id"`
	Name         string     `bun:"name,notnull"`
	Email        *string    `bun:"email"`
	PasswordHash string     `bun:"password_hash,notnull"`
	SerchingID   *string    `bun:"serching_id"`
	AllowExpose  bool       `bun:"allow_expose,notnull,default:false"`
	RoleCD       string     `bun:"role_cd,notnull"`

	// Relations
	Parent   *TUser  `bun:"rel:belongs-to,join:parent_user_id=id"`
	Children []TUser `bun:"rel:has-many,join:id=parent_user_id"`

	Details    *TUserDetail            `bun:"rel:has-one,join:id=t_user_id"`
	Taishitsu  *TTaishitsuQuestionaire `bun:"rel:has-one,join:id=t_user_id"`
	ActionLogs []TUserActionLog        `bun:"rel:has-many,join:id=t_user_id"`
	Favorites  []TFavoriteAlcohol      `bun:"rel:has-many,join:id=t_user_id"`
	Kensakits  []TKensakit             `bun:"rel:has-many,join:id=t_user_id"`
	Sessions   []WSession              `bun:"rel:has-many,join:id=t_user_id"`
}

type TUserDetail struct {
	bun.BaseModel `bun:"table:t_user_details"`

	ID                     int64      `bun:"id,pk"`
	CreatedAt              *time.Time `bun:"created_at"`
	UpdatedAt              *time.Time `bun:"updated_at"`
	DeletedAt              *time.Time `bun:"deleted_at,soft_delete"`
	TUserID                int64      `bun:"t_user_id,notnull"`
	SexCD                  int64      `bun:"sex_cd,notnull"`
	BirthDate              *time.Time `bun:"birth_date"` // date; time part ignored
	ResidenceCD            int64      `bun:"residence_cd,notnull"`
	BirthplaceCD           int64      `bun:"birthplace_cd,notnull"`
	TypeOfWorkCD           string     `bun:"type_of_work_cd,notnull"`
	TypeOfOccupationCD     string     `bun:"type_of_occupation_cd,notnull"`
	AlcoholTaishitsuTypeCD int64      `bun:"alcohol_taishitsu_type_cd,notnull"`

	// Relations
	User *TUser `bun:"rel:belongs-to,join:t_user_id=id"`
}

type TTaishitsuQuestionaire struct {
	bun.BaseModel `bun:"table:t_taishitsu_questionaires"`

	ID                                int64      `bun:"id,pk"`
	CreatedAt                         *time.Time `bun:"created_at"`
	UpdatedAt                         *time.Time `bun:"updated_at"`
	DeletedAt                         *time.Time `bun:"deleted_at,soft_delete"`
	TUserID                           int64      `bun:"t_user_id,notnull"`
	AuditCFrequencyCD                 int64      `bun:"audit_c_frequency_cd,notnull"`
	AuditCQuantityCD                  int64      `bun:"audit_c_quantity_cd,notnull"`
	AuditCExcessiveAlcoholFrequencyCD int64      `bun:"audit_c_excessive_alcohol_frequency_cd,notnull"`
	SelfKnowledgeAboutToleranceCD     int64      `bun:"self_knowledge_about_tolerance_cd,notnull"`
	SelfKnowledgeAboutTypeCD          int64      `bun:"self_knowledge_about_type_cd,notnull"`

	// Relations
	User *TUser `bun:"rel:belongs-to,join:t_user_id=id"`
}

type TUserActionLog struct {
	bun.BaseModel `bun:"table:t_user_action_logs"`

	ID         int64            `bun:"id,pk"`
	CreatedAt  *time.Time       `bun:"created_at"`
	UpdatedAt  *time.Time       `bun:"updated_at"`
	DeletedAt  *time.Time       `bun:"deleted_at,soft_delete"`
	TUserID    int64            `bun:"t_user_id,notnull"`
	ActionType string           `bun:"action_type,notnull"`
	Details    *json.RawMessage `bun:"details,type:jsonb"`
	Message    string           `bun:"message,notnull"`

	// Relations
	User *TUser `bun:"rel:belongs-to,join:t_user_id=id"`
}

type TFavoriteAlcohol struct {
	bun.BaseModel `bun:"table:t_favorite_alcohols"`

	TUserID       int64      `bun:"t_user_id,pk"`
	AlcoholTypeCD int64      `bun:"alcohol_type_cd,pk"`
	CreatedAt     *time.Time `bun:"created_at"`
	UpdatedAt     *time.Time `bun:"updated_at"`

	// Relations
	User *TUser `bun:"rel:belongs-to,join:t_user_id=id"`
}

/* ========== Kensakit domain ========== */

type TKensakit struct {
	bun.BaseModel `bun:"table:t_kensakits"`

	ID            int64      `bun:"id,pk"`
	CreatedAt     *time.Time `bun:"created_at"`
	UpdatedAt     *time.Time `bun:"updated_at"`
	DeletedAt     *time.Time `bun:"deleted_at,soft_delete"`
	TUserID       *int64     `bun:"t_user_id"`
	SerialNumber  string     `bun:"serial_number,notnull"`
	AuthKey1      string     `bun:"auth_key1,notnull"`
	AuthKey2      string     `bun:"auth_key2,notnull"`
	StatusForUser int64      `bun:"status_for_user,notnull"`
	TestResult    *int64     `bun:"test_result"`

	// Relations
	User         *TUser                   `bun:"rel:belongs-to,join:t_user_id=id"`
	StatusEvents []TKensakitStatusHistory `bun:"rel:has-many,join:id=t_kensakit_id"`
	Sanwa        *TSanwaQuestionaire      `bun:"rel:has-one,join:id=t_kensakit_id"`
}

type TKensakitStatusHistory struct {
	bun.BaseModel `bun:"table:t_kensakit_status_histories"`

	ID           int64      `bun:"id,pk"`
	CreatedAt    *time.Time `bun:"created_at"`
	UpdatedAt    *time.Time `bun:"updated_at"`
	DeletedAt    *time.Time `bun:"deleted_at,soft_delete"`
	TKensakitID  int64      `bun:"t_kensakit_id,notnull"`
	StatusCD     string     `bun:"status_cd,notnull"`
	ExecutedTime time.Time  `bun:"executed_time,notnull"`

	// Relations
	Kensakit *TKensakit `bun:"rel:belongs-to,join:t_kensakit_id=id"`
}

type TSanwaQuestionaire struct {
	bun.BaseModel `bun:"table:t_sanwa_questionaires"`

	ID          int64      `bun:"id,pk"`
	CreatedAt   *time.Time `bun:"created_at"`
	UpdatedAt   *time.Time `bun:"updated_at"`
	DeletedAt   *time.Time `bun:"deleted_at,soft_delete"`
	TKensakitID int64      `bun:"t_kensakit_id,notnull"`

	// …質問項目は省略なく定義するなら全フィールドを列挙（ここでは割愛）
	IsQS1Done   *bool `bun:"is_qs1_done"`
	IsQS2Done   *bool `bun:"is_qs2_done"`
	IsQS3_1Done *bool `bun:"is_qs3_1_done"`
	IsQS3_2Done *bool `bun:"is_qs3_2_done"`
	IsQS3_3Done *bool `bun:"is_qs3_3_done"`
	// （残りの qs_*/qn_* カラムも必要に応じて追加してください）

	// Relations
	Kensakit *TKensakit `bun:"rel:belongs-to,join:t_kensakit_id=id"`
}

type TKensakitReceiptImport struct {
	bun.BaseModel `bun:"table:t_kensakit_receipt_imports"`

	ID                                    int64      `bun:"id,pk"`
	CreatedByAdminUserID                  int64      `bun:"created_by_admin_user_id,notnull"` // FK未定義
	AdminMemo                             *string    `bun:"admin_memo"`
	CreatedAt                             *time.Time `bun:"created_at"`
	UpdatedAt                             *time.Time `bun:"updated_at"`
	DeletedAt                             *time.Time `bun:"deleted_at,soft_delete"`
	Status                                string     `bun:"status,notnull"`
	ImportedAt                            time.Time  `bun:"imported_at,notnull"`
	ImportedFileName                      string     `bun:"imported_file_name,notnull"`
	UpdateKensakitStatusBatchAt           *time.Time `bun:"update_kensakit_status_batch_at"`
	UpdateKensakitStatusBatchSuccess      *bool      `bun:"update_kensakit_status_batch_success"`
	UpdateKensakitStatusBatchErrorMessage *string    `bun:"update_kensakit_status_batch_error_message"`
	SendMailBatchAt                       *time.Time `bun:"send_mail_batch_at"`
	SendMailBatchSuccess                  *bool      `bun:"send_mail_batch_success"`
	SendMailBatchErrorMessage             *string    `bun:"send_mail_batch_error_message"`

	// Relations
	Items []TKensakitReceiptImportItem `bun:"rel:has-many,join:id=t_kensakit_receipt_import_id"`
}

type TKensakitReceiptImportItem struct {
	bun.BaseModel `bun:"table:t_kensakit_receipt_import_items"`

	ID                                    int64      `bun:"id,pk"`
	CreatedByAdminUserID                  int64      `bun:"created_by_admin_user_id,notnull"` // FK未定義
	AdminMemo                             *string    `bun:"admin_memo"`
	CreatedAt                             *time.Time `bun:"created_at"`
	UpdatedAt                             *time.Time `bun:"updated_at"`
	DeletedAt                             *time.Time `bun:"deleted_at,soft_delete"`
	TKensakitReceiptImportID              int64      `bun:"t_kensakit_receipt_import_id,notnull"`
	SerialNumber                          string     `bun:"serial_number,notnull"`
	Status                                string     `bun:"status,notnull"`
	UpdateKensakitStatusBatchAt           *time.Time `bun:"update_kensakit_status_batch_at"`
	UpdateKensakitStatusBatchSuccess      *bool      `bun:"update_kensakit_status_batch_success"`
	UpdateKensakitStatusBatchErrorMessage *string    `bun:"update_kensakit_status_batch_error_message"`
	SendMailBatchAt                       *time.Time `bun:"send_mail_batch_at"`
	SendMailBatchSuccess                  *bool      `bun:"send_mail_batch_success"`
	SendMailBatchErrorMessage             *string    `bun:"send_mail_batch_error_message"`
	SendMailTo                            *string    `bun:"send_mail_to"`

	// Relations
	Import *TKensakitReceiptImport `bun:"rel:belongs-to,join:t_kensakit_receipt_import_id=id"`
}

/* ========== Restaurants ========== */

type TRestaurant struct {
	bun.BaseModel `bun:"table:t_restaurants"`

	ID         int64      `bun:"id,pk"`
	CreatedAt  *time.Time `bun:"created_at"`
	UpdatedAt  *time.Time `bun:"updated_at"`
	DeletedAt  *time.Time `bun:"deleted_at,soft_delete"`
	Name       string     `bun:"name,notnull"`
	Prefecture string     `bun:"prefecture,notnull"`
	City       string     `bun:"city,notnull"`
}

/* ========== Android temp users ========== */

type TAndroidTmpUser struct {
	bun.BaseModel `bun:"table:t_android_tmp_users"`

	ID        int64      `bun:"id,pk"`
	CreatedAt *time.Time `bun:"created_at"`
	UpdatedAt *time.Time `bun:"updated_at"`
	DeletedAt *time.Time `bun:"deleted_at,soft_delete"`
	Email     string     `bun:"email,notnull"`
}

/* ========== Sessions & auth ========== */

type WSession struct {
	bun.BaseModel `bun:"table:w_sessions"`

	AccessToken           string     `bun:"access_token,pk"`
	RefreshToken          string     `bun:"refresh_token,notnull"`
	TUserID               int64      `bun:"t_user_id,notnull"`
	AccessTokenExpiredIn  time.Time  `bun:"access_token_expired_in,notnull"`
	RefreshTokenExpiredIn time.Time  `bun:"refresh_token_expired_in,notnull"`
	CreatedAt             *time.Time `bun:"created_at"`
	UpdatedAt             *time.Time `bun:"updated_at"`
	DeviceType            string     `bun:"device_type,notnull,default:'mobile'"`

	// Relations
	User *TUser `bun:"rel:belongs-to,join:t_user_id=id"`
}

type WEmailAuthenticating struct {
	bun.BaseModel `bun:"table:w_email_authenticatings"`

	AccessToken   string     `bun:"access_token,pk"`
	Email         string     `bun:"email,notnull"`
	AuthCode      string     `bun:"auth_code,notnull"`
	Authenticated bool       `bun:"authenticated,notnull"`
	CreatedAt     *time.Time `bun:"created_at"`
	UpdatedAt     *time.Time `bun:"updated_at"`
}
