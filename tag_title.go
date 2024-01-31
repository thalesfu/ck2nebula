package ck2nebula

import (
	"fmt"
	"github.com/thalesfu/nebulagolang"
	"github.com/thalesfu/paradoxtools/CK2/save"
	"github.com/thalesfu/paradoxtools/utils"
)

type Title struct {
	VID                string `nebulakey:"vid" nebulatagname:"title" nebulatagcomment:"title" json:"vid,omitempty"`
	Name               string `nebulaproperty:"name" description:"title name" nebulaindexes:"name" json:"name,omitempty"`
	ID                 string `nebulaproperty:"id" description:"title code" nebulaindexes:"code" json:"code,omitempty"`
	Foa                string `nebulaproperty:"foa" description:"title foa" nebulaindexes:"foa" son:"foa,omitempty"`
	Title              string `nebulaproperty:"title" description:"title" nebulaindexes:"title" json:"title,omitempty"`
	TitleFemale        string `nebulaproperty:"title_female" description:"female title" nebulaindexes:"title_female" json:"title_female,omitempty"`
	Gender             string `nebulaproperty:"gender" description:"title gender" nebulaindexes:"gender" json:"gender,omitempty"`
	Active             bool   `nebulaproperty:"active" description:"title is active" nebulaindexes:"active" json:"active,omitempty"`
	IsCustom           bool   `nebulaproperty:"is_custom" description:"is custom title" nebulaindexes:"is_custom" json:"is_custom,omitempty"`
	IsDynamic          bool   `nebulaproperty:"is_dynamic" mappingalias:"Dynamic" description:"is dynamic title" nebulaindexes:"is_dynamic" json:"is_dynamic,omitempty"`
	Nomad              bool   `nebulaproperty:"nomad" description:"title is nomad" nebulaindexes:"nomad" json:"nomad,omitempty"`
	Landless           bool   `nebulaproperty:"landless" description:"title is land less" nebulaindexes:"landless" json:"landless,omitempty"`
	MajorRevolt        bool   `nebulaproperty:"major_revolt" description:"title is major revolt" nebulaindexes:"major_revolt" json:"major_revolt,omitempty"`
	Mercenary          bool   `nebulaproperty:"mercenary" description:"title is mercenary" nebulaindexes:"mercenary" json:"mercenary,omitempty"`
	Rebels             bool   `nebulaproperty:"rebels" description:"title is rebels" nebulaindexes:"rebels" json:"rebels,omitempty"`
	Temporary          bool   `nebulaproperty:"temporary" description:"title is temporary" nebulaindexes:"temporary" json:"temporary,omitempty"`
	CoatOfArmsData     string `nebulaproperty:"coat_of_arms_data" description:"coat_of_arms_data" nebulaindexes:"coat_of_arms_data" json:"coat_of_arms_data,omitempty"`
	CoatOfArmsReligion string `nebulaproperty:"coat_of_arms_religion" description:"coat_of_arms_religion" nebulaindexes:"coat_of_arms_religion" json:"coat_of_arms_religion,omitempty"`
	StoryID            int    `nebulaproperty:"story_id" mappingalias:"PlayID" description:"game play id" nebulaindexes:"story_id" json:"story_id,omitempty"`
}

func NewTitle(storyId int, titleId string) *Title {
	nebulaTitle := Title{
		VID:     getTitleVid(storyId, titleId),
		ID:      titleId,
		StoryID: storyId,
	}
	return &nebulaTitle
}

func NewTitleByData(title *save.Title) *Title {
	nebulaTitle := utils.Mapping[Title](title)
	nebulaTitle.VID = getTitleVid(nebulaTitle.StoryID, nebulaTitle.ID)

	if title.CoatOfArms != nil {
		if len(title.CoatOfArms.Data) > 0 {
			for _, v := range title.CoatOfArms.Data {
				nebulaTitle.CoatOfArmsData += fmt.Sprintf("%d ", v)
			}

			nebulaTitle.CoatOfArmsData = nebulaTitle.CoatOfArmsData[:len(nebulaTitle.CoatOfArmsData)-1]
		}

		nebulaTitle.CoatOfArmsReligion = title.CoatOfArms.Religion
	}

	return &nebulaTitle
}

func getTitleVid(playId int, titleId string) string {
	return fmt.Sprintf("title.%s.%d", titleId, playId)
}

func (t *Title) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertVertexes(space, t)
}

func (t *Title) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateVertexes(space, t)
}

func (t *Title) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertVertexes(space, t)
}

func (t *Title) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteVertexes(space, t)
}

func (t *Title) DeleteFromNebulaWithEdge(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteVertexesWithEdges(space, t)
}

func (t *Title) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadVertex(space, t)
}

func InsertTitles(space *nebulagolang.Space, ts ...*Title) *nebulagolang.Result {
	return nebulagolang.BatchInsertVertexes(space, 250, ts)
}

func UpdateTitles(space *nebulagolang.Space, ts ...*Title) *nebulagolang.Result {
	return nebulagolang.BatchUpdateVertexes(space, 250, ts)
}

func UpsertTitles(space *nebulagolang.Space, ts ...*Title) *nebulagolang.Result {
	return nebulagolang.BatchUpsertVertexes(space, 250, ts)
}

func DeleteTitles(space *nebulagolang.Space, playId int, ids ...string) *nebulagolang.Result {
	ts := make([]string, len(ids))
	for i, c := range ids {
		ts[i] = getTitleVid(playId, c)
	}
	return nebulagolang.DeleteVertexesByVids(space, ts...)
}

func DeleteTitleWithEdge(space *nebulagolang.Space, playId int, ids ...string) *nebulagolang.Result {
	ts := make([]string, len(ids))
	for i, c := range ids {
		ts[i] = getTitleVid(playId, c)
	}
	return nebulagolang.DeleteVertexesWithEdgesByVids(space, ts...)
}

func DeleteAllTitle(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllVertexesByTag[Title](space)
}

func DeleteAllTitleByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.Result {
	return nebulagolang.DeleteAllVertexesByQuery[Title](space, getPlayIdQuery[Title](playId))
}

func DeleteAllTitleWithEdges(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllVertexesWithEdgesByTag[Title](space)
}

func DeleteAllTitleWithEdgesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.Result {
	return nebulagolang.DeleteAllVertexesWithEdgesByQuery[Title](space, getPlayIdQuery[Title](playId))
}

func GetTitleByVid(space *nebulagolang.Space, vid string) *nebulagolang.ResultT[*Title] {
	return nebulagolang.GetVertexByVid[*Title](space, vid)
}

func GetTitleByID(space *nebulagolang.Space, playId int, id string) *nebulagolang.ResultT[*Title] {
	return nebulagolang.GetVertexByVid[*Title](space, getTitleVid(playId, id))
}

func GetAllTitle(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*Title] {
	return nebulagolang.GetAllVertexesByQuery[*Title](space, "")
}

func GetAllTitleByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]*Title] {
	return nebulagolang.GetAllVertexesByQuery[*Title](space, getPlayIdQuery[Title](playId))
}

func GetAllTitlesByPlayIdAndName(space *nebulagolang.Space, playId int, name string) *nebulagolang.ResultT[map[string]*Title] {
	return nebulagolang.GetAllVertexesByQuery[*Title](space, getPlayIdAndPropertyQuery[Title](playId, "name", name))
}

func GetAllTitleVids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesVIDsByQuery[*Title](space, "")
}

func GetAllTitleVidsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesVIDsByQuery[*Title](space, getPlayIdQuery[Title](playId))
}

func GetAllTitleIds(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesPropertyByQuery[*Title](space, "", "id", "")
}

func GetAllTitleIdsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesPropertyByQuery[*Title](space, getPlayIdQuery[Title](playId), "id", "")
}

func GetAllTitleNames(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesPropertyByQuery[*Title](space, "", "name", "")
}

func GetAllTitleNamesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesPropertyByQuery[*Title](space, getPlayIdQuery[Title](playId), "name", "")
}
