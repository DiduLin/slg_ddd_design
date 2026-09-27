package acl

import (
	"github.com/DiduLin/slg_ddd_design/map/domain/enum"
	"github.com/DiduLin/slg_ddd_design/types"
)

type PathFindGateway interface {
	FindPath(objectID int64, command enum.MarchCommand, targetObjectID int64, start, end *types.Point) error
}
