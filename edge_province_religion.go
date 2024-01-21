package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
)

type Province_Religion struct {
	Province *Province `nebulaedgename:"province_religion" nebulaedgecomment:"Province -> Religion" nebulakey:"edgefrom" json:"from,omitempty"`
	Religion *Religion `nebulakey:"edgeto" json:"to,omitempty"`
	PlayID   int       `nebulaproperty:"play_id" description:"game play id" nebulaindexes:"play_id" json:"play_id,omitempty"`
}

func NewProvince_Religion(p *Province, r *Religion) *Province_Religion {
	return &Province_Religion{
		Province: p,
		Religion: r,
		PlayID:   p.PlayID,
	}
}

func (pr *Province_Religion) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertEdges(space, pr)
}

func (pr *Province_Religion) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateEdges(space, pr)
}

func (pr *Province_Religion) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertEdges(space, pr)
}

func (pr *Province_Religion) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, pr)
}

func (pr *Province_Religion) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadEdge(space, pr)
}

func InsertProvince_Religions(space *nebulagolang.Space, prs ...*Province_Religion) *nebulagolang.Result {
	return nebulagolang.BatchInsertEdges(space, 250, prs)
}

func UpdateProvince_Religions(space *nebulagolang.Space, prs ...*Province_Religion) *nebulagolang.Result {
	return nebulagolang.BatchUpdateEdges(space, 250, prs)
}

func UpsertProvince_Religions(space *nebulagolang.Space, prs ...*Province_Religion) *nebulagolang.Result {
	return nebulagolang.BatchUpsertEdges(space, 250, prs)
}

func DeleteProvince_Religions(space *nebulagolang.Space, prs ...*Province_Religion) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, prs...)
}

func DeleteAllProvince_Religions(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByEdgeType[Province_Religion](space)
}

func DeleteAllProvince_ReligionsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByQuery[Province_Religion](space, getPlayIdQuery[Province_Religion](playId))
}

func GetProvince_ReligionByEid(space *nebulagolang.Space, eid *nebulagolang.EID) *nebulagolang.ResultT[*Province_Religion] {
	return nebulagolang.GetEdgeByEid[*Province_Religion](space, eid)
}

func GetProvince_ReligionByFromAndToIds(space *nebulagolang.Space, fromId string, toId string) *nebulagolang.ResultT[*Province_Religion] {
	return nebulagolang.GetEdgeByEid[*Province_Religion](space, nebulagolang.NewEID(fromId, toId, nebulagolang.GetEdgeName[Province_Religion]()))
}

func GetAllProvince_ReligionsEids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*Province_Religion](space, "")
}

func GetAllProvince_ReligionsEidsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*Province_Religion](space, getPlayIdQuery[Province_Religion](playId))
}

func GetAllProvince_Religions(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*Province_Religion] {
	return nebulagolang.GetAllEdgesByEdgeType[*Province_Religion](space)
}

func GetAllProvince_ReligionsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]*Province_Religion] {
	return nebulagolang.GetAllEdgesByQuery[*Province_Religion](space, getPlayIdQuery[Province_Religion](playId))
}
