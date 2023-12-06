package easemob_go

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Error struct {
	ErrorInfo        string `json:"error"`
	Exception        string `json:"exception"`
	Timestamp        int64  `json:"timestamp"`
	Duration         int    `json:"duration"`
	ErrorDescription string `json:"error_description"`
}

func (e Error) Error() string {
	return e.ErrorDescription
}

type Response struct {
	Path            string      `json:"path,omitempty"`
	URI             string      `json:"uri,omitempty"`
	Timestamp       int64       `json:"timestamp,omitempty"`
	Count           int         `json:"count,omitempty"`
	Action          string      `json:"action,omitempty"`
	Duration        int         `json:"duration,omitempty"`
	Data            bool        `json:"data,omitempty"`
	ApplicationName string      `json:"applicationName,omitempty"`
	Organization    string      `json:"organization,omitempty"`
	Application     string      `json:"application,omitempty"`
	Cursor          string      `json:"cursor,omitempty"`
	Entities        interface{} `json:"entities,omitempty"`
	Properties      interface{} `json:"properties,omitempty"`
}
type ResultResponse struct {
	Data interface{} `json:"data"`
	Response
}

func (c *Client) parseResponse(resp *http.Response, result interface{}) error {
	if resp.Body == nil {
		return errors.New("http body is nil")
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read HTTP response: %w", err)
	}

	if resp.StatusCode != 200 {
		var apiErr Error
		err := json.Unmarshal(b, &apiErr)
		if err != nil {
			return apiErr
		}
		return apiErr
	}

	if _, ok := result.(*Response); !ok {
		// Unmarshal the body only when it is expected.
		err = json.Unmarshal(b, result)
		if err != nil {
			return fmt.Errorf("cannot unmarshal body: %w", err)
		}
	}

	return nil
}

func (c *Client) requestURL(path string, values url.Values) (string, error) {
	u, err := url.Parse(c.baseURL + "/" + path)
	if err != nil {
		return "", errors.New("url.Parse: " + err.Error())
	}
	if values == nil {
		values = make(url.Values)
	}
	u.RawQuery = values.Encode()
	return u.String(), nil
}

func (c *Client) newRequest(ctx context.Context, method, path string, params url.Values, data interface{}) (*http.Request, error) {
	u, err := c.requestURL(path, params)
	if err != nil {
		return nil, err
	}

	r, err := http.NewRequestWithContext(ctx, method, u, http.NoBody)
	if err != nil {
		return nil, err
	}

	switch t := data.(type) {
	case nil:
		r.Body = nil

	case io.ReadCloser:
		r.Body = t

	case io.Reader:
		r.Body = io.NopCloser(t)

	default:
		b, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		r.Body = io.NopCloser(bytes.NewReader(b))
	}

	return r, nil
}

func (c *Client) makeRequest(ctx context.Context, method, path string, params url.Values, data, result interface{}) error {
	r, _ := c.newRequest(ctx, method, path, params, data)

	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", "Bearer "+c.appToken)
	r.Header.Set("User-Agent", "go-sdk")

	// Uncomment the following part to dump a request
	/*
		reqDump, _ := httputil.DumpRequestOut(r, true)
		fmt.Printf("REQUEST:\n%s", string(reqDump))
	*/
	resp, err := c.http.Do(r)

	// Uncomment the following part to dump a request
	/*
		respDump, _ := httputil.DumpResponse(resp, true)
		fmt.Printf("\nRESPONSE:\n%s", string(respDump))
	*/

	if err != nil {
		select {
		case <-ctx.Done():
			// If we got an error, and the context has been canceled,
			// return context's error which is more useful.
			return ctx.Err()
		default:
		}
		return err
	}

	return c.parseResponse(resp, result)
}
func (c *Client) uploadingFile(ctx context.Context, method, path, contentType string, data io.Reader, result interface{}) error {
	u, err := c.requestURL(path, nil)
	r, _ := http.NewRequest(method, u, data)
	r.Header.Set("Authorization", "Bearer "+c.appToken)
	r.Header.Set("Content-Type", contentType)
	r.Header.Set("User-Agent", "go-sdk")
	resp, err := http.DefaultClient.Do(r)
	//resp, err := c.http.Do(r)
	if err != nil {
		select {
		case <-ctx.Done():
			// If we got an error, and the context has been canceled,
			// return context's error which is more useful.
			return ctx.Err()
		default:
		}
		return err
	}

	return c.parseResponse(resp, result)
}
