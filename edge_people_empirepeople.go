package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
)

type People_EmpirePeople struct {
	People       *People `nebulaedgename:"people_empirepeople" nebulaedgecomment:"People -> Empire People" nebulakey:"edgefrom" json:"from,omitempty"`
	EmpirePeople *People `nebulakey:"edgeto" json:"to,omitempty"`
	StoryID      int     `nebulaproperty:"story_id" mappingalias:"PlayID" description:"game play id" nebulaindexes:"story_id" json:"story_id,omitempty"`
}

func NewPeople_EmpirePeople(p *People, ep *People) *People_EmpirePeople {
	return &People_EmpirePeople{
		People:       p,
		EmpirePeople: ep,
		StoryID:      p.StoryID,
	}
}

func (pep *People_EmpirePeople) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertEdges(space, pep)
}

func (pep *People_EmpirePeople) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateEdges(space, pep)
}

func (pep *People_EmpirePeople) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertEdges(space, pep)
}

func (pep *People_EmpirePeople) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, pep)
}

func (pep *People_EmpirePeople) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadEdge(space, pep)
}

func InsertPeople_EmpirePeoples(space *nebulagolang.Space, peps ...*People_EmpirePeople) *nebulagolang.Result {
	return nebulagolang.BatchInsertEdges(space, 250, peps)
}

func UpdatePeople_EmpirePeoples(space *nebulagolang.Space, peps ...*People_EmpirePeople) *nebulagolang.Result {
	return nebulagolang.BatchUpdateEdges(space, 250, peps)
}

func UpsertPeople_EmpirePeoples(space *nebulagolang.Space, peps ...*People_EmpirePeople) *nebulagolang.Result {
	return nebulagolang.BatchUpsertEdges(space, 250, peps)
}

func DeletePeople_EmpirePeoples(space *nebulagolang.Space, peps ...*People_EmpirePeople) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, peps...)
}

func DeleteAllPeople_EmpirePeoples(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByEdgeType[People_EmpirePeople](space)
}

func DeleteAllPeople_EmpireByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByQuery[People_EmpirePeople](space, getPlayIdQuery[People_EmpirePeople](playId))
}

func GetPeople_EmpirePeopleByEid(space *nebulagolang.Space, eid *nebulagolang.EID) *nebulagolang.ResultT[*People_EmpirePeople] {
	return nebulagolang.GetEdgeByEid[*People_EmpirePeople](space, eid)
}

func GetPeople_EmpirePeopleByFromAndToIds(space *nebulagolang.Space, fromId string, toId string) *nebulagolang.ResultT[*People_EmpirePeople] {
	return nebulagolang.GetEdgeByEid[*People_EmpirePeople](space, nebulagolang.NewEID(fromId, toId, nebulagolang.GetEdgeName[People_EmpirePeople]()))
}

func GetAllPeople_EmpirePeoplesEids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*People_EmpirePeople](space, "")
}

func GetAllPeople_EmpirePeoplesEidsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*People_EmpirePeople](space, getPlayIdQuery[People_EmpirePeople](playId))
}

func GetAllPeople_EmpirePeoples(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*People_EmpirePeople] {
	return nebulagolang.GetAllEdgesByEdgeType[*People_EmpirePeople](space)
}

func GetAllPeople_EmpirePeoplesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]*People_EmpirePeople] {
	return nebulagolang.GetAllEdgesByQuery[*People_EmpirePeople](space, getPlayIdQuery[People_EmpirePeople](playId))
}
