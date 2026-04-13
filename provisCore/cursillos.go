package provisCore

import (
	"context"
	"net/http"
	"net/url"
	"provisgo/provisEntities"
	"provisgo/util"
)

func (pe provisExecutor) Cursillos(filterParams *provisEntities.CursillosParams) (*provisEntities.CursillosResponse, *util.ErrorResponse) {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), pe.defaultTimeout)
	defer cancel()
	resultChan := make(chan util.RequestResult[[]provisEntities.Cursillo], 1)
	go func() {
		var params = url.Values{}
		if filterParams != nil {
			params = filterParams.ToURLValues()
		}
		params.Set("installationId", pe.installationId)

		var request *http.Request = pe.config.generateRequest(pe.installationId,
			http.MethodGet, "/api/courses/simple/", params,
			nil)
		request = request.WithContext(ctxWithTimeout)

		result := util.ExecuteRequest[[]provisEntities.Cursillo](ctxWithTimeout, pe.client, request)
		resultChan <- result
	}()

	select {
	case res := <-resultChan:
		return &provisEntities.CursillosResponse{
			Cursillos: res.Response,
		}, res.Error
	case <-ctxWithTimeout.Done():
		return nil, &util.ErrorResponse{
			Code:    http.StatusRequestTimeout,
			Message: "Request timeout: operation cancelled after 10 seconds",
		}
	}
}
