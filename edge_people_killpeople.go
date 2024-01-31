package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
)

type People_KillPeople struct {
	People     *People `nebulaedgename:"people_killpeople" nebulaedgecomment:"People -> Kill People" nebulakey:"edgefrom" json:"from,omitempty"`
	KillPeople *People `nebulakey:"edgeto" json:"to,omitempty"`
	StoryID    int     `nebulaproperty:"story_id" mappingalias:"PlayID" description:"game play id" nebulaindexes:"story_id" json:"story_id,omitempty"`
}

func NewPeople_KillPeople(p *People, kp *People) *People_KillPeople {
	return &People_KillPeople{
		People:     p,
		KillPeople: kp,
		StoryID:    p.StoryID,
	}
}

func (pkp *People_KillPeople) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertEdges(space, pkp)
}

func (pkp *People_KillPeople) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateEdges(space, pkp)
}

func (pkp *People_KillPeople) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertEdges(space, pkp)
}

func (pkp *People_KillPeople) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, pkp)
}

func (pkp *People_KillPeople) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadEdge(space, pkp)
}

func InsertPeople_KillPeoples(space *nebulagolang.Space, pkps ...*People_KillPeople) *nebulagolang.Result {
	return nebulagolang.BatchInsertEdges(space, 250, pkps)
}

func UpdatePeople_KillPeoples(space *nebulagolang.Space, pkps ...*People_KillPeople) *nebulagolang.Result {
	return nebulagolang.BatchUpdateEdges(space, 250, pkps)
}

func UpsertPeople_KillPeoples(space *nebulagolang.Space, pkps ...*People_KillPeople) *nebulagolang.Result {
	return nebulagolang.BatchUpsertEdges(space, 250, pkps)
}

func DeletePeople_KillPeoples(space *nebulagolang.Space, pkps ...*People_KillPeople) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, pkps...)
}

func DeleteAllPeople_KillPeoples(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByEdgeType[People_KillPeople](space)
}

func DeleteAllPeople_KillByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByQuery[People_KillPeople](space, getPlayIdQuery[People_KillPeople](playId))
}

func GetPeople_KillPeopleByEid(space *nebulagolang.Space, eid *nebulagolang.EID) *nebulagolang.ResultT[*People_KillPeople] {
	return nebulagolang.GetEdgeByEid[*People_KillPeople](space, eid)
}

func GetPeople_KillPeopleByFromAndToIds(space *nebulagolang.Space, fromId string, toId string) *nebulagolang.ResultT[*People_KillPeople] {
	return nebulagolang.GetEdgeByEid[*People_KillPeople](space, nebulagolang.NewEID(fromId, toId, nebulagolang.GetEdgeName[People_KillPeople]()))
}

func GetAllPeople_KillPeoplesEids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*People_KillPeople](space, "")
}

func GetAllPeople_KillPeoplesEidsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*People_KillPeople](space, getPlayIdQuery[People_KillPeople](playId))
}

func GetAllPeople_KillPeoples(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*People_KillPeople] {
	return nebulagolang.GetAllEdgesByEdgeType[*People_KillPeople](space)
}

func GetAllPeople_KillPeoplesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]*People_KillPeople] {
	return nebulagolang.GetAllEdgesByQuery[*People_KillPeople](space, getPlayIdQuery[People_KillPeople](playId))
}
