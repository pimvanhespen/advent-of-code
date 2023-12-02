package aoc

import (
	"context"
	"encoding/json"
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

func NewDefaultClient() (*Client, error) {
	cookie, err := getCookie()
	if err != nil {
		return nil, fmt.Errorf("failed to get cookie: %w", err)
	}

	base, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse base url: %w", err)
	}

	return NewClient(http.DefaultClient, base, cookie)
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

func (c *Client) SubmitAnswer(ctx context.Context, year, day, part int, answer string) error {
	u := c.base.JoinPath(strconv.Itoa(year), "day", strconv.Itoa(day), "answer")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u.String(), nil)
	if err != nil {
		return err
	}

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: c.cookie,
	})

	q := req.URL.Query()
	q.Add("level", strconv.Itoa(part))
	q.Add("answer", answer)
	req.URL.RawQuery = q.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code %d", resp.StatusCode)
	}

	return nil
}

func (c *Client) GetLeaderboard(ctx context.Context, year, group int) (*Leaderboard, error) {
	u := c.base.JoinPath(strconv.Itoa(year), "leaderboard", "private", "view", fmt.Sprintf("%d.json", group))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: c.cookie,
	})

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code %d", resp.StatusCode)
	}

	var leaderboard Leaderboard

	err = json.NewDecoder(resp.Body).Decode(&leaderboard)
	if err != nil {
		return nil, fmt.Errorf("failed to decode leaderboard: %w", err)
	}

	return &leaderboard, nil
}

type Leaderboard struct {
	OwnerId int               `json:"owner_id"`
	Members map[string]Member `json:"members"`
	Event   string            `json:"event"`
}

type Member struct {
	Name               string `json:"name"`
	LocalScore         int    `json:"local_score"`
	LastStarTs         int    `json:"last_star_ts"`
	Stars              int    `json:"stars"`
	GlobalScore        int    `json:"global_score"`
	Id                 int    `json:"id"`
	CompletionDayLevel map[int]struct {
		Part1 *Part `json:"1"`
		Part2 *Part `json:"2"`
	} `json:"completion_day_level"`
}

type Part struct {
	GetStarTs int `json:"get_star_ts"`
	StarIndex int `json:"star_index"`
}
