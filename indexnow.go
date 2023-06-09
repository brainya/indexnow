package indexnow

import (
	"context"
	"net/http"
	"strings"

	"github.com/brainya/indexnow/api"
)

type SearchEngineURL = string

const (
	BING   = SearchEngineURL("https://www.bing.com")
	YANDEX = SearchEngineURL("https://www.yandex.com")
)

func NewIndexer(seu SearchEngineURL, host string, key string, keyLocation *string) (*SearchEngine, error) {
	client, err := api.NewClient(seu)
	if err != nil {
		return nil, err
	}
	return &SearchEngine{
		searchEngineURL: seu,
		client:          client,
		host:            strings.TrimPrefix(host, "https://"),
		key:             key,
		keyLocation:     keyLocation,
	}, nil
}

type SearchEngine struct {
	searchEngineURL SearchEngineURL
	client          *api.Client
	host            string
	key             string
	keyLocation     *string
}

func (se *SearchEngine) IndexSinglePage(ctx context.Context, url string) (*http.Response, error) {
	return se.client.GetIndexnow(ctx, &api.GetIndexnowParams{
		Url:         url,
		Key:         se.key,
		KeyLocation: se.keyLocation,
	})
}

func (se *SearchEngine) BulkIndexPages(ctx context.Context, urlList *[]string) (*http.Response, error) {
	return se.client.PostIndexnow(ctx, api.UrlSet{
		Host:        &se.host,
		Key:         &se.key,
		UrlList:     urlList,
		KeyLocation: se.keyLocation,
	})
}
