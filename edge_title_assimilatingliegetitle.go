package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
)

type Title_AssimilatingLiegeTitle struct {
	Title                  *Title `nebulaedgename:"title_liegeassimilatingtitle" nebulaedgecomment:"Title -> Assimilating Dejure Liege Title" nebulakey:"edgefrom" json:"from,omitempty"`
	AssimilatingLiegeTitle *Title `nebulakey:"edgeto" json:"to,omitempty"`
	DeJureAssYears         int    `nebulaproperty:"de_jure_ass_years" description:"assimilating dejure liege years" nebulaindexes:"de_jure_ass_years" json:"de_jure_ass_years,omitempty"`
	StoryID                int    `nebulaproperty:"story_id" mappingalias:"PlayID" description:"game play id" nebulaindexes:"story_id" json:"story_id,omitempty"`
}

func NewTitle_AssimilatingLiegeTitle(t *Title, alt *Title) *Title_AssimilatingLiegeTitle {
	return &Title_AssimilatingLiegeTitle{
		Title:                  t,
		AssimilatingLiegeTitle: alt,
		StoryID:                t.StoryID,
	}
}

func (talt *Title_AssimilatingLiegeTitle) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertEdges(space, talt)
}

func (talt *Title_AssimilatingLiegeTitle) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateEdges(space, talt)
}

func (talt *Title_AssimilatingLiegeTitle) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertEdges(space, talt)
}

func (talt *Title_AssimilatingLiegeTitle) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, talt)
}

func (talt *Title_AssimilatingLiegeTitle) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadEdge(space, talt)
}

func InsertTitle_AssimilatingLiegeTitles(space *nebulagolang.Space, talts ...*Title_AssimilatingLiegeTitle) *nebulagolang.Result {
	return nebulagolang.BatchInsertEdges(space, 250, talts)
}

func UpdateTitle_AssimilatingLiegeTitles(space *nebulagolang.Space, talts ...*Title_AssimilatingLiegeTitle) *nebulagolang.Result {
	return nebulagolang.BatchUpdateEdges(space, 250, talts)
}

func UpsertTitle_AssimilatingLiegeTitles(space *nebulagolang.Space, talts ...*Title_AssimilatingLiegeTitle) *nebulagolang.Result {
	return nebulagolang.BatchUpsertEdges(space, 250, talts)
}

func DeleteTitle_AssimilatingLiegeTitles(space *nebulagolang.Space, talts ...*Title_AssimilatingLiegeTitle) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, talts...)
}

func DeleteAllTitle_AssimilatingLiegeTitles(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByEdgeType[Title_AssimilatingLiegeTitle](space)
}

func DeleteAllTitle_AssimilatingLiegeTitlesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByQuery[Title_AssimilatingLiegeTitle](space, getPlayIdQuery[Title_AssimilatingLiegeTitle](playId))
}

func GetTitle_AssimilatingLiegeTitleByEid(space *nebulagolang.Space, eid *nebulagolang.EID) *nebulagolang.ResultT[*Title_AssimilatingLiegeTitle] {
	return nebulagolang.GetEdgeByEid[*Title_AssimilatingLiegeTitle](space, eid)
}

func GetTitle_AssimilatingLiegeTitleByFromAndToIds(space *nebulagolang.Space, fromId string, toId string) *nebulagolang.ResultT[*Title_AssimilatingLiegeTitle] {
	return nebulagolang.GetEdgeByEid[*Title_AssimilatingLiegeTitle](space, nebulagolang.NewEID(fromId, toId, nebulagolang.GetEdgeName[Title_AssimilatingLiegeTitle]()))
}

func GetAllTitle_AssimilatingLiegeTitlesEids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*Title_AssimilatingLiegeTitle](space, "")
}

func GetAllTitle_AssimilatingLiegeTitlesEidsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*Title_AssimilatingLiegeTitle](space, getPlayIdQuery[Title_AssimilatingLiegeTitle](playId))
}

func GetAllTitle_AssimilatingLiegeTitles(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*Title_AssimilatingLiegeTitle] {
	return nebulagolang.GetAllEdgesByEdgeType[*Title_AssimilatingLiegeTitle](space)
}

func GetAllTitle_AssimilatingLiegeTitlesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]*Title_AssimilatingLiegeTitle] {
	return nebulagolang.GetAllEdgesByQuery[*Title_AssimilatingLiegeTitle](space, getPlayIdQuery[Title_AssimilatingLiegeTitle](playId))
}
