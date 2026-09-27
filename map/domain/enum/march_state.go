package enum

type MarchState int32

const (
	MarchState_Idle MarchState = iota
	MarchState_Moving
	MarchState_Stay
	MarchState_Retreat
	MarchState_Rout
	MarchState_BattleWait
	MarchState_Battle
)

var marchStateName = map[MarchState]string{
	MarchState_Idle:       "Idle",
	MarchState_Moving:     "Moving",
	MarchState_Stay:       "Stay",
	MarchState_Retreat:    "Retreat",
	MarchState_Rout:       "Rout",
	MarchState_BattleWait: "BattleWait",
	MarchState_Battle:     "Battle",
}

func (p MarchState) String() string {
	return marchStateName[p]
}
