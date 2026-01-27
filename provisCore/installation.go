package provisCore

import (
	"context"
	"net/http"
	"net/url"
	"provisgo/provisEntities"
	"provisgo/util"
)

func (pe provisExecutor) Installations() (any, *provisEntities.ErrorResponse) {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), pe.defaultTimeout)
	defer cancel()
	resultChan := make(chan util.RequestResult, 1)
	go func() {
		var params = url.Values{}
		// params.Set("installationid", pe.installationId)

		var request *http.Request = pe.config.generateRequest(pe.installationId,
			http.MethodGet, "/api/installations/byappkey?extended=true", params,
			nil)
		request = request.WithContext(ctxWithTimeout)
		var responseArray any
		result := util.ExecuteRequest(ctxWithTimeout, pe.client, request, &responseArray)
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
			return nil, &provisEntities.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: "Invalid response format",
			}
		}
		return nil, res.Error
	case <-ctxWithTimeout.Done():
		return nil, &provisEntities.ErrorResponse{
			Code:    http.StatusRequestTimeout,
			Message: "Request timeout: operation cancelled after 10 seconds",
		}
	}
}
