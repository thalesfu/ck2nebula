package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
)

type Story_Title struct {
	Story  *Story `nebulaedgename:"story_title" nebulaedgecomment:"Story -> Title" nebulakey:"edgefrom" json:"from,omitempty"`
	Title  *Title `nebulakey:"edgeto" json:"to,omitempty"`
	PlayID int    `nebulaproperty:"play_id" description:"game play id" nebulaindexes:"play_id" json:"play_id,omitempty"`
}

func NewStory_Title(s *Story, t *Title) *Story_Title {
	return &Story_Title{
		Story:  s,
		Title:  t,
		PlayID: s.PlayID,
	}
}

func (st *Story_Title) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertEdges(space, st)
}

func (st *Story_Title) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateEdges(space, st)
}

func (st *Story_Title) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertEdges(space, st)
}

func (st *Story_Title) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, st)
}

func (st *Story_Title) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadEdge(space, st)
}

func InsertStory_Titles(space *nebulagolang.Space, sts ...*Story_Title) *nebulagolang.Result {
	return nebulagolang.BatchInsertEdges(space, 250, sts)
}

func UpdateStory_Titles(space *nebulagolang.Space, sts ...*Story_Title) *nebulagolang.Result {
	return nebulagolang.BatchUpdateEdges(space, 250, sts)
}

func UpsertStory_Titles(space *nebulagolang.Space, sts ...*Story_Title) *nebulagolang.Result {
	return nebulagolang.BatchUpsertEdges(space, 250, sts)
}

func DeleteStory_Titles(space *nebulagolang.Space, sts ...*Story_Title) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, sts...)
}

func DeleteAllStory_Titles(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByEdgeType[Story_Title](space)
}

func DeleteAllStory_TitlesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByQuery[Story_Title](space, getPlayIdQuery[Story_Title](playId))
}

func GetStory_TitleByEid(space *nebulagolang.Space, eid *nebulagolang.EID) *nebulagolang.ResultT[*Story_Title] {
	return nebulagolang.GetEdgeByEid[*Story_Title](space, eid)
}

func GetStory_TitleByFromAndToIds(space *nebulagolang.Space, fromId string, toId string) *nebulagolang.ResultT[*Story_Title] {
	return nebulagolang.GetEdgeByEid[*Story_Title](space, nebulagolang.NewEID(fromId, toId, nebulagolang.GetEdgeName[Story_Title]()))
}

func GetAllStory_TitlesEids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*Story_Title](space, "")
}

func GetAllStory_TitlesEidsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*Story_Title](space, getPlayIdQuery[Story_Title](playId))
}

func GetAllStory_Titles(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*Story_Title] {
	return nebulagolang.GetAllEdgesByEdgeType[*Story_Title](space)
}

func GetAllStory_TitlesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]*Story_Title] {
	return nebulagolang.GetAllEdgesByQuery[*Story_Title](space, getPlayIdQuery[Story_Title](playId))
}
