package ck2nebula

import (
	"fmt"
	"github.com/thalesfu/nebulagolang"
	"github.com/thalesfu/paradoxtools/CK2/culture"
	"github.com/thalesfu/paradoxtools/CK2/localisation"
	"github.com/thalesfu/paradoxtools/CK2/religion"
	"github.com/thalesfu/paradoxtools/CK2/save"
	"strings"
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
}

func GetStory(path string, savePath string) *StoryDetail {
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
		return fmt.Sprintf("%s.play_id==%d", itemName, playId)
	}

	pnvs := make([]string, len(propertiesNamesAndValues))
	i := 0
	for propertyName, propertyValue := range propertiesNamesAndValues {
		pnvs[i] = fmt.Sprintf(" AND %s.%s==\"%s\"", itemName, propertyName, propertyValue)
		i++
	}

	return fmt.Sprintf("%s.play_id==%d%s", itemName, playId, strings.Join(pnvs, ""))
}

func LoadAndUpdateStory(path string, savePath string) (*StoryUpdateDetail, *nebulagolang.Result) {
	s := GetStory(path, savePath)

	if s == nil {
		return nil, nil
	}

	result := &StoryUpdateDetail{}

	usr, csr := nebulagolang.CompareAndUpdateNebulaEntityBySliceAndQuery(SPACE, []*Story{s.Story}, getPlayIdQuery[Story](s.Story.PlayID))

	if !usr.Ok {
		return result, usr
	}

	result.Story = csr

	utr, ctr := nebulagolang.CompareAndUpdateNebulaEntityBySliceAndQuery(SPACE, s.Titles, getPlayIdQuery[Title](s.Story.PlayID))

	if !utr.Ok {
		return result, utr
	}

	result.Titles = ctr

	utbtr, ctbtr := nebulagolang.CompareAndUpdateNebulaEntityBySliceAndQuery(SPACE, s.Title_BaseTitles, getPlayIdQuery[Title_BaseTitle](s.Story.PlayID))

	if !utbtr.Ok {
		return result, utbtr
	}

	result.Title_BaseTitles = ctbtr

	utltr, ctltr := nebulagolang.CompareAndUpdateNebulaEntityBySliceAndQuery(SPACE, s.Title_LiegeTitles, getPlayIdQuery[Title_LiegeTitle](s.Story.PlayID))

	if !utltr.Ok {
		return result, utltr
	}

	result.Title_LiegeTitles = ctltr

	utdltr, ctdltr := nebulagolang.CompareAndUpdateNebulaEntityBySliceAndQuery(SPACE, s.Title_DejureLiegeTitles, getPlayIdQuery[Title_DejureLiegeTitle](s.Story.PlayID))

	if !utdltr.Ok {
		return result, utdltr
	}

	result.Title_DejureLiegeTitles = ctdltr

	utaltr, ctaltr := nebulagolang.CompareAndUpdateNebulaEntityBySliceAndQuery(SPACE, s.Title_AssimilatingLiegeTitles, getPlayIdQuery[Title_AssimilatingLiegeTitle](s.Story.PlayID))

	if !utaltr.Ok {
		return result, utaltr
	}

	result.Title_AssimilatingLiegeTitles = ctaltr

	utdr, ctdr := nebulagolang.CompareAndUpdateNebulaEntityBySliceAndQuery(SPACE, s.Title_Dynasties, getPlayIdQuery[Title_Dynasty](s.Story.PlayID))

	if !utdr.Ok {
		return result, utdr
	}

	result.Title_Dynasties = ctdr

	utp, ctp := nebulagolang.CompareAndUpdateNebulaEntityBySliceAndQuery(SPACE, s.Title_People, getPlayIdQuery[Title_People](s.Story.PlayID))

	if !utp.Ok {
		return result, utp
	}

	result.Title_People = ctp

	ustr, cstr := nebulagolang.CompareAndUpdateNebulaEntityBySliceAndQuery(SPACE, s.Story_Titles, getPlayIdQuery[Story_Title](s.Story.PlayID))

	if !ustr.Ok {
		return result, ustr
	}

	result.Story_Titles = cstr

	upr, cpr := nebulagolang.CompareAndUpdateNebulaEntityBySliceAndQuery(SPACE, s.Provinces, getPlayIdQuery[Province](s.Story.PlayID))

	if !upr.Ok {
		return result, upr
	}

	result.Provinces = cpr

	uspr, cspr := nebulagolang.CompareAndUpdateNebulaEntityBySliceAndQuery(SPACE, s.Story_Provinces, getPlayIdQuery[Story_Province](s.Story.PlayID))

	if !uspr.Ok {
		return result, uspr
	}

	result.Story_Provinces = cspr

	upmr, cpmr := nebulagolang.CompareAndUpdateNebulaEntityBySliceAndQuery(SPACE, s.Province_Modifiers, getPlayIdQuery[Province_Modifier](s.Story.PlayID))

	if !upmr.Ok {
		return result, upmr
	}

	result.Province_Modifiers = cpmr

	upcr, cpcr := nebulagolang.CompareAndUpdateNebulaEntityBySliceAndQuery(SPACE, s.Province_Cultures, getPlayIdQuery[Province_Culture](s.Story.PlayID))

	if !upcr.Ok {
		return result, upcr
	}

	result.Province_Cultures = cpcr

	uprr, cprr := nebulagolang.CompareAndUpdateNebulaEntityBySliceAndQuery(SPACE, s.Province_Religions, getPlayIdQuery[Province_Religion](s.Story.PlayID))

	if !uprr.Ok {
		return result, uprr
	}

	result.Province_Religions = cprr

	uptr, cptr := nebulagolang.CompareAndUpdateNebulaEntityBySliceAndQuery(SPACE, s.Province_Titles, getPlayIdQuery[Province_Title](s.Story.PlayID))

	if !uptr.Ok {
		return result, uptr
	}

	result.Province_Titles = cptr

	ubr, cbr := nebulagolang.CompareAndUpdateNebulaEntityBySliceAndQuery(SPACE, s.Barons, getPlayIdQuery[Baron](s.Story.PlayID))

	if !ubr.Ok {
		return result, ubr
	}

	result.Barons = cbr

	upbr, cpbr := nebulagolang.CompareAndUpdateNebulaEntityBySliceAndQuery(SPACE, s.Province_Barons, getPlayIdQuery[Province_Baron](s.Story.PlayID))

	if !upbr.Ok {
		return result, upbr
	}

	result.Province_Barons = cpbr

	ubbr, cbbr := nebulagolang.CompareAndUpdateNebulaEntityBySliceAndQuery(SPACE, s.Baron_Buildings, getPlayIdQuery[Baron_Building](s.Story.PlayID))

	if !ubbr.Ok {
		return result, ubbr
	}

	result.Baron_Buildings = cbbr

	ubtr, cbtr := nebulagolang.CompareAndUpdateNebulaEntityBySliceAndQuery(SPACE, s.Baron_Titles, getPlayIdQuery[Baron_Title](s.Story.PlayID))

	if !ubtr.Ok {
		return result, ubtr
	}

	result.Baron_Titles = cbtr

	usb, csb := nebulagolang.CompareAndUpdateNebulaEntityBySliceAndQuery(SPACE, s.Story_Barons, getPlayIdQuery[Story_Baron](s.Story.PlayID))

	if !usb.Ok {
		return result, usb
	}

	result.Story_Barons = csb

	udr, cdr := nebulagolang.CompareAndUpdateNebulaEntityBySliceAndQuery(SPACE, s.Dynasties, getPlayIdQuery[Dynasty](s.Story.PlayID))

	if !udr.Ok {
		return result, udr
	}

	result.Dynasties = cdr

	udcr, cdcr := nebulagolang.CompareAndUpdateNebulaEntityBySliceAndQuery(SPACE, s.Dynasty_Cultures, getPlayIdQuery[Dynasty_Culture](s.Story.PlayID))

	if !udcr.Ok {
		return result, udcr
	}

	result.Dynasty_Cultures = cdcr

	udrr, cdrr := nebulagolang.CompareAndUpdateNebulaEntityBySliceAndQuery(SPACE, s.Dynasty_Religions, getPlayIdQuery[Dynasty_Religion](s.Story.PlayID))

	if !udrr.Ok {
		return result, udrr
	}

	result.Dynasty_Religions = cdrr

	usdr, csdr := nebulagolang.CompareAndUpdateNebulaEntityBySliceAndQuery(SPACE, s.Story_Dynasties, getPlayIdQuery[Story_Dynasty](s.Story.PlayID))

	if !usdr.Ok {
		return result, usdr
	}

	result.Story_Dynasties = csdr

	upeopler, cpeopler := nebulagolang.CompareAndUpdateNebulaEntityBySliceAndQuery(SPACE, s.People, getPlayIdQuery[People](s.Story.PlayID))

	if !upeopler.Ok {
		return result, upeopler
	}

	result.People = cpeopler

	upeople_cr, cpeople_cr := nebulagolang.CompareAndUpdateNebulaEntityBySliceAndQuery(SPACE, s.People_Cultures, getPlayIdQuery[People_Culture](s.Story.PlayID))

	if !upeople_cr.Ok {
		return result, upeople_cr
	}

	result.People_Cultures = cpeople_cr

	upeople_gfxcr, cpeople_gfxcr := nebulagolang.CompareAndUpdateNebulaEntityBySliceAndQuery(SPACE, s.People_GFXCultures, getPlayIdQuery[People_GFXCulture](s.Story.PlayID))

	if !upeople_gfxcr.Ok {
		return result, upeople_gfxcr
	}

	result.People_GFXCultures = cpeople_gfxcr

	upeople_rr, cpeople_rr := nebulagolang.CompareAndUpdateNebulaEntityBySliceAndQuery(SPACE, s.People_Religions, getPlayIdQuery[People_Religion](s.Story.PlayID))

	if !upeople_rr.Ok {
		return result, upeople_rr
	}

	result.People_Religions = cpeople_rr

	upeople_srr, cpeople_srr := nebulagolang.CompareAndUpdateNebulaEntityBySliceAndQuery(SPACE, s.People_SecretReligions, getPlayIdQuery[People_SecretReligion](s.Story.PlayID))

	if !upeople_srr.Ok {
		return result, upeople_srr
	}

	result.People_SecretReligions = cpeople_srr

	u_story_people_r, c_story_people_r := nebulagolang.CompareAndUpdateNebulaEntityBySliceAndQuery(SPACE, s.Story_People, getPlayIdQuery[Story_People](s.Story.PlayID))

	if !u_story_people_r.Ok {
		return result, u_story_people_r
	}

	result.Story_People = c_story_people_r

	upeople_tr, cpeople_tr := nebulagolang.CompareAndUpdateNebulaEntityBySliceAndQuery(SPACE, s.People_Traits, getPlayIdQuery[People_Trait](s.Story.PlayID))

	if !upeople_tr.Ok {
		return result, upeople_tr
	}

	result.People_Traits = cpeople_tr

	upeople_mr, cpeople_mr := nebulagolang.CompareAndUpdateNebulaEntityBySliceAndQuery(SPACE, s.People_Modifiers, getPlayIdQuery[People_Modifier](s.Story.PlayID))

	if !upeople_mr.Ok {
		return result, upeople_mr
	}

	result.People_Modifiers = cpeople_mr

	upeople_ctr, cpeople_ctr := nebulagolang.CompareAndUpdateNebulaEntityBySliceAndQuery(SPACE, s.People_ClaimTitles, getPlayIdQuery[People_ClaimTitle](s.Story.PlayID))

	if !upeople_ctr.Ok {
		return result, upeople_ctr
	}

	result.People_ClaimTitles = cpeople_ctr

	upeople_dr, cpeople_dr := nebulagolang.CompareAndUpdateNebulaEntityBySliceAndQuery(SPACE, s.People_Dynasties, getPlayIdQuery[People_Dynasty](s.Story.PlayID))

	if !upeople_dr.Ok {
		return result, upeople_dr
	}

	result.People_Dynasties = cpeople_dr

	upeople_fr, cpeople_fr := nebulagolang.CompareAndUpdateNebulaEntityBySliceAndQuery(SPACE, s.People_Families, getPlayIdQuery[People_FamilyPeople](s.Story.PlayID))

	if !upeople_fr.Ok {
		return result, upeople_fr
	}

	result.People_Families = cpeople_fr

	upeople_hr, cpeople_hr := nebulagolang.CompareAndUpdateNebulaEntityBySliceAndQuery(SPACE, s.People_Hosts, getPlayIdQuery[People_HostPeople](s.Story.PlayID))

	if !upeople_hr.Ok {
		return result, upeople_hr
	}

	result.People_Hosts = cpeople_hr

	upeople_er, cpeople_er := nebulagolang.CompareAndUpdateNebulaEntityBySliceAndQuery(SPACE, s.People_Empires, getPlayIdQuery[People_EmpirePeople](s.Story.PlayID))

	if !upeople_er.Ok {
		return result, upeople_er
	}

	result.People_Empires = cpeople_er

	upeople_kr, cpeople_kr := nebulagolang.CompareAndUpdateNebulaEntityBySliceAndQuery(SPACE, s.People_Killers, getPlayIdQuery[People_KillPeople](s.Story.PlayID))

	if !upeople_kr.Ok {
		return result, upeople_kr
	}

	result.People_Killers = cpeople_kr

	upeople_relarte_r, cpeople_r_relate_r := nebulagolang.CompareAndUpdateNebulaEntityBySliceAndQuery(SPACE, s.People_Relates, getPlayIdQuery[People_RelatePeople](s.Story.PlayID))

	if !upeople_relarte_r.Ok {
		return result, upeople_relarte_r
	}

	result.People_Relates = cpeople_r_relate_r

	upeople_lr, cpeople_lr := nebulagolang.CompareAndUpdateNebulaEntityBySliceAndQuery(SPACE, s.People_Lovers, getPlayIdQuery[People_LoverPeople](s.Story.PlayID))

	if !upeople_lr.Ok {
		return result, upeople_lr
	}

	result.People_Lovers = cpeople_lr

	upeople_gr, cpeople_gr := nebulagolang.CompareAndUpdateNebulaEntityBySliceAndQuery(SPACE, s.People_Guardians, getPlayIdQuery[People_GuardianPeople](s.Story.PlayID))

	if !upeople_gr.Ok {
		return result, upeople_gr
	}

	result.People_Guardians = cpeople_gr

	upeople_ar, cpeople_ar := nebulagolang.CompareAndUpdateNebulaEntityBySliceAndQuery(SPACE, s.People_Ambtions, getPlayIdQuery[People_Ambition](s.Story.PlayID))

	if !upeople_ar.Ok {
		return result, upeople_ar
	}

	result.People_Ambtions = cpeople_ar

	u_people_focuses_r, c_people_focuses_r := nebulagolang.CompareAndUpdateNebulaEntityBySliceAndQuery(SPACE, s.People_Focuses, getPlayIdQuery[People_Focus](s.Story.PlayID))

	if !u_people_focuses_r.Ok {
		return result, u_people_focuses_r
	}

	result.People_Focuses = c_people_focuses_r

	return result, ustr
}

func BuildStory(path string, savePath string) {
	story, result := LoadAndUpdateStory(path, savePath)

	if !result.Ok {
		fmt.Println(result.Err)
	} else {
		fmt.Println("Story added:", len(story.Story.Added))
		fmt.Println("Story updated:", len(story.Story.Updated))
		fmt.Println("Story deleted:", len(story.Story.Deleted))
		fmt.Println("Story kept:", len(story.Story.Kept))
		fmt.Println("Title added:", len(story.Titles.Added))
		fmt.Println("Title updated:", len(story.Titles.Updated))
		fmt.Println("Title deleted:", len(story.Titles.Deleted))
		fmt.Println("Title kept:", len(story.Titles.Kept))
		fmt.Println("Title_BaseTitle added:", len(story.Title_BaseTitles.Added))
		fmt.Println("Title_BaseTitle updated:", len(story.Title_BaseTitles.Updated))
		fmt.Println("Title_BaseTitle deleted:", len(story.Title_BaseTitles.Deleted))
		fmt.Println("Title_BaseTitle kept:", len(story.Title_BaseTitles.Kept))
		fmt.Println("Title_LiegeTitle added:", len(story.Title_LiegeTitles.Added))
		fmt.Println("Title_LiegeTitle updated:", len(story.Title_LiegeTitles.Updated))
		fmt.Println("Title_LiegeTitle deleted:", len(story.Title_LiegeTitles.Deleted))
		fmt.Println("Title_LiegeTitle kept:", len(story.Title_LiegeTitles.Kept))
		fmt.Println("Title_DejureLiegeTitle added:", len(story.Title_DejureLiegeTitles.Added))
		fmt.Println("Title_DejureLiegeTitle updated:", len(story.Title_DejureLiegeTitles.Updated))
		fmt.Println("Title_DejureLiegeTitle deleted:", len(story.Title_DejureLiegeTitles.Deleted))
		fmt.Println("Title_DejureLiegeTitle kept:", len(story.Title_DejureLiegeTitles.Kept))
		fmt.Println("Title_AssimilatingLiegeTitle added:", len(story.Title_AssimilatingLiegeTitles.Added))
		fmt.Println("Title_AssimilatingLiegeTitle updated:", len(story.Title_AssimilatingLiegeTitles.Updated))
		fmt.Println("Title_AssimilatingLiegeTitle deleted:", len(story.Title_AssimilatingLiegeTitles.Deleted))
		fmt.Println("Title_AssimilatingLiegeTitle kept:", len(story.Title_AssimilatingLiegeTitles.Kept))
		fmt.Println("Title_Dynasty added:", len(story.Title_Dynasties.Added))
		fmt.Println("Title_Dynasty updated:", len(story.Title_Dynasties.Updated))
		fmt.Println("Title_Dynasty deleted:", len(story.Title_Dynasties.Deleted))
		fmt.Println("Title_Dynasty kept:", len(story.Title_Dynasties.Kept))
		fmt.Println("Title_People added:", len(story.Title_People.Added))
		fmt.Println("Title_People updated:", len(story.Title_People.Updated))
		fmt.Println("Title_People deleted:", len(story.Title_People.Deleted))
		fmt.Println("Title_People kept:", len(story.Title_People.Kept))
		fmt.Println("Story_Title added:", len(story.Story_Titles.Added))
		fmt.Println("Story_Title updated:", len(story.Story_Titles.Updated))
		fmt.Println("Story_Title deleted:", len(story.Story_Titles.Deleted))
		fmt.Println("Story_Title kept:", len(story.Story_Titles.Kept))
		fmt.Println("Province added:", len(story.Provinces.Added))
		fmt.Println("Province updated:", len(story.Provinces.Updated))
		fmt.Println("Province deleted:", len(story.Provinces.Deleted))
		fmt.Println("Province kept:", len(story.Provinces.Kept))
		fmt.Println("Story_Province added:", len(story.Story_Provinces.Added))
		fmt.Println("Story_Province updated:", len(story.Story_Provinces.Updated))
		fmt.Println("Story_Province deleted:", len(story.Story_Provinces.Deleted))
		fmt.Println("Story_Province kept:", len(story.Story_Provinces.Kept))
		fmt.Println("Province_Modifier added:", len(story.Province_Modifiers.Added))
		fmt.Println("Province_Modifier updated:", len(story.Province_Modifiers.Updated))
		fmt.Println("Province_Modifier deleted:", len(story.Province_Modifiers.Deleted))
		fmt.Println("Province_Modifier kept:", len(story.Province_Modifiers.Kept))
		fmt.Println("Province_Culture added:", len(story.Province_Cultures.Added))
		fmt.Println("Province_Culture updated:", len(story.Province_Cultures.Updated))
		fmt.Println("Province_Culture deleted:", len(story.Province_Cultures.Deleted))
		fmt.Println("Province_Culture kept:", len(story.Province_Cultures.Kept))
		fmt.Println("Province_Religion added:", len(story.Province_Religions.Added))
		fmt.Println("Province_Religion updated:", len(story.Province_Religions.Updated))
		fmt.Println("Province_Religion deleted:", len(story.Province_Religions.Deleted))
		fmt.Println("Province_Religion kept:", len(story.Province_Religions.Kept))
		fmt.Println("Province_Title added:", len(story.Province_Titles.Added))
		fmt.Println("Province_Title updated:", len(story.Province_Titles.Updated))
		fmt.Println("Province_Title deleted:", len(story.Province_Titles.Deleted))
		fmt.Println("Province_Title kept:", len(story.Province_Titles.Kept))
		fmt.Println("Baron added:", len(story.Barons.Added))
		fmt.Println("Baron updated:", len(story.Barons.Updated))
		fmt.Println("Baron deleted:", len(story.Barons.Deleted))
		fmt.Println("Baron kept:", len(story.Barons.Kept))
		fmt.Println("Province_Baron added:", len(story.Province_Barons.Added))
		fmt.Println("Province_Baron updated:", len(story.Province_Barons.Updated))
		fmt.Println("Province_Baron deleted:", len(story.Province_Barons.Deleted))
		fmt.Println("Province_Baron kept:", len(story.Province_Barons.Kept))
		fmt.Println("Baron_Building added:", len(story.Baron_Buildings.Added))
		fmt.Println("Baron_Building updated:", len(story.Baron_Buildings.Updated))
		fmt.Println("Baron_Building deleted:", len(story.Baron_Buildings.Deleted))
		fmt.Println("Baron_Building kept:", len(story.Baron_Buildings.Kept))
		fmt.Println("Baron_Title added:", len(story.Baron_Titles.Added))
		fmt.Println("Baron_Title updated:", len(story.Baron_Titles.Updated))
		fmt.Println("Baron_Title deleted:", len(story.Baron_Titles.Deleted))
		fmt.Println("Baron_Title kept:", len(story.Baron_Titles.Kept))
		fmt.Println("Story_Baron added:", len(story.Story_Barons.Added))
		fmt.Println("Story_Baron updated:", len(story.Story_Barons.Updated))
		fmt.Println("Story_Baron deleted:", len(story.Story_Barons.Deleted))
		fmt.Println("Story_Baron kept:", len(story.Story_Barons.Kept))
		fmt.Println("Dynasty added:", len(story.Dynasties.Added))
		fmt.Println("Dynasty updated:", len(story.Dynasties.Updated))
		fmt.Println("Dynasty deleted:", len(story.Dynasties.Deleted))
		fmt.Println("Dynasty kept:", len(story.Dynasties.Kept))
		fmt.Println("Dynasty_Culture added:", len(story.Dynasty_Cultures.Added))
		fmt.Println("Dynasty_Culture updated:", len(story.Dynasty_Cultures.Updated))
		fmt.Println("Dynasty_Culture deleted:", len(story.Dynasty_Cultures.Deleted))
		fmt.Println("Dynasty_Culture kept:", len(story.Dynasty_Cultures.Kept))
		fmt.Println("Dynasty_Religion added:", len(story.Dynasty_Religions.Added))
		fmt.Println("Dynasty_Religion updated:", len(story.Dynasty_Religions.Updated))
		fmt.Println("Dynasty_Religion deleted:", len(story.Dynasty_Religions.Deleted))
		fmt.Println("Dynasty_Religion kept:", len(story.Dynasty_Religions.Kept))
		fmt.Println("Story_Dynasty added:", len(story.Story_Dynasties.Added))
		fmt.Println("Story_Dynasty updated:", len(story.Story_Dynasties.Updated))
		fmt.Println("Story_Dynasty deleted:", len(story.Story_Dynasties.Deleted))
		fmt.Println("Story_Dynasty kept:", len(story.Story_Dynasties.Kept))
		fmt.Println("People added:", len(story.People.Added))
		fmt.Println("People updated:", len(story.People.Updated))
		fmt.Println("People deleted:", len(story.People.Deleted))
		fmt.Println("People kept:", len(story.People.Kept))
		fmt.Println("People_Culture added:", len(story.People_Cultures.Added))
		fmt.Println("People_Culture updated:", len(story.People_Cultures.Updated))
		fmt.Println("People_Culture deleted:", len(story.People_Cultures.Deleted))
		fmt.Println("People_Culture kept:", len(story.People_Cultures.Kept))
		fmt.Println("People_GFXCulture added:", len(story.People_GFXCultures.Added))
		fmt.Println("People_GFXCulture updated:", len(story.People_GFXCultures.Updated))
		fmt.Println("People_GFXCulture deleted:", len(story.People_GFXCultures.Deleted))
		fmt.Println("People_GFXCulture kept:", len(story.People_GFXCultures.Kept))
		fmt.Println("People_Religion added:", len(story.People_Religions.Added))
		fmt.Println("People_Religion updated:", len(story.People_Religions.Updated))
		fmt.Println("People_Religion deleted:", len(story.People_Religions.Deleted))
		fmt.Println("People_Religion kept:", len(story.People_Religions.Kept))
		fmt.Println("People_SecretReligion added:", len(story.People_SecretReligions.Added))
		fmt.Println("People_SecretReligion updated:", len(story.People_SecretReligions.Updated))
		fmt.Println("People_SecretReligion deleted:", len(story.People_SecretReligions.Deleted))
		fmt.Println("People_SecretReligion kept:", len(story.People_SecretReligions.Kept))
		fmt.Println("Story_People added:", len(story.Story_People.Added))
		fmt.Println("Story_People updated:", len(story.Story_People.Updated))
		fmt.Println("Story_People deleted:", len(story.Story_People.Deleted))
		fmt.Println("Story_People kept:", len(story.Story_People.Kept))
		fmt.Println("People_Trait added:", len(story.People_Traits.Added))
		fmt.Println("People_Trait updated:", len(story.People_Traits.Updated))
		fmt.Println("People_Trait deleted:", len(story.People_Traits.Deleted))
		fmt.Println("People_Trait kept:", len(story.People_Traits.Kept))
		fmt.Println("People_Modifier added:", len(story.People_Modifiers.Added))
		fmt.Println("People_Modifier updated:", len(story.People_Modifiers.Updated))
		fmt.Println("People_Modifier deleted:", len(story.People_Modifiers.Deleted))
		fmt.Println("People_Modifier kept:", len(story.People_Modifiers.Kept))
		fmt.Println("People_ClaimTitle added:", len(story.People_ClaimTitles.Added))
		fmt.Println("People_ClaimTitle updated:", len(story.People_ClaimTitles.Updated))
		fmt.Println("People_ClaimTitle deleted:", len(story.People_ClaimTitles.Deleted))
		fmt.Println("People_ClaimTitle kept:", len(story.People_ClaimTitles.Kept))
		fmt.Println("People_Dynasty added:", len(story.People_Dynasties.Added))
		fmt.Println("People_Dynasty updated:", len(story.People_Dynasties.Updated))
		fmt.Println("People_Dynasty deleted:", len(story.People_Dynasties.Deleted))
		fmt.Println("People_Dynasty kept:", len(story.People_Dynasties.Kept))
		fmt.Println("People_FamilyPeople added:", len(story.People_Families.Added))
		fmt.Println("People_FamilyPeople updated:", len(story.People_Families.Updated))
		fmt.Println("People_FamilyPeople deleted:", len(story.People_Families.Deleted))
		fmt.Println("People_FamilyPeople kept:", len(story.People_Families.Kept))
		fmt.Println("People_HostPeople added:", len(story.People_Hosts.Added))
		fmt.Println("People_HostPeople updated:", len(story.People_Hosts.Updated))
		fmt.Println("People_HostPeople deleted:", len(story.People_Hosts.Deleted))
		fmt.Println("People_HostPeople kept:", len(story.People_Hosts.Kept))
		fmt.Println("People_EmpirePeople added:", len(story.People_Empires.Added))
		fmt.Println("People_EmpirePeople updated:", len(story.People_Empires.Updated))
		fmt.Println("People_EmpirePeople deleted:", len(story.People_Empires.Deleted))
		fmt.Println("People_EmpirePeople kept:", len(story.People_Empires.Kept))
		fmt.Println("People_KillPeople added:", len(story.People_Killers.Added))
		fmt.Println("People_KillPeople updated:", len(story.People_Killers.Updated))
		fmt.Println("People_KillPeople deleted:", len(story.People_Killers.Deleted))
		fmt.Println("People_KillPeople kept:", len(story.People_Killers.Kept))
		fmt.Println("People_RelatePeople added:", len(story.People_Relates.Added))
		fmt.Println("People_RelatePeople updated:", len(story.People_Relates.Updated))
		fmt.Println("People_RelatePeople deleted:", len(story.People_Relates.Deleted))
		fmt.Println("People_RelatePeople kept:", len(story.People_Relates.Kept))
		fmt.Println("People_LoverPeople added:", len(story.People_Lovers.Added))
		fmt.Println("People_LoverPeople updated:", len(story.People_Lovers.Updated))
		fmt.Println("People_LoverPeople deleted:", len(story.People_Lovers.Deleted))
		fmt.Println("People_LoverPeople kept:", len(story.People_Lovers.Kept))
		fmt.Println("People_GuardianPeople added:", len(story.People_Guardians.Added))
		fmt.Println("People_GuardianPeople updated:", len(story.People_Guardians.Updated))
		fmt.Println("People_GuardianPeople deleted:", len(story.People_Guardians.Deleted))
		fmt.Println("People_GuardianPeople kept:", len(story.People_Guardians.Kept))
		fmt.Println("People_Ambition added:", len(story.People_Ambtions.Added))
		fmt.Println("People_Ambition updated:", len(story.People_Ambtions.Updated))
		fmt.Println("People_Ambition deleted:", len(story.People_Ambtions.Deleted))
		fmt.Println("People_Ambition kept:", len(story.People_Ambtions.Kept))
		fmt.Println("People_Focus added:", len(story.People_Focuses.Added))
		fmt.Println("People_Focus updated:", len(story.People_Focuses.Updated))
		fmt.Println("People_Focus deleted:", len(story.People_Focuses.Deleted))
		fmt.Println("People_Focus kept:", len(story.People_Focuses.Kept))
	}
}

func DeleteStoryData(playId int) *nebulagolang.Result {
	r := DeleteAllStory_TitlesByPlayId(SPACE, playId)

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
