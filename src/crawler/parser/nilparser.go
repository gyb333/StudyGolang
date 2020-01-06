package parser

import "crawler/engine"

type NilParser struct {}

func (NilParser) Parse([] byte) engine.ParseResult {
	return engine.ParseResult{}
}
