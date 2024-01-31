package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
)

type Story_People struct {
	Story   *Story  `nebulaedgename:"story_people" nebulaedgecomment:"Story -> People" nebulakey:"edgefrom" json:"from,omitempty"`
	People  *People `nebulakey:"edgeto" json:"to,omitempty"`
	StoryID int     `nebulaproperty:"story_id" mappingalias:"PlayID" description:"game play id" nebulaindexes:"story_id" json:"story_id,omitempty"`
}

func NewStory_People(s *Story, p *People) *Story_People {
	return &Story_People{
		Story:   s,
		People:  p,
		StoryID: s.StoryId,
	}
}

func (sp *Story_People) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertEdges(space, sp)
}

func (sp *Story_People) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateEdges(space, sp)
}

func (sp *Story_People) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertEdges(space, sp)
}

func (sp *Story_People) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, sp)
}

func (sp *Story_People) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadEdge(space, sp)
}

func InsertStory_Peoples(space *nebulagolang.Space, sps ...*Story_People) *nebulagolang.Result {
	return nebulagolang.BatchInsertEdges(space, 250, sps)
}

func UpdateStory_Peoples(space *nebulagolang.Space, sps ...*Story_People) *nebulagolang.Result {
	return nebulagolang.BatchUpdateEdges(space, 250, sps)
}

func UpsertStory_Peoples(space *nebulagolang.Space, sps ...*Story_People) *nebulagolang.Result {
	return nebulagolang.BatchUpsertEdges(space, 250, sps)
}

func DeleteStory_Peoples(space *nebulagolang.Space, sps ...*Story_People) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, sps...)
}

func DeleteAllStory_Peoples(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByEdgeType[Story_People](space)
}

func DeleteAllStory_PeopleByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByQuery[Story_People](space, getPlayIdQuery[Story_People](playId))
}

func GetStory_PeopleByEid(space *nebulagolang.Space, eid *nebulagolang.EID) *nebulagolang.ResultT[*Story_People] {
	return nebulagolang.GetEdgeByEid[*Story_People](space, eid)
}

func GetStory_PeopleByFromAndToIds(space *nebulagolang.Space, fromId string, toId string) *nebulagolang.ResultT[*Story_People] {
	return nebulagolang.GetEdgeByEid[*Story_People](space, nebulagolang.NewEID(fromId, toId, nebulagolang.GetEdgeName[Story_People]()))
}

func GetAllStory_PeoplesEids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*Story_People](space, "")
}

func GetAllStory_PeoplesEidsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*Story_People](space, getPlayIdQuery[Story_People](playId))
}

func GetAllStory_Peoples(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*Story_People] {
	return nebulagolang.GetAllEdgesByEdgeType[*Story_People](space)
}

func GetAllStory_PeoplesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]*Story_People] {
	return nebulagolang.GetAllEdgesByQuery[*Story_People](space, getPlayIdQuery[Story_People](playId))
}
