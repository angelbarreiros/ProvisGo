package proviscore

import (
	"context"
	"net/http"
	"net/url"

	"github.com/angelbarreiros/ProvisGo/provisentities"
	"github.com/angelbarreiros/ProvisGo/util"
)

func (pe provisExecutor) AccessByDate(filterParams *provisentities.AccessByDateParams) (*provisentities.AccessByDateResponse, *provisentities.ErrorResponse) {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), pe.defaultTimeout)
	defer cancel()
	resultChan := make(chan util.RequestResult, 1)
	go func() {
		params := url.Values{}
		if filterParams != nil {
			params = filterParams.ToURLValues()
		}

		request := pe.config.generateRequest(pe.installationId,
			http.MethodGet, "/api/access/bydate/", params,
			nil)
		request = request.WithContext(ctxWithTimeout)

		responseBody := &provisentities.AccessByDateResponse{}
		result := util.ExecuteRequest(ctxWithTimeout, pe.client, request, pe.config.Debug, responseBody)
		resultChan <- result
	}()

	select {
	case res := <-resultChan:
		if res.Response != nil {
			responseAccesses := res.Response.(*provisentities.AccessByDateResponse)
			return responseAccesses, res.Error
		}
		return nil, res.Error
	case <-ctxWithTimeout.Done():
		return nil, &provisentities.ErrorResponse{
			Code:    http.StatusRequestTimeout,
			Message: "Request timeout: operation cancelled after 30 seconds",
		}
	}
}
