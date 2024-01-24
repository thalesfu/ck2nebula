package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
)

type People_HostPeople struct {
	People     *People `nebulaedgename:"people_hostpeople" nebulaedgecomment:"People -> Host People" nebulakey:"edgefrom" json:"from,omitempty"`
	HostPeople *People `nebulakey:"edgeto" json:"to,omitempty"`
	PlayID     int     `nebulaproperty:"play_id" description:"game play id" nebulaindexes:"play_id" json:"play_id,omitempty"`
}

func NewPeople_HostPeople(p *People, hp *People) *People_HostPeople {
	return &People_HostPeople{
		People:     p,
		HostPeople: hp,
		PlayID:     p.PlayID,
	}
}

func (php *People_HostPeople) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertEdges(space, php)
}

func (php *People_HostPeople) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateEdges(space, php)
}

func (php *People_HostPeople) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertEdges(space, php)
}

func (php *People_HostPeople) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, php)
}

func (php *People_HostPeople) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadEdge(space, php)
}

func InsertPeople_HostPeoples(space *nebulagolang.Space, phps ...*People_HostPeople) *nebulagolang.Result {
	return nebulagolang.BatchInsertEdges(space, 250, phps)
}

func UpdatePeople_HostPeoples(space *nebulagolang.Space, phps ...*People_HostPeople) *nebulagolang.Result {
	return nebulagolang.BatchUpdateEdges(space, 250, phps)
}

func UpsertPeople_HostPeoples(space *nebulagolang.Space, phps ...*People_HostPeople) *nebulagolang.Result {
	return nebulagolang.BatchUpsertEdges(space, 250, phps)
}

func DeletePeople_HostPeoples(space *nebulagolang.Space, phps ...*People_HostPeople) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, phps...)
}

func DeleteAllPeople_HostPeoples(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByEdgeType[People_HostPeople](space)
}

func DeleteAllPeople_HostByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByQuery[People_HostPeople](space, getPlayIdQuery[People_HostPeople](playId))
}

func GetPeople_HostPeopleByEid(space *nebulagolang.Space, eid *nebulagolang.EID) *nebulagolang.ResultT[*People_HostPeople] {
	return nebulagolang.GetEdgeByEid[*People_HostPeople](space, eid)
}

func GetPeople_HostPeopleByFromAndToIds(space *nebulagolang.Space, fromId string, toId string) *nebulagolang.ResultT[*People_HostPeople] {
	return nebulagolang.GetEdgeByEid[*People_HostPeople](space, nebulagolang.NewEID(fromId, toId, nebulagolang.GetEdgeName[People_HostPeople]()))
}

func GetAllPeople_HostPeoplesEids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*People_HostPeople](space, "")
}

func GetAllPeople_HostPeoplesEidsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*People_HostPeople](space, getPlayIdQuery[People_HostPeople](playId))
}

func GetAllPeople_HostPeoples(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*People_HostPeople] {
	return nebulagolang.GetAllEdgesByEdgeType[*People_HostPeople](space)
}

func GetAllPeople_HostPeoplesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]*People_HostPeople] {
	return nebulagolang.GetAllEdgesByQuery[*People_HostPeople](space, getPlayIdQuery[People_HostPeople](playId))
}
