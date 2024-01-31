package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
)

type People_Dynasty struct {
	People  *People  `nebulaedgename:"people_dynasty" nebulaedgecomment:"People -> Dynasty" nebulakey:"edgefrom" json:"from,omitempty"`
	Dynasty *Dynasty `nebulakey:"edgeto" json:"to,omitempty"`
	StoryID int      `nebulaproperty:"story_id" mappingalias:"PlayID" description:"game play id" nebulaindexes:"story_id" json:"story_id,omitempty"`
}

func NewPeople_Dynasty(p *People, c *Dynasty) *People_Dynasty {
	return &People_Dynasty{
		People:  p,
		Dynasty: c,
		StoryID: p.StoryID,
	}
}

func (pd *People_Dynasty) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertEdges(space, pd)
}

func (pd *People_Dynasty) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateEdges(space, pd)
}

func (pd *People_Dynasty) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertEdges(space, pd)
}

func (pd *People_Dynasty) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, pd)
}

func (pd *People_Dynasty) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadEdge(space, pd)
}

func InsertPeople_Dynastys(space *nebulagolang.Space, pds ...*People_Dynasty) *nebulagolang.Result {
	return nebulagolang.BatchInsertEdges(space, 250, pds)
}

func UpdatePeople_Dynastys(space *nebulagolang.Space, pds ...*People_Dynasty) *nebulagolang.Result {
	return nebulagolang.BatchUpdateEdges(space, 250, pds)
}

func UpsertPeople_Dynastys(space *nebulagolang.Space, pds ...*People_Dynasty) *nebulagolang.Result {
	return nebulagolang.BatchUpsertEdges(space, 250, pds)
}

func DeletePeople_Dynastys(space *nebulagolang.Space, pds ...*People_Dynasty) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, pds...)
}

func DeleteAllPeople_Dynastys(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByEdgeType[People_Dynasty](space)
}

func DeleteAllPeople_DynastiesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByQuery[People_Dynasty](space, getPlayIdQuery[People_Dynasty](playId))
}

func GetPeople_DynastyByEid(space *nebulagolang.Space, eid *nebulagolang.EID) *nebulagolang.ResultT[*People_Dynasty] {
	return nebulagolang.GetEdgeByEid[*People_Dynasty](space, eid)
}

func GetPeople_DynastyByFromAndToIds(space *nebulagolang.Space, fromId string, toId string) *nebulagolang.ResultT[*People_Dynasty] {
	return nebulagolang.GetEdgeByEid[*People_Dynasty](space, nebulagolang.NewEID(fromId, toId, nebulagolang.GetEdgeName[People_Dynasty]()))
}

func GetAllPeople_DynastysEids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*People_Dynasty](space, "")
}

func GetAllPeople_DynastysEidsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*People_Dynasty](space, getPlayIdQuery[People_Dynasty](playId))
}

func GetAllPeople_Dynastys(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*People_Dynasty] {
	return nebulagolang.GetAllEdgesByEdgeType[*People_Dynasty](space)
}

func GetAllPeople_DynastysByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]*People_Dynasty] {
	return nebulagolang.GetAllEdgesByQuery[*People_Dynasty](space, getPlayIdQuery[People_Dynasty](playId))
}
