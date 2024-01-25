package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
)

type People_Focus struct {
	People *People    `nebulaedgename:"people_focus" nebulaedgecomment:"People -> Focus" nebulakey:"edgefrom" json:"from,omitempty"`
	Focus  *Objective `nebulakey:"edgeto" json:"to,omitempty"`
	PlayID int        `nebulaproperty:"play_id" description:"game play id" nebulaindexes:"play_id" json:"play_id,omitempty"`
}

func NewPeople_Focus(p *People, f *Objective) *People_Focus {
	return &People_Focus{
		People: p,
		Focus:  f,
		PlayID: p.PlayID,
	}
}

func (pf *People_Focus) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertEdges(space, pf)
}

func (pf *People_Focus) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateEdges(space, pf)
}

func (pf *People_Focus) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertEdges(space, pf)
}

func (pf *People_Focus) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, pf)
}

func (pf *People_Focus) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadEdge(space, pf)
}

func InsertPeople_Focuss(space *nebulagolang.Space, pfs ...*People_Focus) *nebulagolang.Result {
	return nebulagolang.BatchInsertEdges(space, 250, pfs)
}

func UpdatePeople_Focuss(space *nebulagolang.Space, pfs ...*People_Focus) *nebulagolang.Result {
	return nebulagolang.BatchUpdateEdges(space, 250, pfs)
}

func UpsertPeople_Focuss(space *nebulagolang.Space, pfs ...*People_Focus) *nebulagolang.Result {
	return nebulagolang.BatchUpsertEdges(space, 250, pfs)
}

func DeletePeople_Focuss(space *nebulagolang.Space, pfs ...*People_Focus) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, pfs...)
}

func DeleteAllPeople_Focuss(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByEdgeType[People_Focus](space)
}

func DeleteAllPeople_FocusesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByQuery[People_Focus](space, getPlayIdQuery[People_Focus](playId))
}

func GetPeople_FocusByEid(space *nebulagolang.Space, eid *nebulagolang.EID) *nebulagolang.ResultT[*People_Focus] {
	return nebulagolang.GetEdgeByEid[*People_Focus](space, eid)
}

func GetPeople_FocusByFromAndToIds(space *nebulagolang.Space, fromId string, toId string) *nebulagolang.ResultT[*People_Focus] {
	return nebulagolang.GetEdgeByEid[*People_Focus](space, nebulagolang.NewEID(fromId, toId, nebulagolang.GetEdgeName[People_Focus]()))
}

func GetAllPeople_FocussEids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*People_Focus](space, "")
}

func GetAllPeople_FocussEidsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*People_Focus](space, getPlayIdQuery[People_Focus](playId))
}

func GetAllPeople_Focuss(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*People_Focus] {
	return nebulagolang.GetAllEdgesByEdgeType[*People_Focus](space)
}

func GetAllPeople_FocussByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]*People_Focus] {
	return nebulagolang.GetAllEdgesByQuery[*People_Focus](space, getPlayIdQuery[People_Focus](playId))
}
