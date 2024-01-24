package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
)

type People_GFXCulture struct {
	People     *People  `nebulaedgename:"people_gfx_culture" nebulaedgecomment:"People -> Gfx Culture" nebulakey:"edgefrom" json:"from,omitempty"`
	GFXCulture *Culture `nebulakey:"edgeto" json:"to,omitempty"`
	PlayID     int      `nebulaproperty:"play_id" description:"game play id" nebulaindexes:"play_id" json:"play_id,omitempty"`
}

func NewPeople_GFXCulture(p *People, gfx_c *Culture) *People_GFXCulture {
	return &People_GFXCulture{
		People:     p,
		GFXCulture: gfx_c,
		PlayID:     p.PlayID,
	}
}

func (pgfxc *People_GFXCulture) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertEdges(space, pgfxc)
}

func (pgfxc *People_GFXCulture) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateEdges(space, pgfxc)
}

func (pgfxc *People_GFXCulture) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertEdges(space, pgfxc)
}

func (pgfxc *People_GFXCulture) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, pgfxc)
}

func (pgfxc *People_GFXCulture) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadEdge(space, pgfxc)
}

func InsertPeople_GFXCultures(space *nebulagolang.Space, pgfxcs ...*People_GFXCulture) *nebulagolang.Result {
	return nebulagolang.BatchInsertEdges(space, 250, pgfxcs)
}

func UpdatePeople_GFXCultures(space *nebulagolang.Space, pgfxcs ...*People_GFXCulture) *nebulagolang.Result {
	return nebulagolang.BatchUpdateEdges(space, 250, pgfxcs)
}

func UpsertPeople_GFXCultures(space *nebulagolang.Space, pgfxcs ...*People_GFXCulture) *nebulagolang.Result {
	return nebulagolang.BatchUpsertEdges(space, 250, pgfxcs)
}

func DeletePeople_GFXCultures(space *nebulagolang.Space, pgfxcs ...*People_GFXCulture) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, pgfxcs...)
}

func DeleteAllPeople_GFXCultures(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByEdgeType[People_GFXCulture](space)
}

func DeleteAllPeople_GFXCulturesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByQuery[People_GFXCulture](space, getPlayIdQuery[People_GFXCulture](playId))
}

func GetPeople_GFXCultureByEid(space *nebulagolang.Space, eid *nebulagolang.EID) *nebulagolang.ResultT[*People_GFXCulture] {
	return nebulagolang.GetEdgeByEid[*People_GFXCulture](space, eid)
}

func GetPeople_GFXCultureByFromAndToIds(space *nebulagolang.Space, fromId string, toId string) *nebulagolang.ResultT[*People_GFXCulture] {
	return nebulagolang.GetEdgeByEid[*People_GFXCulture](space, nebulagolang.NewEID(fromId, toId, nebulagolang.GetEdgeName[People_GFXCulture]()))
}

func GetAllPeople_GFXCulturesEids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*People_GFXCulture](space, "")
}

func GetAllPeople_GFXCulturesEidsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*People_GFXCulture](space, getPlayIdQuery[People_GFXCulture](playId))
}

func GetAllPeople_GFXCultures(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*People_GFXCulture] {
	return nebulagolang.GetAllEdgesByEdgeType[*People_GFXCulture](space)
}

func GetAllPeople_GFXCulturesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]*People_GFXCulture] {
	return nebulagolang.GetAllEdgesByQuery[*People_GFXCulture](space, getPlayIdQuery[People_GFXCulture](playId))
}
