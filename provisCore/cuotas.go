package provisCore

import (
	"context"
	"net/http"
	"net/url"
	"provisgo/provisEntities"
	"provisgo/util"
)

func (pe provisExecutor) Cuotas(filterParams *provisEntities.CuotasParams) (*provisEntities.CursillosResponse, *provisEntities.ErrorResponse) {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), pe.defaultTimeout)
	defer cancel()
	resultChan := make(chan util.RequestResult, 1)
	go func() {
		var params = url.Values{}
		if filterParams != nil {
			params = filterParams.ToURLValues()
		}
		params.Set("idInstallation", pe.installationId)
		// Default values if not provided via filterParams
		if params.Get("fechaInicio") == "" {
			params.Set("fechaInicio", "2023-12-12T14:46:30")
		}
		if params.Get("fechaFin") == "" {
			params.Set("fechaFin", "2027-12-12T14:46:30")
		}
		if params.Get("personid") == "" {
			params.Set("personid", "0")
		}

		var request *http.Request = pe.config.generateRequest(pe.installationId,
			http.MethodGet, "/api/enrollments/reservation/", params,
			nil)
		request = request.WithContext(ctxWithTimeout)
		var responseArray []any = make([]any, 0)
		var responseBody = &responseArray
		result := util.ExecuteRequest(ctxWithTimeout, pe.client, request, responseBody)
		resultChan <- result
	}()

	select {
	case res := <-resultChan:
		if res.Error == nil {
			return &provisEntities.CursillosResponse{}, res.Error
		}
		return nil, res.Error
	case <-ctxWithTimeout.Done():
		return nil, &provisEntities.ErrorResponse{
			Code:    http.StatusRequestTimeout,
			Message: "Request timeout: operation cancelled after 10 seconds",
		}
	}
}
