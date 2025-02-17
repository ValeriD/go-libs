package api

import (
	"net/url"
	"strconv"
	"time"

	"github.com/trustwallet/go-libs/client"
)

// Client is a binance API client
type Client struct {
	req client.Request
}

func InitClient(url string, errorHandler client.HttpErrorHandler) Client {
	request := client.InitJSONClient(url, errorHandler)

	return Client{
		req: request,
	}
}

func (c *Client) GetTransactionsByAddress(address string, limit int) ([]Tx, error) {
	startTime := strconv.Itoa(int(time.Now().AddDate(0, 0, -7).Unix() * 1000))
	endTime := strconv.Itoa(int(time.Now().Unix() * 1000))
	params := url.Values{
		"address":   {address},
		"startTime": {startTime},
		"endTime":   {endTime},
		"limit":     {strconv.Itoa(limit)},
	}

	var result TransactionsResponse

	err := c.req.Get(&result, "bc/api/v1/txs", params)
	return result.Tx, err
}
