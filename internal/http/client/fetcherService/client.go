package fetcherService

import (
	"bytes"
	"context"
	"encoding/json"
	"holidays-seeker/cmd/config"
	"holidays-seeker/internal/core/domain/fetcher"
	"holidays-seeker/internal/core/domain/holiday"
	"io"
	"net/http"
)

type apiClientHTTP struct {
	url string
}

func (c *apiClientHTTP) RetrieveHolidays(ctx context.Context) ([]holiday.Holiday, error) {
	url := c.url

	result, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := io.ReadAll(result.Body)
	if err != nil {
		return nil, err
	}

	// The error referring to ï is because the UTF-8 BOM interpreted as an ISO-8859-1 string will produce the characters ï»¿.
	bodyBytes = bytes.TrimPrefix(bodyBytes, []byte("\xef\xbb\xbf")) // Or []byte{239, 187, 191}

	var holidays holiday.Holidays

	err = json.Unmarshal(bodyBytes, &holidays)
	if err != nil {
		return nil, err
	}

	holidateout := holidays.Holidays
	return holidateout, nil
}

func NewFetcherService(c config.Config) fetcher.Service {
	return &apiClientHTTP{url: c.Api.Holiday.Url}
}
