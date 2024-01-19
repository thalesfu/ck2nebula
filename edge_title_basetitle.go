package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
	"time"
)

type Title_BaseTitle struct {
	Title     *Title    `nebulaedgename:"title_basetitle" nebulaedgecomment:"Title -> Base Title" nebulakey:"edgefrom" json:"from,omitempty"`
	BaseTitle *Title    `nebulakey:"edgeto" json:"to,omitempty"`
	PlayID    int       `nebulaproperty:"play_id" description:"game play id" nebulaindexes:"play_id" json:"play_id,omitempty"`
	PlayDate  time.Time `nebulaproperty:"play_date" nebulatype:"Date" description:"game play date" nebulaindexes:"play_date" json:"play_date,omitempty"`
}

func NewTitle_BaseTitle(t *Title, bt *Title) *Title_BaseTitle {
	return &Title_BaseTitle{
		Title:     t,
		BaseTitle: bt,
	}
}

func (tbt *Title_BaseTitle) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertEdges(space, tbt)
}

func (tbt *Title_BaseTitle) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateEdges(space, tbt)
}

func (tbt *Title_BaseTitle) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertEdges(space, tbt)
}

func (tbt *Title_BaseTitle) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, tbt)
}

func (tbt *Title_BaseTitle) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadEdge(space, tbt)
}

func InsertTitle_BaseTitles(space *nebulagolang.Space, tbts ...*Title_BaseTitle) *nebulagolang.Result {
	return nebulagolang.BatchInsertEdges(space, 10000, tbts)
}

func UpdateTitle_BaseTitles(space *nebulagolang.Space, tbts ...*Title_BaseTitle) *nebulagolang.Result {
	return nebulagolang.BatchUpdateEdges(space, 5000, tbts)
}

func UpsertTitle_BaseTitles(space *nebulagolang.Space, tbts ...*Title_BaseTitle) *nebulagolang.Result {
	return nebulagolang.BatchUpsertEdges(space, 5000, tbts)
}

func DeleteTitle_BaseTitles(space *nebulagolang.Space, tbts ...*Title_BaseTitle) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, tbts...)
}

func DeleteAllTitle_BaseTitles(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByEdgeType[Title_BaseTitle](space)
}

func GetTitle_BaseTitleByEid(space *nebulagolang.Space, eid *nebulagolang.EID) *nebulagolang.ResultT[*Title_BaseTitle] {
	return nebulagolang.GetEdgeByEid[*Title_BaseTitle](space, eid)
}

func GetTitle_BaseTitleByFromAndToIds(space *nebulagolang.Space, fromId string, toId string) *nebulagolang.ResultT[*Title_BaseTitle] {
	return nebulagolang.GetEdgeByEid[*Title_BaseTitle](space, nebulagolang.NewEID(fromId, toId, nebulagolang.GetEdgeName[Title_BaseTitle]()))
}

func GetAllTitle_BaseTitlesEids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*Title_BaseTitle](space, "")
}

func GetAllTitle_BaseTitles(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*Title_BaseTitle] {
	return nebulagolang.GetAllEdgesByEdgeType[*Title_BaseTitle](space)
}
