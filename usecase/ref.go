package usecase

import (
	"fmt"

	"github.com/datsukan/datsukan-blog-good-core/pkg"
	"github.com/datsukan/datsukan-blog-good-core/repo"
	"github.com/datsukan/datsukan-blog-good-core/repoif"
	"github.com/guregu/dynamo"
)

// Ref は、指定した記事のいいね数を参照する。
func Ref(articleID string) error {
	db := pkg.NewDynamoDBClient()
	repo := repoif.BlogGoodRepository(repo.NewBlogGoodRepository(db))

	bg, err := repo.Read(articleID)
	if err != nil && err != dynamo.ErrNotFound {
		return err
	}
	if bg == nil {
		fmt.Println("not found")
		return nil
	}

	fmt.Printf("ArticleID: %s, Amount: %d\n", bg.ArticleID, bg.Amount)
	return nil
}
