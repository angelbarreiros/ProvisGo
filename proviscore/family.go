package proviscore

import (
	"context"
	"github.com/angelbarreiros/ProvisGo/provisentities"
	"github.com/angelbarreiros/ProvisGo/util"
	"net/http"
	"net/url"
)

func (pe provisExecutor) Families(clientID string, filterParams *provisentities.FamiliesParams) (*provisentities.Familie, *provisentities.ErrorResponse) {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), pe.defaultTimeout)
	defer cancel()
	resultChan := make(chan util.RequestResult, 1)
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
		var responseArray = make([]*provisentities.FamilyPerson, 0)
		result := util.ExecuteRequest(ctxWithTimeout, pe.client, request, pe.config.Debug, &responseArray)
		resultChan <- result
	}()

	select {
	case res := <-resultChan:
		if res.Error == nil {
			if persons, ok := res.Response.(*[]*provisentities.FamilyPerson); ok {
				if persons == nil {
					persons = &[]*provisentities.FamilyPerson{}
				}
				return &provisentities.Familie{
					FamilyPersons: *persons,
				}, res.Error
			}
			return nil, &provisentities.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: "Invalid response format",
			}
		}
		return nil, res.Error
	case <-ctxWithTimeout.Done():
		return nil, &provisentities.ErrorResponse{
			Code:    http.StatusRequestTimeout,
			Message: "Request timeout: operation cancelled after 10 seconds",
		}
	}
}
