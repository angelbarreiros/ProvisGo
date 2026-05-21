package util

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"provisgo/provisEntities"
	"sort"
	"strings"
)

// RequestResult encapsulates the possible outcomes of an API request
type RequestResult struct {
	Response interface{}
	Error    *provisEntities.ErrorResponse
}
type ProvisErrorResponse struct {
	Code          int    `json:"code"`
	Message       string `json:"message"`
	CustomMessage struct {
		Id         any            `json:"Id"`
		Parameters map[string]any `json:"Parameters"`
	} `json:"CustomMessage"`
	Details string `json:"Details"`
	Result  string `json:"Result"`
}

// ExecuteRequest handles common HTTP request execution pattern including error handling and response processing.
// It takes a context, http client, request, debug flag, and a target interface to unmarshal the response into.
func ExecuteRequest(ctx context.Context, client *http.Client, request *http.Request, debug bool, target any) RequestResult {
	if debug {
		curlCommand, curlErr := formatCurlCommand(request)
		if curlErr != nil {
			log.Printf("Failed to format cURL command: %v", curlErr)
		} else {
			fmt.Println(curlCommand)
		}
	}

	response, clientErr := client.Do(request)
	if clientErr != nil {
		return RequestResult{
			Response: nil,
			Error: &provisEntities.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: "Failed to execute request: " + clientErr.Error(),
			},
		}
	}
	defer response.Body.Close()

	if debug {
		fmt.Printf("Response status: %s\n", response.Status)
		for _, key := range sortedHeaderKeys(response.Header) {
			for _, value := range response.Header.Values(key) {
				fmt.Printf("Response header %s: %s\n", key, value)
			}
		}
	}

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return RequestResult{
			Response: nil,
			Error: &provisEntities.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: "Failed to read response body: " + err.Error(),
			},
		}
	}
	if debug {
		fmt.Printf("Body: %s\n", string(bodyBytes))
	}

	// If we received a non-success status code, return an error after saving the body.
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return RequestResult{
			Response: nil,
			Error: &provisEntities.ErrorResponse{
				Code:    response.StatusCode,
				Message: "HTTP " + response.Status + " response: " + string(bodyBytes),
			},
		}
	}
	if target != nil && len(bodyBytes) > 0 {
		err = json.Unmarshal(bodyBytes, target)
		if err != nil {
			return RequestResult{
				Response: nil,
				Error: &provisEntities.ErrorResponse{
					Code:    http.StatusInternalServerError,
					Message: "Failed to unmarshal response: " + err.Error(),
				},
			}
		}
		return RequestResult{
			Response: target,
			Error:    nil,
		}
	}

	return RequestResult{
		Response: nil,
		Error:    nil,
	}
}

func formatCurlCommand(request *http.Request) (string, error) {
	var builder strings.Builder
	builder.WriteString("curl -X ")
	builder.WriteString(request.Method)
	builder.WriteString(" \\\n'")
	builder.WriteString(shellSingleQuote(request.URL.String()))
	builder.WriteString("'")

	for _, key := range sortedHeaderKeys(request.Header) {
		for _, value := range request.Header.Values(key) {
			builder.WriteString(" \\\n-H '")
			builder.WriteString(shellSingleQuote(key))
			builder.WriteString(": ")
			builder.WriteString(shellSingleQuote(value))
			builder.WriteString("'")
		}
	}

	if request.Body != nil {
		bodyBytes, err := io.ReadAll(request.Body)
		if err != nil {
			return "", err
		}
		request.Body.Close()
		request.Body = io.NopCloser(bytes.NewReader(bodyBytes))

		if len(bodyBytes) > 0 {
			builder.WriteString(" \\\n-d '")
			builder.WriteString(shellSingleQuote(string(bodyBytes)))
			builder.WriteString("'")
		}
	}

	return builder.String(), nil
}

func sortedHeaderKeys(header http.Header) []string {
	preferred := []string{"Authorization", "Cache-Control", "Content-Type", "Accept", "User-Agent", "Timestamp"}
	keys := make([]string, 0, len(header))
	seen := make(map[string]bool, len(header))

	for _, key := range preferred {
		if _, ok := header[key]; ok {
			keys = append(keys, key)
			seen[key] = true
		}
	}

	var remaining []string
	for key := range header {
		if !seen[key] {
			remaining = append(remaining, key)
		}
	}
	sort.Strings(remaining)

	return append(keys, remaining...)
}

func shellSingleQuote(value string) string {
	return strings.ReplaceAll(value, "'", "'\"'\"'")
}
