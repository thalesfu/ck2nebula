package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
	"time"
)

type Title_DejureLiegeTitle struct {
	Title          *Title    `nebulaedgename:"title_dejureliegetitle" nebulaedgecomment:"Title -> Dejure Liege Title" nebulakey:"edgefrom" json:"from,omitempty"`
	LiegeJureTitle *Title    `nebulakey:"edgeto" json:"to,omitempty"`
	PlayID         int       `nebulaproperty:"play_id" description:"game play id" nebulaindexes:"play_id" json:"play_id,omitempty"`
	PlayDate       time.Time `nebulaproperty:"play_date" nebulatype:"Date" description:"game play date" nebulaindexes:"play_date" json:"play_date,omitempty"`
}

func NewTitle_DejureLiegeTitle(t *Title, dlt *Title) *Title_DejureLiegeTitle {
	return &Title_DejureLiegeTitle{
		Title:          t,
		LiegeJureTitle: dlt,
	}
}

func (tdlts *Title_DejureLiegeTitle) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertEdges(space, tdlts)
}

func (tdlts *Title_DejureLiegeTitle) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateEdges(space, tdlts)
}

func (tdlts *Title_DejureLiegeTitle) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertEdges(space, tdlts)
}

func (tdlts *Title_DejureLiegeTitle) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, tdlts)
}

func (tdlts *Title_DejureLiegeTitle) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadEdge(space, tdlts)
}

func InsertTitle_DejureLiegeTitles(space *nebulagolang.Space, tdlts ...*Title_DejureLiegeTitle) *nebulagolang.Result {
	return nebulagolang.BatchInsertEdges(space, 10000, tdlts)
}

func UpdateTitle_DejureLiegeTitles(space *nebulagolang.Space, tdlts ...*Title_DejureLiegeTitle) *nebulagolang.Result {
	return nebulagolang.BatchUpdateEdges(space, 5000, tdlts)
}

func UpsertTitle_DejureLiegeTitles(space *nebulagolang.Space, tdlts ...*Title_DejureLiegeTitle) *nebulagolang.Result {
	return nebulagolang.BatchUpsertEdges(space, 5000, tdlts)
}

func DeleteTitle_DejureLiegeTitles(space *nebulagolang.Space, tdlts ...*Title_DejureLiegeTitle) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, tdlts...)
}

func DeleteAllTitle_DejureLiegeTitles(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByEdgeType[Title_DejureLiegeTitle](space)
}

func GetTitle_DejureLiegeTitleByEid(space *nebulagolang.Space, eid *nebulagolang.EID) *nebulagolang.ResultT[*Title_DejureLiegeTitle] {
	return nebulagolang.GetEdgeByEid[*Title_DejureLiegeTitle](space, eid)
}

func GetTitle_DejureLiegeTitleByFromAndToIds(space *nebulagolang.Space, fromId string, toId string) *nebulagolang.ResultT[*Title_DejureLiegeTitle] {
	return nebulagolang.GetEdgeByEid[*Title_DejureLiegeTitle](space, nebulagolang.NewEID(fromId, toId, nebulagolang.GetEdgeName[Title_DejureLiegeTitle]()))
}

func GetAllTitle_DejureLiegeTitlesEids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*Title_DejureLiegeTitle](space, "")
}

func GetAllTitle_DejureLiegeTitles(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*Title_DejureLiegeTitle] {
	return nebulagolang.GetAllEdgesByEdgeType[*Title_DejureLiegeTitle](space)
}
