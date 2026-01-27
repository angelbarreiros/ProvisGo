package provisCore

import (
	"context"
	"net/http"
	"net/url"
	"provisgo/provisEntities"
	"provisgo/util"
)

func (pe provisExecutor) Workers() (*provisEntities.ProvisWorkers, *provisEntities.ErrorResponse) {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), pe.defaultTimeout)
	defer cancel()
	resultChan := make(chan util.RequestResult, 1)
	go func() {
		var params = url.Values{}
		params.Set("installationid", pe.installationId)

		var request *http.Request = pe.config.generateRequest(pe.installationId,
			http.MethodGet, "/api/staff/list/", params,
			nil)
		request = request.WithContext(ctxWithTimeout)
		var responseArray = make([]*provisEntities.ProvisWorker, 0)
		result := util.ExecuteRequest(ctxWithTimeout, pe.client, request, &responseArray)
		resultChan <- result
	}()

	select {
	case res := <-resultChan:
		if res.Error == nil {
			if workers, ok := res.Response.(*[]*provisEntities.ProvisWorker); ok {
				if workers == nil {
					workers = &[]*provisEntities.ProvisWorker{}
				}
				return &provisEntities.ProvisWorkers{
					Workers: *workers,
				}, res.Error
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
