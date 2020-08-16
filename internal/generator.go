package main

//go:generate go run generator.go

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strings"
	"text/template"
)

type jsonModel struct {
	CommandName string         `json:"command_name,omitempty"`
	Enabled     bool           `json:"enabled"`
	Options     []jsonCmdModel `json:"options"`
}

type jsonCmdModel struct {
	MethodName  string `json:"method_name,omitempty"`
	Argument    string `json:"argument"`
	Arguments   string `json:"arguments"`
	Description string `json:"description"`
}

type genCmdModel struct {
	Name      string
	ImportFMT bool
	Metas     []cmdMeta
}

type cmdMeta struct {
	Type       string
	Method     string
	Argument   string
	Cmd        string
	Comments   []string
	CmdComment string
}

// byMethodName sort method by name.
type byMethodName []cmdMeta

func (r byMethodName) Len() int           { return len(r) }
func (r byMethodName) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r byMethodName) Less(i, j int) bool { return r[i].Method < r[j].Method }

const (
	fileTemplate = `package {{ .Name }}

// CODE GENERATED AUTOMATICALLY
// THIS FILE MUST NOT BE EDITED BY HAND

{{if .ImportFMT }}import (
	"fmt"

	"github.com/ldez/go-git-cmd-wrapper/v2/types"
)
{{else}}import "github.com/ldez/go-git-cmd-wrapper/v2/types"
{{end -}}

{{range .Metas -}}
{{if eq .Type "SIMPLE"}}{{template "templateCmdSimple" .}}
{{else if eq .Type "EQUAL_NO_OPTIONAL" }}{{template "templateCmdEqualNoOptional" .}}
{{else if eq .Type "EQUAL_WITHOUT_NAME" }}{{template "templateCmdEqualNoOptional" .}}
{{else if eq .Type "EQUAL_OPTIONAL_WITHOUT_NAME" }}{{template "templateCmdEqualOptional" .}}
{{else if eq .Type "EQUAL_OPTIONAL_WITH_NAME" }}{{template "templateCmdEqualOptional" .}}
{{else if eq .Type "WITH_PARAMETER" }}{{template "templateCmdWithParameter" .}}
{{else if eq .Type "WITH_OPTIONAL_PARAMETER" }}{{template "templateCmdWithOptionalParameter" .}}
{{else}} BUG
{{end -}}
{{end -}}
`
	templateCmdSimple = `{{- range $index, $element := .Comments}}
// {{if eq $index 0 }}{{ $.Method }} {{end}}{{ $element }}{{end}}
// {{ .CmdComment }}
func {{ .Method }}(g *types.Cmd) {
	g.AddOptions("{{ .Cmd }}")
}`

	templateCmdEqualNoOptional = `{{- range $index, $element := .Comments}}
// {{if eq $index 0 }}{{ $.Method }} {{end}}{{ $element }}{{end}}
// {{ .CmdComment }}
func {{ .Method }}({{ .Argument }} string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("{{ .Cmd }}=%s", {{ .Argument }}))
	}
}`

	templateCmdEqualOptional = `{{- range $index, $element := .Comments}}
// {{if eq $index 0 }}{{ $.Method }} {{end}}{{ $element }}{{end}}
// {{ .CmdComment }}
func {{ .Method }}({{ .Argument }} string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		if len({{ .Argument }}) == 0 {
			g.AddOptions("{{ .Cmd }}")
		} else {
			g.AddOptions(fmt.Sprintf("{{ .Cmd }}=%s", {{ .Argument }}))
		}
	}
}`

	templateCmdWithParameter = `{{- range $index, $element := .Comments}}
// {{if eq $index 0 }}{{ $.Method }} {{end}}{{ $element }}{{end}}
// {{ .CmdComment }}
func {{ .Method }}({{ .Argument }} string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("{{ .Cmd }}")
		g.AddOptions({{ .Argument }})
	}
}`

	templateCmdWithOptionalParameter = `{{- range $index, $element := .Comments}}
// {{if eq $index 0 }}{{ $.Method }} {{end}}{{ $element }}{{end}}
// {{ .CmdComment }}
func {{ .Method }}({{ .Argument }} string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("{{ .Cmd }}")
		if len({{ .Argument }}) != 0 {
			g.AddOptions({{ .Argument }})
		}
	}
}`
)

var (
	// --quiet
	expCmdSimple = regexp.MustCompile(`^(-{1,2}([\w\d\-]+))$`)

	// --strategy=<strategy>
	expCmdEqualNoOptional = regexp.MustCompile(`^(-{1,2}([\w\d\-]+))=<([\w\d\- ]+)>$`)

	// --no-recurse-submodules[=yes|on-demand|no]
	expCmdEqualOptionalWithoutName = regexp.MustCompile(`^(-{1,2}([\w\d\-]+))\[=[\w\d\-()|]+]$`)

	// --recurse-submodules-default=[yes|on-demand]
	// --sign=(true|false|if-asked)
	expCmdEqualWithoutName = regexp.MustCompile(`^(-{1,2}([\w\d\-]+))=[\[(][\w\d\-|()]+[])]$`)

	// --log[=<n>]
	expCmdEqualOptionalWithName = regexp.MustCompile(`^(-{1,2}([\w\d\-]+))\[=<([\w\d\-)]+)>]$`)

	// --foo <bar>
	expCmdWithParameter = regexp.MustCompile(`^(-{1,2}([\w\d\-]+)) ?<([\w\d\-)]+)>$`)

	// --foo [bar]
	expCmdWithOptionalParameter = regexp.MustCompile(`^(-{1,2}([\w\d\-]+)) ?\[([\w\d\-)]+)]$`)
)

func main() {
	filePath := "descriptions.json"

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	var jsonModels []jsonModel
	err = json.Unmarshal(file, &jsonModels)
	if err != nil {
		log.Fatal(err)
	}

	for _, jsonModel := range jsonModels {
		if len(jsonModel.CommandName) != 0 && jsonModel.Enabled {
			cmdModel := newGenCmdModel(jsonModel)

			data, err := generateFileContent(cmdModel)
			if err != nil {
				log.Fatal(err)
			}

			genFilePath := fmt.Sprintf("../%[1]s/%[1]s_gen.go", jsonModel.CommandName)

			fmt.Println(genFilePath)

			// gofmt
			source, err := format.Source([]byte(data))
			if err != nil {
				log.Fatal(err)
			}

			err = ioutil.WriteFile(genFilePath, source, 0644)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func generateFileContent(model genCmdModel) (string, error) {
	base := template.New(model.Name)

	_, err := base.New("templateCmdSimple").Parse(templateCmdSimple)
	if err != nil {
		return "", err
	}
	_, err = base.New("templateCmdEqualNoOptional").Parse(templateCmdEqualNoOptional)
	if err != nil {
		return "", err
	}
	_, err = base.New("templateCmdEqualOptional").Parse(templateCmdEqualOptional)
	if err != nil {
		return "", err
	}
	_, err = base.New("templateCmdWithParameter").Parse(templateCmdWithParameter)
	if err != nil {
		return "", err
	}
	_, err = base.New("templateCmdWithOptionalParameter").Parse(templateCmdWithOptionalParameter)
	if err != nil {
		return "", err
	}

	tmpl := template.Must(base.Parse(fileTemplate))

	b := &bytes.Buffer{}
	err = tmpl.Execute(b, model)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}

func newGenCmdModel(jsonModel jsonModel) genCmdModel {
	return genCmdModel{
		Name:      jsonModel.CommandName,
		Metas:     jsonCmdModelToCmdMetas(jsonModel.Options),
		ImportFMT: hasImportFMT(jsonModel.Options),
	}
}

func hasImportFMT(jsonCmdModels []jsonCmdModel) bool {
	for _, jsonCmdModel := range jsonCmdModels {
		if expCmdEqualNoOptional.MatchString(jsonCmdModel.Argument) ||
			expCmdEqualOptionalWithoutName.MatchString(jsonCmdModel.Argument) ||
			expCmdEqualWithoutName.MatchString(jsonCmdModel.Argument) ||
			expCmdEqualOptionalWithName.MatchString(jsonCmdModel.Argument) {
			return true
		}
	}
	return false
}

func jsonCmdModelToCmdMetas(jsonCmdModels []jsonCmdModel) []cmdMeta {
	var metas []cmdMeta

	for _, jsonCmdModel := range jsonCmdModels {
		switch {
		case expCmdSimple.MatchString(jsonCmdModel.Argument):
			metas = append(metas, newMetaCmdSimple(jsonCmdModel))
		case expCmdEqualNoOptional.MatchString(jsonCmdModel.Argument):
			metas = append(metas, newMetaCmdEqualNoOptional(jsonCmdModel))
		case expCmdEqualOptionalWithoutName.MatchString(jsonCmdModel.Argument):
			metas = append(metas, newMetaCmdEqualOptionalWithoutName(jsonCmdModel))
		case expCmdEqualWithoutName.MatchString(jsonCmdModel.Argument):
			metas = append(metas, newMetaCmdEqualWithoutName(jsonCmdModel))
		case expCmdEqualOptionalWithName.MatchString(jsonCmdModel.Argument):
			metas = append(metas, newMetaCmdEqualOptionalWithName(jsonCmdModel))
		case expCmdWithParameter.MatchString(jsonCmdModel.Argument):
			metas = append(metas, newMetaCmdWithParameter(jsonCmdModel))
		case expCmdWithOptionalParameter.MatchString(jsonCmdModel.Argument):
			metas = append(metas, newMetaCmdWithOptionalParameter(jsonCmdModel))
		default:
			log.Println("fail", jsonCmdModel)
		}
	}

	sort.Sort(byMethodName(metas))

	return metas
}

func newMetaCmdWithOptionalParameter(jsonCmdModel jsonCmdModel) cmdMeta {
	return newMetaCmd(expCmdWithOptionalParameter, jsonCmdModel, func(subMatch []string) cmdMeta {
		return newMeta(
			methodName(subMatch[2], jsonCmdModel),
			subMatch[3],
			subMatch[1],
			"WITH_OPTIONAL_PARAMETER",
			jsonCmdModel.Description,
			jsonCmdModel.Arguments)
	})
}

func newMetaCmdSimple(jsonCmdModel jsonCmdModel) cmdMeta {
	return newMetaCmd(expCmdSimple, jsonCmdModel, func(subMatch []string) cmdMeta {
		return newMeta(
			methodName(subMatch[2], jsonCmdModel),
			"",
			subMatch[1],
			"SIMPLE",
			jsonCmdModel.Description,
			jsonCmdModel.Arguments)
	})
}

func newMetaCmdEqualNoOptional(jsonCmdModel jsonCmdModel) cmdMeta {
	return newMetaCmd(expCmdEqualNoOptional, jsonCmdModel, func(subMatch []string) cmdMeta {
		return newMeta(
			methodName(subMatch[2], jsonCmdModel),
			subMatch[3],
			subMatch[1],
			"EQUAL_NO_OPTIONAL",
			jsonCmdModel.Description,
			jsonCmdModel.Arguments)
	})
}

func newMetaCmdEqualOptionalWithoutName(jsonCmdModel jsonCmdModel) cmdMeta {
	return newMetaCmd(expCmdEqualOptionalWithoutName, jsonCmdModel, func(subMatch []string) cmdMeta {
		return newMeta(
			methodName(subMatch[2], jsonCmdModel),
			"value",
			subMatch[1],
			"EQUAL_OPTIONAL_WITHOUT_NAME",
			jsonCmdModel.Description,
			jsonCmdModel.Arguments)
	})
}

func newMetaCmdEqualWithoutName(jsonCmdModel jsonCmdModel) cmdMeta {
	return newMetaCmd(expCmdEqualWithoutName, jsonCmdModel, func(subMatch []string) cmdMeta {
		return newMeta(
			methodName(subMatch[2], jsonCmdModel),
			"value",
			subMatch[1],
			"EQUAL_WITHOUT_NAME",
			jsonCmdModel.Description,
			jsonCmdModel.Arguments)
	})
}

func newMetaCmdEqualOptionalWithName(jsonCmdModel jsonCmdModel) cmdMeta {
	return newMetaCmd(expCmdEqualOptionalWithName, jsonCmdModel, func(subMatch []string) cmdMeta {
		return newMeta(
			methodName(subMatch[2], jsonCmdModel),
			subMatch[3],
			subMatch[1],
			"EQUAL_OPTIONAL_WITH_NAME",
			jsonCmdModel.Description,
			jsonCmdModel.Arguments)
	})
}

func newMetaCmdWithParameter(jsonCmdModel jsonCmdModel) cmdMeta {
	return newMetaCmd(expCmdWithParameter, jsonCmdModel, func(subMatch []string) cmdMeta {
		return newMeta(
			methodName(subMatch[2], jsonCmdModel),
			subMatch[3],
			subMatch[1],
			"WITH_PARAMETER",
			jsonCmdModel.Description,
			jsonCmdModel.Arguments)
	})
}

func methodName(raw string, jsonCmdModel jsonCmdModel) string {
	if len(jsonCmdModel.MethodName) == 0 {
		return raw
	}
	return jsonCmdModel.MethodName
}

type metaBuilder func(subMatch []string) cmdMeta

func newMetaCmd(regexp *regexp.Regexp, jsonCmdModel jsonCmdModel, builder metaBuilder) cmdMeta {
	subMatch := regexp.FindStringSubmatch(jsonCmdModel.Argument)
	return builder(subMatch)
}

func newMeta(rawMethodName, rawArg, cmd, cmdType, description, arguments string) cmdMeta {
	method := toGoName(rawMethodName, true)

	var arg string
	if len(rawArg) != 0 {
		arg = toGoName(rawArg, false)
		if strings.EqualFold(method, arg) {
			arg = "value"
		}
	}

	return cmdMeta{
		Type:       cmdType,
		Method:     method,
		Argument:   arg,
		Cmd:        cmd,
		Comments:   strings.Split(description, "\n"),
		CmdComment: arguments,
	}
}

func toGoName(kebab string, upperFirst bool) string {
	kebabTrim := strings.NewReplacer(" ", "-", "_", "-").Replace(strings.TrimSpace(kebab))

	var camelCase string

	isToUpper := false
	for i, runeValue := range kebabTrim {
		switch {
		case i == 0 && upperFirst:
			camelCase += strings.ToUpper(string(runeValue))
		case isToUpper:
			camelCase += strings.ToUpper(string(runeValue))
			isToUpper = false
		case runeValue == '-':
			isToUpper = true
		default:
			camelCase += string(runeValue)
		}
	}
	return camelCase
}
