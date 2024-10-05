package http

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/volvofixthis/pow-server/internal/adapters/handlers/dtos"
	"github.com/volvofixthis/pow-server/internal/core/ports"
)

func NewPowTaskAdapter(powService ports.PowService) *PowTaskAdapter {
	return &PowTaskAdapter{powService: powService}
}

type PowTaskAdapter struct {
	powService ports.PowService
}

func (pjh *PowTaskAdapter) CreateTask(c echo.Context) error {
	task, err := pjh.powService.Create(context.Background())
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	resp := &dtos.PowTaskResp{}
	resp.FromDomain(task)
	return c.JSON(http.StatusOK, resp)
}
