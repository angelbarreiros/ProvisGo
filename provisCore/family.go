package provisCore

import (
	"context"
	"net/http"
	"net/url"
	"provisgo/provisEntities"
	"provisgo/util"
)

func (pe provisExecutor) Families(clientID string, filterParams *provisEntities.FamiliesParams) (*provisEntities.Familie, *util.ErrorResponse) {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), pe.defaultTimeout)
	defer cancel()
	resultChan := make(chan util.RequestResult[provisEntities.Familie], 1)
	go func() {
		var params = url.Values{}
		if filterParams != nil {
			params = filterParams.ToURLValues()
		}
		params.Set("installationId", pe.installationId)
		params.Set("clientID", clientID)

		var request *http.Request = pe.config.generateRequest(pe.installationId,
			http.MethodGet, "/api/person/family/byperson/", params,
			nil)
		request = request.WithContext(ctxWithTimeout)

		result := util.ExecuteRequest[provisEntities.Familie](ctxWithTimeout, pe.client, request)
		resultChan <- result
	}()

	select {
	case res := <-resultChan:

		return &res.Response, res.Error
	case <-ctxWithTimeout.Done():
		return nil, &util.ErrorResponse{
			Code:    http.StatusRequestTimeout,
			Message: "Request timeout: operation cancelled after 10 seconds",
		}
	}
}
