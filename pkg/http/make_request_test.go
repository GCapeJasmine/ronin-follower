package http

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/GCapeJasmine/ronin-follower/pkg/test"
	"github.com/stretchr/testify/require"
)

func TestMakeRequest_GET(t *testing.T) {
	t.Parallel()
	req, err := http.NewRequest("GET", "http://test.com", nil)
	require.NoError(t, err)
	fakeHTTPClient := test.NewTestHTTPClient(func(req *http.Request) *http.Response {
		body := `success`
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewBufferString(body)),
		}
	})
	body, err := MakeRequest(context.TODO(), fakeHTTPClient, req, 4, 0)
	require.NoError(t, err)
	require.Equal(t, `success`, string(body))
}

func TestMakeRequest_POST(t *testing.T) {
	t.Parallel()
	req, err := http.NewRequest("POST", "http://test.com", bytes.NewBufferString("data"))
	require.NoError(t, err)
	fakeHTTPClient := test.NewTestHTTPClient(func(req *http.Request) *http.Response {
		body, err1 := io.ReadAll(req.Body)
		require.NoError(t, err1)
		require.Equal(t, "data", string(body))
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewBufferString(`success`)),
		}
	})
	body, err := MakeRequest(context.TODO(), fakeHTTPClient, req, 4, 0)
	require.NoError(t, err)
	require.Equal(t, `success`, string(body))
}

func TestMakeRequest_GETRetry(t *testing.T) {
	t.Parallel()
	req, err := http.NewRequest("GET", "http://test.com", nil)
	require.NoError(t, err)
	retryCount := 0
	fakeHTTPClient := test.NewTestHTTPClient(func(req *http.Request) *http.Response {
		body := `failed`
		retryCount++
		return &http.Response{
			StatusCode: 400,
			Body:       io.NopCloser(bytes.NewBufferString(body)),
		}
	})
	body, err := MakeRequest(context.TODO(), fakeHTTPClient, req, 4, 0)
	require.Error(t, err)
	require.Equal(t, "failed", string(body))
	require.Equal(t, 4, retryCount)
}

func TestMakeRequest_POSTRetry(t *testing.T) {
	t.Parallel()
	req, err := http.NewRequest("POST", "http://test.com", bytes.NewBufferString("data"))
	require.NoError(t, err)
	retryCount := 0
	fakeHTTPClient := test.NewTestHTTPClient(func(req *http.Request) *http.Response {
		body, err1 := io.ReadAll(req.Body)
		require.NoError(t, err1)
		require.Equal(t, "data", string(body))
		retryCount++
		return &http.Response{
			StatusCode: 500,
			Body:       io.NopCloser(bytes.NewBufferString(`success`)),
		}
	})
	body, err := MakeRequest(context.TODO(), fakeHTTPClient, req, 4, 0)
	require.Error(t, err)
	require.Equal(t, `success`, string(body))
	require.Equal(t, 4, retryCount)
}
