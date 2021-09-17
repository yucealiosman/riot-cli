package pkg

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	urlpkg "net/url"
)

type HttpMethod string

const (
	GetMethod    HttpMethod = "GET"
	PostMethod   HttpMethod = "POST"
	PutMethod    HttpMethod = "PUT"
	PatchMethod  HttpMethod = "PATCH"
	DeleteMethod HttpMethod = "DELETE"
)

type Request struct {
	Method      HttpMethod
	Headers     map[string]string
	QueryParams map[string]string
	Body        []byte
	Path        string
}

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type Client struct {
	BaseURL    string
	HttpClient HTTPClient
}

type Response struct {
	StatusCode int
	Body       string
	Headers    map[string][]string
}

type HttpError struct {
	StatusCode int
	Err        error
}

func (h *HttpError) Error() string {
	text := fmt.Sprintf("Http Client Error Occured. Status Code: %d: Error Message: %s", h.StatusCode, h.Err.Error())

	return text
}

func addQueryParameters(baseURL string, queryParams map[string]string) string {
	baseURL += "?"
	params := urlpkg.Values{}
	for key, value := range queryParams {
		params.Add(key, value)
	}
	return baseURL + params.Encode()
}

func buildUrl(baseUrl string, urlPath string, queryParams map[string]string) string {
	parsedUrl, err := urlpkg.Parse(baseUrl)
	if err != nil {
		log.Fatal(err)
	}

	path, err := urlpkg.Parse(urlPath)
	if err != nil {
		log.Fatal(err)
	}
	url := (parsedUrl.ResolveReference(path)).String()

	if len(queryParams) != 0 {
		url = addQueryParameters(url, queryParams)
	}
	return url
}

func builHttpRequestObject(request *Request, baseUrl string) (*http.Request, error) {
	url := buildUrl(baseUrl, request.Path, request.QueryParams)

	req, err := http.NewRequest(string(request.Method), url, bytes.NewBuffer(request.Body))
	if err != nil {
		return req, err
	}
	for key, value := range request.Headers {
		req.Header.Set(key, value)
	}
	_, exists := req.Header["Content-Type"]
	if len(request.Body) > 0 && !exists {
		req.Header.Set("Content-Type", "application/json")
	}
	return req, err
}

func buildResponse(res *http.Response) (*Response, error) {
	body, err := ioutil.ReadAll(res.Body)
	response := Response{
		StatusCode: res.StatusCode,
		Body:       string(body),
		Headers:    res.Header,
	}
	res.Body.Close()
	return &response, err
}

func (c *Client) makeRequest(req *http.Request) (*http.Response, error) {
	return c.HttpClient.Do(req)
}

func (c *Client) Send(request *Request) (*Response, error) {
	req, err := builHttpRequestObject(request, c.BaseURL)
	if err != nil {
		return nil, err
	}
	rawResponse, err := c.makeRequest(req)

	if err != nil {
		return nil, err
	}
	defer rawResponse.Body.Close()

	response, err := buildResponse(rawResponse)
	if err != nil {
		return nil, err
	}
	if !((http.StatusOK <= response.StatusCode) && (response.StatusCode < http.StatusBadRequest)) {
		return nil, &HttpError{
			StatusCode: response.StatusCode,
			Err:        errors.New(response.Body),
		}
	}

	return response, nil
}

func NewGetRequest(headers map[string]string, queryParams map[string]string, path string) (*Request, error) {
	return &Request{
		Method:      http.MethodGet,
		Headers:     headers,
		QueryParams: queryParams,
		Path:        path,
	}, nil
}

func NewRequest(method HttpMethod, headers map[string]string, queryParams map[string]string,
	data map[string]string, path string) (*Request, error) {

	body, err := json.Marshal(data)

	return &Request{
		Method:      method,
		Headers:     headers,
		QueryParams: queryParams,
		Body:        body,
		Path:        path,
	}, err
}

func NewClient(baseUrl string, httpClient HTTPClient) (Client, error) {
	return Client{
		BaseURL:    baseUrl,
		HttpClient: httpClient,
	}, nil
}
