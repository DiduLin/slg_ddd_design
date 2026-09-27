package march

import (
	"testing"
	"time"

	"github.com/DiduLin/slg_ddd_design/map/infra/test"
	"github.com/DiduLin/slg_ddd_design/types"
)

func TestMarch_StartMarchStay(t *testing.T) {
	clock := test.NewManualClock(time.Now())
	march := NewMarch(1, &test.PrintEventbus{}, clock, nil)

	march.StartMarchStay(time.Second*6, types.NewPoint(2, 2), []*types.Point{
		types.NewPoint(0, 0),
		types.NewPoint(1, 1),
		types.NewPoint(2, 2),
	})

	for i := range 60000 {
		clock.After(time.Millisecond)

		if i == 3000 {
			march.TrapedInBattle(clock.Now())
		}

		if i == 5000 {
			march.BattleEnd(clock.Now())
		}

		march.Update(clock.Now())
	}
}
