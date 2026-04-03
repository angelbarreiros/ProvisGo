package provisCore

import (
	"context"
	"net/http"
	"net/url"
	"provisgo/provisEntities"
	"provisgo/util"
)

func (pe provisExecutor) Personaldata(personId string, filterParams *provisEntities.PersonalDataParams) (*provisEntities.FamilyPerson, *provisEntities.ErrorResponse) {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), pe.defaultTimeout)
	defer cancel()
	resultChan := make(chan util.RequestResult, 1)
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
		var response *provisEntities.FamilyPerson = new(provisEntities.FamilyPerson)
		result := util.ExecuteRequest(ctxWithTimeout, pe.client, request, response)
		resultChan <- result
	}()

	select {
	case res := <-resultChan:
		if res.Error == nil {
			var response = res.Response.(*provisEntities.FamilyPerson)
			return response, res.Error
		}
		return nil, res.Error
	case <-ctxWithTimeout.Done():
		return nil, &provisEntities.ErrorResponse{
			Code:    http.StatusRequestTimeout,
			Message: "Request timeout: operation cancelled after 10 seconds",
		}
	}
}
