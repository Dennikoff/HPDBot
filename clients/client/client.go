package client

import (
	"encoding/json"
	"github.com/denis0108/BotHpd/lib/e"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

const getUpdatesMethod = "getUpdates"

type Client struct {
	host     string
	baseBath string
	client   http.Client
}

func New(host string, token string) Client {
	return Client{
		host:     host,
		baseBath: newBasePath(token),
		client:   http.Client{},
	}
}

func newBasePath(token string) string {
	return "bot" + token
}

//Update возвращает все updates от бота, offset смещение сообщений, limit - предел updates для 1 обращения
func (c *Client) Update(offset int, limit int) ([]Update, error) {
	const errMsg = "Can't get Updates"
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))

	data, err := c.doRequest(getUpdatesMethod, q)

	if err != nil {
		return nil, e.Wrap(errMsg, err)
	}

	var res UpdateResponse

	if err := json.Unmarshal(data, &res); err != nil {
		return nil, e.Wrap(errMsg, err)
	}
	return res.Result, nil
}

func (c *Client) SendMessage() {

}

func (c *Client) doRequest(method string, query url.Values) ([]byte, error) {
	const errMsg = "Can't do request"
	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.baseBath, method),
	}
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, e.Wrap(errMsg, err)
	}
	req.URL.RawQuery = query.Encode()

	resp, err := c.client.Do(req)

	if err != nil {
		return nil, e.Wrap(errMsg, err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, e.Wrap(errMsg, err)
	}

	return body, nil
}
