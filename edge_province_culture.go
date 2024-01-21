package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
)

type Province_Culture struct {
	Province *Province `nebulaedgename:"province_culture" nebulaedgecomment:"Province -> Culture" nebulakey:"edgefrom" json:"from,omitempty"`
	Culture  *Culture  `nebulakey:"edgeto" json:"to,omitempty"`
	PlayID   int       `nebulaproperty:"play_id" description:"game play id" nebulaindexes:"play_id" json:"play_id,omitempty"`
}

func NewProvince_Culture(p *Province, c *Culture) *Province_Culture {
	return &Province_Culture{
		Province: p,
		Culture:  c,
		PlayID:   p.PlayID,
	}
}

func (pc *Province_Culture) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertEdges(space, pc)
}

func (pc *Province_Culture) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateEdges(space, pc)
}

func (pc *Province_Culture) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertEdges(space, pc)
}

func (pc *Province_Culture) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, pc)
}

func (pc *Province_Culture) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadEdge(space, pc)
}

func InsertProvince_Cultures(space *nebulagolang.Space, pcs ...*Province_Culture) *nebulagolang.Result {
	return nebulagolang.BatchInsertEdges(space, 250, pcs)
}

func UpdateProvince_Cultures(space *nebulagolang.Space, pcs ...*Province_Culture) *nebulagolang.Result {
	return nebulagolang.BatchUpdateEdges(space, 250, pcs)
}

func UpsertProvince_Cultures(space *nebulagolang.Space, pcs ...*Province_Culture) *nebulagolang.Result {
	return nebulagolang.BatchUpsertEdges(space, 250, pcs)
}

func DeleteProvince_Cultures(space *nebulagolang.Space, pcs ...*Province_Culture) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, pcs...)
}

func DeleteAllProvince_Cultures(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByEdgeType[Province_Culture](space)
}

func DeleteAllProvince_CulturesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByQuery[Province_Culture](space, getPlayIdQuery[Province_Culture](playId))
}

func GetProvince_CultureByEid(space *nebulagolang.Space, eid *nebulagolang.EID) *nebulagolang.ResultT[*Province_Culture] {
	return nebulagolang.GetEdgeByEid[*Province_Culture](space, eid)
}

func GetProvince_CultureByFromAndToIds(space *nebulagolang.Space, fromId string, toId string) *nebulagolang.ResultT[*Province_Culture] {
	return nebulagolang.GetEdgeByEid[*Province_Culture](space, nebulagolang.NewEID(fromId, toId, nebulagolang.GetEdgeName[Province_Culture]()))
}

func GetAllProvince_CulturesEids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*Province_Culture](space, "")
}

func GetAllProvince_CulturesEidsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*Province_Culture](space, getPlayIdQuery[Province_Culture](playId))
}

func GetAllProvince_Cultures(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*Province_Culture] {
	return nebulagolang.GetAllEdgesByEdgeType[*Province_Culture](space)
}

func GetAllProvince_CulturesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]*Province_Culture] {
	return nebulagolang.GetAllEdgesByQuery[*Province_Culture](space, getPlayIdQuery[Province_Culture](playId))
}
