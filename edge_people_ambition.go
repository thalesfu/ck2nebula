package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
)

type People_Ambition struct {
	People   *People    `nebulaedgename:"people_ambition" nebulaedgecomment:"People -> Ambition" nebulakey:"edgefrom" json:"from,omitempty"`
	Ambition *Objective `nebulakey:"edgeto" json:"to,omitempty"`
	StoryID  int        `nebulaproperty:"story_id" mappingalias:"PlayID" description:"game play id" nebulaindexes:"story_id" json:"story_id,omitempty"`
}

func NewPeople_Ambition(p *People, a *Objective) *People_Ambition {
	return &People_Ambition{
		People:   p,
		Ambition: a,
		StoryID:  p.StoryID,
	}
}

func (pa *People_Ambition) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertEdges(space, pa)
}

func (pa *People_Ambition) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateEdges(space, pa)
}

func (pa *People_Ambition) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertEdges(space, pa)
}

func (pa *People_Ambition) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, pa)
}

func (pa *People_Ambition) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadEdge(space, pa)
}

func InsertPeople_Ambitions(space *nebulagolang.Space, pas ...*People_Ambition) *nebulagolang.Result {
	return nebulagolang.BatchInsertEdges(space, 250, pas)
}

func UpdatePeople_Ambitions(space *nebulagolang.Space, pas ...*People_Ambition) *nebulagolang.Result {
	return nebulagolang.BatchUpdateEdges(space, 250, pas)
}

func UpsertPeople_Ambitions(space *nebulagolang.Space, pas ...*People_Ambition) *nebulagolang.Result {
	return nebulagolang.BatchUpsertEdges(space, 250, pas)
}

func DeletePeople_Ambitions(space *nebulagolang.Space, pas ...*People_Ambition) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, pas...)
}

func DeleteAllPeople_Ambitions(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByEdgeType[People_Ambition](space)
}

func DeleteAllPeople_AmbitionsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByQuery[People_Ambition](space, getPlayIdQuery[People_Ambition](playId))
}

func GetPeople_AmbitionByEid(space *nebulagolang.Space, eid *nebulagolang.EID) *nebulagolang.ResultT[*People_Ambition] {
	return nebulagolang.GetEdgeByEid[*People_Ambition](space, eid)
}

func GetPeople_AmbitionByFromAndToIds(space *nebulagolang.Space, fromId string, toId string) *nebulagolang.ResultT[*People_Ambition] {
	return nebulagolang.GetEdgeByEid[*People_Ambition](space, nebulagolang.NewEID(fromId, toId, nebulagolang.GetEdgeName[People_Ambition]()))
}

func GetAllPeople_AmbitionsEids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*People_Ambition](space, "")
}

func GetAllPeople_AmbitionsEidsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*People_Ambition](space, getPlayIdQuery[People_Ambition](playId))
}

func GetAllPeople_Ambitions(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*People_Ambition] {
	return nebulagolang.GetAllEdgesByEdgeType[*People_Ambition](space)
}

func GetAllPeople_AmbitionsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]*People_Ambition] {
	return nebulagolang.GetAllEdgesByQuery[*People_Ambition](space, getPlayIdQuery[People_Ambition](playId))
}
