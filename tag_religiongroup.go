package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
	"github.com/thalesfu/paradoxtools/CK2/religion"
	"github.com/thalesfu/paradoxtools/utils"
)

type ReligionGroup struct {
	VID  string `nebulakey:"vid" nebulatagname:"religiongroup" nebulatagcomment:"religion group" json:"vid,omitempty"`
	Name string `nebulaproperty:"name" description:"religion group name" nebulaindexes:"name" json:"name,omitempty"`
	Code string `nebulaproperty:"code" description:"religion group code" nebulaindexes:"code" json:"code,omitempty"`
}

func NewReligionGroup(code string) *ReligionGroup {
	nebulaReligionGroup := ReligionGroup{
		VID:  getReligionGroupVid(code),
		Code: code,
	}
	return &nebulaReligionGroup
}

func NewReligionGroupByData(group *religion.ReligionGroup) *ReligionGroup {
	nebulaReligionGroup := utils.Mapping[ReligionGroup](group)
	nebulaReligionGroup.VID = getReligionGroupVid(nebulaReligionGroup.Code)
	return &nebulaReligionGroup
}

func getReligionGroupVid(code string) string {
	return "religiongroup." + code
}

func (rg *ReligionGroup) NewReligionByCode(code string) (*Religion, *ReligionGroup_Religion) {
	return rg.NewReligion(NewReligion(code))
}

func (rg *ReligionGroup) NewReligion(nebulaReligion *Religion) (*Religion, *ReligionGroup_Religion) {
	nebulaReligionGroup_Religion := NewReligionGroup_Religion(rg, nebulaReligion)
	return nebulaReligion, nebulaReligionGroup_Religion
}

func (rg *ReligionGroup) NewReligionByData(religion *religion.Religion) (*Religion, *ReligionGroup_Religion) {
	return rg.NewReligion(NewReligionByData(religion))
}

func (rg *ReligionGroup) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertVertexes(space, rg)
}

func (rg *ReligionGroup) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateVertexes(space, rg)
}

func (rg *ReligionGroup) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertVertexes(space, rg)
}

func (rg *ReligionGroup) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteVertexes(space, rg)
}

func (rg *ReligionGroup) DeleteFromNebulaWithEdge(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteVertexesWithEdges(space, rg)
}

func (rg *ReligionGroup) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadVertex(space, rg)
}

func DeleteReligionGroups(space *nebulagolang.Space, codes ...string) *nebulagolang.Result {
	cs := make([]string, len(codes))
	for i, c := range codes {
		cs[i] = getReligionGroupVid(c)
	}
	return nebulagolang.DeleteVertexesByVids(space, cs...)
}

func DeleteReligionGroupsWithEdge(space *nebulagolang.Space, codes ...string) *nebulagolang.Result {
	cs := make([]string, len(codes))
	for i, c := range codes {
		cs[i] = getReligionGroupVid(c)
	}
	return nebulagolang.DeleteVertexesWithEdgesByVids(space, cs...)
}

func DeleteAllReligionGroup(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllVertexesByTag[ReligionGroup](space)
}

func DeleteAllReligionGroupWithEdges(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllVertexesWithEdgesByTag[ReligionGroup](space)
}

func InsertReligionGroups(space *nebulagolang.Space, cgs ...*ReligionGroup) *nebulagolang.Result {
	return nebulagolang.BatchInsertVertexes(space, 10000, cgs)
}

func UpdateReligionGroups(space *nebulagolang.Space, cgs ...*ReligionGroup) *nebulagolang.Result {
	return nebulagolang.BatchUpdateVertexes(space, 5000, cgs)
}

func UpsertReligionGroups(space *nebulagolang.Space, cgs ...*ReligionGroup) *nebulagolang.Result {
	return nebulagolang.BatchUpsertVertexes(space, 5000, cgs)
}

func GetReligionGroupByVid(space *nebulagolang.Space, vid string) *nebulagolang.ResultT[*ReligionGroup] {
	return nebulagolang.GetVertexByVid[*ReligionGroup](space, vid)
}

func GetReligionGroupByCode(space *nebulagolang.Space, code string) *nebulagolang.ResultT[*ReligionGroup] {
	return nebulagolang.GetVertexByVid[*ReligionGroup](space, getReligionGroupVid(code))
}

func GetAllReligionGroup(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*ReligionGroup] {
	return nebulagolang.GetAllVertexesByQuery[*ReligionGroup](space, "")
}

func GetAllReligionGroupVids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesVIDsByQuery[*ReligionGroup](space, "")
}

func GetAllReligionGroupCode(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesPropertyByQuery[*ReligionGroup](space, "", "code", "")
}
