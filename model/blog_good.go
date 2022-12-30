package model

// BlogGood は、いいね数を保持するテーブルの構造体。
type BlogGood struct {
	ArticleID string `dynamo:"ArticleID,hash"`
	Amount    int    `dynamo:"Amount"`
}
