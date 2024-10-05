package passage

import (
	"context"
	"math/rand"

	"github.com/volvofixthis/pow-server/internal/core/models"
)

var passages []string = []string{
	"The beginning of wisdom is this: Get wisdom. Though it cost all you have, get understanding.",
	"When pride comes, then comes disgrace, but with humility comes wisdom.",
	"How much better to get wisdom than gold, to get insight rather than silver!",
	"Even fools are thought wise if they keep silent, and discerning if they hold their tongues.",
	"The one who gets wisdom loves life; the one who cherishes understanding will soon prosper.",
	"Where there is strife, there is pride, but wisdom is found in those who take advice.",
}

func NewPassageRepository() *PassageRepository {
	return &PassageRepository{}
}

type PassageRepository struct {
}

func (pr *PassageRepository) Get(ctx context.Context) (*models.Passage, error) {
	randomIndex := rand.Intn(len(passages))
	return &models.Passage{Text: passages[randomIndex]}, nil
}
