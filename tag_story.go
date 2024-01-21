package ck2nebula

import (
	"fmt"
	"github.com/thalesfu/nebulagolang"
	"github.com/thalesfu/paradoxtools/CK2/save"
	"time"
)

type Story struct {
	VID            string    `nebulakey:"vid" nebulatagname:"story" nebulatagcomment:"story" json:"vid,omitempty"`
	Version        string    `nebulaproperty:"version" description:"version" nebulaindexes:"version" json:"version,omitempty"`
	PlayID         int       `nebulaproperty:"play_id" description:"game play id" nebulaindexes:"play_id" json:"play_id,omitempty"`
	PlayDate       time.Time `nebulaproperty:"play_date" nebulatype:"Date" description:"game play date" nebulaindexes:"play_date" json:"play_date,omitempty"`
	PlayerID       int       `nebulaproperty:"player_id" description:"player id" nebulaindexes:"player id" json:"player_id,omitempty"`
	PlayerName     string    `nebulaproperty:"player_name" description:"player name" nebulaindexes:"player name" json:"player_name,omitempty"`
	PlayerAge      int       `nebulaproperty:"player_age" description:"player age" nebulaindexes:"player age" json:"player_age,omitempty"`
	Religion       string    `nebulaproperty:"religion" description:"religion" nebulaindexes:"religion" json:"religion,omitempty"`
	ReligionName   string    `nebulaproperty:"religion_name" description:"religion name" nebulaindexes:"religion_name" json:"religion_name,omitempty"`
	Culture        string    `nebulaproperty:"culture" description:"culture" nebulaindexes:"culture" json:"culture,omitempty"`
	CultureName    string    `nebulaproperty:"culture_name" description:"culture name" nebulaindexes:"culture_name" json:"culture_name,omitempty"`
	Government     string    `nebulaproperty:"government" description:"government" nebulaindexes:"government" json:"government,omitempty"`
	GovernmentName string    `nebulaproperty:"government_name" description:"government name" nebulaindexes:"government_name" json:"government_name,omitempty"`
	Dynasty        string    `nebulaproperty:"dynasty" description:"dynasty" nebulaindexes:"dynasty" json:"dynasty,omitempty"`
}

func NewStory(id int) *Story {
	nebulaStory := Story{
		VID:    getStoryVid(id),
		PlayID: id,
	}
	return &nebulaStory
}

func NewStoryByData(story *save.SaveFile) *Story {
	nebulaStory := NewStory(story.PlayThroughID)
	nebulaStory.PlayDate = time.Time(story.Date)
	nebulaStory.PlayerID = story.Player.ID
	nebulaStory.PlayerName = story.PlayerName
	nebulaStory.PlayerAge = story.PlayerAge
	nebulaStory.Version = story.Version
	nebulaStory.Religion = story.PlayerPortrait.Religion
	nebulaStory.Culture = story.PlayerPortrait.Culture
	nebulaStory.Government = story.PlayerPortrait.Government
	if d, ok := story.Dynasties[story.PlayerPortrait.Dynasty]; ok {
		nebulaStory.Dynasty = d.Name
	}

	return nebulaStory
}

func getStoryVid(id int) string {
	return fmt.Sprintf("story.%d", id)
}

func (s *Story) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertVertexes(space, s)
}

func (s *Story) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateVertexes(space, s)
}

func (s *Story) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertVertexes(space, s)
}

func (s *Story) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteVertexes(space, s)
}

func (s *Story) DeleteFromNebulaWithEdge(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteVertexesWithEdges(space, s)
}

func (s *Story) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadVertex(space, s)
}

func InsertStories(space *nebulagolang.Space, ses ...*Story) *nebulagolang.Result {
	return nebulagolang.BatchInsertVertexes(space, 250, ses)
}

func UpdateStories(space *nebulagolang.Space, ses ...*Story) *nebulagolang.Result {
	return nebulagolang.BatchUpdateVertexes(space, 250, ses)
}

func UpsertStories(space *nebulagolang.Space, ses ...*Story) *nebulagolang.Result {
	return nebulagolang.BatchUpsertVertexes(space, 250, ses)
}

func DeleteStories(space *nebulagolang.Space, ses ...int) *nebulagolang.Result {
	ss := make([]string, len(ses))
	for i, c := range ses {
		ss[i] = getStoryVid(c)
	}
	return nebulagolang.DeleteVertexesByVids(space, ss...)
}

func DeleteStoriesWithEdge(space *nebulagolang.Space, ids ...int) *nebulagolang.Result {
	ss := make([]string, len(ids))
	for i, c := range ids {
		ss[i] = getStoryVid(c)
	}
	return nebulagolang.DeleteVertexesWithEdgesByVids(space, ss...)
}

func DeleteAllStories(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllVertexesByTag[Story](space)
}

func DeleteAllStoriesWithEdges(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllVertexesWithEdgesByTag[Story](space)
}

func GetStoryByVid(space *nebulagolang.Space, vid string) *nebulagolang.ResultT[*Story] {
	return nebulagolang.GetVertexByVid[*Story](space, vid)
}

func GetStoryByID(space *nebulagolang.Space, id int) *nebulagolang.ResultT[*Story] {
	return nebulagolang.GetVertexByVid[*Story](space, getStoryVid(id))
}

func GetAllStories(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*Story] {
	return nebulagolang.GetAllVertexesByQuery[*Story](space, "")
}

func GetAllStoriesVids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesVIDsByQuery[*Story](space, "")
}
