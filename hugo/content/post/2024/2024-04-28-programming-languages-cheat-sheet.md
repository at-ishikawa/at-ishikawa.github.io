---
date: "2024-04-28T00:00:00Z"
last_modified_at: "2025-01-19"
tags:
- programming
title: Cheat sheet for programming languages

# Page variables
table_columns: &default_columns
- code: php
  text: PHP
- code: python
  text: Python
- code: golang
  text: Golang
- code: typescript
  text: TypeScript

table_basics:
  columns: *default_columns
  rows:
    - description: Null
      is_plain_text: true
      languages:
        php: "null"
        python: None
        golang: nil
        typescript: "null"
    - description: Undefined key
      is_plain_text: true
      languages:
        typescript: undefined
    - description: Equal
      languages:
        php: $var1 == $var2
        python: var1 == var2
        golang: var1 == var2
        typescript: var1 == var2
    - description: Strict equal
      languages:
        php: $var1 === $var2
        typescript: var1 === var2
    - description: Null Coalescence
      languages:
        php: $var ?? 'default'
        typescript: var ?? 'default'
    - description: Null Assertion (Ignore null value)
      languages:
        typescript: var!

table_type:
  columns: *default_columns
  rows:
    - description: Type Alias
      languages:
        golang: type NewType OldType
        typescript: |
          type VariableType = OldType
          type FunctionType = (var: number) => void
    - description: Interface
      languages:
        golang: |
          interface Interface {
            method(var int) void
          }
        typescript: |
          interface VariableInterface {
            var: number;
          }
          interface FunctionInterface {
            (var: number): void;
          }

table_control_flows:
  columns: *default_columns
  rows:
    - description: if, else if, else
      languages:
        php: |
          if ($condition1) {
              // logic 1
          } else if ($condition2) {
              // logic 2
          } else {
              // logic 3
          }
        python: |
          if $condition1:
              # logic 1
          elif $condition2:
              # logic 2
          else:
              # logic 3
        golang: |
          if condition1 {
            // logic 1
          } else if condition 2 {
            // logic 2
          } else {
            // logic 3
          }
    - description: for
      languages:
        php: |
          for ($i = 0; $i < 10; $i++) {
            # do something
          }
        python: |
          for i in range(10):
            # do something
        golang: |
          for i := 0; i < 10; i++ {
            // do something
          }
    - description: foreach for an array without a index
      languages:
        php: |
          foreach ($array as $value) {
            // do something
          }
        python: |
          for value in list:
            # do something
        golang: |
          for _, value := range slice {
            // do something
          }
    - description: foreach for an array with an index
      languages:
        php: |
          foreach ($array as $index => $value) {
            // do something
          }
        python: |
          for index, value in enumerate(list):
            # do something
        golang: |
          for index, value := range slice {
            // do something
          }
          // Or
          for index := range slice {
            value := slice[index]
            // do something
          }
    - description: foreach for a hash
      languages:
        php: |
          foreach ($array as $key => $value) {
            // do something
          }
        python: |
          for key, value in dictionary.items():
            # do something
        golang: |
          for key, value := range map {
            // do something
          }

table_numbers:
  columns: *default_columns
  rows:
    - description: Decimal to binary
      languages:
        golang: strconv.FormatInt(decimal, 2)

table_list:
  columns: *default_columns
  rows:
    - description: Type name
      is_plain_text: true
      languages:
        php: array
        python: List
        golang: Slice
    - description: Initialize
      languages:
        php: "$list = []"
        python: "list = []"
        golang: "slice := make([]int, length) or slice := make([]int, length, capacity)"
    - description: Add an element
      languages:
        php: "$list[] = $element"
        python: "list.append(element)"
        golang: "slice = append(slice, element)"
    - description: Sort (ascending)
      languages:
        php: "sort($list)"
        python: |
          list.sort()
          # or
          sortedList = sorted(list)
        golang: |
          sort.Ints(s)
          # or
          sort.Slice(slice, func (i, j int) bool {
            return slice[i] < slice[j]
          })
    - description: Sort (descending)
      languages:
        php: "rsort($list)"
        python: |
          list.sort(reverse=True)
          # or
          sortedList = sorted(list, reverse=True)
        golang: "sort.Reverse(sort.Sort(slice))"

table_hash:
  columns: *default_columns
  rows:
    - description: Type
      is_plain_text: true
      languages:
        php: array
        python: Dictionary
        golang: Map
    - description: Initialize
      languages:
        php: "$hash = []"
        python: "dictionary = {}"
        golang: |
          m := make(map[string]int)
          # or
          slice := make(map[string]int, capacity)
    - description: Add
      languages:
        php: "$hash[$key] = $value"
        python: "dictionary[key] = value"
        golang: "m[key] = value"
    - description: Check if a key exists
      languages:
        php: "array_key_exists($key, $hash)"
        python: "key in dictionary"
        golang: "value, exists := m[key]"

table_queue:
  columns: *default_columns
  rows:
    - description: Type
      is_plain_text: true
      languages:
        python: "collections.deque"
        golang: Slice
    - description: Initialize
      languages:
        python: "queue = deque([1, 2])"
    - description: Enqueue
      languages:
        python: "queue.append(element)"
    - description: Dequeue
      languages:
        python: "queue.pop()"
    - description: Length
      languages:
        python: "len(queue)"
    - description: Empty
      languages:
        python: "bool(queue)"

table_heap:
  columns: *default_columns
  rows:
    - description: Type
      is_plain_text: true
      languages:
        python: "heap.Interface"
        golang: "[container.heap.Interface](https://pkg.go.dev/container/heap)"
    - description: Initialize
      languages:
        python: "heap.Init(h heap.Interface)"
    - description: Enqueue
      languages:
        python: "heap.Push(h heap.Interface, value any)"
    - description: Dequeue
      languages:
        python: "heap.Pop(h heap.Interface) any"

table_package_managers:
  columns:
    - code: php
      description: PHP
      children:
        - code: composer
          description: Composer
    - code: python
      description: Python
      children:
        - code: pdm
          description: "[PDM](https://pdm-project.org/en/latest/)"
        - code: poetry
          description: "[Poerty](https://python-poetry.org/)"
        - code: hatch
          description: "[Hatch](https://hatch.pypa.io/latest/)"
        - code: conda
          description: "[conda](https://docs.conda.io/projects/conda/en/stable/user-guide/install/index.html)"
        - code: venv
          description: "[venv module](https://docs.python.org/3/library/venv.html)"
    - code: golang
      description: Golang
      children:
        - code: "go module"
          description: go module
    - code: typescript
      description: TypeScript
      children:
        - code: npm
          description: npm
        - code: yarn
          description: yarn
        - code: pnpm
          description: pnpm
  rows:
    - description: Support different language versions
      values:
        composer: "No"
        venv: "No"
        conda: "Yes"
    - description: How to set up
      highlight_lang: bash
      values:
        venv: |
          python -m venv /path/to/venv
          source /path/to/venv/bin/activate.$SHELL_NAME
---


## Basics
{{< render_codes_table.inline columns="table_columns" table="table_basics" >}}
{{ $columns := index $.Page.Params (.Get "columns") }}
{{ $table := index $.Page.Params (.Get "table" )}}
{{ $rows := index $table "rows" }}

<table>
  <thead>
    <tr>
      <th></th>
      {{- range $columns -}}
        <th>{{ .text }}</th>
      {{- end -}}
    </tr>
  </thead>
  <tbody>
    {{- range $row := $rows -}}
    <tr>
      <td>{{ $row.description }}</td>
      {{ range $column := $columns }}
        <td>
          {{- $cell := index $row.languages $column.code -}}
          {{- if $cell -}}
            {{- if $row.is_plain_text -}}
              {{- $cell | page.RenderString -}}
            {{- else -}}
              {{- highlight $cell $column.code -}}
            {{- end -}}
          {{- end -}}
        </td>
      {{ end }}
    </tr>
    {{ end }}
  </tbody>
</table>
{{< /render_codes_table.inline >}}


- Strict equal in TypeScript
    - 2 variables have the same types

### Syntax for types
{{< render_codes_table.inline columns="table_columns" table="table_type" />}}

The difference of type and interface between TypeScript:
- The type can be used for a type alias, but not for interface
- The class/interface is static, so it cannot extend the union of types
- An interface can be defined multiple times
- An interface isn't used to rename primitives
- An interface has a performance advantage. See [this article](https://github.com/microsoft/TypeScript/wiki/Performance#preferring-interfaces-over-intersections)

These were discussed in the following articles:
- [Stackoverflow](https://stackoverflow.com/a/52682220)
- [Typescript official doc](https://www.typescriptlang.org/docs/handbook/2/everyday-types.html#differences-between-type-aliases-and-interfaces)


### Syntax for control flow
{{< render_codes_table.inline columns="table_columns" table="table_control_flows" />}}

## Operations for common types
### Numbers
{{< render_codes_table.inline columns="table_columns" table="table_numbers" />}}

### List
{{< render_codes_table.inline columns="table_columns" table="table_list" />}}

### Hash
{{< render_codes_table.inline columns="table_columns" table="table_hash" />}}

### Queue
{{< render_codes_table.inline columns="table_columns" table="table_queue" />}}

### Heap
{{< render_codes_table.inline columns="table_columns" table="table_heap" />}}

#### Golang implementation for Heap

Golang needs to implement the interface `[container.heap.Interface]` to use a heap or a priority queue.

```golang
// An Item is something we manage in a priority queue.
type Item struct {
	value    string
	priority int

}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	*pq = old[0 : n-1]
	return item
}
```

# Tools
## Package managers / Virtual environments

{{< render_package_managers.inline >}}
{{ $table := index $.Page.Params "table_package_managers"}}
{{ $rows := index $table "rows" }}

<table>
  <thead>
    <tr>
      <th></th>
      {{- range $column := $table.columns -}}
        <th colspan={{ len $column.children }}>{{ $column.description}}</th>
      {{- end -}}
    </tr>
    <tr>
      <th></th>
      {{- range $column := $table.columns -}}
        {{- range $child := $column.children -}}
          <th>{{ $child.description | page.RenderString }}</th>
        {{- end -}}
      {{- end -}}
    </tr>
  </thead>

  <tbody>
    {{- range $row := $table.rows -}}
    <tr>
      <td>{{ $row.description }}</td>
      {{- range $column := $table.columns -}}
        {{- range $child := $column.children -}}
          <td>
            {{- $value := index $row.values $child.code -}}
            {{- if $value -}}
              {{- if (index $row "highlight_lang") -}}
                {{- highlight $value (index $row "highlight_lang") -}}
              {{- else -}}
                {{- $value | page.RenderString -}}
              {{- end -}}
            {{- end -}}
          </td>
        {{- end -}}
      {{- end -}}
    </tr>
    {{- end -}}
  </tbody>

</table>
{{< /render_package_managers.inline >}}

For Python, see some articles like [this dev.to article](https://dev.to/adamghill/python-package-manager-comparison-1g98) for better comparisons.
