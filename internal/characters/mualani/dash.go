package mualani

import (
	"github.com/genshinsim/gcsim/pkg/core/action"
)

func (c *char) Dash(p map[string]int) (action.Info, error) {
	if c.nightsoulPoints > 0 {
		c.reduceNightsoulPoints(10)
	}
	return c.Character.Dash(p)
}
