// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

// +build ignore

package main

import (
	"bytes"
	"flag"
	"go/format"
	"io/ioutil"
	"log"
	"reflect"
	"strings"
	"text/template"

	libconfd "."
)

var flagOutput = flag.String("output", "tmpl_funcs_zz.go", "set output file")

type FuncInfo struct {
	GoFuncName   string
	TmplFuncName string
}

func main() {
	flag.Parse()

	var funcs []FuncInfo
	var tmplFuncType = reflect.TypeOf(libconfd.TemplateFunc{})
	for i := 0; i < tmplFuncType.NumMethod(); i++ {
		method := tmplFuncType.Method(i)
		funcs = append(funcs, FuncInfo{
			GoFuncName:   method.Name,
			TmplFuncName: strings.ToLower(method.Name[:1]) + method.Name[1:],
		})
	}

	code := genFuncCode(funcs)

	data, err := format.Source([]byte(code))
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(*flagOutput, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func genFuncCode(funcs []FuncInfo) string {
	var buf bytes.Buffer
	t := template.Must(template.New("").Funcs(tmplFuncMap).Parse(tmpl))
	err := t.Execute(&buf, funcs)
	if err != nil {
		log.Fatal(err)
	}
	return buf.String()
}

var tmplFuncMap = template.FuncMap{
	"strings_Title":      strings.Title,
	"strings_ToLower":    strings.ToLower,
	"strings_ToUpper":    strings.ToUpper,
	"strings_TrimPrefix": strings.TrimPrefix,
	"strings_TrimSuffix": strings.TrimSuffix,
}

const tmpl = `
{{- $FuncList := . -}}

// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

// Auto generated by go generate, DO NOT EDIT!!!

package libconfd

import (
	"text/template"
)

func init() {
	_TemplateFunc_initFuncMap = func(p *TemplateFunc) {
		p.FuncMap = template.FuncMap{
			{{range $_, $FuncInfo := $FuncList -}}
				"{{$FuncInfo.TmplFuncName}}": p.{{$FuncInfo.GoFuncName}},
			{{end -}}
		}
	}
}
`
