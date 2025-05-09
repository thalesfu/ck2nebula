package ck2nebula

import (
	"github.com/thalesfu/golangutils"
	"github.com/thalesfu/nebulagolang"
	"github.com/thalesfu/paradoxtools/CK2/culture"
)

type CultureGroup struct {
	VID  string `nebulakey:"vid" nebulatagname:"culturegroup" nebulatagcomment:"culture group" json:"vid,omitempty"`
	Name string `nebulaproperty:"name" description:"culture group name" nebulaindexes:"name" json:"name,omitempty"`
	Code string `nebulaproperty:"code" description:"culture group code" nebulaindexes:"code" json:"code,omitempty"`
}

func NewCultureGroup(code string) *CultureGroup {
	nebulaCultureGroup := CultureGroup{
		VID:  getCultureGroupVid(code),
		Code: code,
	}
	return &nebulaCultureGroup
}

func NewCultureGroupByData(group *culture.CultureGroup) *CultureGroup {
	nebulaCultureGroup := golangutils.Mapping[CultureGroup](group)
	nebulaCultureGroup.VID = getCultureGroupVid(nebulaCultureGroup.Code)
	return &nebulaCultureGroup
}

func getCultureGroupVid(code string) string {
	return "culturegroup." + code
}

func (cg *CultureGroup) NewCultureByCode(code string) (*Culture, *CultureGroup_Culture) {
	return cg.NewCulture(NewCulture(code))
}

func (cg *CultureGroup) NewCulture(nebulaCulture *Culture) (*Culture, *CultureGroup_Culture) {
	nebulaCultureGroup_Culture := NewCultureGroup_Culture(cg, nebulaCulture)
	return nebulaCulture, nebulaCultureGroup_Culture
}

func (cg *CultureGroup) NewCultureByData(culture *culture.Culture) (*Culture, *CultureGroup_Culture) {
	return cg.NewCulture(NewCultureByData(culture))
}

func (cg *CultureGroup) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertVertexes(space, cg)
}

func (cg *CultureGroup) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateVertexes(space, cg)
}

func (cg *CultureGroup) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertVertexes(space, cg)
}

func (cg *CultureGroup) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteVertexes(space, cg)
}

func (cg *CultureGroup) DeleteFromNebulaWithEdge(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteVertexesWithEdges(space, cg)
}

func (cg *CultureGroup) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadVertex(space, cg)
}

func DeleteCultureGroups(space *nebulagolang.Space, codes ...string) *nebulagolang.Result {
	cs := make([]string, len(codes))
	for i, c := range codes {
		cs[i] = getCultureGroupVid(c)
	}
	return nebulagolang.DeleteVertexesByVids(space, cs...)
}

func DeleteCultureGroupsWithEdge(space *nebulagolang.Space, codes ...string) *nebulagolang.Result {
	cs := make([]string, len(codes))
	for i, c := range codes {
		cs[i] = getCultureGroupVid(c)
	}
	return nebulagolang.DeleteVertexesWithEdgesByVids(space, cs...)
}

func DeleteAllCultureGroup(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllVertexesByTag[CultureGroup](space)
}

func DeleteAllCultureGroupWithEdges(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllVertexesWithEdgesByTag[CultureGroup](space)
}

func InsertCultureGroups(space *nebulagolang.Space, cgs ...*CultureGroup) *nebulagolang.Result {
	return nebulagolang.BatchInsertVertexes(space, 250, cgs)
}

func UpdateCultureGroups(space *nebulagolang.Space, cgs ...*CultureGroup) *nebulagolang.Result {
	return nebulagolang.BatchUpdateVertexes(space, 250, cgs)
}

func UpsertCultureGroups(space *nebulagolang.Space, cgs ...*CultureGroup) *nebulagolang.Result {
	return nebulagolang.BatchUpsertVertexes(space, 250, cgs)
}

func GetCultureGroupByVid(space *nebulagolang.Space, vid string) *nebulagolang.ResultT[*CultureGroup] {
	return nebulagolang.GetVertexByVid[*CultureGroup](space, vid)
}

func GetCultureGroupByCode(space *nebulagolang.Space, code string) *nebulagolang.ResultT[*CultureGroup] {
	return nebulagolang.GetVertexByVid[*CultureGroup](space, getCultureGroupVid(code))
}

func GetAllCultureGroup(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*CultureGroup] {
	return nebulagolang.GetAllVertexesByQuery[*CultureGroup](space, "")
}

func GetAllCultureGroupVids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesVIDsByQuery[*CultureGroup](space, "")
}

func GetAllCultureGroupCode(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesPropertyByQuery[*CultureGroup](space, "", "code", "")
}
