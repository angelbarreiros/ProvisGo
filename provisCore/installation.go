package provisCore

import (
	"context"
	"net/http"
	"net/url"
	"provisgo/provisEntities"
	"provisgo/util"
)

func (pe provisExecutor) Installations(filterParams *provisEntities.InstallationsParams) (any, *util.ErrorResponse) {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), pe.defaultTimeout)
	defer cancel()
	resultChan := make(chan util.RequestResult[any], 1)
	go func() {
		var params = url.Values{}
		if filterParams != nil {
			params = filterParams.ToURLValues()
		}
		// Build URL path with optional extended parameter
		urlPath := "/api/installations/byappkey"
		// if filterParams != nil && filterParams.Extended != nil && *filterParams.Extended {
		// 	urlPath += "?extended=true"
		// } else {
		// 	urlPath += "?extended=true" // Default behavior
		// }

		var request *http.Request = pe.config.generateRequest(pe.installationId,
			http.MethodGet, urlPath, params,
			nil)
		request = request.WithContext(ctxWithTimeout)

		result := util.ExecuteRequest[any](ctxWithTimeout, pe.client, request)
		resultChan <- result
	}()

	select {
	case res := <-resultChan:
		if res.Error == nil {
			if workers, ok := res.Response.([]any); ok {
				if workers == nil {
					workers = nil
				}
				return res.Response, res.Error

			}
			return nil, &util.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: "Invalid response format",
			}
		}
		return nil, res.Error
	case <-ctxWithTimeout.Done():
		return nil, &util.ErrorResponse{
			Code:    http.StatusRequestTimeout,
			Message: "Request timeout: operation cancelled after 10 seconds",
		}
	}
}
