package ck2nebula

import (
	"fmt"
	"github.com/thalesfu/nebulagolang"
	"github.com/thalesfu/paradoxtools/CK2/save"
	"github.com/thalesfu/paradoxtools/utils"
)

type Dynasty struct {
	VID                string `nebulakey:"vid" nebulatagname:"dynasty" nebulatagcomment:"dynasty" json:"vid,omitempty"`
	ID                 int    `nebulaproperty:"id" description:"id" nebulaindexes:"id" json:"id,omitempty"`
	Name               string `nebulaproperty:"name" description:"name" nebulaindexes:"name" json:"name,omitempty"`
	Culture            string `nebulaproperty:"culture" description:"culture" nebulaindexes:"culture" json:"culture,omitempty"`
	Religion           string `nebulaproperty:"religion" description:"religion" nebulaindexes:"religion" json:"religion,omitempty"`
	CoatOfArmsData     string `nebulaproperty:"coat_of_arms_data" description:"coat_of_arms_data" nebulaindexes:"coat_of_arms_data" json:"coat_of_arms_data,omitempty"`
	CoatOfArmsReligion string `nebulaproperty:"coat_of_arms_religion" description:"coat_of_arms_religion" nebulaindexes:"coat_of_arms_religion" json:"coat_of_arms_religion,omitempty"`
	PlayID             int    `nebulaproperty:"play_id" description:"game play id" nebulaindexes:"play_id" json:"play_id,omitempty"`
}

func NewDynasty(playId int, dynastyId int) *Dynasty {
	nebulaDynasty := Dynasty{
		VID:    getDynastyVid(playId, dynastyId),
		ID:     dynastyId,
		PlayID: playId,
	}
	return &nebulaDynasty
}

func NewDynastyByData(dynasty *save.Dynasty) *Dynasty {
	nebulaDynasty := utils.Mapping[Dynasty](dynasty)
	nebulaDynasty.VID = getDynastyVid(nebulaDynasty.PlayID, nebulaDynasty.ID)

	if dynasty.CoatOfArms != nil {
		if len(dynasty.CoatOfArms.Data) > 0 {
			for _, v := range dynasty.CoatOfArms.Data {
				nebulaDynasty.CoatOfArmsData += fmt.Sprintf("%d ", v)
			}

			nebulaDynasty.CoatOfArmsData = nebulaDynasty.CoatOfArmsData[:len(nebulaDynasty.CoatOfArmsData)-1]
		}

		nebulaDynasty.CoatOfArmsReligion = dynasty.CoatOfArms.Religion
	}

	return &nebulaDynasty
}

func getDynastyVid(playId int, dynastyId int) string {
	return fmt.Sprintf("dynasty.%d.%d", dynastyId, playId)
}

func (p *Dynasty) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertVertexes(space, p)
}

func (p *Dynasty) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateVertexes(space, p)
}

func (p *Dynasty) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertVertexes(space, p)
}

func (p *Dynasty) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteVertexes(space, p)
}

func (p *Dynasty) DeleteFromNebulaWithEdge(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteVertexesWithEdges(space, p)
}

func (p *Dynasty) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadVertex(space, p)
}

func InsertDynastys(space *nebulagolang.Space, ps ...*Dynasty) *nebulagolang.Result {
	return nebulagolang.BatchInsertVertexes(space, 250, ps)
}

func UpdateDynastys(space *nebulagolang.Space, ps ...*Dynasty) *nebulagolang.Result {
	return nebulagolang.BatchUpdateVertexes(space, 250, ps)
}

func UpsertDynastys(space *nebulagolang.Space, ps ...*Dynasty) *nebulagolang.Result {
	return nebulagolang.BatchUpsertVertexes(space, 250, ps)
}

func DeleteDynastys(space *nebulagolang.Space, playId int, ids ...int) *nebulagolang.Result {
	ts := make([]string, len(ids))
	for i, c := range ids {
		ts[i] = getDynastyVid(playId, c)
	}
	return nebulagolang.DeleteVertexesByVids(space, ts...)
}

func DeleteDynastyWithEdge(space *nebulagolang.Space, playId int, ids ...int) *nebulagolang.Result {
	ts := make([]string, len(ids))
	for i, c := range ids {
		ts[i] = getDynastyVid(playId, c)
	}
	return nebulagolang.DeleteVertexesWithEdgesByVids(space, ts...)
}

func DeleteAllDynasty(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllVertexesByTag[Dynasty](space)
}

func DeleteAllDynastyByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.Result {
	return nebulagolang.DeleteAllVertexesByQuery[Dynasty](space, getPlayIdQuery[Dynasty](playId))
}

func DeleteAllDynastyWithEdges(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllVertexesWithEdgesByTag[Dynasty](space)
}

func DeleteAllDynastyWithEdgesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.Result {
	return nebulagolang.DeleteAllVertexesWithEdgesByQuery[Dynasty](space, getPlayIdQuery[Dynasty](playId))
}

func GetDynastyByVid(space *nebulagolang.Space, vid string) *nebulagolang.ResultT[*Dynasty] {
	return nebulagolang.GetVertexByVid[*Dynasty](space, vid)
}

func GetDynastyByID(space *nebulagolang.Space, playId int, id int) *nebulagolang.ResultT[*Dynasty] {
	return nebulagolang.GetVertexByVid[*Dynasty](space, getDynastyVid(playId, id))
}

func GetAllDynastys(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*Dynasty] {
	return nebulagolang.GetAllVertexesByQuery[*Dynasty](space, "")
}

func GetAllDynastysByPlayIdAndName(space *nebulagolang.Space, playId int, name string) *nebulagolang.ResultT[map[string]*Dynasty] {
	return nebulagolang.GetAllVertexesByQuery[*Dynasty](space, getPlayIdAndPropertyQuery[Dynasty](playId, "name", name))
}

func GetAllDynastysByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]*Dynasty] {
	return nebulagolang.GetAllVertexesByQuery[*Dynasty](space, getPlayIdQuery[Dynasty](playId))
}

func GetAllDynastysVids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesVIDsByQuery[*Dynasty](space, "")
}

func GetAllDynastysVidsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesVIDsByQuery[*Dynasty](space, getPlayIdQuery[Dynasty](playId))
}

func GetAllDynastysIds(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesPropertyByQuery[*Dynasty](space, "", "id", "")
}

func GetAllDynastysIdsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesPropertyByQuery[*Dynasty](space, getPlayIdQuery[Dynasty](playId), "id", "")
}

func GetAllDynastysNames(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesPropertyByQuery[*Dynasty](space, "", "name", "")
}

func GetAllDynastyNamesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesPropertyByQuery[*Dynasty](space, getPlayIdQuery[Dynasty](playId), "name", "")
}

func GetAllDynastysCodes(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesPropertyByQuery[*Dynasty](space, "", "code", "")
}

func GetAllDynastyCodesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesPropertyByQuery[*Dynasty](space, getPlayIdQuery[Dynasty](playId), "code", "")
}
