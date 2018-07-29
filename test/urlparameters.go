package test

import "github.com/matrosov-nikita/test2doc/doc/parse"

func RegisterURLVarExtractor(fn parse.URLVarExtractor) {
	parse.SetURLVarExtractor(&fn)
}
