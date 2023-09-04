package model

import (
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

/*
* ブックマークテーブルのカラム定義
 */
type Bookmark struct {
	Bookmark_id int       `gorm:"primary_key"`               // ブックマークID
	User_id     int       `gorm:"not null"`                  // ユーザーID
	Item_id     string    `gorm:"not null"`                  // 商品ID
	Book_info   string    `gorm:"type:text;not null"`        // 書籍データ
	Create_date time.Time `gorm:"default:CURRENT_TIMESTAMP"` // 登録日時
}

/*
* ブックマークデータの単一検索
* 【リクエスト】
* 	db: データベース接続オブジェクト
* 	conditions: 検索条件
* 	orderBy: ソート順
*
* 【レスポンス】
* 	*Bookmark: 検索結果
*		error: エラー内容
 */
func GetBookmarkData(db *gorm.DB, conditions []interface{}, orderBy string) (*Bookmark, error) {
	query := db
	var bookmark_data Bookmark

	// 検索条件セット
	if len(conditions) > 0 {
		query = query.Where(conditions[0], conditions[1:]...)
	}

	// ソート順セット
	if orderBy != "" {
		query.Order(orderBy)
	} else {
		// ソート順がない場合、bookmark_id の昇順とする
		query.Order("bookmark_id desc")
	}

	// データ検索
	if err := query.First(&bookmark_data).Error; err != nil {
		return nil, err
	}

	return &bookmark_data, nil
}

/*
* ブックマークデータの一括検索
* 【リクエスト】
*		db: データベース接続オブジェクト
*		search_column: 検索カラム
*		conditions: 検索条件
*		target_data: 検索条件に対応するカラムデータ
*
* 【レスポンス】
*		[]string: 検索結果
*		error: エラー内容
 */
func GetBookmarkDatas(db *gorm.DB, conditions string, target_data ...interface{}) ([]Bookmark, error) {
	var bookmark_data []Bookmark

	// 特定のカラムデータのみを検索
	if err := db.Where(conditions, target_data).Table("bookmarks").Order("bookmark_id desc").Find(&bookmark_data).Error; err != nil {
		return nil, err
	}

	return bookmark_data, nil
}

/*
* 特定カラムデータに紐づくブックマークデータの検索
* 【リクエスト】
*		db: データベース接続オブジェクト
*		search_column: 検索カラム
*		conditions: 検索条件
*		target_data: 検索条件に対応するカラムデータ
*
* 【レスポンス】
*		[]string: 検索結果
*		error: エラー内容
 */
func GetBookmarkSpecificData(db *gorm.DB, search_column string, conditions string, target_data ...interface{}) ([]string, error) {
	var bookmark_data []string

	// 特定のカラムデータのみを検索
	if err := db.Model(&Bookmark{}).Where(conditions, target_data).Pluck(search_column, &bookmark_data).Error; err != nil {
		return nil, err
	}
	return bookmark_data, nil
}

/*
*	ブックマークデータの新規登録
* 【リクエスト】
*		db: データベース接続オブジェクト
*		bookmark_data: 登録データセット
*
* 【レスポンス】
*		*Bookmark: 登録結果
*		error: エラー内容
 */
func CreateBookmarkData(db *gorm.DB, bookmark_data ...string) (*Bookmark, error) {
	// int への型変換
	user_id, err := strconv.Atoi(bookmark_data[0])
	if err != nil {
		return nil, err
	}

	// 登録データをセット
	newBookmark := &Bookmark{
		User_id:   user_id,
		Item_id:   bookmark_data[1],
		Book_info: bookmark_data[2],
	}

	// ブックマークデータを新規登録
	result := db.Create(newBookmark)
	if result.Error != nil {
		return nil, result.Error
	}

	return newBookmark, nil
}

/*
* ブックマークデータの更新
* 【リクエスト】
*		db: データベース接続オブジェクト
*		bookmark_id: 更新対象のブックマークID
*		updates: 更新データセット
*
* 【レスポンス】
*		*Bookmark: 更新結果
*		error: エラー内容
 */
func UpdateBookmarkData(db *gorm.DB, bookmark_id int, updates map[string]interface{}) (*Bookmark, error) {
	// ブックマークデータを検索
	var bookmark Bookmark
	result := db.First(&bookmark, bookmark_id)
	if result.Error != nil {
		return nil, result.Error
	}

	// ブックマークデータを更新
	result = db.Model(&bookmark).Updates(updates)
	if result.Error != nil {
		return nil, result.Error
	}

	return &bookmark, nil
}

/*
* ブックマークデータの削除
* 【リクエスト】
*		db: データベース接続オブジェクト
*		conditions: 検索条件
*		target_data: 検索条件に対応するカラムデータ
*
* 【レスポンス】
*		error: エラー内容
 */
func DeleteBookmarkData(db *gorm.DB, condition string, target_data ...interface{}) error {
	var bookmark Bookmark

	// ブックマークデータを削除
	if err := db.Unscoped().Where(condition, target_data).Delete(&bookmark).Error; err != nil {
		return err
	}
	return nil
}
