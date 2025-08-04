package provisCore

import (
	"context"
	"net/http"
	"net/url"
	"provisgo/provisEntities"
	"provisgo/util"
)

func (pe provisExecutor) Cuotas(ctx context.Context) (*provisEntities.CursillosResponse, *provisEntities.ErrorResponse) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, pe.defaultTimeout)
	defer cancel()
	resultChan := make(chan util.RequestResult, 1)
	go func() {
		var params = url.Values{}
		params.Set("idInstallation", pe.installationId)
		params.Set("fechaInicio", "2024-11-12T14:46:30")
		params.Set("fechaFin", "2024-12-16T14:46:30")

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
