package provisCore

import (
	"context"
	"github.com/angelbarreiros/ProvisGo/provisEntities"
	"github.com/angelbarreiros/ProvisGo/util"
	"net/http"
	"net/url"
)

func (pe provisExecutor) Groups(courseGroupId string, dateToConsult string, filterParams *provisEntities.GroupsParams) (*provisEntities.GroupsResponse, *provisEntities.ErrorResponse) {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), pe.defaultTimeout)
	defer cancel()
	resultChan := make(chan util.RequestResult, 1)
	go func() {
		var params = url.Values{}
		if filterParams != nil {
			params = filterParams.ToURLValues()
		}
		params.Set("coursegroupid", courseGroupId)
		params.Set("installationId", pe.installationId)
		params.Set("datetoconsult", dateToConsult)

		var request *http.Request = pe.config.generateRequest(pe.installationId,
			http.MethodGet, "/api/courses/reservation/personlist", params,
			nil)
		request = request.WithContext(ctxWithTimeout)

		// Use AcceptLanguage from filterParams if provided, otherwise default to "en"
		acceptLanguage := "en"
		if filterParams != nil && filterParams.AcceptLanguage != nil {
			acceptLanguage = *filterParams.AcceptLanguage
		}
		request.Header.Add("Accept-Language", acceptLanguage)
		var responseBody = &provisEntities.GroupsResponse{}
		result := util.ExecuteRequest(ctxWithTimeout, pe.client, request, pe.config.Debug, responseBody)
		resultChan <- result
	}()

	select {
	case res := <-resultChan:
		if res.Response != nil {
			var responseGroups = res.Response.(*provisEntities.GroupsResponse)
			return responseGroups, res.Error
		}
		return nil, res.Error
	case <-ctxWithTimeout.Done():
		return nil, &provisEntities.ErrorResponse{
			Code:    http.StatusRequestTimeout,
			Message: "Request timeout: operation cancelled after 30 seconds",
		}
	}
}
