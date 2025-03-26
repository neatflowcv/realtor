package naver

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

type TransactionType string

const (
	TransactionTypeSale      TransactionType = "A1" // 매매
	TransactionTypeJeonse    TransactionType = "B1" // 전세
	TransactionTypeMonthly   TransactionType = "B2" // 월세
	TransactionTypeShortTerm TransactionType = "B3" // 단기임대
)

type PropertyType string

const (
	PropertyTypeOfficetel PropertyType = "OPST"
	PropertyTypeApartment PropertyType = "VL"
)

type GetArticleListOptions struct {
	Page             int // 페이지 번호, 1번부터 시작
	MaxDeposit       int // 보증금, 만원 단위
	MaxRent          int // 월세, 만원 단위
	PropertyTypes    []PropertyType
	TransactionTypes []TransactionType
}

func WithPropertyTypes(propertyTypes ...PropertyType) GetArticleListOption {
	return func(o *GetArticleListOptions) {
		o.PropertyTypes = propertyTypes
	}
}

func WithTransactionTypes(transactionTypes ...TransactionType) GetArticleListOption {
	return func(o *GetArticleListOptions) {
		o.TransactionTypes = transactionTypes
	}
}

func NewGetArticleListOptions() *GetArticleListOptions {
	return &GetArticleListOptions{
		Page:             1,
		MaxDeposit:       0,
		MaxRent:          0,
		PropertyTypes:    nil,
		TransactionTypes: nil,
	}
}

type GetArticleListOption func(*GetArticleListOptions)

func WithPage(page int) GetArticleListOption {
	return func(o *GetArticleListOptions) {
		o.Page = page
	}
}

func WithMaxDeposit(maxDeposit int) GetArticleListOption {
	return func(o *GetArticleListOptions) {
		o.MaxDeposit = maxDeposit
	}
}

func WithMaxRent(maxRent int) GetArticleListOption {
	return func(o *GetArticleListOptions) {
		o.MaxRent = maxRent
	}
}

func (c *Client) GetArticleList(ctx context.Context, opts ...GetArticleListOption) (*ArticleList, error) {
	options := NewGetArticleListOptions()
	for _, opt := range opts {
		opt(options)
	}
	const host = "https://m.land.naver.com"
	const uri = "/cluster/ajax/articleList?"
	values := url.Values{}
	if len(options.PropertyTypes) > 0 {
		var arr []string
		for _, propertyType := range options.PropertyTypes {
			arr = append(arr, string(propertyType))
		}
		rletTpCd := strings.Join(arr, ":")
		values.Add("rletTpCd", rletTpCd)
	}
	if len(options.TransactionTypes) > 0 {
		var arr []string
		for _, transactionType := range options.TransactionTypes {
			arr = append(arr, string(transactionType))
		}
		tradTpCd := strings.Join(arr, ":")
		values.Add("tradTpCd", tradTpCd)
	}
	values.Add("z", "12") // 줌인 레벨
	values.Add("lat", "37.484265")
	values.Add("lon", "126.98864")
	values.Add("btm", "37.3645709")
	values.Add("lft", "126.6666033")
	values.Add("top", "37.6037676")
	values.Add("rgt", "127.3106767")
	values.Add("showR0", "") // FIXME: 뭔지 모르겠음
	values.Add("totCnt", "10875")
	values.Add("cortarNo", "1165000000")
	if options.MaxDeposit > 0 {
		wprcMax := strconv.FormatInt(int64(options.MaxDeposit), 10)
		values.Add("wprcMax", wprcMax)
	}
	if options.MaxRent > 0 {
		rprcMax := strconv.FormatInt(int64(options.MaxRent), 10)
		values.Add("rprcMax", rprcMax)
	}
	query := values.Encode()

	// page=2 1부터 시작한다.
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, host+uri+query, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:136.0) Gecko/20100101 Firefox/136.0")
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br, zstd")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Referer", "https://m.land.naver.com/")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("DNT", "1")
	var client http.Client
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %d, body: %s", resp.StatusCode, string(bodyText))
	}
	var article ArticleList
	err = json.Unmarshal(bodyText, &article)
	if err != nil {
		return nil, err
	}
	return &article, nil
}
