package device_test

import (
	"flag"
	"testing"

	"github.com/freeconf/manage/device"
	"github.com/freeconf/manage/testdata"
	"github.com/freeconf/yang/c2"
	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/nodes"
)

var update = flag.Bool("update", false, "update golden test files")

func TestYangLibNode(t *testing.T) {
	d, _ := testdata.BirdDevice(`{"bird":[{
		"name" : "robin"
	},{
		"name" : "blue jay"
	}]}`)
	moduleNameAsAddress := func(m *meta.Module) string {
		return m.Ident()
	}
	if err := d.Add("ietf-yang-library", device.LocalDeviceYangLibNode(moduleNameAsAddress, d)); err != nil {
		t.Error(err)
	}
	b, err := d.Browser("ietf-yang-library")
	if err != nil {
		t.Error(err)
		return
	}
	if b == nil {
		t.Error("no browser")
		return
	}
	actual, err := nodes.WritePrettyJSON(b.Root())
	if err != nil {
		t.Error(err)
	}
	c2.Gold(t, *update, []byte(actual), "gold/yang_lib.json")
}
