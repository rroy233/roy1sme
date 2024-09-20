package roy1sme

import (
	"encoding/json"
	"errors"
	"github.com/valyala/fasthttp"
	"time"
)

type Client struct {
	apiKey     string
	httpClient *fasthttp.Client
}

// NewApiEndpoint Custom backend address
func NewApiEndpoint(apiServerAddr string) {
	apiEndpoint = apiServerAddr
}

// NewClient Initializes an API client
func NewClient(apiKey string) *Client {
	c := new(Client)

	c.apiKey = apiKey

	readTimeout, _ := time.ParseDuration("2000ms")
	writeTimeout, _ := time.ParseDuration("2000ms")
	maxIdleConnDuration, _ := time.ParseDuration("1h")

	c.httpClient = &fasthttp.Client{
		ReadTimeout:                   readTimeout,
		WriteTimeout:                  writeTimeout,
		MaxIdleConnDuration:           maxIdleConnDuration,
		NoDefaultUserAgentHeader:      true,
		DisableHeaderNamesNormalizing: true,
		DisablePathNormalizing:        true,
		Dial: (&fasthttp.TCPDialer{
			Concurrency:      4096,
			DNSCacheDuration: time.Hour,
		}).Dial,
	}
	return c
}

// CreateUrl Creates a short URL
// url is the input long URL
// expire is the expiration option (ExpireOneDay, ExpireOneWeek, ExpireOneMonth, ExpireNever)
//
// Returns a ShortUrl object
func (c *Client) CreateUrl(url string, expire UrlLife) (ShortUrl, error) {
	req := ReqCreate{
		Url:         url,
		CustomToken: "",
		ExpireID:    expire,
	}
	resByte, err := json.Marshal(req)
	if err != nil {
		return ShortUrl{}, errors.New("Failed to format request:" + err.Error())
	}

	respData, err := c.sendPostRequest(apiCreate, resByte)
	if err != nil {
		return ShortUrl{}, errors.New("Failed to CreateUrl:" + err.Error())
	}

	res := new(RespCreate)
	if err = json.Unmarshal(respData, res); err != nil {
		return ShortUrl{}, errors.New("Failed to Parse Response:" + err.Error())
	}

	if res.Status != 0 {
		return ShortUrl{}, errorMap[res.Status]
	}

	return res.Data, nil
}

// CreateCustomUrl Creates a custom short URL
//
// Requires user plan support
func (c *Client) CreateCustomUrl(url string, CustomUrl string, expire UrlLife) (ShortUrl, error) {
	req := ReqCreate{
		Url:         url,
		CustomToken: CustomUrl,
		ExpireID:    expire,
	}
	resByte, err := json.Marshal(req)
	if err != nil {
		return ShortUrl{}, errors.New("Failed to format request:" + err.Error())
	}

	respData, err := c.sendPostRequest(apiCreate, resByte)
	if err != nil {
		return ShortUrl{}, errors.New("Failed to CreateUrl:" + err.Error())
	}

	res := new(RespCreate)
	if err = json.Unmarshal(respData, res); err != nil {
		return ShortUrl{}, errors.New("Failed to Parse Response:" + err.Error())
	}

	if res.Status != 0 {
		return ShortUrl{}, errorMap[res.Status]
	}

	return res.Data, nil
}

// GetHistory Retrieves URL creation history
func (c *Client) GetHistory() ([]UserHistoryItem, error) {
	respData, err := c.sendGetRequest(apiHistory)
	if err != nil {
		return nil, errors.New("Failed to GetHistory:" + err.Error())
	}

	res := new(RespUserHistory)
	if err = json.Unmarshal(respData, res); err != nil {
		return nil, errors.New("Failed to Parse Response:" + err.Error())
	}

	if res.Status != 0 {
		return nil, errorMap[res.Status]
	}

	return res.Data, nil
}

// sendGetRequest Helper to send a GET request
func (c *Client) sendGetRequest(path string) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(apiEndpoint + path)
	req.Header.SetMethod(fasthttp.MethodGet)
	req.Header.Set("user-agent", ua)
	req.Header.Set("X-API-KEY", c.apiKey)
	resp := fasthttp.AcquireResponse()

	err := c.httpClient.Do(req, resp)
	fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)
	if err != nil {
		return nil, err
	}

	return resp.Body(), nil
}

// sendPostRequest Helper to send a POST request
func (c *Client) sendPostRequest(path string, reqEntityBytes []byte) ([]byte, error) {
	reqTimeout := time.Duration(2000) * time.Millisecond

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(apiEndpoint + path)
	req.Header.SetMethod(fasthttp.MethodPost)
	req.Header.SetContentTypeBytes([]byte("application/json"))
	req.Header.Set("user-agent", ua)
	req.Header.Set("X-API-KEY", c.apiKey)

	req.SetBodyRaw(reqEntityBytes)

	resp := fasthttp.AcquireResponse()
	err := c.httpClient.DoTimeout(req, resp, reqTimeout)
	fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	if err != nil {
		return nil, err
	}

	return resp.Body(), nil
}
