package main

import (
	"fmt"
	"github.com/thalesfu/ck2nebula"
)

const ck2Folder = "R:\\Thales\\Game\\SteamLibrary\\steamapps\\common\\Crusader Kings II"
const saveFile = "T:\\OneDrive\\fu.thales@live.com\\OneDrive\\MyDocument\\Paradox Interactive\\Crusader Kings II\\save games\\酒泉771_02_14dd.ck2"
const saveFile2 = "T:\\OneDrive\\fu.thales@live.com\\OneDrive\\MyDocument\\Paradox Interactive\\Crusader Kings II\\save games\\閰掓硥796_12_01.ck2"

func main() {
	//cgs, cs, cg_c_s := ck2nebula.GenerateCultureData("R:\\Thales\\Game\\SteamLibrary\\steamapps\\common\\Crusader Kings II")
	//
	//fmt.Println(utils.MarshalJSON(cgs))
	//fmt.Println(utils.MarshalJSON(cs))
	//fmt.Println(utils.MarshalJSON(cg_c_s))
	//
	//result := ck2nebula.InsertCultureGroups(ck2nebula.SPACE, cgs...)
	//
	//if !result.Ok {
	//	fmt.Println(result.Err.Error())
	//} else {
	//	fmt.Println("Insert culture groups success")
	//}
	//
	//result = ck2nebula.InsertCultures(ck2nebula.SPACE, cs...)
	//
	//if !result.Ok {
	//	fmt.Println(result.Err.Error())
	//} else {
	//	fmt.Println("Insert cultures success")
	//}
	//
	//result = ck2nebula.InsertCultureGroup_Cultures(ck2nebula.SPACE, cg_c_s...)
	//
	//if !result.Ok {
	//	fmt.Println(result.Err.Error())
	//} else {
	//	fmt.Println("Insert culturegroup_culture success")
	//}

	//rgs, rs, rg_r_s := ck2nebula.GenerateReligionData("R:\\Thales\\Game\\SteamLibrary\\steamapps\\common\\Crusader Kings II")
	//
	//fmt.Println(utils.MarshalJSON(rgs))
	//fmt.Println(utils.MarshalJSON(rs))
	//fmt.Println(utils.MarshalJSON(rg_r_s))
	//
	//result := ck2nebula.InsertReligionGroups(ck2nebula.SPACE, rgs...)
	//
	//if !result.Ok {
	//	fmt.Println(result.Err.Error())
	//} else {
	//	fmt.Println("Insert religion groups success")
	//}
	//
	//result = ck2nebula.InsertReligions(ck2nebula.SPACE, rs...)
	//
	//if !result.Ok {
	//	fmt.Println(result.Err.Error())
	//} else {
	//	fmt.Println("Insert religion success")
	//}
	//
	//result = ck2nebula.InsertReligionGroup_Religions(ck2nebula.SPACE, rg_r_s...)
	//
	//if !result.Ok {
	//	fmt.Println(result.Err.Error())
	//} else {
	//	fmt.Println("Insert religiongroup_religion success")
	//}

	//_, _, _, ldlts_, _ := ck2nebula.GetTitlesFromFile(ck2Folder, saveFile)
	//
	//r, cr := nebulagolang.CompareAndUpdateNebulaEntitySlice(ck2nebula.SPACE, ldlts_)
	//
	//if !r.Ok {
	//	fmt.Println(r.Err.Error())
	//	logWriter, deferFunc := utils.CreateLogWriter("compareandupdate")
	//	defer deferFunc()
	//	logWriter.WriteString(r.Err.Error())
	//	logWriter.Flush()
	//} else {
	//	fmt.Println(utils.MarshalJSON(cr.Added))
	//}

	//s := ck2nebula.GetStory(ck2Folder, saveFile)

	//result := ck2nebula.InsertTitles(ck2nebula.SPACE, s.Titles...)
	//
	//if !result.Ok {
	//	fmt.Println(result.Err.Error())
	//} else {
	//	fmt.Println("Insert titles success")
	//}
	//
	//result = ck2nebula.InsertTitle_BaseTitles(ck2nebula.SPACE, s.Title_BaseTitles...)
	//
	//if !result.Ok {
	//	fmt.Println(result.Err.Error())
	//} else {
	//	fmt.Println("Insert titles -> base tiles success")
	//}
	//
	//result = ck2nebula.InsertTitle_LiegeTitles(ck2nebula.SPACE, s.Title_LiegeTitles...)
	//
	//if !result.Ok {
	//	fmt.Println(result.Err.Error())
	//} else {
	//	fmt.Println("Insert titles -> liege tiles success")
	//}
	//
	//result = ck2nebula.InsertTitle_DejureLiegeTitles(ck2nebula.SPACE, s.Title_DejureLiegeTitles...)
	//
	//if !result.Ok {
	//	fmt.Println(result.Err.Error())
	//} else {
	//	fmt.Println("Insert titles -> dejure liege tiles success")
	//}
	//
	//result = ck2nebula.InsertTitle_AssimilatingLiegeTitles(ck2nebula.SPACE, s.Title_AssimilatingLiegeTitles...)
	//
	//if !result.Ok {
	//	fmt.Println(result.Err.Error())
	//} else {
	//	fmt.Println("Insert titles -> assimilating dejure liege tiles success")
	//}

	//result := s.Story.InsertToNebula(ck2nebula.SPACE)
	//
	//if !result.Ok {
	//	fmt.Println(result.Err.Error())
	//} else {
	//	fmt.Println("Insert story success")
	//}
	//
	//result = ck2nebula.InsertStory_Titles(ck2nebula.SPACE, s.Story_Titles...)
	//
	//if !result.Ok {
	//	fmt.Println(result.Err.Error())
	//} else {
	//	fmt.Println("Insert story -> title success")
	//}

	//result := ck2nebula.DeleteStoryData(199229416)
	//
	//if !result.Ok {
	//	fmt.Println(result.Err.Error())
	//} else {
	//	fmt.Println("Delete story data success")
	//}

	story, result := ck2nebula.LoadAndUpdateStory(ck2Folder, saveFile2)

	//
	//resultT = ck2nebula.DeleteAllTitle_BaseTitles(ck2nebula.SPACE)
	//resultT = ck2nebula.DeleteAllTitle_LiegeTitles(ck2nebula.SPACE)
	//resultT = ck2nebula.DeleteAllTitle_DejureLiegeTitles(ck2nebula.SPACE)
	//resultT = ck2nebula.DeleteAllTitle_AssimilatingLiegeTitles(ck2nebula.SPACE)
	//sid, _ := strconv.Atoi(s.Story.PlayID)
	//resultT := ck2nebula.GetStoryByID(ck2nebula.SPACE, sid)

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
