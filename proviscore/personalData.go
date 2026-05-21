package proviscore

import (
	"context"
	"github.com/angelbarreiros/ProvisGo/provisentities"
	"github.com/angelbarreiros/ProvisGo/util"
	"net/http"
	"net/url"
)

func (pe provisExecutor) Personaldata(personId string, filterParams *provisentities.PersonalDataParams) (*provisentities.FamilyPerson, *provisentities.ErrorResponse) {
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
		var response *provisentities.FamilyPerson = new(provisentities.FamilyPerson)
		result := util.ExecuteRequest(ctxWithTimeout, pe.client, request, pe.config.Debug, response)
		resultChan <- result
	}()

	select {
	case res := <-resultChan:
		if res.Error == nil {
			var response = res.Response.(*provisentities.FamilyPerson)
			return response, res.Error
		}
		return nil, res.Error
	case <-ctxWithTimeout.Done():
		return nil, &provisentities.ErrorResponse{
			Code:    http.StatusRequestTimeout,
			Message: "Request timeout: operation cancelled after 10 seconds",
		}
	}
}
