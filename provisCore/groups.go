package provisCore

import (
	"context"
	"net/http"
	"net/url"
	"provisgo/provisEntities"
	"provisgo/util"
)

func (pe provisExecutor) Groups(courseGroupId string, dateToConsult string, filterParams *provisEntities.GroupsParams) (*provisEntities.GroupsResponse, *util.ErrorResponse) {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), pe.defaultTimeout)
	defer cancel()
	resultChan := make(chan util.RequestResult[provisEntities.GroupsResponse], 1)
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
		result := util.ExecuteRequest[provisEntities.GroupsResponse](ctxWithTimeout, pe.client, request)
		resultChan <- result
	}()

	select {
	case res := <-resultChan:

		return &res.Response, nil
	case <-ctxWithTimeout.Done():
		return nil, &util.ErrorResponse{
			Code:    http.StatusRequestTimeout,
			Message: "Request timeout: operation cancelled after 30 seconds",
		}
	}
}
