package ck2nebula

import (
	"fmt"
	"github.com/thalesfu/golangutils"
	"github.com/thalesfu/nebulagolang"
	"github.com/thalesfu/paradoxtools/CK2/save"
)

type Dynasty struct {
	VID                string `nebulakey:"vid" nebulatagname:"dynasty" nebulatagcomment:"dynasty" json:"vid,omitempty"`
	ID                 int    `nebulaproperty:"id" description:"id" nebulaindexes:"id" json:"id,omitempty"`
	Name               string `nebulaproperty:"name" description:"name" nebulaindexes:"name" json:"name,omitempty"`
	Culture            string `nebulaproperty:"culture" description:"culture" nebulaindexes:"culture" json:"culture,omitempty"`
	Religion           string `nebulaproperty:"religion" description:"religion" nebulaindexes:"religion" json:"religion,omitempty"`
	CoatOfArmsData     string `nebulaproperty:"coat_of_arms_data" description:"coat_of_arms_data" nebulaindexes:"coat_of_arms_data" json:"coat_of_arms_data,omitempty"`
	CoatOfArmsReligion string `nebulaproperty:"coat_of_arms_religion" description:"coat_of_arms_religion" nebulaindexes:"coat_of_arms_religion" json:"coat_of_arms_religion,omitempty"`
	StoryID            int    `nebulaproperty:"story_id" mappingalias:"PlayID" description:"game play id" nebulaindexes:"story_id" json:"story_id,omitempty"`
}

func NewDynasty(storyId int, dynastyId int) *Dynasty {
	nebulaDynasty := Dynasty{
		VID:     getDynastyVid(storyId, dynastyId),
		ID:      dynastyId,
		StoryID: storyId,
	}
	return &nebulaDynasty
}

func NewDynastyByData(dynasty *save.Dynasty) *Dynasty {
	nebulaDynasty := golangutils.Mapping[Dynasty](dynasty)
	nebulaDynasty.VID = getDynastyVid(nebulaDynasty.StoryID, nebulaDynasty.ID)

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

func (d *Dynasty) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertVertexes(space, d)
}

func (d *Dynasty) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateVertexes(space, d)
}

func (d *Dynasty) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertVertexes(space, d)
}

func (d *Dynasty) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteVertexes(space, d)
}

func (d *Dynasty) DeleteFromNebulaWithEdge(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteVertexesWithEdges(space, d)
}

func (d *Dynasty) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadVertex(space, d)
}

func (d *Dynasty) GetAlivePeoples(space *nebulagolang.Space) *nebulagolang.ResultT[map[int]*People] {
	command := fmt.Sprintf("go from \"%s\" over people_dynasty reversely where properties($$).isdead==false yield $$ as v", d.VID)

	return d.getPeopleByQuery(space, command)
}

func (d *Dynasty) GetPeoples(space *nebulagolang.Space) *nebulagolang.ResultT[map[int]*People] {
	command := fmt.Sprintf("go from \"%s\" over people_dynasty reversely yield $$ as v", d.VID)

	return d.getPeopleByQuery(space, command)
}

func (d *Dynasty) getPeopleByQuery(space *nebulagolang.Space, command string) *nebulagolang.ResultT[map[int]*People] {
	r := nebulagolang.QueryVertexesByQueryToSlice[*People](space, command)

	if !r.Ok {
		return nebulagolang.NewResultT[map[int]*People](r.Result)
	}

	result := make(map[int]*People)

	for _, f := range r.Data {
		result[f.ID] = f
	}

	return nebulagolang.NewResultTWithData(r.Result, result)
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
