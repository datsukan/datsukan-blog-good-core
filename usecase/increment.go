package usecase

import (
	"fmt"

	"github.com/datsukan/datsukan-blog-good-core/pkg"
	"github.com/datsukan/datsukan-blog-good-core/repo"
	"github.com/datsukan/datsukan-blog-good-core/repoif"
	"github.com/guregu/dynamo"
)

// Increment は、指定した記事のいいね数をインクリメント（+1）する。
func Increment(articleID string) error {
	db := pkg.NewDynamoDBClient()
	repo := repoif.BlogGoodRepository(repo.NewBlogGoodRepository(db))

	bg, err := repo.Read(articleID)
	if err != nil && err != dynamo.ErrNotFound {
		return err
	}

	// レコードが存在する場合。
	if bg != nil {
		bg, err := repo.Add(articleID, 1)
		if err != nil {
			return err
		}

		fmt.Printf("ArticleID: %s, Amount: %d\n", bg.ArticleID, bg.Amount)
		return nil
	}

	// レコードが存在しない場合。
	rbg, err := repo.Create(articleID, 1)
	if err != nil {
		return err
	}

	fmt.Printf("ArticleID: %s, Amount: %d\n", rbg.ArticleID, rbg.Amount)
	return nil
}
