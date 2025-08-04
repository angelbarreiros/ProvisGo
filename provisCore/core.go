package provisCore

import (
	"context"
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
func (pp provisProvider) Cursillos(ctx context.Context, installationId string) (*provisEntities.CursillosResponse, *provisEntities.ErrorResponse) {
	if err := checkInstallationId(installationId); err != nil {
		return nil, err
	}
	executor := pp.getExecutor(installationId)
	defer pp.putExecutor(executor)
	var response, err = executor.Cursillos(ctx)
	return response, err
}
func (pp provisProvider) Cuotas(ctx context.Context, installationId string) (*provisEntities.CursillosResponse, *provisEntities.ErrorResponse) {
	if err := checkInstallationId(installationId); err != nil {
		return nil, err
	}
	executor := pp.getExecutor(installationId)
	defer pp.putExecutor(executor)
	var response, err = executor.Cuotas(ctx)
	return response, err
}
