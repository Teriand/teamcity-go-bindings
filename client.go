package teamcity

import (
	"time"

	"github.com/sethgrid/pester"
)

func New(url, username, password string, concurrencyLimit int) *Client {
	if concurrencyLimit == 0 {
		concurrencyLimit = 1000
	}

	http := pester.New()
	http.Concurrency = concurrencyLimit
	http.MaxRetries = 5
	http.Backoff = pester.ExponentialBackoff
	http.KeepLog = true
	http.Timeout = 60 * time.Second

	client := &Client{
		HTTPClient: http,
		URL:        url,
		Username:   username,
		Password:   password,
		Flow:       make(chan DataFlow, 10000),
		semaphore:  make(chan bool, concurrencyLimit),
	}

	go client.processDataFlow()

	return client
}

func (c *Client) Close() {
	close(c.Flow)
}
