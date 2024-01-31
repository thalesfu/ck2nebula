package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
)

type Province_Title struct {
	Province *Province `nebulaedgename:"province_title" nebulaedgecomment:"Province -> Title" nebulakey:"edgefrom" json:"from,omitempty"`
	Title    *Title    `nebulakey:"edgeto" json:"to,omitempty"`
	StoryID  int       `nebulaproperty:"story_id" mappingalias:"PlayID" description:"game play id" nebulaindexes:"story_id" json:"story_id,omitempty"`
}

func NewProvince_Title(p *Province, t *Title) *Province_Title {
	return &Province_Title{
		Province: p,
		Title:    t,
		StoryID:  p.StoryID,
	}
}

func (pt *Province_Title) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertEdges(space, pt)
}

func (pt *Province_Title) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateEdges(space, pt)
}

func (pt *Province_Title) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertEdges(space, pt)
}

func (pt *Province_Title) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, pt)
}

func (pt *Province_Title) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadEdge(space, pt)
}

func InsertProvince_Titles(space *nebulagolang.Space, pts ...*Province_Title) *nebulagolang.Result {
	return nebulagolang.BatchInsertEdges(space, 250, pts)
}

func UpdateProvince_Titles(space *nebulagolang.Space, pts ...*Province_Title) *nebulagolang.Result {
	return nebulagolang.BatchUpdateEdges(space, 250, pts)
}

func UpsertProvince_Titles(space *nebulagolang.Space, pts ...*Province_Title) *nebulagolang.Result {
	return nebulagolang.BatchUpsertEdges(space, 250, pts)
}

func DeleteProvince_Titles(space *nebulagolang.Space, pts ...*Province_Title) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, pts...)
}

func DeleteAllProvince_Titles(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByEdgeType[Province_Title](space)
}

func DeleteAllProvince_TitlesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByQuery[Province_Title](space, getPlayIdQuery[Province_Title](playId))
}

func GetProvince_TitleByEid(space *nebulagolang.Space, eid *nebulagolang.EID) *nebulagolang.ResultT[*Province_Title] {
	return nebulagolang.GetEdgeByEid[*Province_Title](space, eid)
}

func GetProvince_TitleByFromAndToIds(space *nebulagolang.Space, fromId string, toId string) *nebulagolang.ResultT[*Province_Title] {
	return nebulagolang.GetEdgeByEid[*Province_Title](space, nebulagolang.NewEID(fromId, toId, nebulagolang.GetEdgeName[Province_Title]()))
}

func GetAllProvince_TitlesEids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*Province_Title](space, "")
}

func GetAllProvince_TitlesEidsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*Province_Title](space, getPlayIdQuery[Province_Title](playId))
}

func GetAllProvince_Titles(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*Province_Title] {
	return nebulagolang.GetAllEdgesByEdgeType[*Province_Title](space)
}

func GetAllProvince_TitlesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]*Province_Title] {
	return nebulagolang.GetAllEdgesByQuery[*Province_Title](space, getPlayIdQuery[Province_Title](playId))
}
