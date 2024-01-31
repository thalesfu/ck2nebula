package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
)

type Baron_Title struct {
	Baron   *Baron `nebulaedgename:"baron_title" nebulaedgecomment:"Baron -> Title" nebulakey:"edgefrom" json:"from,omitempty"`
	Title   *Title `nebulakey:"edgeto" json:"to,omitempty"`
	StoryID int    `nebulaproperty:"story_id" mappingalias:"PlayID" description:"game play id" nebulaindexes:"story_id" json:"story_id,omitempty"`
}

func NewBaron_Title(b *Baron, t *Title) *Baron_Title {
	return &Baron_Title{
		Baron:   b,
		Title:   t,
		StoryID: b.StoryID,
	}
}

func (bt *Baron_Title) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertEdges(space, bt)
}

func (bt *Baron_Title) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateEdges(space, bt)
}

func (bt *Baron_Title) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertEdges(space, bt)
}

func (bt *Baron_Title) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, bt)
}

func (bt *Baron_Title) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadEdge(space, bt)
}

func InsertBaron_Titles(space *nebulagolang.Space, bts ...*Baron_Title) *nebulagolang.Result {
	return nebulagolang.BatchInsertEdges(space, 250, bts)
}

func UpdateBaron_Titles(space *nebulagolang.Space, bts ...*Baron_Title) *nebulagolang.Result {
	return nebulagolang.BatchUpdateEdges(space, 250, bts)
}

func UpsertBaron_Titles(space *nebulagolang.Space, bts ...*Baron_Title) *nebulagolang.Result {
	return nebulagolang.BatchUpsertEdges(space, 250, bts)
}

func DeleteBaron_Titles(space *nebulagolang.Space, bts ...*Baron_Title) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, bts...)
}

func DeleteAllBaron_Titles(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByEdgeType[Baron_Title](space)
}

func DeleteAllBaron_TitlesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByQuery[Baron_Title](space, getPlayIdQuery[Baron_Title](playId))
}

func GetBaron_TitleByEid(space *nebulagolang.Space, eid *nebulagolang.EID) *nebulagolang.ResultT[*Baron_Title] {
	return nebulagolang.GetEdgeByEid[*Baron_Title](space, eid)
}

func GetBaron_TitleByFromAndToIds(space *nebulagolang.Space, fromId string, toId string) *nebulagolang.ResultT[*Baron_Title] {
	return nebulagolang.GetEdgeByEid[*Baron_Title](space, nebulagolang.NewEID(fromId, toId, nebulagolang.GetEdgeName[Baron_Title]()))
}

func GetAllBaron_TitlesEids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*Baron_Title](space, "")
}

func GetAllBaron_TitlesEidsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*Baron_Title](space, getPlayIdQuery[Baron_Title](playId))
}

func GetAllBaron_Titles(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*Baron_Title] {
	return nebulagolang.GetAllEdgesByEdgeType[*Baron_Title](space)
}

func GetAllBaron_TitlesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]*Baron_Title] {
	return nebulagolang.GetAllEdgesByQuery[*Baron_Title](space, getPlayIdQuery[Baron_Title](playId))
}
