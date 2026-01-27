package util

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"provisgo/provisEntities"
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

// ExecuteRequest handles common HTTP request execution pattern including error handling and response processing
// It takes a context, http client, request, and a target interface to unmarshal the response into
func ExecuteRequest(ctx context.Context, client *http.Client, request *http.Request, target any) RequestResult {
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

	// If we received a non-success status code, return an error
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return RequestResult{
			Response: nil,
			Error: &provisEntities.ErrorResponse{
				Code:    response.StatusCode,
				Message: "Response: " + string(bodyBytes),
			},
		}
	}

	file, fileErr := os.Create("output.txt")
	if fileErr != nil {
		log.Printf("Failed to create file: %v", fileErr)
		return RequestResult{
			Response: nil,
			Error: &provisEntities.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: "Failed to create file: " + fileErr.Error(),
			},
		}
	}
	defer file.Close()
	// log.Println(string(bodyBytes))
	_, writeErr := file.Write(bodyBytes)
	if writeErr != nil {
		log.Printf("Failed to write to file: %v", writeErr)
		return RequestResult{
			Response: nil,
			Error: &provisEntities.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: "Failed to write to file: " + writeErr.Error(),
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
