package dullback_test

import (
	"testing"

	"github.com/taylormonacelli/dullback/dullback"
)

func Test(t *testing.T) {
	s := dullback.NewStuff(dullback.NewBuiltinLogger())
	s.DoStuff()
}
