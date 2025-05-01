package httpclient

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"math/rand/v2"
	"net/http"
	"time"
)

// Client is a custom HTTP client
type Client struct {
	baseClient     *http.Client
	maxRetries     int
	retryWaitMin   time.Duration
	retryWaitMax   time.Duration
	logger         Logger
	retryableCodes []int
}

// Logger defines the logging interface
type Logger interface {
	Debugf(format string, args ...any)
	Infof(format string, args ...any)
	Warnf(format string, args ...any)
	Errorf(format string, args ...any)
}

// DefaultLogger provides a default logging implementation
type DefaultLogger struct{}

func (l *DefaultLogger) Debugf(format string, args ...any) {
	log.Printf("[DEBUG] "+format, args...)
}

func (l *DefaultLogger) Infof(format string, args ...any) {
	log.Printf("[INFO] "+format, args...)
}

func (l *DefaultLogger) Warnf(format string, args ...any) {
	log.Printf("[WARN] "+format, args...)
}

func (l *DefaultLogger) Errorf(format string, args ...any) {
	log.Printf("[ERROR] "+format, args...)
}

// Option configures the Client
type Option func(*Client)

// New creates a new HTTP client
func New(opts ...Option) *Client {
	c := &Client{
		baseClient:     &http.Client{},
		maxRetries:     3,
		retryWaitMin:   1 * time.Second,
		retryWaitMax:   30 * time.Second,
		logger:         &DefaultLogger{},
		retryableCodes: []int{408, 429, 500, 502, 503, 504},
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// WithMaxRetries sets the maximum retry attempts
func WithMaxRetries(maxRetries int) Option {
	return func(c *Client) {
		c.maxRetries = maxRetries
	}
}

// WithRetryWait sets the retry backoff time range
func WithRetryWait(min, max time.Duration) Option {
	return func(c *Client) {
		c.retryWaitMin = min
		c.retryWaitMax = max
	}
}

// WithLogger sets a custom logger
func WithLogger(logger Logger) Option {
	return func(c *Client) {
		c.logger = logger
	}
}

// WithHTTPClient sets the base HTTP client
func WithHTTPClient(client *http.Client) Option {
	return func(c *Client) {
		c.baseClient = client
	}
}

// WithRetryableCodes sets the retryable status codes
func WithRetryableCodes(codes []int) Option {
	return func(c *Client) {
		c.retryableCodes = codes
	}
}

// Do executes an HTTP request with retry support
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	var resp *http.Response
	var err error
	var body []byte

	// Read request body first for retry attempts
	if req.Body != nil {
		body, err = io.ReadAll(req.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read request body: %w", err)
		}
	}

	for i := 0; i <= c.maxRetries; i++ {
		// Reset body for retry
		if body != nil {
			req.Body = io.NopCloser(bytes.NewReader(body))
		}

		// Log request start
		c.logger.Debugf("Attempt %d: %s %s", i+1, req.Method, req.URL.String())

		resp, err = c.baseClient.Do(req)
		if err != nil {
			c.logger.Warnf("Request failed (attempt %d/%d): %v", i+1, c.maxRetries+1, err)
			if i == c.maxRetries {
				return nil, fmt.Errorf("max retries exceeded, last error: %w", err)
			}
			c.sleep(i)
			continue
		}

		// Check if status code is retryable
		if c.shouldRetry(resp.StatusCode) {
			c.logger.Warnf("Received retryable status code %d (attempt %d/%d)",
				resp.StatusCode, i+1, c.maxRetries+1)

			// Drain and close response body
			io.Copy(io.Discard, resp.Body)
			resp.Body.Close()

			if i == c.maxRetries {
				return nil, fmt.Errorf("max retries exceeded, last status: %d", resp.StatusCode)
			}
			c.sleep(i)
			continue
		}

		// Request succeeded
		c.logger.Debugf("Request succeeded (attempt %d/%d)", i+1, c.maxRetries+1)
		return resp, nil
	}

	return nil, fmt.Errorf("max retries exceeded")
}

// shouldRetry checks if the status code is retryable
func (c *Client) shouldRetry(statusCode int) bool {
	for _, code := range c.retryableCodes {
		if statusCode == code {
			return true
		}
	}
	return false
}

// sleep implements exponential backoff
func (c *Client) sleep(retryCount int) {
	wait := c.calculateWait(retryCount)
	c.logger.Debugf("Sleeping %v before retry...", wait)
	time.Sleep(wait)
}

// calculateWait computes the backoff duration
func (c *Client) calculateWait(retryCount int) time.Duration {
	// Exponential backoff algorithm
	wait := float64(c.retryWaitMin) * math.Pow(2, float64(retryCount))

	// Add jitter
	wait = wait * (0.8 + 0.4*rand.Float64())

	// Cap maximum wait time
	if wait > float64(c.retryWaitMax) {
		wait = float64(c.retryWaitMax)
	}

	return time.Duration(wait)
}

// Get sends a GET request
func (c *Client) Get(ctx context.Context, url string, headers http.Header) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	if headers != nil {
		for key, values := range headers {
			for _, value := range values {
				req.Header.Add(key, value)
			}
		}
	}
	return c.Do(req)
}

// Post sends a POST request
func (c *Client) Post(ctx context.Context, url string, contentType string, body io.Reader, headers http.Header) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, "POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	if headers != nil {
		for key, values := range headers {
			for _, value := range values {
				req.Header.Add(key, value)
			}
		}
	}
	return c.Do(req)
}

// Put sends a PUT request
func (c *Client) Put(ctx context.Context, url string, contentType string, body io.Reader, headers http.Header) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, "PUT", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	if headers != nil {
		for key, values := range headers {
			for _, value := range values {
				req.Header.Add(key, value)
			}
		}
	}
	return c.Do(req)
}

// Delete sends a DELETE request
func (c *Client) Delete(ctx context.Context, url string, headers http.Header) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	if headers != nil {
		for key, values := range headers {
			for _, value := range values {
				req.Header.Add(key, value)
			}
		}
	}
	return c.Do(req)
}

// Patch sends a PATCH request
func (c *Client) Patch(ctx context.Context, url string, contentType string, body io.Reader, headers http.Header) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, "PATCH", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	if headers != nil {
		for key, values := range headers {
			for _, value := range values {
				req.Header.Add(key, value)
			}
		}
	}
	return c.Do(req)
}

// Head sends a HEAD request
func (c *Client) Head(ctx context.Context, url string, headers http.Header) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, "HEAD", url, nil)
	if err != nil {
		return nil, err
	}
	if headers != nil {
		for key, values := range headers {
			for _, value := range values {
				req.Header.Add(key, value)
			}
		}
	}
	return c.Do(req)
}

// Options sends an OPTIONS request
func (c *Client) Options(ctx context.Context, url string, headers http.Header) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, "OPTIONS", url, nil)
	if err != nil {
		return nil, err
	}
	if headers != nil {
		for key, values := range headers {
			for _, value := range values {
				req.Header.Add(key, value)
			}
		}
	}
	return c.Do(req)
}
