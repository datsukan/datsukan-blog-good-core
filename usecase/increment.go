package usecase

import (
	"fmt"

	"github.com/datsukan/datsukan-blog-good-core/pkg"
	"github.com/datsukan/datsukan-blog-good-core/repo"
	"github.com/datsukan/datsukan-blog-good-core/repoif"
	"github.com/guregu/dynamo"
)

// Increment は、指定した記事のいいね数をインクリメント（+1）する。
func Increment(articleID string) (int, error) {
	db := pkg.NewDynamoDBClient()
	r := repoif.BlogGoodRepository(repo.NewBlogGoodRepository(db))

	bg, err := r.Read(articleID)
	if err != nil && err != dynamo.ErrNotFound {
		return 0, err
	}

	// レコードが存在しない場合。
	if bg == nil {
		rbg, err := r.Create(articleID, 1)
		if err != nil {
			return 0, err
		}

		fmt.Printf("ArticleID: %s, Amount: %d\n", rbg.ArticleID, rbg.Amount)
		return rbg.Amount, nil
	}

	// レコードが存在する場合。
	rbg, err := r.Add(articleID, 1)
	if err != nil {
		return 0, err
	}

	fmt.Printf("ArticleID: %s, Amount: %d\n", rbg.ArticleID, rbg.Amount)
	return rbg.Amount, nil
}
