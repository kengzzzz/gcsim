package mualani

import (
	"github.com/genshinsim/gcsim/pkg/core/action"
)

func (c *char) Jump(p map[string]int) (action.Info, error) {
	if c.nightsoulPoints > 0 {
		if c.Core.Player.LastAction.Type == action.ActionDash {
			c.reduceNightsoulPoints(14) // total 24, 10 from dash, 14 from dash jump
		} else {
			c.reduceNightsoulPoints(2)
		}
	}
	return c.Character.Jump(p)
}
