package aoc

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type Client struct {
	client *http.Client
	cookie string
	base   *url.URL
}

func NewClient(client *http.Client, base *url.URL, cookie string) (*Client, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}

	if base == nil {
		return nil, errors.New("url is nil")
	}

	if cookie == "" {
		return nil, errors.New("cookie is empty")
	}

	return &Client{
		client: client,
		cookie: cookie,
		base:   base,
	}, nil
}

func (c *Client) DownloadInput(ctx context.Context, year, day int) (io.ReadCloser, error) {

	u := c.base.JoinPath(strconv.Itoa(year), "day", strconv.Itoa(day), "input")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	// // req.Header.Add("Cookie", fmt.Sprintf("session=%s", cookie))
	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: c.cookie,
	})

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
		return nil, fmt.Errorf("status code %d", resp.StatusCode)
	}

	return resp.Body, nil
}
