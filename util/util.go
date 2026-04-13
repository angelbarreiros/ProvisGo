package util

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// ErrorResponse represents an error response from the API
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// RequestResult encapsulates the possible outcomes of an API request
type RequestResult[T any] struct {
	Response T
	Error    *ErrorResponse
}

// ExecuteRequest handles common HTTP request execution pattern including error handling and response processing
// It takes a context, http client, request, and returns a typed RequestResult
func ExecuteRequest[T any](ctx context.Context, client *http.Client, request *http.Request) RequestResult[T] {
	var zero T
	log.Println(request.URL)

	// Read and store the body (if present), then restore it
	var bodyContent string
	if request.Body != nil {
		bodyBytes, err := io.ReadAll(request.Body)
		if err == nil {
			bodyContent = string(bodyBytes)
			// Restore the body so it can be read again by client.Do
			request.Body = io.NopCloser(bytes.NewReader(bodyBytes))
		}
	}

	// Build curl command
	var curlBuilder strings.Builder
	curlBuilder.WriteString(fmt.Sprintf("curl -X %s '%s'", request.Method, request.URL.String()))

	// Add headers
	for key, values := range request.Header {
		for _, value := range values {
			curlBuilder.WriteString(fmt.Sprintf(" -H '%s: %s'", key, value))
		}
	}

	// Add body if present
	if bodyContent != "" {
		curlBuilder.WriteString(fmt.Sprintf(" -d '%s'", bodyContent))
	}

	curlCommand := curlBuilder.String()

	// Print to stdout
	fmt.Println(curlCommand)

	response, clientErr := client.Do(request)
	if clientErr != nil {
		return RequestResult[T]{
			Response: zero,
			Error: &ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: "Failed to execute request: " + clientErr.Error(),
			},
		}
	}
	defer response.Body.Close()

	responseBodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return RequestResult[T]{
			Response: zero,
			Error: &ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: "Failed to read response body: " + err.Error(),
			},
		}
	}

	responseBody := string(responseBodyBytes)

	// Append to file with curl command and response body
	file, err := os.OpenFile("calls.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		defer file.Close()
		file.WriteString(curlCommand + "\n")
		file.WriteString("Response:" + response.Status + " " + responseBody + "\n")
		file.WriteString("\n")
	}

	// If we received a non-success status code, return an error
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return RequestResult[T]{
			Response: zero,
			Error: &ErrorResponse{
				Code:    response.StatusCode,
				Message: "Response: " + responseBody,
			},
		}
	}

	// Only try to unmarshal if we have response body
	if len(responseBodyBytes) > 0 {
		var target T
		err = json.Unmarshal(responseBodyBytes, &target)
		if err != nil {
			return RequestResult[T]{
				Response: zero,
				Error: &ErrorResponse{
					Code:    http.StatusInternalServerError,
					Message: "Failed to unmarshal response: " + err.Error(),
				},
			}
		}
		return RequestResult[T]{
			Response: target,
			Error:    nil,
		}
	}

	return RequestResult[T]{
		Response: zero,
		Error:    nil,
	}
}

func AddQueryParam(name string, value *string, values *url.Values) {
	if value != nil && *value != "" {
		values.Add(name, *value)
	}

}
