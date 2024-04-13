package ck2nebula

import (
	"fmt"
	"github.com/thalesfu/nebulagolang"
	"github.com/thalesfu/nebulagolang/utils"
	"github.com/thalesfu/paradoxtools/CK2/culture"
	"github.com/thalesfu/paradoxtools/CK2/localisation"
	"github.com/thalesfu/paradoxtools/CK2/religion"
	"github.com/thalesfu/paradoxtools/CK2/save"
	"log"
	"reflect"
	"strings"
	"time"
)

type StoryDetail struct {
	Story                         *Story
	Titles                        []*Title
	Title_BaseTitles              []*Title_BaseTitle
	Title_LiegeTitles             []*Title_LiegeTitle
	Title_DejureLiegeTitles       []*Title_DejureLiegeTitle
	Title_AssimilatingLiegeTitles []*Title_AssimilatingLiegeTitle
	Title_Dynasties               []*Title_Dynasty
	Title_People                  []*Title_People
	Story_Titles                  []*Story_Title
	Provinces                     []*Province
	Story_Provinces               []*Story_Province
	Province_Modifiers            []*Province_Modifier
	Province_Cultures             []*Province_Culture
	Province_Religions            []*Province_Religion
	Province_Titles               []*Province_Title
	Barons                        []*Baron
	Province_Barons               []*Province_Baron
	Baron_Buildings               []*Baron_Building
	Baron_Titles                  []*Baron_Title
	Story_Barons                  []*Story_Baron
	Dynasties                     []*Dynasty
	Dynasty_Cultures              []*Dynasty_Culture
	Dynasty_Religions             []*Dynasty_Religion
	Story_Dynasties               []*Story_Dynasty
	People                        []*People
	People_Cultures               []*People_Culture
	People_GFXCultures            []*People_GFXCulture
	People_Religions              []*People_Religion
	People_SecretReligions        []*People_SecretReligion
	Story_People                  []*Story_People
	People_Traits                 []*People_Trait
	People_Modifiers              []*People_Modifier
	People_ClaimTitles            []*People_ClaimTitle
	People_Dynasties              []*People_Dynasty
	People_Families               []*People_FamilyPeople
	People_Hosts                  []*People_HostPeople
	People_Empires                []*People_EmpirePeople
	People_Killers                []*People_KillPeople
	People_Relates                []*People_RelatePeople
	People_Lovers                 []*People_LoverPeople
	People_Guardians              []*People_GuardianPeople
	People_Ambtions               []*People_Ambition
	People_Focuses                []*People_Focus
	Story_Player                  []*Story_Player
}

type StoryUpdateDetail struct {
	Story                         *nebulagolang.CompareResult[*Story]
	Titles                        *nebulagolang.CompareResult[*Title]
	Title_BaseTitles              *nebulagolang.CompareResult[*Title_BaseTitle]
	Title_LiegeTitles             *nebulagolang.CompareResult[*Title_LiegeTitle]
	Title_DejureLiegeTitles       *nebulagolang.CompareResult[*Title_DejureLiegeTitle]
	Title_AssimilatingLiegeTitles *nebulagolang.CompareResult[*Title_AssimilatingLiegeTitle]
	Title_Dynasties               *nebulagolang.CompareResult[*Title_Dynasty]
	Title_People                  *nebulagolang.CompareResult[*Title_People]
	Story_Titles                  *nebulagolang.CompareResult[*Story_Title]
	Provinces                     *nebulagolang.CompareResult[*Province]
	Story_Provinces               *nebulagolang.CompareResult[*Story_Province]
	Province_Modifiers            *nebulagolang.CompareResult[*Province_Modifier]
	Province_Cultures             *nebulagolang.CompareResult[*Province_Culture]
	Province_Religions            *nebulagolang.CompareResult[*Province_Religion]
	Province_Titles               *nebulagolang.CompareResult[*Province_Title]
	Barons                        *nebulagolang.CompareResult[*Baron]
	Province_Barons               *nebulagolang.CompareResult[*Province_Baron]
	Baron_Buildings               *nebulagolang.CompareResult[*Baron_Building]
	Baron_Titles                  *nebulagolang.CompareResult[*Baron_Title]
	Story_Barons                  *nebulagolang.CompareResult[*Story_Baron]
	Dynasties                     *nebulagolang.CompareResult[*Dynasty]
	Dynasty_Cultures              *nebulagolang.CompareResult[*Dynasty_Culture]
	Dynasty_Religions             *nebulagolang.CompareResult[*Dynasty_Religion]
	Story_Dynasties               *nebulagolang.CompareResult[*Story_Dynasty]
	People                        *nebulagolang.CompareResult[*People]
	People_Cultures               *nebulagolang.CompareResult[*People_Culture]
	People_GFXCultures            *nebulagolang.CompareResult[*People_GFXCulture]
	People_Religions              *nebulagolang.CompareResult[*People_Religion]
	People_SecretReligions        *nebulagolang.CompareResult[*People_SecretReligion]
	Story_People                  *nebulagolang.CompareResult[*Story_People]
	People_Traits                 *nebulagolang.CompareResult[*People_Trait]
	People_Modifiers              *nebulagolang.CompareResult[*People_Modifier]
	People_ClaimTitles            *nebulagolang.CompareResult[*People_ClaimTitle]
	People_Dynasties              *nebulagolang.CompareResult[*People_Dynasty]
	People_Families               *nebulagolang.CompareResult[*People_FamilyPeople]
	People_Hosts                  *nebulagolang.CompareResult[*People_HostPeople]
	People_Empires                *nebulagolang.CompareResult[*People_EmpirePeople]
	People_Killers                *nebulagolang.CompareResult[*People_KillPeople]
	People_Relates                *nebulagolang.CompareResult[*People_RelatePeople]
	People_Lovers                 *nebulagolang.CompareResult[*People_LoverPeople]
	People_Guardians              *nebulagolang.CompareResult[*People_GuardianPeople]
	People_Ambtions               *nebulagolang.CompareResult[*People_Ambition]
	People_Focuses                *nebulagolang.CompareResult[*People_Focus]
	Story_Player                  *nebulagolang.CompareResult[*Story_Player]
}

func GetStory(path string, savePath string) *StoryDetail {
	log.Println("Get story.")

	s, ok, err := save.LoadSave(path, savePath)

	if err != nil {
		return nil
	}

	translations := localisation.LoadAllTranslations(path)

	if !ok {
		return nil
	}

	cgs := culture.LoadAllCultures(path)
	cm := make(map[string]string)
	rm := make(map[string]string)

	for _, cg := range cgs {
		for _, c := range cg.Cultures {
			cm[c.Code] = c.Name
		}
	}

	rgs := religion.LoadAllReligions(path)

	for _, rg := range rgs {
		for _, r := range rg.Religions {
			rm[r.Code] = r.Name
		}
	}

	detail := GenerateStoryDetails(s, cm, rm, translations)

	return detail
}

func GenerateStoryDetails(file *save.SaveFile, cm map[string]string, rm map[string]string, translations map[string]string) *StoryDetail {

	titles,
		baseTitles,
		liegeTitles,
		dejureLiegeTitles,
		assimilatingLiegeTitles,
		title_dynasties,
		title_people := GenerateTitles(file.Titles)

	provinces, province_modifiers, province_cultures, province_religions, province_titles, barons, province_barons, baron_buildings, baron_titles := GenerateProvinces(file.Provinces, cm, rm)

	dynasties, dynasty_cultures, dynasty_religions := GenerateDynasties(file.Dynasties, cm, rm)

	people,
		people_cultures,
		people_gfxCultures,
		people_religions,
		people_secretReligions,
		people_traits,
		people_modifiers,
		people_claimTitles,
		people_dynasties,
		people_families,
		people_hosts,
		people_empires,
		people_killers,
		people_lovers,
		people_guardians,
		people_ambtions,
		people_focuses := GeneratePeople(file, cm, rm)

	people_relations := GenerateRelations(file, translations)

	sd := &StoryDetail{
		Story:                         NewStoryByData(file),
		Titles:                        titles,
		Title_BaseTitles:              baseTitles,
		Title_LiegeTitles:             liegeTitles,
		Title_DejureLiegeTitles:       dejureLiegeTitles,
		Title_AssimilatingLiegeTitles: assimilatingLiegeTitles,
		Title_Dynasties:               title_dynasties,
		Title_People:                  title_people,
		Story_Titles:                  make([]*Story_Title, 0),
		Provinces:                     provinces,
		Story_Provinces:               make([]*Story_Province, 0),
		Province_Modifiers:            province_modifiers,
		Province_Cultures:             province_cultures,
		Province_Religions:            province_religions,
		Province_Titles:               province_titles,
		Barons:                        barons,
		Province_Barons:               province_barons,
		Baron_Buildings:               baron_buildings,
		Baron_Titles:                  baron_titles,
		Story_Barons:                  make([]*Story_Baron, 0),
		Dynasties:                     dynasties,
		Dynasty_Cultures:              dynasty_cultures,
		Dynasty_Religions:             dynasty_religions,
		Story_Dynasties:               make([]*Story_Dynasty, 0),
		People:                        people,
		People_Cultures:               people_cultures,
		People_GFXCultures:            people_gfxCultures,
		People_Religions:              people_religions,
		People_SecretReligions:        people_secretReligions,
		Story_People:                  make([]*Story_People, 0),
		People_Traits:                 people_traits,
		People_Modifiers:              people_modifiers,
		People_ClaimTitles:            people_claimTitles,
		People_Dynasties:              people_dynasties,
		People_Families:               people_families,
		People_Hosts:                  people_hosts,
		People_Empires:                people_empires,
		People_Killers:                people_killers,
		People_Relates:                people_relations,
		People_Lovers:                 people_lovers,
		People_Guardians:              people_guardians,
		People_Ambtions:               people_ambtions,
		People_Focuses:                people_focuses,
	}

	sd.Story.CultureName = cm[sd.Story.Culture]
	sd.Story.ReligionName = rm[sd.Story.Religion]

	sd.Story_Player = append(sd.Story_Player, NewStory_Player(sd.Story, NewPeople(sd.Story.StoryId, sd.Story.PlayerID)))

	for _, title := range titles {
		sd.Story_Titles = append(sd.Story_Titles, NewStory_Title(sd.Story, title))
	}

	for _, province := range provinces {
		sd.Story_Provinces = append(sd.Story_Provinces, NewStory_Province(sd.Story, province))
	}

	for _, baron := range barons {
		sd.Story_Barons = append(sd.Story_Barons, NewStory_Baron(sd.Story, baron))
	}

	for _, dynasty := range dynasties {
		sd.Story_Dynasties = append(sd.Story_Dynasties, NewStory_Dynasty(sd.Story, dynasty))
	}

	for _, p := range people {
		sd.Story_People = append(sd.Story_People, NewStory_People(sd.Story, p))
	}

	return sd
}

func getPlayIdQuery[T interface{}](playId int) string {
	return getPlayIdAndPropertiesQuery[T](playId, nil)
}

func getPlayIdAndPropertyQuery[T interface{}](playId int, propertyName string, propertyValue string) string {
	propertiesNamesAndValues := map[string]string{propertyName: propertyValue}

	return getPlayIdAndPropertiesQuery[T](playId, propertiesNamesAndValues)
}

func getPlayIdAndPropertiesQuery[T interface{}](playId int, propertiesNamesAndValues map[string]string) string {
	itemName := nebulagolang.GetTagName[T]()

	if itemName == "" {
		itemName = nebulagolang.GetEdgeName[T]()
	}

	if len(propertiesNamesAndValues) == 0 {
		return fmt.Sprintf("%s.story_id==%d", itemName, playId)
	}

	pnvs := make([]string, len(propertiesNamesAndValues))
	i := 0
	for propertyName, propertyValue := range propertiesNamesAndValues {
		pnvs[i] = fmt.Sprintf(" AND %s.%s==\"%s\"", itemName, propertyName, propertyValue)
		i++
	}

	return fmt.Sprintf("%s.story_id==%d%s", itemName, playId, strings.Join(pnvs, ""))
}

type CUResult struct {
	Name          string
	UpdateResult  *nebulagolang.Result
	CompareResult interface{}
	StartTime     time.Time
}

func LoadAndUpdateStory(path string, savePath string) (*StoryUpdateDetail, *nebulagolang.Result) {
	s := GetStory(path, savePath)

	if s == nil {
		return nil, nil
	}

	result := &StoryUpdateDetail{}

	rfv := GetStoryCompareAndUpdateDetailGenericFieldReflectValue(result)
	fields := make(map[string]bool)

	chanelDeep := len(rfv)

	sem := make(chan struct{}, 40)

	cuResultChanel := make(chan *CUResult, chanelDeep)

	CRData(cuResultChanel, sem, s.People_Relates, getPlayIdQuery[People_RelatePeople](s.Story.StoryId), fields)
	CRData(cuResultChanel, sem, s.People, getPlayIdQuery[People](s.Story.StoryId), fields)
	CRData(cuResultChanel, sem, s.Story_People, getPlayIdQuery[Story_People](s.Story.StoryId), fields)
	CRData(cuResultChanel, sem, s.People_Cultures, getPlayIdQuery[People_Culture](s.Story.StoryId), fields)
	CRData(cuResultChanel, sem, s.People_GFXCultures, getPlayIdQuery[People_GFXCulture](s.Story.StoryId), fields)
	CRData(cuResultChanel, sem, s.People_Religions, getPlayIdQuery[People_Religion](s.Story.StoryId), fields)
	CRData(cuResultChanel, sem, s.People_SecretReligions, getPlayIdQuery[People_SecretReligion](s.Story.StoryId), fields)
	CRData(cuResultChanel, sem, s.People_Traits, getPlayIdQuery[People_Trait](s.Story.StoryId), fields)
	CRData(cuResultChanel, sem, s.People_Modifiers, getPlayIdQuery[People_Modifier](s.Story.StoryId), fields)
	CRData(cuResultChanel, sem, s.People_ClaimTitles, getPlayIdQuery[People_ClaimTitle](s.Story.StoryId), fields)
	CRData(cuResultChanel, sem, s.People_Dynasties, getPlayIdQuery[People_Dynasty](s.Story.StoryId), fields)
	CRData(cuResultChanel, sem, s.People_Families, getPlayIdQuery[People_FamilyPeople](s.Story.StoryId), fields)
	CRData(cuResultChanel, sem, s.People_Hosts, getPlayIdQuery[People_HostPeople](s.Story.StoryId), fields)
	CRData(cuResultChanel, sem, s.People_Empires, getPlayIdQuery[People_EmpirePeople](s.Story.StoryId), fields)
	CRData(cuResultChanel, sem, s.People_Killers, getPlayIdQuery[People_KillPeople](s.Story.StoryId), fields)
	CRData(cuResultChanel, sem, s.People_Lovers, getPlayIdQuery[People_LoverPeople](s.Story.StoryId), fields)
	CRData(cuResultChanel, sem, s.People_Guardians, getPlayIdQuery[People_GuardianPeople](s.Story.StoryId), fields)
	CRData(cuResultChanel, sem, s.People_Ambtions, getPlayIdQuery[People_Ambition](s.Story.StoryId), fields)
	CRData(cuResultChanel, sem, s.People_Focuses, getPlayIdQuery[People_Focus](s.Story.StoryId), fields)
	CRData(cuResultChanel, sem, s.Titles, getPlayIdQuery[Title](s.Story.StoryId), fields)
	CRData(cuResultChanel, sem, s.Title_BaseTitles, getPlayIdQuery[Title_BaseTitle](s.Story.StoryId), fields)
	CRData(cuResultChanel, sem, s.Title_LiegeTitles, getPlayIdQuery[Title_LiegeTitle](s.Story.StoryId), fields)
	CRData(cuResultChanel, sem, s.Title_DejureLiegeTitles, getPlayIdQuery[Title_DejureLiegeTitle](s.Story.StoryId), fields)
	CRData(cuResultChanel, sem, s.Title_AssimilatingLiegeTitles, getPlayIdQuery[Title_AssimilatingLiegeTitle](s.Story.StoryId), fields)
	CRData(cuResultChanel, sem, s.Title_Dynasties, getPlayIdQuery[Title_Dynasty](s.Story.StoryId), fields)
	CRData(cuResultChanel, sem, s.Title_People, getPlayIdQuery[Title_People](s.Story.StoryId), fields)
	CRData(cuResultChanel, sem, s.Story_Titles, getPlayIdQuery[Story_Title](s.Story.StoryId), fields)
	CRData(cuResultChanel, sem, s.Provinces, getPlayIdQuery[Province](s.Story.StoryId), fields)
	CRData(cuResultChanel, sem, s.Story_Provinces, getPlayIdQuery[Story_Province](s.Story.StoryId), fields)
	CRData(cuResultChanel, sem, s.Province_Modifiers, getPlayIdQuery[Province_Modifier](s.Story.StoryId), fields)
	CRData(cuResultChanel, sem, s.Province_Cultures, getPlayIdQuery[Province_Culture](s.Story.StoryId), fields)
	CRData(cuResultChanel, sem, s.Province_Religions, getPlayIdQuery[Province_Religion](s.Story.StoryId), fields)
	CRData(cuResultChanel, sem, s.Province_Titles, getPlayIdQuery[Province_Title](s.Story.StoryId), fields)
	CRData(cuResultChanel, sem, s.Barons, getPlayIdQuery[Baron](s.Story.StoryId), fields)
	CRData(cuResultChanel, sem, s.Province_Barons, getPlayIdQuery[Province_Baron](s.Story.StoryId), fields)
	CRData(cuResultChanel, sem, s.Baron_Buildings, getPlayIdQuery[Baron_Building](s.Story.StoryId), fields)
	CRData(cuResultChanel, sem, s.Baron_Titles, getPlayIdQuery[Baron_Title](s.Story.StoryId), fields)
	CRData(cuResultChanel, sem, s.Story_Barons, getPlayIdQuery[Story_Baron](s.Story.StoryId), fields)
	CRData(cuResultChanel, sem, s.Dynasties, getPlayIdQuery[Dynasty](s.Story.StoryId), fields)
	CRData(cuResultChanel, sem, s.Dynasty_Cultures, getPlayIdQuery[Dynasty_Culture](s.Story.StoryId), fields)
	CRData(cuResultChanel, sem, s.Dynasty_Religions, getPlayIdQuery[Dynasty_Religion](s.Story.StoryId), fields)
	CRData(cuResultChanel, sem, s.Story_Dynasties, getPlayIdQuery[Story_Dynasty](s.Story.StoryId), fields)
	CRData(cuResultChanel, sem, s.Story_Player, getPlayIdQuery[Story_Player](s.Story.StoryId), fields)
	CRData(cuResultChanel, sem, []*Story{s.Story}, getPlayIdQuery[Story](s.Story.StoryId), fields)

	var updateResult *nebulagolang.Result

	for i := 0; i < chanelDeep; i++ {
		cur := <-cuResultChanel
		ur, ok := UpdateStoryCompareAndUpdateDetail(rfv, cur, fields)
		updateResult = ur

		if !ok {
			return result, ur
		}

		if len(fields) == 0 {
			break
		}
	}

	return result, updateResult
}

func CRData[T interface{}](chanel chan<- *CUResult, sem chan struct{}, data []T, query string, fields map[string]bool) {
	t := utils.GetType[T]()
	fields[t.Name()] = true
	sem <- struct{}{}
	go func() {
		defer func() { <-sem }()
		log.Printf("Compare and update %s.\n", t.Name())
		r := &CUResult{Name: t.Name(), StartTime: time.Now()}
		usr, csr := nebulagolang.CompareAndUpdateNebulaEntityBySliceAndQuery[T](SPACE, data, query, false)
		r.UpdateResult = usr
		r.CompareResult = csr
		diff := time.Now().Sub(r.StartTime)
		log.Printf("%sCompare and update %s done. Take %.0f seconds.%s\n", utils.PrintColorYellow, r.Name, diff.Seconds(), utils.PrintColorReset)
		chanel <- r
	}()
}

func GetStoryCompareAndUpdateDetailGenericFieldReflectValue(detail *StoryUpdateDetail) map[reflect.Type]reflect.Value {
	result := make(map[reflect.Type]reflect.Value)

	fv := utils.IndirectValue(reflect.ValueOf(detail))
	ft := fv.Type()

	for i := 0; i < ft.NumField(); i++ {
		f := ft.Field(i)
		v := fv.Field(i)
		result[f.Type] = v
	}

	return result
}

func UpdateStoryCompareAndUpdateDetail(detail map[reflect.Type]reflect.Value, result *CUResult, fields map[string]bool) (*nebulagolang.Result, bool) {
	if !result.UpdateResult.Ok {
		return result.UpdateResult, false
	}

	t := reflect.TypeOf(result.CompareResult)

	if v, ok := detail[t]; ok {
		v.Set(reflect.ValueOf(result.CompareResult))
	}
	delete(fields, result.Name)

	return result.UpdateResult, true
}

func BuildStory(path string, savePath string) {
	story, result := LoadAndUpdateStory(path, savePath)

	if !result.Ok {
		fmt.Println(result.Err)
	} else {
		printCompareAndUpdatedResult(story.Story)
		printCompareAndUpdatedResult(story.Titles)
		printCompareAndUpdatedResult(story.Title_BaseTitles)
		printCompareAndUpdatedResult(story.Title_LiegeTitles)
		printCompareAndUpdatedResult(story.Title_DejureLiegeTitles)
		printCompareAndUpdatedResult(story.Title_AssimilatingLiegeTitles)
		printCompareAndUpdatedResult(story.Title_Dynasties)
		printCompareAndUpdatedResult(story.Title_People)
		printCompareAndUpdatedResult(story.Story_Titles)
		printCompareAndUpdatedResult(story.Provinces)
		printCompareAndUpdatedResult(story.Story_Provinces)
		printCompareAndUpdatedResult(story.Province_Modifiers)
		printCompareAndUpdatedResult(story.Province_Cultures)
		printCompareAndUpdatedResult(story.Province_Religions)
		printCompareAndUpdatedResult(story.Province_Titles)
		printCompareAndUpdatedResult(story.Barons)
		printCompareAndUpdatedResult(story.Province_Barons)
		printCompareAndUpdatedResult(story.Baron_Buildings)
		printCompareAndUpdatedResult(story.Baron_Titles)
		printCompareAndUpdatedResult(story.Story_Barons)
		printCompareAndUpdatedResult(story.Dynasties)
		printCompareAndUpdatedResult(story.Dynasty_Cultures)
		printCompareAndUpdatedResult(story.Dynasty_Religions)
		printCompareAndUpdatedResult(story.Story_Dynasties)
		printCompareAndUpdatedResult(story.People)
		printCompareAndUpdatedResult(story.People_Cultures)
		printCompareAndUpdatedResult(story.People_GFXCultures)
		printCompareAndUpdatedResult(story.People_Religions)
		printCompareAndUpdatedResult(story.People_SecretReligions)
		printCompareAndUpdatedResult(story.Story_People)
		printCompareAndUpdatedResult(story.People_Traits)
		printCompareAndUpdatedResult(story.People_Modifiers)
		printCompareAndUpdatedResult(story.People_ClaimTitles)
		printCompareAndUpdatedResult(story.People_Dynasties)
		printCompareAndUpdatedResult(story.People_Families)
		printCompareAndUpdatedResult(story.People_Hosts)
		printCompareAndUpdatedResult(story.People_Empires)
		printCompareAndUpdatedResult(story.People_Killers)
		printCompareAndUpdatedResult(story.People_Relates)
		printCompareAndUpdatedResult(story.People_Lovers)
		printCompareAndUpdatedResult(story.People_Guardians)
		printCompareAndUpdatedResult(story.People_Ambtions)
		printCompareAndUpdatedResult(story.People_Focuses)
		printCompareAndUpdatedResult(story.Story_Player)
	}
}

func printCompareAndUpdatedResult[T interface{}](result *nebulagolang.CompareResult[T]) {
	name := utils.GetType[T]().Name()
	if result == nil {
		fmt.Printf("%sNo data for %s%s\n", utils.PrintColorRed, name, utils.PrintColorReset)
		return
	}

	if len(result.Added) > 0 {
		fmt.Printf("%s%s added: %d%s\n", utils.PrintColorGreen, name, result.AddedCount, utils.PrintColorReset)
	}

	if len(result.Updated) > 0 {
		fmt.Printf("%s%s updated: %d%s\n", utils.PrintColorYellow, name, result.UpdatedCount, utils.PrintColorReset)
	}

	if len(result.Deleted) > 0 {
		fmt.Printf("%s%s deleted: %d%s\n", utils.PrintColorRed, name, result.DeletedCount, utils.PrintColorReset)
	}

	if len(result.Kept) > 0 {
		fmt.Printf("%s%s kept: %d%s\n", utils.PrintColorWhite, name, result.DeletedCount, utils.PrintColorReset)
	}
}

func DeleteStoryData(playId int) *nebulagolang.Result {
	r := DeleteAllStory_PlayerByPlayId(SPACE, playId)

	if !r.Ok {
		return r
	}

	r = DeleteAllPeople_AmbitionsByPlayId(SPACE, playId)

	if !r.Ok {
		return r
	}

	r = DeleteAllPeople_FocusesByPlayId(SPACE, playId)

	if !r.Ok {
		return r
	}

	r = DeleteAllPeople_LoverByPlayId(SPACE, playId)

	if !r.Ok {
		return r
	}

	r = DeleteAllPeople_GuardianByPlayId(SPACE, playId)

	if !r.Ok {
		return r
	}

	r = DeleteAllPeople_FamilyByPlayId(SPACE, playId)

	if !r.Ok {
		return r
	}

	r = DeleteAllPeople_HostByPlayId(SPACE, playId)

	if !r.Ok {
		return r
	}

	r = DeleteAllPeople_EmpireByPlayId(SPACE, playId)

	if !r.Ok {
		return r
	}

	r = DeleteAllPeople_KillByPlayId(SPACE, playId)

	if !r.Ok {
		return r
	}

	r = DeleteAllPeople_RelateByPlayId(SPACE, playId)

	if !r.Ok {
		return r
	}

	r = DeleteAllPeople_ClaimTitlesByPlayId(SPACE, playId)

	if !r.Ok {
		return r
	}

	r = DeleteAllPeople_DynastiesByPlayId(SPACE, playId)

	if !r.Ok {
		return r
	}

	r = DeleteAllPeople_ModifiersByPlayId(SPACE, playId)

	if !r.Ok {
		return r
	}

	r = DeleteAllPeople_TraitsByPlayId(SPACE, playId)

	if !r.Ok {
		return r
	}

	r = DeleteAllStory_PeopleByPlayId(SPACE, playId)

	if !r.Ok {
		return r
	}

	r = DeleteAllPeople_SecretReligionsByPlayId(SPACE, playId)

	if !r.Ok {
		return r
	}

	r = DeleteAllPeople_ReligionsByPlayId(SPACE, playId)

	if !r.Ok {
		return r
	}

	r = DeleteAllPeople_GFXCulturesByPlayId(SPACE, playId)

	if !r.Ok {
		return r
	}

	r = DeleteAllPeople_CulturesByPlayId(SPACE, playId)

	if !r.Ok {
		return r
	}

	r = DeleteAllStory_DynastysByPlayId(SPACE, playId)

	if !r.Ok {
		return r
	}

	r = DeleteAllDynasty_ReligionsByPlayId(SPACE, playId)

	if !r.Ok {
		return r
	}

	r = DeleteAllDynasty_CulturesByPlayId(SPACE, playId)

	if !r.Ok {
		return r
	}

	r = DeleteAllStory_BaronsByPlayId(SPACE, playId)

	if !r.Ok {
		return r
	}

	r = DeleteAllBaron_BuildingsByPlayId(SPACE, playId)

	if !r.Ok {
		return r
	}

	r = DeleteAllBaron_TitlesByPlayId(SPACE, playId)

	if !r.Ok {
		return r
	}

	r = DeleteAllProvince_BaronsByPlayId(SPACE, playId)

	if !r.Ok {
		return r
	}

	r = DeleteAllProvince_ModifiersByPlayId(SPACE, playId)

	if !r.Ok {
		return r
	}

	r = DeleteAllProvince_ReligionsByPlayId(SPACE, playId)

	if !r.Ok {
		return r
	}

	r = DeleteAllProvince_CulturesByPlayId(SPACE, playId)

	if !r.Ok {
		return r
	}

	r = DeleteAllProvince_TitlesByPlayId(SPACE, playId)

	if !r.Ok {
		return r
	}

	r = DeleteAllStory_ProvincesByPlayId(SPACE, playId)

	if !r.Ok {
		return r
	}

	r = DeleteAllTitle_AssimilatingLiegeTitlesByPlayId(SPACE, playId)

	if !r.Ok {
		return r
	}

	r = DeleteAllTitle_DejureLiegeTitlesByPlayId(SPACE, playId)

	if !r.Ok {
		return r
	}

	r = DeleteAllTitle_LiegeTitlesByPlayId(SPACE, playId)

	if !r.Ok {
		return r
	}

	r = DeleteAllTitle_BaseTitlesByPlayId(SPACE, playId)

	if !r.Ok {
		return r
	}

	r = DeleteAllStory_TitlesByPlayId(SPACE, playId)

	if !r.Ok {
		return r
	}

	r = DeleteAllTitle_DynastysByPlayId(SPACE, playId)

	if !r.Ok {
		return r
	}

	r = DeleteAllTitle_PeoplesByPlayId(SPACE, playId)

	if !r.Ok {
		return r
	}

	r = DeleteAllPeopleByPlayId(SPACE, playId)

	if !r.Ok {
		return r
	}

	r = DeleteAllBaronByPlayId(SPACE, playId)

	if !r.Ok {
		return r
	}

	r = DeleteAllDynastyByPlayId(SPACE, playId)

	if !r.Ok {
		return r
	}

	r = DeleteAllTitleByPlayId(SPACE, playId)

	if !r.Ok {
		return r
	}

	r = DeleteAllBaronByPlayId(SPACE, playId)

	if !r.Ok {
		return r
	}

	r = DeleteAllProvinceByPlayId(SPACE, playId)

	if !r.Ok {
		return r
	}

	return DeleteStories(SPACE, playId)
}
