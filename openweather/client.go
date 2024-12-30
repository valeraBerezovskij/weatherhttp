package openweather

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Client struct {
	client  *http.Client
}

func NewClient(timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("timeout can't be zero")
	}

	return &Client{
		client: &http.Client{
			Timeout: timeout,
			Transport: &loggingRoundTripper{
				logger: os.Stdout,
				next:   http.DefaultTransport,
			},
		},
	}, nil
}

func (c *Client) doRequest(url string) ([]byte, error) {
	resp, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *Client) GetAssetWeather(api string, lat float64, lon float64) ([]weatherData, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&appid=%s",
		strconv.FormatFloat(lat, 'E', -1, 64), strconv.FormatFloat(lon, 'f', -1, 64), api)

	body, err := c.doRequest(url)
	if err != nil {
		return nil, err
	}

	var r AssetsResponse
	if err = json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	return r.Weather, nil
}

func (c *Client) GetAssetMain(api string, lat float64, lon float64) (MainData, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&appid=%s",
		strconv.FormatFloat(lat, 'f', -1, 64), strconv.FormatFloat(lon, 'f', -1, 64), api)

	body, err := c.doRequest(url)
	if err != nil {
		return MainData{}, err
	}

	var r AssetsResponse
	if err = json.Unmarshal(body, &r); err != nil {
		return MainData{}, err
	}

	return r.Main, nil
}
