package usecase

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/datsukan/datsukan-blog-good-core/pkg"
	"github.com/datsukan/datsukan-blog-good-core/repo"
	"github.com/datsukan/datsukan-blog-good-core/repoif"
	"github.com/guregu/dynamo"
)

type SQSMessage struct {
	ID string `json:"id"`
}

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

	noticeEnqueue(articleID)

	fmt.Printf("ArticleID: %s, Amount: %d\n", rbg.ArticleID, rbg.Amount)
	return rbg.Amount, nil
}

func noticeEnqueue(articleID string) {
	queueURL := os.Getenv("QUEUE_URL")
	sqsSvc := newSQSClient()

	msg := SQSMessage{ID: articleID}
	msgJson, err := json.Marshal(msg)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = sqsSvc.SendMessage(&sqs.SendMessageInput{
		MessageBody: aws.String(string(msgJson)),
		QueueUrl:    &queueURL,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("notification successful")
}

func newSQSClient() *sqs.SQS {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1"),
	}))
	return sqs.New(sess)
}
