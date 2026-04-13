package provisCore

import (
	"context"
	"net/http"
	"net/url"
	"provisgo/provisEntities"
	"provisgo/util"
)

func (pe provisExecutor) Cuotas(filterParams *provisEntities.CuotasParams) (*provisEntities.CursillosResponse, *util.ErrorResponse) {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), pe.defaultTimeout)
	defer cancel()
	resultChan := make(chan util.RequestResult[provisEntities.CursillosResponse], 1)
	go func() {
		var params = url.Values{}
		if filterParams != nil {
			params = filterParams.ToURLValues()
		}
		// params.Set("idInstallation", pe.installationId)
		// // Default values if not provided via filterParams
		// if params.Get("fechaInicio") == "" {
		// 	params.Set("fechaInicio", "2023-12-12T14:46:30")
		// }
		// if params.Get("fechaFin") == "" {
		// 	params.Set("fechaFin", "2027-12-12T14:46:30")
		// }
		// if params.Get("personid") == "" {
		// 	params.Set("personid", "0")
		// }

		var request *http.Request = pe.config.generateRequest(pe.installationId,
			http.MethodGet, "/api/subscriptions/basic/", params,
			nil)
		request = request.WithContext(ctxWithTimeout)

		result := util.ExecuteRequest[provisEntities.CursillosResponse](ctxWithTimeout, pe.client, request)
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
