package multireq

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/p1ck0/multiset/entity"
	"github.com/p1ck0/multiset/log"
)

func SendMultiReq(ctx context.Context, requests []entity.Request) {
	logger := log.LoggerFromContext(ctx)
	for _, r := range requests {
		request(ctx, r)
	}

	logger.Info("DONE")
}

func SendMultiReqAsync(ctx context.Context, requests []entity.Request) {
	logger := log.LoggerFromContext(ctx)
	var wg sync.WaitGroup
	for _, r := range requests {
		wg.Add(1)
		go func(r entity.Request) {
			defer wg.Done()
			request(ctx, r)
		}(r)
	}
	wg.Wait()
	logger.Info("DONE")
}

func request(ctx context.Context, r entity.Request) {
	logger := log.LoggerFromContext(ctx)
	logger.Info(fmt.Sprintf("METHOD: %s | URL: %s | HEADERS: %s | BODY: %s", r.Method, r.URL, r.Headers, r.Body))
	bodyReader := bytes.NewReader(r.Body)
	request, err := http.NewRequestWithContext(ctx, r.Method, r.URL, bodyReader)
	if err != nil {
		logger.Error("could not create request")
		return
	}

	request.Header = r.Headers
	logger.Info("SENDING REQUEST ON " + r.URL)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		logger.Error("could not send request on " + r.URL)
		return
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("could not read response body")
	}
	logger.Info(fmt.Sprintf("STATUS CODE FROM %s: %d", r.URL, resp.StatusCode))
	logger.Info(fmt.Sprintf("HEADERS FROM %s: %v", r.URL, resp.Header))
	logger.Info(fmt.Sprintf("BODY FROM %s: %s", r.URL, b))
}
