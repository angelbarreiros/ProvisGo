package provisCore

import (
	"context"
	"net/http"
	"net/url"
	"provisgo/provisEntities"
	"provisgo/util"
)

func (pe provisExecutor) Families(personId string) (*provisEntities.Familie, *provisEntities.ErrorResponse) {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), pe.defaultTimeout)
	defer cancel()
	resultChan := make(chan util.RequestResult, 1)
	go func() {
		var params = url.Values{}
		params.Set("installationId", pe.installationId)
		params.Set("clientID", "18")

		var request *http.Request = pe.config.generateRequest(pe.installationId,
			http.MethodGet, "/api/person/family/byperson/", params,
			nil)
		request = request.WithContext(ctxWithTimeout)
		var responseArray = make([]*provisEntities.FamilyPerson, 0)
		result := util.ExecuteRequest(ctxWithTimeout, pe.client, request, &responseArray)
		resultChan <- result
	}()

	select {
	case res := <-resultChan:
		if res.Error == nil {
			if persons, ok := res.Response.(*[]*provisEntities.FamilyPerson); ok {
				if persons == nil {
					persons = &[]*provisEntities.FamilyPerson{}
				}
				return &provisEntities.Familie{
					FamilyPersons: *persons,
				}, res.Error
			}
			return nil, &provisEntities.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: "Invalid response format",
			}
		}
		return nil, res.Error
	case <-ctxWithTimeout.Done():
		return nil, &provisEntities.ErrorResponse{
			Code:    http.StatusRequestTimeout,
			Message: "Request timeout: operation cancelled after 10 seconds",
		}
	}
}
