package main

import (
	"fmt"
	"github.com/thalesfu/ck2nebula"
	"github.com/thalesfu/nebulagolang/utils"
)

const ck2Folder = "R:\\Thales\\Game\\SteamLibrary\\steamapps\\common\\Crusader Kings II"
const saveFile = "T:\\OneDrive\\fu.thales@live.com\\OneDrive\\MyDocument\\Paradox Interactive\\Crusader Kings II\\save games\\酒泉771_02_14dd.ck2"

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

	ck2nebula.GenerateTitles(ck2Folder, saveFile)

	//_, _, tbts, tdlts, talts := ck2nebula.GenerateTitles(ck2Folder, saveFile)
	//
	//result := ck2nebula.InsertTitles(ck2nebula.SPACE, ts...)
	//
	//if !result.Ok {
	//	fmt.Println(result.Err.Error())
	//} else {
	//	fmt.Println("Insert titles success")
	//}
	//
	//result = ck2nebula.InsertTitle_LiegeTitles(ck2nebula.SPACE, tlts...)
	//
	//if !result.Ok {
	//	fmt.Println(result.Err.Error())
	//} else {
	//	fmt.Println("Insert titles -> liege tiles success")
	//}

	//result := ck2nebula.InsertTitle_BaseTitles(ck2nebula.SPACE, tbts...)
	//
	//if !result.Ok {
	//	fmt.Println(result.Err.Error())
	//} else {
	//	fmt.Println("Insert titles -> base tiles success")
	//}
	//
	//result = ck2nebula.InsertTitle_DejureLiegeTitles(ck2nebula.SPACE, tdlts...)
	//
	//if !result.Ok {
	//	fmt.Println(result.Err.Error())
	//} else {
	//	fmt.Println("Insert titles -> dejure liege tiles success")
	//}
	//
	//result = ck2nebula.InsertTitle_AssimilatingLiegeTitles(ck2nebula.SPACE, talts...)
	//
	//if !result.Ok {
	//	fmt.Println(result.Err.Error())
	//} else {
	//	fmt.Println("Insert titles -> assimilating dejure liege tiles success")
	//}

	//resultT := ck2nebula.DeleteAllTitle(ck2nebula.SPACE)

	resultT := ck2nebula.GetAllTitle_AssimilatingLiegeTitles(ck2nebula.SPACE)

	if !resultT.Ok {
		fmt.Println(resultT.Err)
	} else {
		fmt.Println(utils.MarshalJSON(resultT.Data))
		for _, s := range resultT.Commands {
			fmt.Println(s)
		}
	}

}
