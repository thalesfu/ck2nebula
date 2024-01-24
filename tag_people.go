package ck2nebula

import (
	"fmt"
	"github.com/thalesfu/nebulagolang"
	"github.com/thalesfu/paradoxtools/CK2/save"
	"github.com/thalesfu/paradoxtools/utils"
	"time"
)

type People struct {
	VID                 string    `nebulakey:"vid" nebulatagname:"people" nebulatagcomment:"people" json:"vid,omitempty"`
	ID                  int       `nebulaproperty:"id" description:"id" nebulaindexes:"id" json:"id,omitempty"`
	Name                string    `nebulaproperty:"name" description:"name" nebulaindexes:"name" json:"name,omitempty"`
	Birthday            time.Time `nebulaproperty:"birthday" nebulatype:"Date" description:"birthday" nebulaindexes:"birthday" json:"birthday,omitempty"`
	DeathDay            time.Time `nebulaproperty:"deathday" nebulatype:"Date" description:"deathday" nebulaindexes:"deathday" json:"deathday,omitempty"`
	Age                 int       `nebulaproperty:"age" description:"age" nebulaindexes:"age" json:"age,omitempty"`
	IsDead              bool      `nebulaproperty:"isdead" description:"isdead" nebulaindexes:"isdead" json:"isdead,omitempty"`
	IsBastard           bool      `nebulaproperty:"isbastard" description:"isbastard" nebulaindexes:"isbastard" json:"isbastard,omitempty"`
	Diplomacy           int       `nebulaproperty:"diplomacy" description:"diplomacy" nebulaindexes:"diplomacy" json:"diplomacy,omitempty"`
	Martial             int       `nebulaproperty:"martial" description:"martial" nebulaindexes:"martial" json:"martial,omitempty"`
	Stewardship         int       `nebulaproperty:"stewardship" description:"stewardship" nebulaindexes:"stewardship" json:"stewardship,omitempty"`
	Intrigue            int       `nebulaproperty:"intrigue" description:"intrigue" nebulaindexes:"intrigue" json:"intrigue,omitempty"`
	Learning            int       `nebulaproperty:"learning" description:"learning" nebulaindexes:"learning" json:"learning,omitempty"`
	ModifiedDiplomacy   int       `nebulaproperty:"modified_diplomacy" description:"modified diplomacy" nebulaindexes:"modified_diplomacy" json:"modified_diplomacy,omitempty"`
	ModifiedMartial     int       `nebulaproperty:"modified_martial" description:"modified martial" nebulaindexes:"modified_martial" json:"modified_martial,omitempty"`
	ModifiedStewardship int       `nebulaproperty:"modified_stewardship" description:"modified stewardship" nebulaindexes:"modified_stewardship" json:"modified_stewardship,omitempty"`
	ModifiedIntrigue    int       `nebulaproperty:"modified_intrigue" description:"modified intrigue" nebulaindexes:"modified_intrigue" json:"modified_intrigue,omitempty"`
	ModifiedLearning    int       `nebulaproperty:"modified_learning" description:"modified learning" nebulaindexes:"modified_learning" json:"modified_learning,omitempty"`
	ModifiedSexAppeal   int       `nebulaproperty:"modified_sex_appeal" description:"modified sex appeal" nebulaindexes:"modified_sex_appeal" json:"modified_sex_appeal,omitempty"`
	IsHistory           bool      `nebulaproperty:"ishistory" description:"ishistory" nebulaindexes:"ishistory" json:"ishistory,omitempty"`
	IsFemale            bool      `nebulaproperty:"isfemale" description:"isfemale" nebulaindexes:"isfemale" json:"isfemale,omitempty"`
	InHiding            bool      `nebulaproperty:"in_hiding" description:"in_hiding" nebulaindexes:"in_hiding" json:"in_hiding,omitempty"`
	IsPlayer            bool      `nebulaproperty:"isplayer" description:"isplayer" nebulaindexes:"isplayer" json:"isplayer,omitempty"`
	NickNameCode        string    `nebulaproperty:"nicknamecode" description:"nicknamecode" nebulaindexes:"nicknamecode" json:"nicknamecode,omitempty"`
	NickName            string    `nebulaproperty:"nickname" description:"nickname" nebulaindexes:"nickname" json:"nickname,omitempty"`
	DNA                 string    `nebulaproperty:"dna" description:"dna" nebulaindexes:"dna" json:"dna,omitempty"`
	Government          string    `nebulaproperty:"gov" description:"gov" nebulaindexes:"gov" json:"gov,omitempty"`
	Culture             string    `nebulaproperty:"culture" description:"culture" nebulaindexes:"culture" json:"culture,omitempty"`
	CultureName         string    `nebulaproperty:"culture_name" description:"culture name" nebulaindexes:"culture_name" json:"culture_name,omitempty"`
	GFXCulture          string    `nebulaproperty:"gfx_culture" description:"gfx_culture" nebulaindexes:"gfx_culture" json:"gfx_culture,omitempty"`
	GFXCultureName      string    `nebulaproperty:"gfx_culture_name" description:"gfx_culture_name" nebulaindexes:"gfx_culture_name" json:"gfx_culture_name,omitempty"`
	Religion            string    `nebulaproperty:"religion" description:"religion" nebulaindexes:"religion" json:"religion,omitempty"`
	ReligionName        string    `nebulaproperty:"religion_name" description:"religion name" nebulaindexes:"religion_name" json:"religion_name,omitempty"`
	SecretReligion      string    `nebulaproperty:"secret_religion" description:"secret_religion" nebulaindexes:"secret_religion" json:"secret_religion,omitempty"`
	SecretReligionName  string    `nebulaproperty:"secret_religion_name" description:"secret_religion_name" nebulaindexes:"secret_religion_name" json:"secret_religion_name,omitempty"`
	Fertility           float32   `nebulaproperty:"fer" description:"fer" nebulaindexes:"fer" json:"fer,omitempty"`
	Health              float32   `nebulaproperty:"health" description:"health" nebulaindexes:"health" json:"health,omitempty"`
	Prestige            float32   `nebulaproperty:"prestige" description:"prestige" nebulaindexes:"prestige" json:"prestige,omitempty"`
	Piety               float32   `nebulaproperty:"piety" description:"piety" nebulaindexes:"piety" json:"piety,omitempty"`
	Wealth              float32   `nebulaproperty:"wealth" description:"wealth" nebulaindexes:"wealth" json:"wealth,omitempty"`
	Score               float32   `nebulaproperty:"score" description:"score" nebulaindexes:"score" json:"score,omitempty"`
	Society             int       `nebulaproperty:"society" description:"society" nebulaindexes:"society" json:"society,omitempty"`
	PlayID              int       `nebulaproperty:"play_id" description:"game play id" nebulaindexes:"play_id" json:"play_id,omitempty"`
}

func NewPeople(playId int, peopleId int) *People {
	nebulaPeople := People{
		VID:    getPeopleVid(playId, peopleId),
		ID:     peopleId,
		PlayID: playId,
	}
	return &nebulaPeople
}

func NewPeopleByData(character *save.Character) *People {
	nebulaPeople := utils.Mapping[People](character)
	nebulaPeople.VID = getPeopleVid(nebulaPeople.PlayID, nebulaPeople.ID)
	if len(character.Attributes) > 0 {
		for i, attr := range character.Attributes {
			switch i {
			case 0:

				nebulaPeople.Diplomacy = attr
				nebulaPeople.ModifiedDiplomacy = attr
			case 1:
				nebulaPeople.Martial = attr
				nebulaPeople.ModifiedMartial = attr
			case 2:
				nebulaPeople.Stewardship = attr
				nebulaPeople.ModifiedStewardship = attr
			case 3:
				nebulaPeople.Intrigue = attr
				nebulaPeople.ModifiedIntrigue = attr
			case 4:
				nebulaPeople.Learning = attr
				nebulaPeople.ModifiedLearning = attr
			}
		}
	}

	if !nebulaPeople.DeathDay.IsZero() {
		nebulaPeople.IsDead = true
		nebulaPeople.Age = int(nebulaPeople.DeathDay.Sub(nebulaPeople.Birthday).Hours() / 24.0 / 365.0)
	} else {
		nebulaPeople.Age = int(character.PlayDate.Sub(nebulaPeople.Birthday).Hours() / 24.0 / 365.0)
	}

	return &nebulaPeople
}

func getPeopleVid(playId int, peopleId int) string {
	return fmt.Sprintf("people.%d.%d", peopleId, playId)
}

func (p *People) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertVertexes(space, p)
}

func (p *People) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateVertexes(space, p)
}

func (p *People) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertVertexes(space, p)
}

func (p *People) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteVertexes(space, p)
}

func (p *People) DeleteFromNebulaWithEdge(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteVertexesWithEdges(space, p)
}

func (p *People) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadVertex(space, p)
}

func InsertPeoples(space *nebulagolang.Space, ps ...*People) *nebulagolang.Result {
	return nebulagolang.BatchInsertVertexes(space, 250, ps)
}

func UpdatePeoples(space *nebulagolang.Space, ps ...*People) *nebulagolang.Result {
	return nebulagolang.BatchUpdateVertexes(space, 250, ps)
}

func UpsertPeoples(space *nebulagolang.Space, ps ...*People) *nebulagolang.Result {
	return nebulagolang.BatchUpsertVertexes(space, 250, ps)
}

func DeletePeoples(space *nebulagolang.Space, playId int, ids ...int) *nebulagolang.Result {
	ts := make([]string, len(ids))
	for i, c := range ids {
		ts[i] = getPeopleVid(playId, c)
	}
	return nebulagolang.DeleteVertexesByVids(space, ts...)
}

func DeletePeopleWithEdge(space *nebulagolang.Space, playId int, ids ...int) *nebulagolang.Result {
	ts := make([]string, len(ids))
	for i, c := range ids {
		ts[i] = getPeopleVid(playId, c)
	}
	return nebulagolang.DeleteVertexesWithEdgesByVids(space, ts...)
}

func DeleteAllPeople(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllVertexesByTag[People](space)
}

func DeleteAllPeopleByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.Result {
	return nebulagolang.DeleteAllVertexesByQuery[People](space, getPlayIdQuery[People](playId))
}

func DeleteAllPeopleWithEdges(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllVertexesWithEdgesByTag[People](space)
}

func DeleteAllPeopleWithEdgesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.Result {
	return nebulagolang.DeleteAllVertexesWithEdgesByQuery[People](space, getPlayIdQuery[People](playId))
}

func GetPeopleByVid(space *nebulagolang.Space, vid string) *nebulagolang.ResultT[*People] {
	return nebulagolang.GetVertexByVid[*People](space, vid)
}

func GetPeopleByID(space *nebulagolang.Space, playId int, id int) *nebulagolang.ResultT[*People] {
	return nebulagolang.GetVertexByVid[*People](space, getPeopleVid(playId, id))
}

func GetAllPeoples(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*People] {
	return nebulagolang.GetAllVertexesByQuery[*People](space, "")
}

func GetAllPeoplesByPlayIdAndName(space *nebulagolang.Space, playId int, name string) *nebulagolang.ResultT[map[string]*People] {
	return nebulagolang.GetAllVertexesByQuery[*People](space, getPlayIdAndPropertyQuery[People](playId, "name", name))
}

func GetAllPeoplesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]*People] {
	return nebulagolang.GetAllVertexesByQuery[*People](space, getPlayIdQuery[People](playId))
}

func GetAllPeoplesVids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesVIDsByQuery[*People](space, "")
}

func GetAllPeoplesVidsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesVIDsByQuery[*People](space, getPlayIdQuery[People](playId))
}

func GetAllPeoplesIds(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesPropertyByQuery[*People](space, "", "id", "")
}

func GetAllPeoplesIdsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesPropertyByQuery[*People](space, getPlayIdQuery[People](playId), "id", "")
}

func GetAllPeoplesNames(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesPropertyByQuery[*People](space, "", "name", "")
}

func GetAllPeopleNamesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesPropertyByQuery[*People](space, getPlayIdQuery[People](playId), "name", "")
}

func GetAllPeoplesCodes(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesPropertyByQuery[*People](space, "", "code", "")
}

func GetAllPeopleCodesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesPropertyByQuery[*People](space, getPlayIdQuery[People](playId), "code", "")
}
