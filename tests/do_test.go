package tests

import (
	"kjexam/core"
	"os"
	"testing"

	resty "gopkg.in/resty.v1"
)

func TestdoOne(t testing.T) {
	os.Setenv("GOCACHE", "off ")
	resty.SetDebug(true)
	cc := core.Core{Client: resty.R()}

	cc.DoOneVideo(1636, 24370)
}
