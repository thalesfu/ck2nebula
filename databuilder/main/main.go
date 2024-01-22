package main

import (
	"fmt"
	"github.com/thalesfu/ck2nebula"
	"github.com/thalesfu/paradoxtools/utils"
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

	//
	//resultT = ck2nebula.DeleteAllTitle_BaseTitles(ck2nebula.SPACE)
	//resultT = ck2nebula.DeleteAllTitle_LiegeTitles(ck2nebula.SPACE)
	//resultT = ck2nebula.DeleteAllTitle_DejureLiegeTitles(ck2nebula.SPACE)
	//resultT = ck2nebula.DeleteAllTitle_AssimilatingLiegeTitles(ck2nebula.SPACE)
	//sid, _ := strconv.Atoi(s.Story.PlayID)
	//resultT := ck2nebula.GetStoryByID(ck2nebula.SPACE, sid)

	//ck2nebula.BuildModifiers(ck2Folder)
	//ck2nebula.BuildBuildings(ck2Folder)
	//result := ck2nebula.GetAllBuildingCode(ck2nebula.SPACE)
	//if !result.Ok {
	//	fmt.Println(result.Err.Error())
	//} else {
	//	fmt.Println(utils.MarshalJSON(result.Data))
	//}

	//ck2nebula.DeleteStoryData(199229416)

	ck2nebula.BuildStory(ck2Folder, saveFile)

	result := ck2nebula.GetAllProvince_TitlesEidsByPlayId(ck2nebula.SPACE, 199229416)

	fmt.Println(utils.MarshalJSON(result.Data))

	for _, c := range result.Commands {
		fmt.Println(c)
	}

}
