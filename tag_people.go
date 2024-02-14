package ck2nebula

import (
	"fmt"
	"github.com/thalesfu/nebulagolang"
	"github.com/thalesfu/paradoxtools/CK2/flags"
	"github.com/thalesfu/paradoxtools/CK2/save"
	"github.com/thalesfu/paradoxtools/utils"
	"time"
)

type People struct {
	VID                  string    `nebulakey:"vid" nebulatagname:"people" nebulatagcomment:"people" json:"vid,omitempty"`
	ID                   int       `nebulaproperty:"id" description:"id" nebulaindexes:"id" json:"id,omitempty"`
	Name                 string    `nebulaproperty:"name" description:"name" nebulaindexes:"name" json:"name,omitempty"`
	Dynasty              int       `nebulaproperty:"dynasty" description:"dynasty" nebulaindexes:"dynasty" json:"dynasty,omitempty"`
	DynastyName          string    `nebulaproperty:"dynasty_name" description:"dynasty name" nebulaindexes:"dynasty_name" json:"dynasty_name,omitempty"`
	Birthday             time.Time `nebulaproperty:"birthday" nebulatype:"Date" description:"birthday" nebulaindexes:"birthday" json:"birthday,omitempty"`
	DeathDay             time.Time `nebulaproperty:"deathday" nebulatype:"Date" description:"deathday" nebulaindexes:"deathday" json:"deathday,omitempty"`
	Age                  int       `nebulaproperty:"age" description:"age" nebulaindexes:"age" json:"age,omitempty"`
	IsDead               bool      `nebulaproperty:"isdead" description:"isdead" nebulaindexes:"isdead" json:"isdead,omitempty"`
	IsBastard            bool      `nebulaproperty:"isbastard" description:"isbastard" nebulaindexes:"isbastard" json:"isbastard,omitempty"`
	Diplomacy            int       `nebulaproperty:"diplomacy" description:"diplomacy" nebulaindexes:"diplomacy" json:"diplomacy,omitempty"`
	Martial              int       `nebulaproperty:"martial" description:"martial" nebulaindexes:"martial" json:"martial,omitempty"`
	Stewardship          int       `nebulaproperty:"stewardship" description:"stewardship" nebulaindexes:"stewardship" json:"stewardship,omitempty"`
	Intrigue             int       `nebulaproperty:"intrigue" description:"intrigue" nebulaindexes:"intrigue" json:"intrigue,omitempty"`
	Learning             int       `nebulaproperty:"learning" description:"learning" nebulaindexes:"learning" json:"learning,omitempty"`
	ModifiedDiplomacy    int       `nebulaproperty:"modified_diplomacy" description:"modified diplomacy" nebulaindexes:"modified_diplomacy" json:"modified_diplomacy,omitempty"`
	ModifiedMartial      int       `nebulaproperty:"modified_martial" description:"modified martial" nebulaindexes:"modified_martial" json:"modified_martial,omitempty"`
	ModifiedStewardship  int       `nebulaproperty:"modified_stewardship" description:"modified stewardship" nebulaindexes:"modified_stewardship" json:"modified_stewardship,omitempty"`
	ModifiedIntrigue     int       `nebulaproperty:"modified_intrigue" description:"modified intrigue" nebulaindexes:"modified_intrigue" json:"modified_intrigue,omitempty"`
	ModifiedLearning     int       `nebulaproperty:"modified_learning" description:"modified learning" nebulaindexes:"modified_learning" json:"modified_learning,omitempty"`
	ModifiedSexAppeal    int       `nebulaproperty:"modified_sex_appeal" description:"modified sex appeal" nebulaindexes:"modified_sex_appeal" json:"modified_sex_appeal,omitempty"`
	ModifiedHealth       float32   `nebulaproperty:"modified_health" description:"modified health" nebulaindexes:"modified_health" json:"modified_health,omitempty"`
	ModifiedFertility    float32   `nebulaproperty:"modified_fertility" description:"modified fertility" nebulaindexes:"modified_fertility" json:"modified_fertility,omitempty"`
	ModifiedCombatRating int       `nebulaproperty:"modified_combat_rating" description:"modified_combat_rating" nebulaindexes:"modified_combat_rating" json:"modified_combat_rating,omitempty"`
	IsHistory            bool      `nebulaproperty:"ishistory" description:"ishistory" nebulaindexes:"ishistory" json:"ishistory,omitempty"`
	IsFemale             bool      `nebulaproperty:"isfemale" description:"isfemale" nebulaindexes:"isfemale" json:"isfemale,omitempty"`
	InHiding             bool      `nebulaproperty:"in_hiding" description:"in_hiding" nebulaindexes:"in_hiding" json:"in_hiding,omitempty"`
	IsPregnant           bool      `nebulaproperty:"is_pregnant" description:"怀孕的" nebulaindexes:"is_pregnant" json:"is_pregnant,omitempty"`
	IsPlayer             bool      `nebulaproperty:"isplayer" description:"isplayer" nebulaindexes:"isplayer" json:"isplayer,omitempty"`
	IsBuilt              bool      `nebulaproperty:"isbuilt" description:"isbuilt" nebulaindexes:"isbuilt" json:"isbuilt,omitempty"`
	NickNameCode         string    `nebulaproperty:"nicknamecode" description:"nicknamecode" nebulaindexes:"nicknamecode" json:"nicknamecode,omitempty"`
	NickName             string    `nebulaproperty:"nickname" description:"nickname" nebulaindexes:"nickname" json:"nickname,omitempty"`
	DNA                  string    `nebulaproperty:"dna" description:"dna" nebulaindexes:"dna" json:"dna,omitempty"`
	Government           string    `nebulaproperty:"gov" description:"gov" nebulaindexes:"gov" json:"gov,omitempty"`
	Culture              string    `nebulaproperty:"culture" description:"culture" nebulaindexes:"culture" json:"culture,omitempty"`
	CultureName          string    `nebulaproperty:"culture_name" description:"culture name" nebulaindexes:"culture_name" json:"culture_name,omitempty"`
	GFXCulture           string    `nebulaproperty:"gfx_culture" description:"gfx_culture" nebulaindexes:"gfx_culture" json:"gfx_culture,omitempty"`
	GFXCultureName       string    `nebulaproperty:"gfx_culture_name" description:"gfx_culture_name" nebulaindexes:"gfx_culture_name" json:"gfx_culture_name,omitempty"`
	Religion             string    `nebulaproperty:"religion" description:"religion" nebulaindexes:"religion" json:"religion,omitempty"`
	ReligionName         string    `nebulaproperty:"religion_name" description:"religion name" nebulaindexes:"religion_name" json:"religion_name,omitempty"`
	SecretReligion       string    `nebulaproperty:"secret_religion" description:"secret_religion" nebulaindexes:"secret_religion" json:"secret_religion,omitempty"`
	SecretReligionName   string    `nebulaproperty:"secret_religion_name" description:"secret_religion_name" nebulaindexes:"secret_religion_name" json:"secret_religion_name,omitempty"`
	Fertility            float32   `nebulaproperty:"fertility" description:"fertility" nebulaindexes:"fertility" json:"fertility,omitempty"`
	Health               float32   `nebulaproperty:"health" description:"health" nebulaindexes:"health" json:"health,omitempty"`
	Prestige             float32   `nebulaproperty:"prestige" description:"prestige" nebulaindexes:"prestige" json:"prestige,omitempty"`
	Piety                float32   `nebulaproperty:"piety" description:"piety" nebulaindexes:"piety" json:"piety,omitempty"`
	Wealth               float32   `nebulaproperty:"wealth" description:"wealth" nebulaindexes:"wealth" json:"wealth,omitempty"`
	Score                float32   `nebulaproperty:"score" description:"score" nebulaindexes:"score" json:"score,omitempty"`
	Society              int       `nebulaproperty:"society" description:"society" nebulaindexes:"society" json:"society,omitempty"`
	StoryID              int       `nebulaproperty:"story_id" mappingalias:"PlayID" description:"game play id" nebulaindexes:"story_id" json:"story_id,omitempty"`
}

func NewPeople(storyId int, peopleId int) *People {
	nebulaPeople := People{
		VID:     getPeopleVid(storyId, peopleId),
		ID:      peopleId,
		StoryID: storyId,
	}
	return &nebulaPeople
}

func NewPeopleByData(character *save.Character) *People {
	nebulaPeople := utils.Mapping[People](character)
	nebulaPeople.VID = getPeopleVid(nebulaPeople.StoryID, nebulaPeople.ID)
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

	nebulaPeople.ModifiedHealth = nebulaPeople.Health
	nebulaPeople.ModifiedFertility = nebulaPeople.Fertility

	if !nebulaPeople.DeathDay.IsZero() {
		nebulaPeople.IsDead = true
		nebulaPeople.Age = int(nebulaPeople.DeathDay.Sub(nebulaPeople.Birthday).Hours() / 24.0 / 365.0)
	} else {
		nebulaPeople.Age = int(character.PlayDate.Sub(nebulaPeople.Birthday).Hours() / 24.0 / 365.0)
	}

	_, exist := character.Flags[flags.BuiltPeopleFlag]
	nebulaPeople.IsBuilt = exist

	if len(character.Unborn) > 0 {
		nebulaPeople.IsPregnant = true
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

func (p *People) GetAliveFamilies(space *nebulagolang.Space) *nebulagolang.ResultT[map[int]*People] {
	command := fmt.Sprintf("go from \"%s\" over people_familypeople where properties($$).isdead==false yield $$ as v", p.VID)

	return p.getPeopleByQuery(space, command)
}

func (p *People) GetFamilies(space *nebulagolang.Space) *nebulagolang.ResultT[map[int]*People] {
	command := fmt.Sprintf("go from \"%s\" over people_familypeople yield $$ as v", p.VID)

	return p.getPeopleByQuery(space, command)
}

func (p *People) GetAliveDistanceFamilies(space *nebulagolang.Space) *nebulagolang.ResultT[map[int]*People] {
	commands := []string{
		fmt.Sprintf("fetch prop on people \"%s\" yield id(vertex) as vid, properties(vertex).dynasty as dynasty", p.VID),
		"go 1 to 2 steps from $-.vid over people_familypeople BIDIRECT where properties($$).isdead==false and (properties(edge).relation==\"son\" or properties(edge).relation==\"daughter\") and properties($$).dynasty==$-.dynasty yield distinct id($$) as myfamilyids,$-.vid as vid,$-.dynasty as dynasty",
		"yield collect($-.myfamilyids) as myfamilyids, $-.vid as vid, $-.dynasty as dynasty",
		"go from $-.vid over people_dynasty yield id($$) as vid, $-.myfamilyids as myfamilyids",
		"go from $-.vid over people_dynasty reversely where id($$) not in $-.myfamilyids and properties($$).isdead==false yield $$ as v",
	}

	return p.getPeopleByQuery(space, nebulagolang.CommandPipelineCombine(commands...))
}

func (p *People) GetAliveBrothersAndSisters(space *nebulagolang.Space) *nebulagolang.ResultT[map[int]*People] {
	commands := []string{
		fmt.Sprintf("go from \"%s\" over people_familypeople where people_familypeople.relation==\"father\" or people_familypeople.relation==\"mother\" yield $$ as v", p.VID),
		"yield id($-.v) as vid",
		"go from $-.vid over people_familypeople where (people_familypeople.relation==\"son\" or people_familypeople.relation==\"real_son\" or people_familypeople.relation==\"daughter\" or people_familypeople.relation==\"real_daughter\") and properties($$).isdead==false yield $$ as v",
		fmt.Sprintf("yield distinct $-.v as v where id($-.v)!=\"%s\"", p.VID),
	}

	return p.getPeopleByQuery(space, nebulagolang.CommandPipelineCombine(commands...))
}

func (p *People) GetBrothersAndSisters(space *nebulagolang.Space) *nebulagolang.ResultT[map[int]*People] {
	commands := []string{
		fmt.Sprintf("go from \"%s\" over people_familypeople where people_familypeople.relation==\"father\" or people_familypeople.relation==\"mother\" yield $$ as v", p.VID),
		"yield id($-.v) as vid",
		"go from $-.vid over people_familypeople where people_familypeople.relation==\"son\" or people_familypeople.relation==\"real_son\" or people_familypeople.relation==\"daughter\" or people_familypeople.relation==\"real_daughter\" and properties($$).isdead==false yield $$ as v",
		fmt.Sprintf("yield $-.v as v where id($-.v)!=\"%s\"", p.VID),
	}

	return p.getPeopleByQuery(space, nebulagolang.CommandPipelineCombine(commands...))
}

func (p *People) GetAliveGrandSon(space *nebulagolang.Space) *nebulagolang.ResultT[map[int]*People] {
	commands := []string{
		fmt.Sprintf("go from \"%s\" over people_familypeople where properties(edge).relation==\"son\" or properties(edge).relation==\"real_son\" or properties(edge).relation==\"da\nughter\" or properties(edge).relation==\"real_daughter\" yield id($$) as vid", p.VID),
		fmt.Sprintf("go from $-.vid over people_familypeople where properties(edge).relation==\"son\" and properties($$).dynasty==%d and properties($$).isdead==false yield $$ as v", p.Dynasty),
	}

	return p.getPeopleByQuery(space, nebulagolang.CommandPipelineCombine(commands...))
}

func (p *People) GetGrandSon(space *nebulagolang.Space) *nebulagolang.ResultT[map[int]*People] {
	commands := []string{
		fmt.Sprintf("go from \"%s\" over people_familypeople where properties(edge).relation==\"son\" or properties(edge).relation==\"real_son\" or properties(edge).relation==\"da\nughter\" or properties(edge).relation==\"real_daughter\" yield id($$) as vid", p.VID),
		fmt.Sprintf("go from $-.vid over people_familypeople where properties(edge).relation==\"son\" and properties($$).dynasty==%d yield $$ as v", p.Dynasty),
	}

	return p.getPeopleByQuery(space, nebulagolang.CommandPipelineCombine(commands...))
}

func (p *People) GetRealMother(space *nebulagolang.Space) *nebulagolang.ResultT[map[int]*People] {
	command := fmt.Sprintf("go from \"%s\" over people_familypeople where people_familypeople.relation==\"mother\" yield $$ as v", p.VID)

	return p.getPeopleByQuery(space, command)
}

func (p *People) GetAliveRealDaughter(space *nebulagolang.Space) *nebulagolang.ResultT[map[int]*People] {
	command := fmt.Sprintf("go from \"%s\" over people_familypeople where people_familypeople.relation==\"real_daughter\" and properties($$).isdead==false yield $$ as v", p.VID)

	return p.getPeopleByQuery(space, command)
}

func (p *People) GetRealDaughter(space *nebulagolang.Space) *nebulagolang.ResultT[map[int]*People] {
	command := fmt.Sprintf("go from \"%s\" over people_familypeople where people_familypeople.relation==\"real_daughter\" yield $$ as v", p.VID)

	return p.getPeopleByQuery(space, command)
}

func (p *People) GetAliveDaughter(space *nebulagolang.Space) *nebulagolang.ResultT[map[int]*People] {
	command := fmt.Sprintf("go from \"%s\" over people_familypeople where people_familypeople.relation==\"daughter\" and properties($$).isdead==false yield $$ as v", p.VID)

	return p.getPeopleByQuery(space, command)
}

func (p *People) GetDaughter(space *nebulagolang.Space) *nebulagolang.ResultT[map[int]*People] {
	command := fmt.Sprintf("go from \"%s\" over people_familypeople where people_familypeople.relation==\"daughter\" yield $$ as v", p.VID)

	return p.getPeopleByQuery(space, command)
}

func (p *People) GetAliveRealSon(space *nebulagolang.Space) *nebulagolang.ResultT[map[int]*People] {
	command := fmt.Sprintf("go from \"%s\" over people_familypeople where people_familypeople.relation==\"real_son\" and properties($$).isdead==false yield $$ as v", p.VID)

	return p.getPeopleByQuery(space, command)
}

func (p *People) GetRealSon(space *nebulagolang.Space) *nebulagolang.ResultT[map[int]*People] {
	command := fmt.Sprintf("go from \"%s\" over people_familypeople where people_familypeople.relation==\"real_son\" yield $$ as v", p.VID)

	return p.getPeopleByQuery(space, command)
}

func (p *People) GetAliveSon(space *nebulagolang.Space) *nebulagolang.ResultT[map[int]*People] {
	command := fmt.Sprintf("go from \"%s\" over people_familypeople where people_familypeople.relation==\"son\" and properties($$).isdead==false yield $$ as v", p.VID)

	return p.getPeopleByQuery(space, command)
}

func (p *People) GetSon(space *nebulagolang.Space) *nebulagolang.ResultT[map[int]*People] {
	command := fmt.Sprintf("go from \"%s\" over people_familypeople where people_familypeople.relation==\"son\" yield $$ as v", p.VID)

	return p.getPeopleByQuery(space, command)
}

func (p *People) GetRealFather(space *nebulagolang.Space) *nebulagolang.ResultT[map[int]*People] {
	command := fmt.Sprintf("go from \"%s\" over people_familypeople where people_familypeople.relation==\"real_father\" yield $$ as v", p.VID)

	return p.getPeopleByQuery(space, command)
}

func (p *People) GetFather(space *nebulagolang.Space) *nebulagolang.ResultT[map[int]*People] {
	command := fmt.Sprintf("go from \"%s\" over people_familypeople where people_familypeople.relation==\"father\" yield $$ as v", p.VID)

	return p.getPeopleByQuery(space, command)
}

func (p *People) GetAliveConsortOf(space *nebulagolang.Space) *nebulagolang.ResultT[map[int]*People] {
	command := fmt.Sprintf("go from \"%s\" over people_familypeople where people_familypeople.relation==\"consort_of\" and properties($$).isdead==false yield $$ as v", p.VID)

	return p.getPeopleByQuery(space, command)
}

func (p *People) GetConsortOf(space *nebulagolang.Space) *nebulagolang.ResultT[map[int]*People] {
	command := fmt.Sprintf("go from \"%s\" over people_familypeople where people_familypeople.relation==\"consort_of\" yield $$ as v", p.VID)

	return p.getPeopleByQuery(space, command)
}

func (p *People) GetAliveConsorts(space *nebulagolang.Space) *nebulagolang.ResultT[map[int]*People] {
	command := fmt.Sprintf("go from \"%s\" over people_familypeople where people_familypeople.relation==\"consort\" and properties($$).isdead==false yield $$ as v", p.VID)

	return p.getPeopleByQuery(space, command)
}

func (p *People) GetConsorts(space *nebulagolang.Space) *nebulagolang.ResultT[map[int]*People] {
	command := fmt.Sprintf("go from \"%s\" over people_familypeople where people_familypeople.relation==\"consort\" yield $$ as v", p.VID)

	return p.getPeopleByQuery(space, command)
}

func (p *People) GetAliveSpouses(space *nebulagolang.Space) *nebulagolang.ResultT[map[int]*People] {
	command := fmt.Sprintf("go from \"%s\" over people_familypeople where people_familypeople.relation==\"spouse\" and properties($$).isdead==false yield $$ as v", p.VID)

	return p.getPeopleByQuery(space, command)
}

func (p *People) GetSpouses(space *nebulagolang.Space) *nebulagolang.ResultT[map[int]*People] {
	command := fmt.Sprintf("go from \"%s\" over people_familypeople where people_familypeople.relation==\"spouse\" yield $$ as v", p.VID)

	return p.getPeopleByQuery(space, command)
}

func (p *People) GetFriends(space *nebulagolang.Space) *nebulagolang.ResultT[map[int]*People] {
	command := fmt.Sprintf("go from \"%s\" over people_relatepeople where people_relatepeople.code==\"opinion_friend\" yield $$ as v", p.VID)

	return p.getPeopleByQuery(space, command)
}

func (p *People) getPeopleByQuery(space *nebulagolang.Space, command string) *nebulagolang.ResultT[map[int]*People] {
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

func (p *People) GetTraits(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*Trait] {
	tr := nebulagolang.QueryVertexesByQueryToSlice[*Trait](space, fmt.Sprintf("go from \"%s\" over people_trait yield $$ as v", p.VID))

	if !tr.Ok {
		return nebulagolang.NewResultT[map[string]*Trait](tr.Result)
	}

	result := make(map[string]*Trait)

	for _, t := range tr.Data {
		result[t.Code] = t
	}

	return nebulagolang.NewResultTWithData(tr.Result, result)
}

func (p *People) GetEductionTraits(space *nebulagolang.Space) *nebulagolang.ResultT[[]*Trait] {
	return nebulagolang.QueryVertexesByQueryToSlice[*Trait](space, fmt.Sprintf("go from \"%s\" over people_trait where properties($$).education==true yield $$ as v", p.VID))
}

func (p *People) IsSuccessionMarriage(space *nebulagolang.Space) bool {
	command := fmt.Sprintf("go from \"%s\" over people_relatepeople BIDIRECT where properties(edge).code==\"succession\" yield edge as e", p.VID)
	r := nebulagolang.CountByQuery(space, command)

	if !r.Ok {
		return false
	}

	return r.Data > 0
}

func (p *People) GetVassals(space *nebulagolang.Space) *nebulagolang.ResultT[map[int]*People] {
	commands := []string{
		fmt.Sprintf("go from \"%s\" over title_people reversely yield id($$) as vid", p.VID),
		"go from $-.vid over title_liegetitle reversely yield id($$) as tid",
		fmt.Sprintf("go from $-.tid over title_people where id($$)!=\"%s\" yield distinct $$ as v", p.VID),
	}

	return p.getPeopleByQuery(space, nebulagolang.CommandPipelineCombine(commands...))
}

func (p *People) GetEmpirePeople(space *nebulagolang.Space) *nebulagolang.ResultT[map[int]*People] {
	command := fmt.Sprintf("go from \"%s\" over people_empirepeople reversely yield $$ as v", p.VID)

	return p.getPeopleByQuery(space, command)
}

func (p *People) IsReligionVassal(space *nebulagolang.Space) bool {
	r := p.GetBarons(space)

	if !r.Ok || len(r.Data) == 0 {
		return false
	}

	for _, b := range r.Data {
		if b.Type != "temple" {
			return false
		}
	}

	return true
}

func (p *People) GetBarons(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*Baron] {
	commands := []string{
		fmt.Sprintf("go from \"%s\" over title_people reversely yield id($$) as vid", p.VID),
		"go from $-.vid over baron_title reversely yield distinct $$ as v",
	}

	r := nebulagolang.QueryVertexesByQueryToSlice[*Baron](space, nebulagolang.CommandPipelineCombine(commands...))

	if !r.Ok {
		return nebulagolang.NewResultT[map[string]*Baron](r.Result)
	}

	result := make(map[string]*Baron)

	for _, b := range r.Data {
		result[b.Code] = b
	}

	return nebulagolang.NewResultTWithData(r.Result, result)
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
