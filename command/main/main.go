package main

import (
	"fmt"
	"github.com/thalesfu/ck2nebula"
	"github.com/thalesfu/paradoxtools/utils"
)

const playid = 916505602

func main() {
	resultT := ck2nebula.GetCourtierByPlayIdAndPeopleId(ck2nebula.SPACE, playid, 2609830)

	if !resultT.Ok {
		fmt.Println(resultT.Err.Error())
	} else {
		fmt.Println(utils.MarshalJSON(resultT.Data))

		for _, p := range resultT.Data {
			fmt.Printf("%d\t%s\t%s\n", p.ID, p.DynastyName, p.Name)
		}

		for _, c := range resultT.Commands {
			fmt.Println(c)
		}
	}
}
