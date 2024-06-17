package main

import (
	"fmt"
	"github.com/thalesfu/ck2nebula"
	"github.com/thalesfu/paradoxtools/utils"
)

const ck2Folder = "/Users/thalesfu/Windows/steam/steamapps/common/Crusader Kings II"
const saveFile = "/Users/thalesfu/Documents/Paradox Interactive/Crusader Kings II/save games/酒泉771_02_14dd.ck2"
const saveFile2 = "/Users/thalesfu/Documents/Paradox Interactive/Crusader Kings II/save games/閰掓硥796_12_01.ck2"
const saveFile3 = "/Users/thalesfu/Documents/Paradox Interactive/Crusader Kings II/save games/酒泉796_12_01.ck2"
const saveFile4 = "/Users/thalesfu/Documents/Paradox Interactive/Crusader Kings II/save games/尿普湖769_01_01.ck2"
const saveFile5 = "/Users/thalesfu/Documents/Paradox Interactive/Crusader Kings II/save games/岳普湖769_01_01.ck2"

func main() {
	cgs, cs, cg_c_s := ck2nebula.GenerateCultureData("/Users/thalesfu/Windows/steam/steamapps/common/Crusader Kings II")

	fmt.Println(utils.MarshalJSON(cgs))
	fmt.Println(utils.MarshalJSON(cs))
	fmt.Println(utils.MarshalJSON(cg_c_s))

	result := ck2nebula.InsertCultureGroups(ck2nebula.SPACE, cgs...)

	if !result.Ok {
		fmt.Println(result.Err.Error())
	} else {
		fmt.Println("Insert culture groups success")
	}

	result = ck2nebula.InsertCultures(ck2nebula.SPACE, cs...)

	if !result.Ok {
		fmt.Println(result.Err.Error())
	} else {
		fmt.Println("Insert cultures success")
	}

	result = ck2nebula.InsertCultureGroup_Cultures(ck2nebula.SPACE, cg_c_s...)

	//if !result.Ok {
	//	fmt.Println(result.Err.Error())
	//} else {
	//	fmt.Println("Insert culturegroup_culture success")
	//}
	//
	//rgs, rs, rg_r_s := ck2nebula.GenerateReligionData("/Users/thalesfu/Windows/steam/steamapps/common/Crusader Kings II")
	//
	//fmt.Println(utils.MarshalJSON(rgs))
	//fmt.Println(utils.MarshalJSON(rs))
	//fmt.Println(utils.MarshalJSON(rg_r_s))
	//
	//result = ck2nebula.InsertReligionGroups(ck2nebula.SPACE, rgs...)
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

	//ck2nebula.BuildModifiers(ck2Folder)
	//ck2nebula.BuildObjectives(ck2Folder)
	//ck2nebula.DeleteAllObjective(ck2nebula.SPACE)
	//ck2nebula.BuildBuildings(ck2Folder)
	//ck2nebula.BuildTraits(ck2Folder)
	//result := ck2nebula.GetTraitById(ck2nebula.SPACE, 56)
	//if !result.Ok {
	//	fmt.Println(result.Err.Error())
	//} else {
	//	fmt.Println(utils.MarshalJSON(result.Data))
	//}

	//ck2nebula.DeleteStoryData(916505602)
	//ck2nebula.DeleteAllModifierWithEdges(ck2nebula.SPACE)
	//result := ck2nebula.DeleteAllPeople_RelatePeoples(ck2nebula.SPACE)

	//ck2nebula.BuildStory(ck2Folder, saveFile)
	//result := ck2nebula.GetAllPeople_RelatePeoplesByPlayId(ck2nebula.SPACE, 199229416)
	//

	//ck2nebula.BuildModifiers(ck2Folder)

	//ck2nebula.BuildModifiers(ck2Folder)
	//if !result.Ok {
	//	fmt.Println(result.Err.Error())
	//} else {
	//	//fmt.Println(utils.MarshalJSON(result.Data))
	//
	//	for _, c := range result.Commands {
	//		fmt.Println(c)
	//	}
	//}
}
