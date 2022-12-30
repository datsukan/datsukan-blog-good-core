package repoif

import "github.com/datsukan/datsukan-blog-good-core/model"

// BlogGoodRepository は、BlogGood テーブルを操作するための Repository インターフェイス。
type BlogGoodRepository interface {
	Read(articleID string) (*model.BlogGood, error)
	Create(articleID string, amount int) (*model.BlogGood, error)
	Update(articleID string, amount int) (*model.BlogGood, error)
	Add(articleID string, amount int) (*model.BlogGood, error)
	Subtract(articleID string, amount int) (*model.BlogGood, error)
	Delete(articleID string) error
}
