// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"reflect"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/iancoleman/strcase"
)

type M map[string]any

var (
	funcMap = mergeFuncMaps(
		sprig.TxtFuncMap(),
		template.FuncMap{
			"urlenc": url.QueryEscape,
			"last": func(x int, a interface{}) bool {
				return x == reflect.ValueOf(a).Len()-1 // if last index.
			},

			// https://github.com/iancoleman/strcase?tab=readme-ov-file#example
			"to_snake":             strcase.ToSnake,           // any_kind_of_string
			"to_snake_with_ignore": strcase.ToSnakeWithIgnore, // ".", any_kind.of_string
			"to_screaming_snake":   strcase.ToScreamingSnake,  // ANY_KIND_OF_STRING
			"to_kebab":             strcase.ToKebab,           // any-kind-of-string
			"to_screaming_kebab":   strcase.ToScreamingKebab,  // ANY-KIND-OF-STRING
			"to_camel":             strcase.ToCamel,           // AnyKindOfString
			"to_lower_camel":       strcase.ToLowerCamel,      // anyKindOfString
		},
	)

	optionGroupTmpl = template.Must(
		template.New("option_group.gotmpl").
			Funcs(funcMap).
			ParseFiles(
				"cmd/codegen/templates/option_group.gotmpl",
			), // TODO: split out into multiple files.
	)

	optionGroupReplacer = strings.NewReplacer(
		" Options", "",
		" and ", " ",
	)
)

func mergeFuncMaps(maps ...template.FuncMap) template.FuncMap {
	out := template.FuncMap{}

	for _, m := range maps {
		for k, v := range m {
			out[k] = v
		}
	}

	return out
}

func createTemplateFile(name string, tmpl *template.Template, data any) {
	// Check if the file exists first, and if it does, panic.
	if _, err := os.Stat(name); err == nil {
		panic(fmt.Sprintf("file %s already exists, not doing anything", name))
	}

	f, err := os.Create(name)
	if err != nil {
		panic(err)
	}

	if err := tmpl.Execute(f, data); err != nil {
		panic(err)
	}

	f.Close()
}

func main() {
	if len(os.Args) < 2 {
		panic("missing command data file")
	}

	var data CommandData

	commandDataFile, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer commandDataFile.Close()

	if err := json.NewDecoder(commandDataFile).Decode(&data); err != nil {
		panic(err)
	}

	data.Generate()

	for _, group := range data.OptionGroups {
		createTemplateFile(fmt.Sprintf("%s.gen.go", strcase.ToSnake(group.Name)), optionGroupTmpl, group)
	}
}