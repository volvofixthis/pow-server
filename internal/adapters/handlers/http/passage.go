package http

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/volvofixthis/pow-server/internal/adapters/handlers/dtos"
	"github.com/volvofixthis/pow-server/internal/core/ports"
)

func NewPassageAdapter(passageService ports.PassageService, powService ports.PowService) *PassageAdapter {
	return &PassageAdapter{
		passageService: passageService,
		powService:     powService,
	}
}

type PassageAdapter struct {
	powService     ports.PowService
	passageService ports.PassageService
}

func (pa *PassageAdapter) GetPassage(c echo.Context) error {
	req := dtos.PassageReq{}
	if err := c.Bind(&req); err != nil {
		return err
	}
	if err := pa.powService.Verify(context.Background(), req.ToDomain()); err != nil {
		return c.String(http.StatusForbidden, err.Error())
	}
	passage, err := pa.passageService.Get(context.Background())
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	resp := &dtos.PassageResp{}
	resp.FromDomain(passage)
	return c.JSON(http.StatusOK, resp)
}
