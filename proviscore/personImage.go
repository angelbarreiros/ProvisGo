package proviscore

import (
	"context"
	"net/http"
	"net/url"

	"github.com/angelbarreiros/ProvisGo/provisentities"
	"github.com/angelbarreiros/ProvisGo/util"
)

func (pe provisExecutor) PersonImage(personId string, filterParams *provisentities.PersonImageParams) (*provisentities.PersonImageResponse, *provisentities.ErrorResponse) {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), pe.defaultTimeout)
	defer cancel()
	resultChan := make(chan util.RequestResult, 1)
	go func() {
		var params = url.Values{}
		if filterParams != nil {
			params = filterParams.ToURLValues()
		}
		params.Set("installationid", pe.installationId)
		params.Set("id", personId)

		var request *http.Request = pe.config.generateRequest(pe.installationId,
			http.MethodGet, "/api/person/id/image/", params,
			nil)
		request = request.WithContext(ctxWithTimeout)

		var image string
		result := util.ExecuteRequest(ctxWithTimeout, pe.client, request, pe.config.Debug, &image)
		resultChan <- result
	}()

	select {
	case res := <-resultChan:
		if res.Error == nil {
			image, ok := res.Response.(*string)
			if !ok || image == nil {
				return nil, &provisentities.ErrorResponse{
					Code:    http.StatusInternalServerError,
					Message: "Invalid response format",
				}
			}
			return &provisentities.PersonImageResponse{Image: *image}, nil
		}
		return nil, res.Error
	case <-ctxWithTimeout.Done():
		return nil, &provisentities.ErrorResponse{
			Code:    http.StatusRequestTimeout,
			Message: "Request timeout: operation cancelled after 30 seconds",
		}
	}
}
