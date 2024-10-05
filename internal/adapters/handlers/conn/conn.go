package conn

import (
	"context"
	"encoding/json"

	"github.com/volvofixthis/pow-server/internal/adapters/handlers/dtos"
	"github.com/volvofixthis/pow-server/internal/core/ports"
)

func NewConnAdapter(powService ports.PowService, passageService ports.PassageService) *ConnAdapter {
	return &ConnAdapter{
		powService:     powService,
		passageService: passageService,
	}
}

type ConnAdapter struct {
	powService     ports.PowService
	passageService ports.PassageService
}

func (ca *ConnAdapter) Handle(ctx context.Context, conn ports.Conn) error {
	defer conn.Close()

	helloReq := &dtos.PowHelloReq{}
	err := json.NewDecoder(conn).Decode(helloReq)
	if err != nil {
		return err
	}

	if helloReq.State == dtos.RequestState {
		powTaskResp := &dtos.PowTaskResp{}
		task, err := ca.powService.Create(ctx)
		if err != nil {
			return err
		}
		powTaskResp.FromDomain(task)
		body, err := json.Marshal(powTaskResp)
		if err != nil {
			return err
		}
		if _, err := conn.Write(body); err != nil {
			return err
		}
	}

	passageReq := &dtos.PassageReq{}
	err = json.NewDecoder(conn).Decode(passageReq)
	if err != nil {
		return err
	}

	if err := ca.powService.Verify(ctx, passageReq.ToDomain()); err != nil {
		return err
	}

	passageResp := &dtos.PassageResp{}
	passage, err := ca.passageService.Get(ctx)
	if err != nil {
		return err
	}
	passageResp.FromDomain(passage)
	body, err := json.Marshal(passageResp)
	if err != nil {
		return err
	}
	if _, err := conn.Write(body); err != nil {
		return err
	}

	return nil
}
