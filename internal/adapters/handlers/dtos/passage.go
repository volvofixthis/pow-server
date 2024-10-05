package dtos

import (
	"github.com/volvofixthis/pow-server/internal/core/models"
)

type PassageReq struct {
	Hash []byte `json:"hash"`
}

func (pr *PassageReq) ToDomain() *models.PowResult {
	return &models.PowResult{Hash: pr.Hash}
}

type PassageResp struct {
	Text string `json:"text"`
}

func (ptr *PassageResp) FromDomain(passage *models.Passage) {
	ptr.Text = passage.Text
}
