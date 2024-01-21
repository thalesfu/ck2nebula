package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
	"time"
)

type Province_Modifier struct {
	Province *Province `nebulaedgename:"province_modifier" nebulaedgecomment:"Province -> Modifier" nebulakey:"edgefrom" json:"from,omitempty"`
	Modifier *Modifier `nebulakey:"edgeto" json:"to,omitempty"`
	Abate    time.Time `nebulaproperty:"abate" nebulatype:"Date"  description:"abate date" nebulaindexes:"abate" json:"abate,omitempty"`
	Hidden   bool      `nebulaproperty:"hidden" description:"hidden" nebulaindexes:"hidden" json:"hidden,omitempty"`
	PlayID   int       `nebulaproperty:"play_id" description:"game play id" nebulaindexes:"play_id" json:"play_id,omitempty"`
}

func NewProvince_Modifier(p *Province, m *Modifier) *Province_Modifier {
	return &Province_Modifier{
		Province: p,
		Modifier: m,
		PlayID:   p.PlayID,
	}
}

func (pm *Province_Modifier) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertEdges(space, pm)
}

func (pm *Province_Modifier) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateEdges(space, pm)
}

func (pm *Province_Modifier) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertEdges(space, pm)
}

func (pm *Province_Modifier) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, pm)
}

func (pm *Province_Modifier) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadEdge(space, pm)
}

func InsertProvince_Modifiers(space *nebulagolang.Space, pms ...*Province_Modifier) *nebulagolang.Result {
	return nebulagolang.BatchInsertEdges(space, 250, pms)
}

func UpdateProvince_Modifiers(space *nebulagolang.Space, pms ...*Province_Modifier) *nebulagolang.Result {
	return nebulagolang.BatchUpdateEdges(space, 250, pms)
}

func UpsertProvince_Modifiers(space *nebulagolang.Space, pms ...*Province_Modifier) *nebulagolang.Result {
	return nebulagolang.BatchUpsertEdges(space, 250, pms)
}

func DeleteProvince_Modifiers(space *nebulagolang.Space, pms ...*Province_Modifier) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, pms...)
}

func DeleteAllProvince_Modifiers(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByEdgeType[Province_Modifier](space)
}

func DeleteAllProvince_ModifiersByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByQuery[Province_Modifier](space, getPlayIdQuery[Province_Modifier](playId))
}

func GetProvince_ModifierByEid(space *nebulagolang.Space, eid *nebulagolang.EID) *nebulagolang.ResultT[*Province_Modifier] {
	return nebulagolang.GetEdgeByEid[*Province_Modifier](space, eid)
}

func GetProvince_ModifierByFromAndToIds(space *nebulagolang.Space, fromId string, toId string) *nebulagolang.ResultT[*Province_Modifier] {
	return nebulagolang.GetEdgeByEid[*Province_Modifier](space, nebulagolang.NewEID(fromId, toId, nebulagolang.GetEdgeName[Province_Modifier]()))
}

func GetAllProvince_ModifiersEids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*Province_Modifier](space, "")
}

func GetAllProvince_ModifiersEidsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*Province_Modifier](space, getPlayIdQuery[Province_Modifier](playId))
}

func GetAllProvince_Modifiers(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*Province_Modifier] {
	return nebulagolang.GetAllEdgesByEdgeType[*Province_Modifier](space)
}

func GetAllProvince_ModifiersByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]*Province_Modifier] {
	return nebulagolang.GetAllEdgesByQuery[*Province_Modifier](space, getPlayIdQuery[Province_Modifier](playId))
}
