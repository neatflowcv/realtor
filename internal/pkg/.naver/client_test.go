package naver_test

import (
	"context"
	"testing"
	"time"

	"github.com/neatflowcv/realtor/internal/pkg/naver"
	"github.com/stretchr/testify/require"
)

func TestGetArticleList(t *testing.T) {
	t.Parallel()
	client := naver.NewClient()
	ctx := context.Background()

	var articles []*naver.ArticleList
	for i := 2; i <= 100; i++ {
		article, err := client.GetArticleList(
			ctx,
			naver.WithPropertyTypes(naver.PropertyTypeOfficetel),
			naver.WithTransactionTypes(naver.TransactionTypeMonthly, naver.TransactionTypeShortTerm),
			naver.WithMaxDeposit(1000),
			naver.WithMaxRent(70),
			naver.WithPage(i),
		)
		require.NoError(t, err, "page: %d", i)
		articles = append(articles, article)
		require.Equal(t, "success", article.Code)
		require.NotEmpty(t, article.Bodies)
		if !article.More {
			break
		}
		time.Sleep(1 * time.Second)
	}
	t.Log()
}
