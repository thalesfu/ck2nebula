package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
)

type People_LoverPeople struct {
	People      *People `nebulaedgename:"people_loverpeople" nebulaedgecomment:"People -> Lover People" nebulakey:"edgefrom" json:"from,omitempty"`
	LoverPeople *People `nebulakey:"edgeto" json:"to,omitempty"`
	PlayID      int     `nebulaproperty:"play_id" description:"game play id" nebulaindexes:"play_id" json:"play_id,omitempty"`
}

func NewPeople_LoverPeople(p *People, lp *People) *People_LoverPeople {
	return &People_LoverPeople{
		People:      p,
		LoverPeople: lp,
		PlayID:      p.PlayID,
	}
}

func (plp *People_LoverPeople) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertEdges(space, plp)
}

func (plp *People_LoverPeople) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateEdges(space, plp)
}

func (plp *People_LoverPeople) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertEdges(space, plp)
}

func (plp *People_LoverPeople) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, plp)
}

func (plp *People_LoverPeople) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadEdge(space, plp)
}

func InsertPeople_LoverPeoples(space *nebulagolang.Space, plps ...*People_LoverPeople) *nebulagolang.Result {
	return nebulagolang.BatchInsertEdges(space, 250, plps)
}

func UpdatePeople_LoverPeoples(space *nebulagolang.Space, plps ...*People_LoverPeople) *nebulagolang.Result {
	return nebulagolang.BatchUpdateEdges(space, 250, plps)
}

func UpsertPeople_LoverPeoples(space *nebulagolang.Space, plps ...*People_LoverPeople) *nebulagolang.Result {
	return nebulagolang.BatchUpsertEdges(space, 250, plps)
}

func DeletePeople_LoverPeoples(space *nebulagolang.Space, plps ...*People_LoverPeople) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, plps...)
}

func DeleteAllPeople_LoverPeoples(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByEdgeType[People_LoverPeople](space)
}

func DeleteAllPeople_LoverByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByQuery[People_LoverPeople](space, getPlayIdQuery[People_LoverPeople](playId))
}

func GetPeople_LoverPeopleByEid(space *nebulagolang.Space, eid *nebulagolang.EID) *nebulagolang.ResultT[*People_LoverPeople] {
	return nebulagolang.GetEdgeByEid[*People_LoverPeople](space, eid)
}

func GetPeople_LoverPeopleByFromAndToIds(space *nebulagolang.Space, fromId string, toId string) *nebulagolang.ResultT[*People_LoverPeople] {
	return nebulagolang.GetEdgeByEid[*People_LoverPeople](space, nebulagolang.NewEID(fromId, toId, nebulagolang.GetEdgeName[People_LoverPeople]()))
}

func GetAllPeople_LoverPeoplesEids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*People_LoverPeople](space, "")
}

func GetAllPeople_LoverPeoplesEidsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*People_LoverPeople](space, getPlayIdQuery[People_LoverPeople](playId))
}

func GetAllPeople_LoverPeoples(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*People_LoverPeople] {
	return nebulagolang.GetAllEdgesByEdgeType[*People_LoverPeople](space)
}

func GetAllPeople_LoverPeoplesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]*People_LoverPeople] {
	return nebulagolang.GetAllEdgesByQuery[*People_LoverPeople](space, getPlayIdQuery[People_LoverPeople](playId))
}
