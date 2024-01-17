package main

import (
	"fmt"
	"github.com/thalesfu/ck2nebula"
	"github.com/thalesfu/nebulagolang/utils"
)

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

	resultT := ck2nebula.GetAllCultureGroup_Cultures(ck2nebula.SPACE)

	if !resultT.Ok {
		fmt.Println(resultT.Err)
	} else {
		fmt.Println(utils.MarshalJSON(resultT.Data))
		for _, s := range resultT.Commands {
			fmt.Println(s)
		}
	}

}
