package gitlab

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	apiBaseURL = "/api/v4/"

	//customUserAgent used to instrument usage of the release-cli
	// https://gitlab.com/gitlab-org/gitlab/-/issues/296612
	customUserAgent = "GitLab-release-cli"
)

type HTTPClient interface {
	Send(*http.Request) (*http.Response, error)
}

// Client is used to send requests to the GitLab API. Normally created with the `New` function
type Client struct {
	fullURL      string
	jobToken     string
	privateToken string // used outside of CI
	httpClient   *http.Client
}

// New creates a new GitLab Client
func NewGitlabClient(fullURL, jobToken, privateToken string) (*Client, error) {
	if jobToken == "" && privateToken == "" {
		return nil, fmt.Errorf("%s", "access token not provided")
	}

	uri, err := url.Parse(fullURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse url: %w", err)
	}

	if !strings.Contains(uri.Path, "projects") && !strings.Contains(uri.Path, "variables") {
		return nil, fmt.Errorf("%s %s", "the url does not contains segment projects or variables", uri.Path)
	}

	client := &http.Client{}
	client.Timeout = time.Second * 15

	return &Client{
		fullURL:      uri.String(),
		jobToken:     jobToken,
		privateToken: privateToken,
		httpClient:   client,
	}, nil
}

// request creates a new request
func (gc *Client) Request(ctx context.Context, method string, value string) (*http.Request, error) {
	reader := strings.NewReader(fmt.Sprint("value=", value))
	req, err := http.NewRequest(method, gc.fullURL, reader)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	// if PRIVATE-TOKEN takes precedence over JOB-TOKEN
	if gc.privateToken != "" {
		req.Header.Set("PRIVATE-TOKEN", gc.privateToken)
	} else {
		req.Header.Set("JOB-TOKEN", gc.jobToken)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", customUserAgent)

	return req, nil
}

func (gc *Client) Send(req *http.Request) ([]byte, error) {
	resp, err := gc.httpClient.Do(req)

	if err != nil {
		return nil, fmt.Errorf("send request failed with %s", err)
	}

	data, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return nil, fmt.Errorf("read response body failed with %s", err)
	}

	return data, nil
}
