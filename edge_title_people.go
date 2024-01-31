package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
)

type Title_People struct {
	Title   *Title  `nebulaedgename:"title_people" nebulaedgecomment:"Title -> People" nebulakey:"edgefrom" json:"from,omitempty"`
	People  *People `nebulakey:"edgeto" json:"to,omitempty"`
	StoryID int     `nebulaproperty:"story_id" mappingalias:"PlayID" description:"game play id" nebulaindexes:"story_id" json:"story_id,omitempty"`
}

func NewTitle_People(t *Title, d *People) *Title_People {
	return &Title_People{
		Title:   t,
		People:  d,
		StoryID: t.StoryID,
	}
}

func (tp *Title_People) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertEdges(space, tp)
}

func (tp *Title_People) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateEdges(space, tp)
}

func (tp *Title_People) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertEdges(space, tp)
}

func (tp *Title_People) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, tp)
}

func (tp *Title_People) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadEdge(space, tp)
}

func InsertTitle_Peoples(space *nebulagolang.Space, tps ...*Title_People) *nebulagolang.Result {
	return nebulagolang.BatchInsertEdges(space, 250, tps)
}

func UpdateTitle_Peoples(space *nebulagolang.Space, tps ...*Title_People) *nebulagolang.Result {
	return nebulagolang.BatchUpdateEdges(space, 250, tps)
}

func UpsertTitle_Peoples(space *nebulagolang.Space, tps ...*Title_People) *nebulagolang.Result {
	return nebulagolang.BatchUpsertEdges(space, 250, tps)
}

func DeleteTitle_Peoples(space *nebulagolang.Space, tps ...*Title_People) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, tps...)
}

func DeleteAllTitle_Peoples(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByEdgeType[Title_People](space)
}

func DeleteAllTitle_PeoplesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByQuery[Title_People](space, getPlayIdQuery[Title_People](playId))
}

func GetTitle_PeopleByEid(space *nebulagolang.Space, eid *nebulagolang.EID) *nebulagolang.ResultT[*Title_People] {
	return nebulagolang.GetEdgeByEid[*Title_People](space, eid)
}

func GetTitle_PeopleByFromAndToIds(space *nebulagolang.Space, fromId string, toId string) *nebulagolang.ResultT[*Title_People] {
	return nebulagolang.GetEdgeByEid[*Title_People](space, nebulagolang.NewEID(fromId, toId, nebulagolang.GetEdgeName[Title_People]()))
}

func GetAllTitle_PeoplesEids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*Title_People](space, "")
}

func GetAllTitle_PeoplesEidsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*Title_People](space, getPlayIdQuery[Title_People](playId))
}

func GetAllTitle_Peoples(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*Title_People] {
	return nebulagolang.GetAllEdgesByEdgeType[*Title_People](space)
}

func GetAllTitle_PeoplesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]*Title_People] {
	return nebulagolang.GetAllEdgesByQuery[*Title_People](space, getPlayIdQuery[Title_People](playId))
}
