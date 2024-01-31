package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
)

type People_Religion struct {
	People   *People   `nebulaedgename:"people_religion" nebulaedgecomment:"People -> Religion" nebulakey:"edgefrom" json:"from,omitempty"`
	Religion *Religion `nebulakey:"edgeto" json:"to,omitempty"`
	StoryID  int       `nebulaproperty:"story_id" mappingalias:"PlayID" description:"game play id" nebulaindexes:"story_id" json:"story_id,omitempty"`
}

func NewPeople_Religion(p *People, r *Religion) *People_Religion {
	return &People_Religion{
		People:   p,
		Religion: r,
		StoryID:  p.StoryID,
	}
}

func (pr *People_Religion) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertEdges(space, pr)
}

func (pr *People_Religion) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateEdges(space, pr)
}

func (pr *People_Religion) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertEdges(space, pr)
}

func (pr *People_Religion) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, pr)
}

func (pr *People_Religion) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadEdge(space, pr)
}

func InsertPeople_Religions(space *nebulagolang.Space, prs ...*People_Religion) *nebulagolang.Result {
	return nebulagolang.BatchInsertEdges(space, 250, prs)
}

func UpdatePeople_Religions(space *nebulagolang.Space, prs ...*People_Religion) *nebulagolang.Result {
	return nebulagolang.BatchUpdateEdges(space, 250, prs)
}

func UpsertPeople_Religions(space *nebulagolang.Space, prs ...*People_Religion) *nebulagolang.Result {
	return nebulagolang.BatchUpsertEdges(space, 250, prs)
}

func DeletePeople_Religions(space *nebulagolang.Space, prs ...*People_Religion) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, prs...)
}

func DeleteAllPeople_Religions(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByEdgeType[People_Religion](space)
}

func DeleteAllPeople_ReligionsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByQuery[People_Religion](space, getPlayIdQuery[People_Religion](playId))
}

func GetPeople_ReligionByEid(space *nebulagolang.Space, eid *nebulagolang.EID) *nebulagolang.ResultT[*People_Religion] {
	return nebulagolang.GetEdgeByEid[*People_Religion](space, eid)
}

func GetPeople_ReligionByFromAndToIds(space *nebulagolang.Space, fromId string, toId string) *nebulagolang.ResultT[*People_Religion] {
	return nebulagolang.GetEdgeByEid[*People_Religion](space, nebulagolang.NewEID(fromId, toId, nebulagolang.GetEdgeName[People_Religion]()))
}

func GetAllPeople_ReligionsEids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*People_Religion](space, "")
}

func GetAllPeople_ReligionsEidsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*People_Religion](space, getPlayIdQuery[People_Religion](playId))
}

func GetAllPeople_Religions(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*People_Religion] {
	return nebulagolang.GetAllEdgesByEdgeType[*People_Religion](space)
}

func GetAllPeople_ReligionsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]*People_Religion] {
	return nebulagolang.GetAllEdgesByQuery[*People_Religion](space, getPlayIdQuery[People_Religion](playId))
}
