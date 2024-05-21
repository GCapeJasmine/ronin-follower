package http

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// MakeRequest make an outbound request and handle retry
func MakeRequest(ctx context.Context, httpClient *http.Client, req *http.Request,
	maxRetry int, retrySleepSec int) ([]byte, error) {
	return makeRequestRecursive(ctx, httpClient, req, 1, maxRetry, retrySleepSec)
}

func makeRequestRecursive(ctx context.Context, httpClient *http.Client, req *http.Request, retryCount int,
	maxRetry int, retrySleepSec int) ([]byte, error) {
	log.Printf("Making outbound call to %s...", req.URL)
	var reqBody io.ReadCloser
	if req.Body != nil {
		reqBody, _ = req.GetBody() // keep this body for future retry
	}

	// set body to req after logging
	if reqBody != nil {
		req.Body = reqBody
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Printf("Error making call to %v. Error: %v", req.URL, err)
		return []byte{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return []byte{}, err
	}
	defer resp.Body.Close()

	// check http code fail : <200 && >=300
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		if retryCount >= maxRetry {
			log.Printf("Return status is not 200. Retried %d times. Got %v", retryCount, resp.StatusCode)
			return body, errors.New(fmt.Sprintf("response status code is %d", resp.StatusCode))
		}
		log.Printf("Return status code is %d not 200. RetryCount='%d' retrying...", resp.StatusCode, retryCount)
		time.Sleep(time.Duration(retrySleepSec) * time.Second)
		if reqBody != nil {
			// reset body, as the old one has been read
			req.Body = io.NopCloser(reqBody)
		}
		return makeRequestRecursive(ctx, httpClient, req, retryCount+1, maxRetry, retrySleepSec)
	}

	log.Printf("Outbound call completed with statusCode %v", resp.StatusCode)
	return body, nil
}
