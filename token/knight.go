package token

import (
	"github.com/cragcraig/chess/core"
)

type knight struct {
	protoToken
}

func (t *knight) GetMoves(curPos core.Position, b *Board) []core.Position {
	pos := []core.Position{
		curPos.Add(core.Offset{1, 2}),
		curPos.Add(core.Offset{2, 1}),
		curPos.Add(core.Offset{1, -2}),
		curPos.Add(core.Offset{2, -1}),
		curPos.Add(core.Offset{-1, 2}),
		curPos.Add(core.Offset{-2, 1}),
		curPos.Add(core.Offset{-1, -2}),
		curPos.Add(core.Offset{-2, -1}),
	}
	return b.trimInvalidPos(pos, t.GetPlayer())
}

func Knight(p core.Player) Token {
	return &knight{protoToken{p, 0, 'N', "knight", -1}}
}
