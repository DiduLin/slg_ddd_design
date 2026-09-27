package event

import "github.com/DiduLin/slg_ddd_design/types"

type MarchMoved struct {
	ObjectID int64
	From     types.Point
	To       types.Point
}
