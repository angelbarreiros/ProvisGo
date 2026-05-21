package proviscore

import (
	"context"
	"net/http"
	"net/url"

	"github.com/angelbarreiros/ProvisGo/provisentities"
	"github.com/angelbarreiros/ProvisGo/util"
)

func (pe provisExecutor) Installations(filterParams *provisentities.InstallationsParams) (*provisentities.InstallationsResponse, *provisentities.ErrorResponse) {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), pe.defaultTimeout)
	defer cancel()
	resultChan := make(chan util.RequestResult, 1)
	go func() {
		var params = url.Values{}
		if filterParams != nil {
			params = filterParams.ToURLValues()
		}
		// Build URL path with optional extended parameter
		urlPath := "/api/installations/byappkey"

		var request *http.Request = pe.config.generateRequest(pe.installationId,
			http.MethodGet, urlPath, params,
			nil)
		request = request.WithContext(ctxWithTimeout)
		var responseArray = make([]provisentities.Installation, 0)
		result := util.ExecuteRequest(ctxWithTimeout, pe.client, request, pe.config.Debug, &responseArray)
		resultChan <- result
	}()

	select {
	case res := <-resultChan:
		if res.Error == nil {
			if installations, ok := res.Response.(*[]provisentities.Installation); ok {
				if installations == nil {
					installations = &[]provisentities.Installation{}
				}
				return &provisentities.InstallationsResponse{
					Installations: *installations,
				}, res.Error
			}
			return nil, &provisentities.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: "Invalid response format",
			}
		}
		return nil, res.Error
	case <-ctxWithTimeout.Done():
		return nil, &provisentities.ErrorResponse{
			Code:    http.StatusRequestTimeout,
			Message: "Request timeout: operation cancelled after 10 seconds",
		}
	}
}
