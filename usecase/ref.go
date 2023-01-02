package usecase

import (
	"fmt"

	"github.com/datsukan/datsukan-blog-good-core/pkg"
	"github.com/datsukan/datsukan-blog-good-core/repo"
	"github.com/datsukan/datsukan-blog-good-core/repoif"
	"github.com/guregu/dynamo"
)

// Ref は、指定した記事のいいね数を参照する。
func Ref(articleID string) (int, error) {
	db := pkg.NewDynamoDBClient()
	r := repoif.BlogGoodRepository(repo.NewBlogGoodRepository(db))

	bg, err := r.Read(articleID)
	if err != nil && err != dynamo.ErrNotFound {
		return 0, err
	}
	if bg == nil {
		fmt.Println("not found")
		return 0, nil
	}

	fmt.Printf("ArticleID: %s, Amount: %d\n", bg.ArticleID, bg.Amount)
	return bg.Amount + 200, nil
}
