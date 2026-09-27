package enum

type MarchCommand int32

const (
	MarchCommand_Stay MarchCommand = iota
	MarchCommand_Retreat
	MarchCommand_Occupy
	MarchCommand_Siege
	MarchCommand_Garrison
	MarchCommand_Redeploy
	MarchCommand_Crusade
	MarchCommand_Sweep
	MarchCommand_Entrust
	MarchCommand_Assemble
	MarchCommand_Develop
	MarchCommand_Stockpile
	MarchCommand_Mining
	MarchCommand_Fire
	MarchCommand_Outfire
	MarchCommand_SOS
	MarchCommand_Traps
	MarchCommand_ChangeArmy
	MarchCommand_Run
	MarchCommand_Control
	MarchCommand_Pickup
)
