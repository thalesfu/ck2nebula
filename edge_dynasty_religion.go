package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
)

type Dynasty_Religion struct {
	Dynasty  *Dynasty  `nebulaedgename:"dynasty_religion" nebulaedgecomment:"Dynasty -> Religion" nebulakey:"edgefrom" json:"from,omitempty"`
	Religion *Religion `nebulakey:"edgeto" json:"to,omitempty"`
	PlayID   int       `nebulaproperty:"play_id" description:"game play id" nebulaindexes:"play_id" json:"play_id,omitempty"`
}

func NewDynasty_Religion(d *Dynasty, r *Religion) *Dynasty_Religion {
	return &Dynasty_Religion{
		Dynasty:  d,
		Religion: r,
		PlayID:   d.PlayID,
	}
}

func (dr *Dynasty_Religion) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertEdges(space, dr)
}

func (dr *Dynasty_Religion) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateEdges(space, dr)
}

func (dr *Dynasty_Religion) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertEdges(space, dr)
}

func (dr *Dynasty_Religion) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, dr)
}

func (dr *Dynasty_Religion) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadEdge(space, dr)
}

func InsertDynasty_Religions(space *nebulagolang.Space, drs ...*Dynasty_Religion) *nebulagolang.Result {
	return nebulagolang.BatchInsertEdges(space, 250, drs)
}

func UpdateDynasty_Religions(space *nebulagolang.Space, drs ...*Dynasty_Religion) *nebulagolang.Result {
	return nebulagolang.BatchUpdateEdges(space, 250, drs)
}

func UpsertDynasty_Religions(space *nebulagolang.Space, drs ...*Dynasty_Religion) *nebulagolang.Result {
	return nebulagolang.BatchUpsertEdges(space, 250, drs)
}

func DeleteDynasty_Religions(space *nebulagolang.Space, drs ...*Dynasty_Religion) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, drs...)
}

func DeleteAllDynasty_Religions(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByEdgeType[Dynasty_Religion](space)
}

func DeleteAllDynasty_ReligionsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByQuery[Dynasty_Religion](space, getPlayIdQuery[Dynasty_Religion](playId))
}

func GetDynasty_ReligionByEid(space *nebulagolang.Space, eid *nebulagolang.EID) *nebulagolang.ResultT[*Dynasty_Religion] {
	return nebulagolang.GetEdgeByEid[*Dynasty_Religion](space, eid)
}

func GetDynasty_ReligionByFromAndToIds(space *nebulagolang.Space, fromId string, toId string) *nebulagolang.ResultT[*Dynasty_Religion] {
	return nebulagolang.GetEdgeByEid[*Dynasty_Religion](space, nebulagolang.NewEID(fromId, toId, nebulagolang.GetEdgeName[Dynasty_Religion]()))
}

func GetAllDynasty_ReligionsEids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*Dynasty_Religion](space, "")
}

func GetAllDynasty_ReligionsEidsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*Dynasty_Religion](space, getPlayIdQuery[Dynasty_Religion](playId))
}

func GetAllDynasty_Religions(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*Dynasty_Religion] {
	return nebulagolang.GetAllEdgesByEdgeType[*Dynasty_Religion](space)
}

func GetAllDynasty_ReligionsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]*Dynasty_Religion] {
	return nebulagolang.GetAllEdgesByQuery[*Dynasty_Religion](space, getPlayIdQuery[Dynasty_Religion](playId))
}
