package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
	"time"
)

type People_Modifier struct {
	People   *People   `nebulaedgename:"people_modifier" nebulaedgecomment:"People -> Modifier" nebulakey:"edgefrom" json:"from,omitempty"`
	Modifier *Modifier `nebulakey:"edgeto" json:"to,omitempty"`
	Abate    time.Time `nebulaproperty:"abate" nebulatype:"Date"  description:"abate date" nebulaindexes:"abate" json:"abate,omitempty"`
	Hidden   bool      `nebulaproperty:"hidden" description:"hidden" nebulaindexes:"hidden" json:"hidden,omitempty"`
	PlayID   int       `nebulaproperty:"play_id" description:"game play id" nebulaindexes:"play_id" json:"play_id,omitempty"`
}

func NewPeople_Modifier(p *People, m *Modifier) *People_Modifier {
	return &People_Modifier{
		People:   p,
		Modifier: m,
		PlayID:   p.PlayID,
	}
}

func (pm *People_Modifier) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertEdges(space, pm)
}

func (pm *People_Modifier) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateEdges(space, pm)
}

func (pm *People_Modifier) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertEdges(space, pm)
}

func (pm *People_Modifier) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, pm)
}

func (pm *People_Modifier) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadEdge(space, pm)
}

func InsertPeople_Modifiers(space *nebulagolang.Space, pms ...*People_Modifier) *nebulagolang.Result {
	return nebulagolang.BatchInsertEdges(space, 250, pms)
}

func UpdatePeople_Modifiers(space *nebulagolang.Space, pms ...*People_Modifier) *nebulagolang.Result {
	return nebulagolang.BatchUpdateEdges(space, 250, pms)
}

func UpsertPeople_Modifiers(space *nebulagolang.Space, pms ...*People_Modifier) *nebulagolang.Result {
	return nebulagolang.BatchUpsertEdges(space, 250, pms)
}

func DeletePeople_Modifiers(space *nebulagolang.Space, pms ...*People_Modifier) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, pms...)
}

func DeleteAllPeople_Modifiers(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByEdgeType[People_Modifier](space)
}

func DeleteAllPeople_ModifiersByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByQuery[People_Modifier](space, getPlayIdQuery[People_Modifier](playId))
}

func GetPeople_ModifierByEid(space *nebulagolang.Space, eid *nebulagolang.EID) *nebulagolang.ResultT[*People_Modifier] {
	return nebulagolang.GetEdgeByEid[*People_Modifier](space, eid)
}

func GetPeople_ModifierByFromAndToIds(space *nebulagolang.Space, fromId string, toId string) *nebulagolang.ResultT[*People_Modifier] {
	return nebulagolang.GetEdgeByEid[*People_Modifier](space, nebulagolang.NewEID(fromId, toId, nebulagolang.GetEdgeName[People_Modifier]()))
}

func GetAllPeople_ModifiersEids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*People_Modifier](space, "")
}

func GetAllPeople_ModifiersEidsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*People_Modifier](space, getPlayIdQuery[People_Modifier](playId))
}

func GetAllPeople_Modifiers(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*People_Modifier] {
	return nebulagolang.GetAllEdgesByEdgeType[*People_Modifier](space)
}

func GetAllPeople_ModifiersByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]*People_Modifier] {
	return nebulagolang.GetAllEdgesByQuery[*People_Modifier](space, getPlayIdQuery[People_Modifier](playId))
}
