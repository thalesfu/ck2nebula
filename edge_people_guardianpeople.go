package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
)

type People_GuardianPeople struct {
	People         *People `nebulaedgename:"people_guardianpeople" nebulaedgecomment:"People -> Guardian People" nebulakey:"edgefrom" json:"from,omitempty"`
	GuardianPeople *People `nebulakey:"edgeto" json:"to,omitempty"`
	PlayID         int     `nebulaproperty:"play_id" description:"game play id" nebulaindexes:"play_id" json:"play_id,omitempty"`
}

func NewPeople_GuardianPeople(p *People, gp *People) *People_GuardianPeople {
	return &People_GuardianPeople{
		People:         p,
		GuardianPeople: gp,
		PlayID:         p.PlayID,
	}
}

func (pgp *People_GuardianPeople) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertEdges(space, pgp)
}

func (pgp *People_GuardianPeople) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateEdges(space, pgp)
}

func (pgp *People_GuardianPeople) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertEdges(space, pgp)
}

func (pgp *People_GuardianPeople) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, pgp)
}

func (pgp *People_GuardianPeople) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadEdge(space, pgp)
}

func InsertPeople_GuardianPeoples(space *nebulagolang.Space, pgps ...*People_GuardianPeople) *nebulagolang.Result {
	return nebulagolang.BatchInsertEdges(space, 250, pgps)
}

func UpdatePeople_GuardianPeoples(space *nebulagolang.Space, pgps ...*People_GuardianPeople) *nebulagolang.Result {
	return nebulagolang.BatchUpdateEdges(space, 250, pgps)
}

func UpsertPeople_GuardianPeoples(space *nebulagolang.Space, pgps ...*People_GuardianPeople) *nebulagolang.Result {
	return nebulagolang.BatchUpsertEdges(space, 250, pgps)
}

func DeletePeople_GuardianPeoples(space *nebulagolang.Space, pgps ...*People_GuardianPeople) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, pgps...)
}

func DeleteAllPeople_GuardianPeoples(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByEdgeType[People_GuardianPeople](space)
}

func DeleteAllPeople_GuardianByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByQuery[People_GuardianPeople](space, getPlayIdQuery[People_GuardianPeople](playId))
}

func GetPeople_GuardianPeopleByEid(space *nebulagolang.Space, eid *nebulagolang.EID) *nebulagolang.ResultT[*People_GuardianPeople] {
	return nebulagolang.GetEdgeByEid[*People_GuardianPeople](space, eid)
}

func GetPeople_GuardianPeopleByFromAndToIds(space *nebulagolang.Space, fromId string, toId string) *nebulagolang.ResultT[*People_GuardianPeople] {
	return nebulagolang.GetEdgeByEid[*People_GuardianPeople](space, nebulagolang.NewEID(fromId, toId, nebulagolang.GetEdgeName[People_GuardianPeople]()))
}

func GetAllPeople_GuardianPeoplesEids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*People_GuardianPeople](space, "")
}

func GetAllPeople_GuardianPeoplesEidsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*People_GuardianPeople](space, getPlayIdQuery[People_GuardianPeople](playId))
}

func GetAllPeople_GuardianPeoples(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*People_GuardianPeople] {
	return nebulagolang.GetAllEdgesByEdgeType[*People_GuardianPeople](space)
}

func GetAllPeople_GuardianPeoplesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]*People_GuardianPeople] {
	return nebulagolang.GetAllEdgesByQuery[*People_GuardianPeople](space, getPlayIdQuery[People_GuardianPeople](playId))
}
