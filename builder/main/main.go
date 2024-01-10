package main

import (
	"fmt"
	"github.com/thalesfu/ck2nebula"
	"github.com/thalesfu/nebulagolang"
	"github.com/thalesfu/nebulagolang/basictype"
	"github.com/thalesfu/nebulagolang/build"
)

func main() {
	db, ok := nebulagolang.LoadDB()

	if !ok {
		return
	}

	_, ok, err := db.CreateSpace("ck2", basictype.FixedString(200), 20, 1)

	if !ok {
		fmt.Println(err)
	}

	space := db.Use("ck2")

	build.RebuildTagWithIndexes[ck2nebula.CultureGroup](space)
}
