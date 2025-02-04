package client

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Azure/go-ntlmssp"
	"io"
	"log"
	"net/http"
	"net/url"
)

type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	Host   *url.URL
	Client HttpRequestDoer
}

type MiddlewareDoer struct {
	Doer       HttpRequestDoer
	Middleware func(req *http.Request) error
}

func (m *MiddlewareDoer) Do(req *http.Request) (*http.Response, error) {
	if m.Middleware != nil {
		if err := m.Middleware(req); err != nil {
			return nil, fmt.Errorf("middleware execution error: %w", err)
		}
	}

	return m.Doer.Do(req)
}

type AuthMiddleware func(req *http.Request) error

type ClientBuilder struct {
	host           string
	authMiddleware AuthMiddleware
	httpClient     HttpRequestDoer
}

func NewClientBuilder(host string) *ClientBuilder {
	return &ClientBuilder{
		host:       host,
		httpClient: &http.Client{},
	}
}

func (b *ClientBuilder) UseBasicAuth(username string, password string) *ClientBuilder {
	b.authMiddleware = func(req *http.Request) error {
		req.SetBasicAuth(username, password)
		return nil
	}

	return b
}

func (b *ClientBuilder) UseNtlmAuth(username string, password string) *ClientBuilder {
	b.httpClient = &http.Client{
		Transport: ntlmssp.Negotiator{
			RoundTripper: &http.Transport{},
		},
	}

	b.authMiddleware = func(req *http.Request) error {
		req.SetBasicAuth(username, password)
		return nil
	}

	return b
}

func (b *ClientBuilder) Build() (*Client, error) {
	hostUri, err := url.Parse(b.host)
	if err != nil {
		return nil, fmt.Errorf("error parsing host: %w", err)
	}

	return &Client{
		Host: hostUri,
		Client: &MiddlewareDoer{
			Doer:       b.httpClient,
			Middleware: b.authMiddleware,
		},
	}, nil
}

func (c *Client) HostInstancesGet(ctx context.Context) (*[]HostInstance, error) {
	return makeGetRequest[[]HostInstance](c, ctx, "/HostInstances")
}

func (c *Client) OperationalDataGetInstances(ctx context.Context) (*[]Instance, error) {
	return makeGetRequest[[]Instance](c, ctx, "/OperationalData/Instances")
}

func (c *Client) OperationalDataGetTrackedMessageEvent(ctx context.Context) (*[]TrackedMessageEvent, error) {
	return makeGetRequest[[]TrackedMessageEvent](c, ctx, "/OperationalData/TrackedMessageEvents")
}

func (c *Client) OperationalDataGetTrackedServiceInstances(ctx context.Context) (*[]TrackedServiceInstance, error) {
	return makeGetRequest[[]TrackedServiceInstance](c, ctx, "/OperationalData/TrackedServiceInstances")
}

func (c *Client) OrchestrationsGet(ctx context.Context) (*[]Orchestration, error) {
	return makeGetRequest[[]Orchestration](c, ctx, "/Orchestrations")
}

func (c *Client) ReceiveLocationsGet(ctx context.Context) (*[]ReceiveLocation, error) {
	return makeGetRequest[[]ReceiveLocation](c, ctx, "/ReceiveLocations")
}

func (c *Client) ReceivePortsGet(ctx context.Context) (*[]ReceivePort, error) {
	return makeGetRequest[[]ReceivePort](c, ctx, "/ReceivePorts")
}

func (c *Client) SendPortGroupsGet(ctx context.Context) (*[]SendPortGroup, error) {
	return makeGetRequest[[]SendPortGroup](c, ctx, "/SendPortGroups")
}

func (c *Client) SendPortsGet(ctx context.Context) (*[]SendPort, error) {
	return makeGetRequest[[]SendPort](c, ctx, "/SendPorts")
}

func makeGetRequest[T any](c *Client, ctx context.Context, endpoint string) (*T, error) {
	fullUri, err := c.Host.Parse("BizTalkManagementService" + endpoint)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", fullUri.String(), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	res, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("error: calling %s returned empty response", fullUri.String())
	}

	responseData, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("error closing response body: %+v", err)
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error calling %s:\nstatus: %s\nresponseData: %s", fullUri.String(), res.Status, responseData)
	}

	var responseObject T
	err = json.Unmarshal(responseData, &responseObject)

	if err != nil {
		log.Printf("error unmarshaling response: %+v", err)
		return nil, err
	}

	return &responseObject, nil
}
