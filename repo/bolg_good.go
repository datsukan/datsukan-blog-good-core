package repo

import (
	"fmt"

	"github.com/datsukan/datsukan-blog-good-core/model"
	"github.com/guregu/dynamo"
)

// BlogGoodRepository は、DynamoDB 用の DB の構造体。
type BlogGoodRepository struct {
	Table dynamo.Table
}

// NewBlogGoodRepository は、 BlogGoodRepository のインスタンスを生成する。
func NewBlogGoodRepository(db *dynamo.DB) *BlogGoodRepository {
	return &BlogGoodRepository{Table: db.Table("DatsukanBlogGood")}
}

// Read は、いいね数のレコードを取得する。
func (r *BlogGoodRepository) Read(articleID string) (*model.BlogGood, error) {
	var b model.BlogGood
	err := r.Table.Get("ArticleID", articleID).One(&b)
	if err != nil {
		fmt.Printf("Failed to get item[%v]\n", err)
		return nil, err
	}

	return &b, nil
}

// Create は、いいね数のレコードを作成する。
func (r *BlogGoodRepository) Create(articleID string, amount int) (*model.BlogGood, error) {
	b := &model.BlogGood{
		ArticleID: articleID,
		Amount:    amount,
	}

	if err := r.Table.Put(b).Run(); err != nil {
		fmt.Printf("Failed to put item[%v]\n", err)
		return nil, err
	}

	return b, nil
}

// Update は、いいね数を更新する。
func (r *BlogGoodRepository) Update(articleID string, amount int) (*model.BlogGood, error) {
	var b *model.BlogGood
	err := r.Table.Update("ArticleID", articleID).Set("Amount", amount).Value(b)
	if err != nil {
		fmt.Printf("Failed to update item[%v]\n", err)
		return nil, err
	}

	return b, nil
}

// Add は、いいね数を加算する。
func (r *BlogGoodRepository) Add(articleID string, amount int) (*model.BlogGood, error) {
	var b model.BlogGood
	err := r.Table.Update("ArticleID", articleID).Add("Amount", amount).Value(&b)
	if err != nil {
		fmt.Printf("Failed to add item[%v]\n", err)
		return nil, err
	}

	return &b, nil
}

// Delete は、いいね数のレコードを削除する。
func (r *BlogGoodRepository) Delete(articleID string) error {
	err := r.Table.Delete("ArticleID", articleID).Run()
	if err != nil {
		fmt.Printf("Failed to delete item[%v]\n", err)
		return nil
	}

	return nil
}
