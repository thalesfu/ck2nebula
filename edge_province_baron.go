package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
)

type Province_Baron struct {
	Province *Province `nebulaedgename:"province_baron" nebulaedgecomment:"Province -> Baron" nebulakey:"edgefrom" json:"from,omitempty"`
	Baron    *Baron    `nebulakey:"edgeto" json:"to,omitempty"`
	PlayID   int       `nebulaproperty:"play_id" description:"game play id" nebulaindexes:"play_id" json:"play_id,omitempty"`
}

func NewProvince_Baron(p *Province, b *Baron) *Province_Baron {
	return &Province_Baron{
		Province: p,
		Baron:    b,
		PlayID:   p.PlayID,
	}
}

func (pb *Province_Baron) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertEdges(space, pb)
}

func (pb *Province_Baron) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateEdges(space, pb)
}

func (pb *Province_Baron) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertEdges(space, pb)
}

func (pb *Province_Baron) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, pb)
}

func (pb *Province_Baron) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadEdge(space, pb)
}

func InsertProvince_Barons(space *nebulagolang.Space, pbs ...*Province_Baron) *nebulagolang.Result {
	return nebulagolang.BatchInsertEdges(space, 250, pbs)
}

func UpdateProvince_Barons(space *nebulagolang.Space, pbs ...*Province_Baron) *nebulagolang.Result {
	return nebulagolang.BatchUpdateEdges(space, 250, pbs)
}

func UpsertProvince_Barons(space *nebulagolang.Space, pbs ...*Province_Baron) *nebulagolang.Result {
	return nebulagolang.BatchUpsertEdges(space, 250, pbs)
}

func DeleteProvince_Barons(space *nebulagolang.Space, pbs ...*Province_Baron) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, pbs...)
}

func DeleteAllProvince_Barons(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByEdgeType[Province_Baron](space)
}

func DeleteAllProvince_BaronsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByQuery[Province_Baron](space, getPlayIdQuery[Province_Baron](playId))
}

func GetProvince_BaronByEid(space *nebulagolang.Space, eid *nebulagolang.EID) *nebulagolang.ResultT[*Province_Baron] {
	return nebulagolang.GetEdgeByEid[*Province_Baron](space, eid)
}

func GetProvince_BaronByFromAndToIds(space *nebulagolang.Space, fromId string, toId string) *nebulagolang.ResultT[*Province_Baron] {
	return nebulagolang.GetEdgeByEid[*Province_Baron](space, nebulagolang.NewEID(fromId, toId, nebulagolang.GetEdgeName[Province_Baron]()))
}

func GetAllProvince_BaronsEids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*Province_Baron](space, "")
}

func GetAllProvince_BaronsEidsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*Province_Baron](space, getPlayIdQuery[Province_Baron](playId))
}

func GetAllProvince_Barons(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*Province_Baron] {
	return nebulagolang.GetAllEdgesByEdgeType[*Province_Baron](space)
}

func GetAllProvince_BaronsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]*Province_Baron] {
	return nebulagolang.GetAllEdgesByQuery[*Province_Baron](space, getPlayIdQuery[Province_Baron](playId))
}
