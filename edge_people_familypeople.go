package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
)

type People_FamilyPeople struct {
	People       *People `nebulaedgename:"people_familypeople" nebulaedgecomment:"People -> Family People" nebulakey:"edgefrom" json:"from,omitempty"`
	FamilyPeople *People `nebulakey:"edgeto" json:"to,omitempty"`
	Relation     string  `nebulaproperty:"relation" description:"relation" nebulaindexes:"relation" json:"relation,omitempty"`
	StoryID      int     `nebulaproperty:"story_id" mappingalias:"PlayID" description:"game play id" nebulaindexes:"story_id" json:"story_id,omitempty"`
}

func NewPeople_FamilyPeople(p *People, fp *People) *People_FamilyPeople {
	return &People_FamilyPeople{
		People:       p,
		FamilyPeople: fp,
		StoryID:      p.StoryID,
	}
}

func (pfp *People_FamilyPeople) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertEdges(space, pfp)
}

func (pfp *People_FamilyPeople) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateEdges(space, pfp)
}

func (pfp *People_FamilyPeople) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertEdges(space, pfp)
}

func (pfp *People_FamilyPeople) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, pfp)
}

func (pfp *People_FamilyPeople) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadEdge(space, pfp)
}

func InsertPeople_FamilyPeoples(space *nebulagolang.Space, pfps ...*People_FamilyPeople) *nebulagolang.Result {
	return nebulagolang.BatchInsertEdges(space, 250, pfps)
}

func UpdatePeople_FamilyPeoples(space *nebulagolang.Space, pfps ...*People_FamilyPeople) *nebulagolang.Result {
	return nebulagolang.BatchUpdateEdges(space, 250, pfps)
}

func UpsertPeople_FamilyPeoples(space *nebulagolang.Space, pfps ...*People_FamilyPeople) *nebulagolang.Result {
	return nebulagolang.BatchUpsertEdges(space, 250, pfps)
}

func DeletePeople_FamilyPeoples(space *nebulagolang.Space, pfps ...*People_FamilyPeople) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, pfps...)
}

func DeleteAllPeople_FamilyPeoples(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByEdgeType[People_FamilyPeople](space)
}

func DeleteAllPeople_FamilyByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByQuery[People_FamilyPeople](space, getPlayIdQuery[People_FamilyPeople](playId))
}

func GetPeople_FamilyPeopleByEid(space *nebulagolang.Space, eid *nebulagolang.EID) *nebulagolang.ResultT[*People_FamilyPeople] {
	return nebulagolang.GetEdgeByEid[*People_FamilyPeople](space, eid)
}

func GetPeople_FamilyPeopleByFromAndToIds(space *nebulagolang.Space, fromId string, toId string) *nebulagolang.ResultT[*People_FamilyPeople] {
	return nebulagolang.GetEdgeByEid[*People_FamilyPeople](space, nebulagolang.NewEID(fromId, toId, nebulagolang.GetEdgeName[People_FamilyPeople]()))
}

func GetAllPeople_FamilyPeoplesEids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*People_FamilyPeople](space, "")
}

func GetAllPeople_FamilyPeoplesEidsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*People_FamilyPeople](space, getPlayIdQuery[People_FamilyPeople](playId))
}

func GetAllPeople_FamilyPeoples(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*People_FamilyPeople] {
	return nebulagolang.GetAllEdgesByEdgeType[*People_FamilyPeople](space)
}

func GetAllPeople_FamilyPeoplesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]*People_FamilyPeople] {
	return nebulagolang.GetAllEdgesByQuery[*People_FamilyPeople](space, getPlayIdQuery[People_FamilyPeople](playId))
}
