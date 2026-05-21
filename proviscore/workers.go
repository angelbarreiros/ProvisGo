package proviscore

import (
	"context"
	"github.com/angelbarreiros/ProvisGo/provisentities"
	"github.com/angelbarreiros/ProvisGo/util"
	"net/http"
	"net/url"
)

func (pe provisExecutor) Workers(filterParams *provisentities.WorkersParams) (*provisentities.ProvisWorkers, *provisentities.ErrorResponse) {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), pe.defaultTimeout)
	defer cancel()
	resultChan := make(chan util.RequestResult, 1)
	go func() {
		var params = url.Values{}
		if filterParams != nil {
			params = filterParams.ToURLValues()
		}
		params.Set("installationid", pe.installationId)

		var request *http.Request = pe.config.generateRequest(pe.installationId,
			http.MethodGet, "/api/staff/list/", params,
			nil)
		request = request.WithContext(ctxWithTimeout)
		var responseArray = make([]*provisentities.ProvisWorker, 0)
		result := util.ExecuteRequest(ctxWithTimeout, pe.client, request, pe.config.Debug, &responseArray)
		resultChan <- result
	}()

	select {
	case res := <-resultChan:
		if res.Error == nil {
			if workers, ok := res.Response.(*[]*provisentities.ProvisWorker); ok {
				if workers == nil {
					workers = &[]*provisentities.ProvisWorker{}
				}
				return &provisentities.ProvisWorkers{
					Workers: *workers,
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
