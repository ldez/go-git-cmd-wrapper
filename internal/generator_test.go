package main

import (
	"reflect"
	"testing"
)

func Test_newMeta(t *testing.T) {
	testCases := []struct {
		name          string
		rawMethodName string
		rawArg        string
		cmd           string
		cmdType       string
		description   string
		arguments     string
		expectedMeta  cmdMeta
	}{
		{
			name:          "with all parameters",
			rawMethodName: "foo-bar",
			rawArg:        "fii-bir",
			cmd:           "--foo-bar",
			cmdType:       "TYPE",
			description:   "line1\nline2",
			arguments:     "arg",
			expectedMeta: cmdMeta{
				Method:     "FooBar",
				Argument:   "fiiBir",
				Cmd:        "--foo-bar",
				Type:       "TYPE",
				CmdComment: "arg",
				Comments:   []string{"line1", "line2"},
			},
		},
		{
			name:          "without argument",
			rawMethodName: "foo-bar",
			rawArg:        "",
			cmd:           "--foo-bar",
			cmdType:       "TYPE",
			description:   "line1\nline2",
			arguments:     "arg",
			expectedMeta: cmdMeta{
				Method:     "FooBar",
				Argument:   "",
				Cmd:        "--foo-bar",
				Type:       "TYPE",
				CmdComment: "arg",
				Comments:   []string{"line1", "line2"},
			},
		},
		{
			name:          "with same name for argument and method",
			rawMethodName: "foo-bar",
			rawArg:        "foo-bar",
			cmd:           "--foo-bar",
			cmdType:       "TYPE",
			description:   "line1\nline2",
			arguments:     "arg",
			expectedMeta: cmdMeta{
				Method:     "FooBar",
				Argument:   "value",
				Cmd:        "--foo-bar",
				Type:       "TYPE",
				CmdComment: "arg",
				Comments:   []string{"line1", "line2"},
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			meta := newMeta(test.rawMethodName, test.rawArg, test.cmd, test.cmdType, test.description, test.arguments)

			if !reflect.DeepEqual(meta, test.expectedMeta) {
				t.Fatalf("Got: %+v, expected: %+v.", meta, test.expectedMeta)
			}
		})
	}
}

func Test_toGoName_should(t *testing.T) {
	testCases := []struct {
		name         string
		rawName      string
		upperFirst   bool
		expectedName string
	}{
		{
			name:         "return empty name when raw name empty and upperFirst",
			rawName:      "",
			upperFirst:   true,
			expectedName: "",
		},
		{
			name:         "return empty name when raw name empty and not upperFirst",
			rawName:      "",
			upperFirst:   false,
			expectedName: "",
		},
		{
			name:         "return a UpperCamelCase name when raw name empty and upperFirst",
			rawName:      "foo-bar-fii-bir",
			upperFirst:   true,
			expectedName: "FooBarFiiBir",
		},
		{
			name:         "return a lowerCamelCase name when raw name empty and not upperFirst",
			rawName:      "foo-bar-fii-bir",
			upperFirst:   false,
			expectedName: "fooBarFiiBir",
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			name := toGoName(test.rawName, test.upperFirst)

			assertEquals(t, name, test.expectedName)
		})
	}
}

func assertEquals(t *testing.T, value, expectedValue string) {
	t.Helper()

	if value != expectedValue {
		t.Fatalf("Got: %s, expected: %s.", value, expectedValue)
	}
}
