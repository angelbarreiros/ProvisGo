package proviscore

import (
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/angelbarreiros/ProvisGo/provisentities"
)

var provisProviderInstace *ProvisProvider = nil

type ProvisProvider struct {
	providers *sync.Pool
}
type provisExecutor struct {
	config         *provisConfig
	client         *http.Client
	defaultTimeout time.Duration
	installationId string
}

func Init(cfg *provisConfig) *ProvisProvider {

	if provisProviderInstace == nil {
		provisProviderInstace = &ProvisProvider{
			providers: &sync.Pool{
				New: func() any {
					return &provisExecutor{config: cfg, client: http.DefaultClient, defaultTimeout: 30 * time.Second}
				},
			},
		}
	}
	return provisProviderInstace
}
func (pp ProvisProvider) getExecutor(installationId string) *provisExecutor {
	var executor = pp.providers.Get().(*provisExecutor)
	executor.installationId = installationId
	return executor
}
func (pp ProvisProvider) putExecutor(executor *provisExecutor) {
	pp.providers.Put(executor)
}
func (pp ProvisProvider) Close() {
	provisProviderInstace = nil
}
func checkInstallationId(installationId string) *provisentities.ErrorResponse {
	if strings.TrimSpace(installationId) == "" {
		return &provisentities.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Installation ID cannot be empty",
		}
	}

	return nil
}
func (pp ProvisProvider) Cursillos(installationId string, filterParams *provisentities.CursillosParams) (*provisentities.CursillosResponse, *provisentities.ErrorResponse) {
	if err := checkInstallationId(installationId); err != nil {
		return nil, err
	}
	executor := pp.getExecutor(installationId)
	defer pp.putExecutor(executor)
	var response, err = executor.Cursillos(filterParams)
	return response, err
}
func (pp ProvisProvider) Cuotas(installationId string, filterParams *provisentities.CuotasParams) (*provisentities.CursillosResponse, *provisentities.ErrorResponse) {
	if err := checkInstallationId(installationId); err != nil {
		return nil, err
	}
	executor := pp.getExecutor(installationId)
	defer pp.putExecutor(executor)
	var response, err = executor.Cuotas(filterParams)
	return response, err
}
func (pp ProvisProvider) Workers(installationId string, filterParams *provisentities.WorkersParams) (*provisentities.ProvisWorkers, *provisentities.ErrorResponse) {
	if err := checkInstallationId(installationId); err != nil {
		return nil, err
	}
	executor := pp.getExecutor(installationId)
	defer pp.putExecutor(executor)
	var response, err = executor.Workers(filterParams)
	return response, err
}
func (pp ProvisProvider) Personaldata(installationId string, personId string, filterParams *provisentities.PersonalDataParams) (*provisentities.FamilyPerson, *provisentities.ErrorResponse) {
	if err := checkInstallationId(installationId); err != nil {
		return nil, err
	}
	executor := pp.getExecutor(installationId)
	defer pp.putExecutor(executor)
	var response, err = executor.Personaldata(personId, filterParams)
	return response, err
}
func (pp ProvisProvider) Families(installationId string, personId string, filterParams *provisentities.FamiliesParams) (*provisentities.Familie, *provisentities.ErrorResponse) {
	if err := checkInstallationId(installationId); err != nil {
		return nil, err
	}
	executor := pp.getExecutor(installationId)
	defer pp.putExecutor(executor)
	var response, err = executor.Families(personId, filterParams)
	return response, err
}
func (pp ProvisProvider) Installations(installationId string, filterParams *provisentities.InstallationsParams) (any, *provisentities.ErrorResponse) {
	executor := pp.getExecutor(installationId)
	defer pp.putExecutor(executor)
	var response, err = executor.Installations(filterParams)
	return response, err
}
func (pp ProvisProvider) Groups(installationId string, courseGroupId string, dateToConsult string, filterParams *provisentities.GroupsParams) (*provisentities.GroupsResponse, *provisentities.ErrorResponse) {
	if err := checkInstallationId(installationId); err != nil {
		return nil, err
	}
	executor := pp.getExecutor(installationId)
	defer pp.putExecutor(executor)
	var response, err = executor.Groups(courseGroupId, dateToConsult, filterParams)
	return response, err
}
