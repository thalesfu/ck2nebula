package main

import "github.com/thalesfu/ck2nebula"

const ck2Folder = "/Users/thalesfu/Windows/steam/steamapps/common/Crusader Kings II"
const saveFile = "/Users/thalesfu/Documents/Paradox Interactive/Crusader Kings II/save games/酒泉771_02_14dd.ck2"
const saveFile2 = "/Users/thalesfu/Documents/Paradox Interactive/Crusader Kings II/save games/閰掓硥796_12_01.ck2"
const saveFile3 = "/Users/thalesfu/Documents/Paradox Interactive/Crusader Kings II/save games/酒泉796_12_01.ck2"
const saveFile4 = "/Users/thalesfu/Documents/Paradox Interactive/Crusader Kings II/save games/尿普湖769_01_01.ck2"
const saveFile5 = "/Users/thalesfu/Documents/Paradox Interactive/Crusader Kings II/save games/岳普湖769_01_01.ck2"

func main() {

	ck2nebula.BuildModifiers(ck2Folder)
	ck2nebula.BuildObjectives(ck2Folder)

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
