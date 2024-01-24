package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
	"time"
)

type People_RelatePeople struct {
	People       *People   `nebulaedgename:"people_relatepeople" nebulaedgecomment:"People -> Relate People" nebulakey:"edgefrom" json:"from,omitempty"`
	RelatePeople *People   `nebulakey:"edgeto" json:"to,omitempty"`
	Rank         int       `nebulakey:"edgerank" json:"rank,omitempty"`
	Code         string    `nebulaproperty:"code" description:"code" nebulaindexes:"code" json:"code,omitempty"`
	Name         string    `nebulaproperty:"name" description:"name" nebulaindexes:"name" json:"name,omitempty"`
	Due          time.Time `nebulaproperty:"due" nebulatype:"Date" description:"due" nebulaindexes:"due" json:"due,omitempty"`
	Detail       string    `nebulaproperty:"detail" description:"detail" nebulaindexes:"detail" json:"detail,omitempty"`
	PlayID       int       `nebulaproperty:"play_id" description:"game play id" nebulaindexes:"play_id" json:"play_id,omitempty"`
}

func NewPeople_RelatePeople(p *People, rp *People) *People_RelatePeople {
	return &People_RelatePeople{
		People:       p,
		RelatePeople: rp,
		PlayID:       p.PlayID,
	}
}

func (prp *People_RelatePeople) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertEdges(space, prp)
}

func (prp *People_RelatePeople) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateEdges(space, prp)
}

func (prp *People_RelatePeople) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertEdges(space, prp)
}

func (prp *People_RelatePeople) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, prp)
}

func (prp *People_RelatePeople) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadEdge(space, prp)
}

func InsertPeople_RelatePeoples(space *nebulagolang.Space, prps ...*People_RelatePeople) *nebulagolang.Result {
	return nebulagolang.BatchInsertEdges(space, 250, prps)
}

func UpdatePeople_RelatePeoples(space *nebulagolang.Space, prps ...*People_RelatePeople) *nebulagolang.Result {
	return nebulagolang.BatchUpdateEdges(space, 250, prps)
}

func UpsertPeople_RelatePeoples(space *nebulagolang.Space, prps ...*People_RelatePeople) *nebulagolang.Result {
	return nebulagolang.BatchUpsertEdges(space, 250, prps)
}

func DeletePeople_RelatePeoples(space *nebulagolang.Space, prps ...*People_RelatePeople) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, prps...)
}

func DeleteAllPeople_RelatePeoples(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByEdgeType[People_RelatePeople](space)
}

func DeleteAllPeople_RelateByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByQuery[People_RelatePeople](space, getPlayIdQuery[People_RelatePeople](playId))
}

func GetPeople_RelatePeopleByEid(space *nebulagolang.Space, eid *nebulagolang.EID) *nebulagolang.ResultT[*People_RelatePeople] {
	return nebulagolang.GetEdgeByEid[*People_RelatePeople](space, eid)
}

func GetPeople_RelatePeopleByFromAndToIds(space *nebulagolang.Space, fromId string, toId string) *nebulagolang.ResultT[*People_RelatePeople] {
	return nebulagolang.GetEdgeByEid[*People_RelatePeople](space, nebulagolang.NewEID(fromId, toId, nebulagolang.GetEdgeName[People_RelatePeople]()))
}

func GetAllPeople_RelatePeoplesEids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*People_RelatePeople](space, "")
}

func GetAllPeople_RelatePeoplesEidsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*People_RelatePeople](space, getPlayIdQuery[People_RelatePeople](playId))
}

func GetAllPeople_RelatePeoples(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*People_RelatePeople] {
	return nebulagolang.GetAllEdgesByEdgeType[*People_RelatePeople](space)
}

func GetAllPeople_RelatePeoplesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]*People_RelatePeople] {
	return nebulagolang.GetAllEdgesByQuery[*People_RelatePeople](space, getPlayIdQuery[People_RelatePeople](playId))
}
