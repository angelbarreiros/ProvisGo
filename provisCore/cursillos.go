package provisCore

import (
	"context"
	"github.com/angelbarreiros/ProvisGo/provisEntities"
	"github.com/angelbarreiros/ProvisGo/util"
	"net/http"
	"net/url"
)

func (pe provisExecutor) Cursillos(filterParams *provisEntities.CursillosParams) (*provisEntities.CursillosResponse, *provisEntities.ErrorResponse) {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), pe.defaultTimeout)
	defer cancel()
	resultChan := make(chan util.RequestResult, 1)
	go func() {
		var params url.Values = url.Values{}
		if filterParams != nil {
			params = filterParams.ToURLValues()
		}

		params.Set("installationId", pe.installationId)

		var request *http.Request = pe.config.generateRequest(pe.installationId,
			http.MethodGet, "/api/courses/simple/", params,
			nil)
		request = request.WithContext(ctxWithTimeout)

		var responseArray []provisEntities.Cursillo = make([]provisEntities.Cursillo, 0)
		var responseBody = &responseArray
		result := util.ExecuteRequest(ctxWithTimeout, pe.client, request, pe.config.Debug, responseBody)
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
