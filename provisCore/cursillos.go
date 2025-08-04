package provisCore

import (
	"context"
	"net/http"
	"net/url"
	"provisgo/provisEntities"
	"provisgo/util"
)

func (pe provisExecutor) Cursillos(ctx context.Context) (*provisEntities.CursillosResponse, *provisEntities.ErrorResponse) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, pe.defaultTimeout)
	defer cancel()
	resultChan := make(chan util.RequestResult, 1)
	go func() {
		var params = url.Values{}
		params.Set("installationId", pe.installationId)
		var request *http.Request = pe.config.GenerateRequest(pe.installationId,
			"/api/courses/simple/", params, http.MethodGet,
			nil)
		request = request.WithContext(ctxWithTimeout)
		var responseArray []provisEntities.Cursillo = make([]provisEntities.Cursillo, 0)
		var responseBody = &responseArray
		result := util.ExecuteRequest(ctxWithTimeout, pe.client, request, responseBody)
		resultChan <- result
	}()

	select {
	case res := <-resultChan:
		if res.Response != nil {
			var responseCursillos = res.Response.(*[]provisEntities.Cursillo)
			return &provisEntities.CursillosResponse{Cursillos: *responseCursillos}, res.Error
		}
		return nil, res.Error
	case <-ctxWithTimeout.Done():
		return nil, &provisEntities.ErrorResponse{
			Code:    http.StatusRequestTimeout,
			Message: "Request timeout: operation cancelled after 10 seconds",
		}
	}
}
