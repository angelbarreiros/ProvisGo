package provisCore

import (
	"context"
	"net/http"
	"net/url"
	"provisgo/provisEntities"
	"provisgo/util"
)

func (pe provisExecutor) Personaldata(personId string, filterParams *provisEntities.PersonalDataParams) (*provisEntities.FamilyPerson, *util.ErrorResponse) {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), pe.defaultTimeout)
	defer cancel()
	resultChan := make(chan util.RequestResult[provisEntities.FamilyPerson], 1)
	go func() {
		var params = url.Values{}
		if filterParams != nil {
			params = filterParams.ToURLValues()
		}
		params.Set("installationid", pe.installationId)
		params.Set("idPersona", personId)
		var request *http.Request = pe.config.generateRequest(pe.installationId,
			http.MethodGet, "/api/person/personaldata", params,
			nil)
		request = request.WithContext(ctxWithTimeout)

		result := util.ExecuteRequest[provisEntities.FamilyPerson](ctxWithTimeout, pe.client, request)
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
