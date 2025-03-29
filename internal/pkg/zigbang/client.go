package zigbang

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/andybalholm/brotli"
)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) GetCatalogList(ctx context.Context, code string, maxDeposit uint64) (*CatalogList, error) {
	const host = "https://apis.zigbang.com"
	values := url.Values{}
	values.Add("tranTypeIn[0]", "trade")
	values.Add("tranTypeIn[1]", "charter")
	values.Add("tranTypeIn[2]", "rental")
	values.Add("includeOfferItem", "true")
	values.Add("offset", "0")
	values.Add("limit", "10")
	if maxDeposit > 0 {
		values.Add("maxSalesDeposit", strconv.FormatUint(maxDeposit, 10))
		values.Add("maxRentDeposit", strconv.FormatUint(maxDeposit, 10))
	}
	query := values.Encode()
	url := fmt.Sprintf("%v/apt/locals/%v/item-catalogs?%v", host, code, query)
	req, err := newRequest(ctx, url)
	if err != nil {
		return nil, err
	}
	bodyText, statusCode, err := c.readAll(req)
	if err != nil {
		return nil, err
	}
	if statusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: status code: %d, body: %s", ErrNotFound, statusCode, string(bodyText))
	}
	var catalog CatalogList
	err = json.Unmarshal(bodyText, &catalog)
	if err != nil {
		return nil, err
	}
	return &catalog, nil
}

func newRequest(ctx context.Context, url string) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:136.0) Gecko/20100101 Firefox/136.0")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br, zstd")
	req.Header.Set("sdk-version", "0.87.0")
	req.Header.Set("X-Zigbang-Platform", "www")
	req.Header.Set("Origin", "https://www.zigbang.com")
	req.Header.Set("DNT", "1")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Referer", "https://www.zigbang.com/")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-site")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("TE", "trailers")
	return req, nil
}

func (c *Client) readAll(req *http.Request) ([]byte, int, error) {
	var client http.Client
	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	reader := brotli.NewReader(resp.Body)
	content, err := io.ReadAll(reader)
	if err != nil {
		return nil, 0, err
	}
	return content, resp.StatusCode, nil
}

//go:embed code.csv
var codeContents []byte

func (c *Client) ListCodes() []*Code {
	lines := bytes.Split(codeContents, []byte("\n"))
	lines = lines[1:] // remove header

	var codes []*Code
	for _, line := range lines {
		fields := bytes.Split(line, []byte(","))
		if len(fields) < 2 {
			panic("invalid line: " + string(line))
		}
		id := string(fields[0])
		location := string(fields[1])
		splits := strings.Split(location, " ")
		if len(splits) != 3 {
			continue
		}
		codes = append(codes, &Code{
			ID:       id,
			Location: location,
		})
	}

	return codes
}
