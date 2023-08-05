---
title: Getting Started with Golang text/template
tags:
  - go
---

# Template package

## First example

- It outputs the result to Stdout
- Insert variables into a template

Use [Go playground](https://go.dev/play/p/gvXME1EA4Yc)

```golang
	type TemplateVariables struct {
		Name string
	}
	tmpl := template.Must(template.New("template").Parse(`{{ .Name }}`))
	err := tmpl.Execute(os.Stdout, TemplateVariables{
		Name: "name inserted",
	})
	if err != nil {
		return err
	}
	return nil
```

## Define a custom function

Use `template.FuncMap` and `template.Template.Func` to define a Custom Func.
In order to use the function in a template, just use that function

- [Go Playground](https://go.dev/play/p/35aRF73MrE4)

```golang
	type TemplateVariables struct {
		Name string
	}

	tmpl := template.Must(template.New("template").Funcs(template.FuncMap{
		"CustomFunc": func(arg string) string {
			return "return " + arg
		},
	}).Parse(`{{ .Name }}
{{ CustomFunc "argument" }}`))
	err := tmpl.Execute(os.Stdout, TemplateVariables{
		Name: "name inserted",
	})
	if err != nil {
		return err
	}
	return nil
```

## Read a file

Use `template.Template.ParseFiles` or `template.ParseFiles`
If you use `template.Template.ParseFiles`, then use the base name of a template file.


# Template features

## Basic syntax

- To access an element of a slice by an index, use `index $var $index`

## To extend a template like inheritance

- Use `block` or `template` on a base file, and `define` in the main file
- Pass these two files to `template.Template.ParseFiles`
- Use the base filename for `template.New`
- In the `define`, the variables defined outside of the block cannot be used. It's out of scope


It was not possible to define a variable in the main file and use it on the base file, probably.
