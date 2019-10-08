package handler

import (
	"context"
	"knowledgemap_backend/{{ .PackagePath }}/api"
)

type {{ .ServiceName}}Service struct{}

func (u *{{ .ServiceName}}Service) Ping(ctx context.Context, req *api.Empty ,rsp *api.Empty) error {
	return nil
}
