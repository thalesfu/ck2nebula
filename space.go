package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
	"log"
)

var SPACE *nebulagolang.Space

func init() {
	db, ok := nebulagolang.LoadDB()

	if !ok {
		log.Fatal("Fail to load database")
	}

	SPACE = db.Use("ck2")
}
