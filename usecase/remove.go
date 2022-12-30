package usecase

import (
	"github.com/datsukan/datsukan-blog-good-core/pkg"
	"github.com/datsukan/datsukan-blog-good-core/repo"
	"github.com/datsukan/datsukan-blog-good-core/repoif"
)

// Remove は、指定した記事のいいね数を削除する。
func Remove(articleID string) error {
	db := pkg.NewDynamoDBClient()
	repo := repoif.BlogGoodRepository(repo.NewBlogGoodRepository(db))

	if err := repo.Delete(articleID); err != nil {
		return err
	}

	return nil
}
