package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddAttr(t *testing.T) {
	data := "<div dir=\"rtl\"><a href=\"111\">111</a></div>"
	expected := "<div id=\"__gowrapper__\"><div dir=\"rtl\"><a href=\"111\" target=\"_blank\">111</a></div></div>"
	editor := new(HtmlEditor)
	actual := editor.Init(data)

	if !assert.EqualValues(t, expected, actual) {
		t.Error("don't added target attr")
	}
}

func TestNotAddUnnecessaryAttr(t *testing.T) {
	data := "<div dir=\"rtl\"><a href=\"111\" target=\"_blank\">111</a></div>"
	expected := "<div id=\"__gowrapper__\"><div dir=\"rtl\"><a href=\"111\" target=\"_blank\">111</a></div></div>"
	editor := new(HtmlEditor)
	actual := editor.Init(data)

	if !assert.EqualValues(t, expected, actual) {
		t.Error("don't added target attr")
	}
}

func TestFewTagsAttr(t *testing.T) {
	data := "<div dir=\"rtl\"><a href=\"111\" target=\"_blank\">111</a> <a href=\"222\">222</a></div>"
	expected := "<div id=\"__gowrapper__\"><div dir=\"rtl\"><a href=\"111\" target=\"_blank\">111</a> <a href=\"222\" target=\"_blank\">222</a></div></div>"
	editor := new(HtmlEditor)
	actual := editor.Init(data)

	if !assert.EqualValues(t, expected, actual) {
		t.Error("don't added target attr")
	}
}
