package ck2nebula

import "github.com/thalesfu/nebulagolang"

type CultureGroup_Culture struct {
	CultureGroup *CultureGroup `nebulaedgename:"culturegroup_culture" nebulaedgecomment:"culture group -> culture" nebulakey:"edgefrom" json:"from,omitempty"`
	Culture      *Culture      `nebulakey:"edgeto" json:"to,omitempty"`
}

func NewCultureGroup_Culture(cg *CultureGroup, c *Culture) *CultureGroup_Culture {
	return &CultureGroup_Culture{
		CultureGroup: cg,
		Culture:      c,
	}
}

func (cgc *CultureGroup_Culture) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertEdges(space, cgc)
}

func (cgc *CultureGroup_Culture) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateEdges(space, cgc)
}

func (cgc *CultureGroup_Culture) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertEdges(space, cgc)
}

func (cgc *CultureGroup_Culture) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, cgc)
}

func (cgc *CultureGroup_Culture) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadEdge(space, cgc)
}

func InsertCultureGroup_Cultures(space *nebulagolang.Space, cgcs ...*CultureGroup_Culture) *nebulagolang.Result {
	return nebulagolang.BatchInsertEdges(space, 10000, cgcs)
}

func UpdateCultureGroup_Cultures(space *nebulagolang.Space, cgcs ...*CultureGroup_Culture) *nebulagolang.Result {
	return nebulagolang.BatchUpdateEdges(space, 5000, cgcs)
}

func UpsertCultureGroup_Cultures(space *nebulagolang.Space, cgcs ...*CultureGroup_Culture) *nebulagolang.Result {
	return nebulagolang.BatchUpsertEdges(space, 5000, cgcs)
}

func DeleteCultureGroup_Cultures(space *nebulagolang.Space, cgcs ...*CultureGroup_Culture) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, cgcs...)
}

func DeleteAllCultureGroup_Cultures(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByEdgeType[CultureGroup_Culture](space)
}

func GetCultureGroup_CultureByEid(space *nebulagolang.Space, eid *nebulagolang.EID) *nebulagolang.ResultT[*CultureGroup_Culture] {
	return nebulagolang.GetEdgeByEid[*CultureGroup_Culture](space, eid)
}

func GetCultureGroup_CultureByFromAndToIds(space *nebulagolang.Space, fromId string, toId string) *nebulagolang.ResultT[*CultureGroup_Culture] {
	return nebulagolang.GetEdgeByEid[*CultureGroup_Culture](space, nebulagolang.NewEID(fromId, toId, nebulagolang.GetEdgeName[CultureGroup_Culture]()))
}

func GetAllCultureGroup_CulturesEids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*CultureGroup_Culture](space, "")
}

func GetAllCultureGroup_Cultures(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*CultureGroup_Culture] {
	return nebulagolang.GetAllEdgesByEdgeType[*CultureGroup_Culture](space)
}
