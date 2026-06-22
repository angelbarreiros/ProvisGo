package proviscore

import (
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/angelbarreiros/ProvisGo/provisentities"
)

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
	return &ProvisProvider{
		providers: &sync.Pool{
			New: func() any {
				return &provisExecutor{config: cfg, client: http.DefaultClient, defaultTimeout: 30 * time.Second}
			},
		},
	}
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
func (pp ProvisProvider) PersonImage(installationId string, personId string, filterParams *provisentities.PersonImageParams) (*provisentities.PersonImageResponse, *provisentities.ErrorResponse) {
	if err := checkInstallationId(installationId); err != nil {
		return nil, err
	}
	executor := pp.getExecutor(installationId)
	defer pp.putExecutor(executor)
	var response, err = executor.PersonImage(personId, filterParams)
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
func (pp ProvisProvider) Installations(installationId string, filterParams *provisentities.InstallationsParams) (*provisentities.InstallationsResponse, *provisentities.ErrorResponse) {
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
func (pp ProvisProvider) AccessByDate(installationId string, filterParams *provisentities.AccessByDateParams) (*provisentities.AccessByDateResponse, *provisentities.ErrorResponse) {
	if err := checkInstallationId(installationId); err != nil {
		return nil, err
	}
	executor := pp.getExecutor(installationId)
	defer pp.putExecutor(executor)
	var response, err = executor.AccessByDate(filterParams)
	return response, err
}
