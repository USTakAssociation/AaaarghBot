package ai

import (
	"github.com/USTakAssociation/AaaarghBot/tak"
	"context"
)

type TakPlayer interface {
	GetMove(ctx context.Context, p *tak.Position) tak.Move
}
