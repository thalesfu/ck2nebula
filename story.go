package ck2nebula

import (
	"fmt"
	"github.com/thalesfu/nebulagolang"
	"github.com/thalesfu/paradoxtools/CK2/culture"
	"github.com/thalesfu/paradoxtools/CK2/religion"
	"github.com/thalesfu/paradoxtools/CK2/save"
)

type StoryDetail struct {
	Story                         *Story
	Titles                        []*Title
	Title_BaseTitles              []*Title_BaseTitle
	Title_LiegeTitles             []*Title_LiegeTitle
	Title_DejureLiegeTitles       []*Title_DejureLiegeTitle
	Title_AssimilatingLiegeTitles []*Title_AssimilatingLiegeTitle
	Story_Titles                  []*Story_Title
}

type StoryUpdateDetail struct {
	Story                         *nebulagolang.CompareResult[*Story]
	Titles                        *nebulagolang.CompareResult[*Title]
	Title_BaseTitles              *nebulagolang.CompareResult[*Title_BaseTitle]
	Title_LiegeTitles             *nebulagolang.CompareResult[*Title_LiegeTitle]
	Title_DejureLiegeTitles       *nebulagolang.CompareResult[*Title_DejureLiegeTitle]
	Title_AssimilatingLiegeTitles *nebulagolang.CompareResult[*Title_AssimilatingLiegeTitle]
	Story_Titles                  *nebulagolang.CompareResult[*Story_Title]
}

func GetStory(path string, savePath string) *StoryDetail {
	s, ok := save.LoadSave(path, savePath)

	if !ok {
		return nil
	}

	detail := GenerateStoryDetails(s)

	cgs := culture.LoadAllCultures(path)

	for _, cg := range cgs {
		if c, ok := cg.Cultures[detail.Story.Culture]; ok {
			detail.Story.Culture = c.Name
		}
	}

	rgs := religion.LoadAllReligions(path)

	for _, rg := range rgs {
		if r, ok := rg.Religions[detail.Story.Religion]; ok {
			detail.Story.Religion = r.Name
		}
	}

	return detail
}

func GenerateStoryDetails(file *save.SaveFile) *StoryDetail {

	titles, baseTitles, liegeTitles, dejureLiegeTitles, assimilatingLiegeTitles := GenerateTitles(file.Titles)

	sd := &StoryDetail{
		Story:                         NewStoryByData(file),
		Titles:                        titles,
		Title_BaseTitles:              baseTitles,
		Title_LiegeTitles:             liegeTitles,
		Title_DejureLiegeTitles:       dejureLiegeTitles,
		Title_AssimilatingLiegeTitles: assimilatingLiegeTitles,
		Story_Titles:                  make([]*Story_Title, 0),
	}

	for _, title := range titles {
		sd.Story_Titles = append(sd.Story_Titles, NewStory_Title(sd.Story, title))
	}

	return sd
}

func getPlayIdQuery[T interface{}](playId int) string {
	itemName := nebulagolang.GetTagName[T]()

	if itemName == "" {
		itemName = nebulagolang.GetEdgeName[T]()
	}

	return fmt.Sprintf("%s.play_id==%d", itemName, playId)
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

	ustr, cstr := nebulagolang.CompareAndUpdateNebulaEntityBySliceAndQuery(SPACE, s.Story_Titles, getPlayIdQuery[Story_Title](s.Story.PlayID))

	if !ustr.Ok {
		return result, ustr
	}

	result.Story_Titles = cstr

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
		fmt.Println("Story_Title added:", len(story.Story_Titles.Added))
		fmt.Println("Story_Title updated:", len(story.Story_Titles.Updated))
		fmt.Println("Story_Title deleted:", len(story.Story_Titles.Deleted))
		fmt.Println("Story_Title kept:", len(story.Story_Titles.Kept))

	}
}

func DeleteStoryData(playId int) *nebulagolang.Result {
	r := DeleteAllStory_TitlesByPlayId(SPACE, playId)

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

	r = DeleteAllTitleByPlayId(SPACE, playId)

	if !r.Ok {
		return r
	}

	r = DeleteStories(SPACE, playId)

	return r
}
