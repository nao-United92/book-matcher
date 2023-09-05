package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

/*
* ユーザーテーブルのカラム定義
 */
type User struct {
	User_id      int        `gorm:"primary_key"` // ユーザーID
	User_name    string     // ユーザー名
	Address      string     // メールアドレス
	Password     string     // パスワード
	Image        string     // ユーザーのプロフィール画像
	Introduction string     `gorm:"type:text"`                 // 自己紹介文
	Create_date  time.Time  `gorm:"default:CURRENT_TIMESTAMP"` // 登録日時
	Update_date  time.Time  `gorm:"default:CURRENT_TIMESTAMP"` // 更新日時
	Delete_date  *time.Time // 削除日時
	Delete_flag  int        `gorm:"default:0"` // 削除フラグ
}

/*
* ユーザーデータの単一検索
* 【リクエスト】
*   db: データベース接続オブジェクト
*   conditions: 検索条件
*   orderBy: ソート順
*
* 【レスポンス】
*   *User: 検索結果
*   error: エラー内容
 */
func GetUserData(db *gorm.DB, conditions []interface{}, orderBy string) (*User, error) {
	query := db
	var users_data User

	// 検索条件セット
	if len(conditions) > 0 {
		query = query.Where(conditions[0], conditions[1:]...)
	}

	// ソート順セット
	if orderBy != "" {
		query = query.Order(orderBy)
	} else {
		// ソート順がない場合、user_id の昇順とする
		query = query.Order("user_id asc")
	}

	// データ検索
	if err := query.First(&users_data).Error; err != nil {
		return nil, err
	}

	return &users_data, nil
}

/*
* 特定カラムデータに紐づくユーザーデータの検索
* 【リクエスト】
*   db: データベース接続オブジェクト
*   search_column: 検索カラム
*   conditions: 検索条件
*   target_data: 検索条件に対応するカラムデータ
*
* 【レスポンス】
*   []string: 検索結果
*   error: エラー内容
 */
func GetUserSpecificData(db *gorm.DB, search_column string, conditions string, target_data ...interface{}) ([]string, error) {
	var user_data []string

	// 特定のカラムデータのみを検索
	if err := db.Model(&User{}).Where(conditions, target_data).Pluck(search_column, &user_data).Error; err != nil {
		return nil, err
	}
	return user_data, nil
}

/*
* ユーザーデータの新規登録
* 【リクエスト】
*   db: データベース接続オブジェクト
*   user_data: 登録データセット
*
* 【レスポンス】
*   *User: 登録結果
*   error: エラー内容
 */
func CreateUserData(db *gorm.DB, user_data ...string) (*User, error) {
	// 登録データをセット
	newUser := &User{
		User_name: user_data[0],
		Address:   user_data[1],
		Password:  user_data[2],
	}

	// ユーザーデータを新規登録
	result := db.Create(newUser)
	if result.Error != nil {
		return nil, result.Error
	}

	return newUser, nil
}

/*
* ユーザーデータの更新
* 【リクエスト】
*   db: データベース接続オブジェクト
*   user_id: 更新対象のユーザーID
*   updates: 更新データセット
*
* 【レスポンス】
*   *User: 更新結果
*   error: エラー内容
 */
func UpdateUserData(db *gorm.DB, user_id int, updates map[string]interface{}) (*User, error) {
	// ユーザーデータを検索
	var user User
	result := db.First(&user, user_id)
	if result.Error != nil {
		return nil, result.Error
	}

	// ユーザーデータを更新
	result = db.Model(&user).Updates(updates)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
