package provisCore

import (
	"net/http"
	"provisgo/provisEntities"
	"strings"
	"sync"
	"time"
)

var provisProviderInstace *provisProvider = nil

type provisProvider struct {
	providers *sync.Pool
}
type provisExecutor struct {
	config         *provisConfig
	client         *http.Client
	defaultTimeout time.Duration
	installationId string
}

func Init(cfg *provisConfig) *provisProvider {

	if provisProviderInstace == nil {
		provisProviderInstace = &provisProvider{
			providers: &sync.Pool{
				New: func() any {
					return &provisExecutor{config: cfg, client: http.DefaultClient, defaultTimeout: 30 * time.Second}
				},
			},
		}
	}
	return provisProviderInstace
}
func (pp provisProvider) getExecutor(installationId string) *provisExecutor {
	var executor = pp.providers.Get().(*provisExecutor)
	executor.installationId = installationId
	return executor
}
func (pp provisProvider) putExecutor(executor *provisExecutor) {
	pp.providers.Put(executor)
}
func (pp provisProvider) Close() {
	provisProviderInstace = nil
}
func checkInstallationId(installationId string) *provisEntities.ErrorResponse {
	if strings.TrimSpace(installationId) == "" {
		return &provisEntities.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Installation ID cannot be empty",
		}
	}

	return nil
}
func (pp provisProvider) Cursillos(installationId string, filterParams *provisEntities.CursillosParams) (*provisEntities.CursillosResponse, *provisEntities.ErrorResponse) {
	if err := checkInstallationId(installationId); err != nil {
		return nil, err
	}
	executor := pp.getExecutor(installationId)
	defer pp.putExecutor(executor)
	var response, err = executor.Cursillos(filterParams)
	return response, err
}
func (pp provisProvider) Cuotas(installationId string, filterParams *provisEntities.CuotasParams) (*provisEntities.CursillosResponse, *provisEntities.ErrorResponse) {
	if err := checkInstallationId(installationId); err != nil {
		return nil, err
	}
	executor := pp.getExecutor(installationId)
	defer pp.putExecutor(executor)
	var response, err = executor.Cuotas(filterParams)
	return response, err
}
func (pp provisProvider) Workers(installationId string, filterParams *provisEntities.WorkersParams) (*provisEntities.ProvisWorkers, *provisEntities.ErrorResponse) {
	if err := checkInstallationId(installationId); err != nil {
		return nil, err
	}
	executor := pp.getExecutor(installationId)
	defer pp.putExecutor(executor)
	var response, err = executor.Workers(filterParams)
	return response, err
}
func (pp provisProvider) Personaldata(installationId string, personId string, filterParams *provisEntities.PersonalDataParams) (*provisEntities.FamilyPerson, *provisEntities.ErrorResponse) {
	if err := checkInstallationId(installationId); err != nil {
		return nil, err
	}
	executor := pp.getExecutor(installationId)
	defer pp.putExecutor(executor)
	var response, err = executor.Personaldata(personId, filterParams)
	return response, err
}
func (pp provisProvider) Families(installationId string, personId string, filterParams *provisEntities.FamiliesParams) (*provisEntities.Familie, *provisEntities.ErrorResponse) {
	if err := checkInstallationId(installationId); err != nil {
		return nil, err
	}
	executor := pp.getExecutor(installationId)
	defer pp.putExecutor(executor)
	var response, err = executor.Families(personId, filterParams)
	return response, err
}
func (pp provisProvider) Installations(filterParams *provisEntities.InstallationsParams) (any, *provisEntities.ErrorResponse) {
	executor := pp.getExecutor("")
	defer pp.putExecutor(executor)
	var response, err = executor.Installations(filterParams)
	return response, err
}
func (pp provisProvider) Groups(installationId string, courseGroupId string, dateToConsult string, filterParams *provisEntities.GroupsParams) (*provisEntities.GroupsResponse, *provisEntities.ErrorResponse) {
	if err := checkInstallationId(installationId); err != nil {
		return nil, err
	}
	executor := pp.getExecutor(installationId)
	defer pp.putExecutor(executor)
	var response, err = executor.Groups(courseGroupId, dateToConsult, filterParams)
	return response, err
}
