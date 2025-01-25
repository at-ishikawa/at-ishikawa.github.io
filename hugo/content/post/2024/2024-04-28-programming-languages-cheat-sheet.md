---
date: "2024-04-28T00:00:00Z"
last_modified_at: "2025-01-19"
table_basics:
  columns:
  - php: PHP
  - python: Python
  - golang: Golang
  - typescript: TypeScript
  rows:
  - description: null
    is_plain_text: true
    languages:
      golang: nil
      php: "null"
      python: None
      typescript: "null"
  - description: Undefined key
    is_plain_text: true
    languages:
      typescript: undefined
  - description: Equal
    languages:
      golang: var1 == var2
      php: $var1 == $var2
      python: var1 == var2
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
table_columns:
- php: PHP
- python: Python
- golang: Golang
- typescript: TypeScript
table_control_flows:
  columns:
  - php: PHP
  - python: Python
  - golang: Golang
  - typescript: TypeScript
  rows:
  - description: if, else if, else
    languages:
      golang: |
        if condition1 {
          // logic 1
        } else if condition 2 {
          // logic 2
        } else {
          // logic 3
        }
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
  - description: for
    languages:
      golang: |
        for i := 0; i < 10; i++ {
          // do something
        }
      php: |
        for ($i = 0; $i < 10; $i++) {
          # do something
        }
      python: |
        for i in range(10):
          # do something
  - description: foreach for an array without a index
    languages:
      golang: |
        for _, value := range slice {
          // do something
        }
      php: |
        foreach ($array as $value) {
          // do something
        }
      python: |
        for value in list:
          # do something
  - description: foreach for an array with an index
    languages:
      golang: |
        for index, value := range slice {
          // do something
        }
        // Or
        for index := range slice {
          value := slice[index]
          // do something
        }
      php: |
        foreach ($array as $index => $value) {
          // do something
        }
      python: |
        for index, value in enumerate(list):
          # do something
  - description: foreach for a hash
    languages:
      golang: |
        for key, value := range map {
          // do something
        }
      php: |
        foreach ($array as $key => $value) {
          // do something
        }
      python: |
        for key, value in dictionary.items():
          # do something
table_hash:
  columns:
  - php: PHP
  - python: Python
  - golang: Golang
  - typescript: TypeScript
  rows:
  - description: Type
    is_plain_text: true
    languages:
      golang: Map
      php: array
      python: Dictionary
  - description: Initialize
    languages:
      golang: |
        m := make(map[string]int)
        # or
        slice := make(map[string]int, capacity)
      php: $hash = []
      python: dictionary = {}
  - description: Add
    languages:
      golang: m[key] = value
      php: $hash[$key] = $value
      python: dictionary[key] = value
  - description: Check if a key exists
    languages:
      golang: value, exists := m[key]
      php: array_key_exists($key, $hash)
      python: key in dictionary
table_heap:
  columns:
  - php: PHP
  - python: Python
  - golang: Golang
  - typescript: TypeScript
  rows:
  - description: Type
    is_plain_text: true
    languages:
      golang: '[container.heap.Interface](https://pkg.go.dev/container/heap)'
      python: heap.Interface
  - description: Initialize
    languages:
      python: heap.Init(h heap.Interface)
  - description: Enqueue
    languages:
      python: heap.Push(h heap.Interface, value any)
  - description: Dequeue
    languages:
      python: heap.Pop(h heap.Interface) any
table_list:
  columns:
  - php: PHP
  - python: Python
  - golang: Golang
  - typescript: TypeScript
  rows:
  - description: Type name
    is_plain_text: true
    languages:
      golang: Slice
      php: array
      python: List
  - description: Initialize
    languages:
      golang: slice := make([]int, length) or slice := make([]int, length, capacity)
      php: $list = []
      python: list = []
  - description: Add an element
    languages:
      golang: slice = append(slice, element)
      php: $list[] = $element
      python: list.append(element)
  - description: Sort (ascending)
    languages:
      golang: |
        sort.Ints(s)
        # or
        sort.Slice(slice, func (i, j int) bool {
          return slice[i] < slice[j]
        })
      php: sort($list)
      python: |
        list.sort()
        # or
        sortedList = sorted(list)
  - description: Sort (descending)
    languages:
      golang: sort.Reverse(sort.Sort(slice))
      php: rsort($list)
      python: |
        list.sort(reverse=True)
        # or
        sortedList = sorted(list, reverse=True)
table_numbers:
  columns:
  - php: PHP
  - python: Python
  - golang: Golang
  - typescript: TypeScript
  rows:
  - description: Decimal to binary
    languages:
      golang: strconv.FormatInt(decimal, 2)
table_package_managers:
  columns:
  - php:
      children:
      - composer: Composer
      description: PHP
  - python:
      children:
      - pdm: '[PDM](https://pdm-project.org/en/latest/)'
      - poetry: '[Poerty](https://python-poetry.org/)'
      - hatch: '[Hatch](https://hatch.pypa.io/latest/)'
      - conda: '[conda](https://docs.conda.io/projects/conda/en/stable/user-guide/install/index.html)'
      - venv: '[venv module](https://docs.python.org/3/library/venv.html)'
      description: Python
  - golang:
      children:
      - go module: go module
      description: Golang
  - typescript:
      children:
      - npm: npm
      - yarn: yarn
      - pnpm: pnpm
      description: TypeScript
  rows:
  - description: Support different language versions
    values:
      composer: "No"
      conda: "Yes"
      venv: "No"
  - description: How to set up
    highlight_lang: bash
    values:
      venv: |
        python -m venv /path/to/venv
        source /path/to/venv/bin/activate.$SHELL_NAME
table_queue:
  columns:
  - php: PHP
  - python: Python
  - golang: Golang
  - typescript: TypeScript
  rows:
  - description: Type
    is_plain_text: true
    languages:
      golang: Slice
      python: collections.deque
  - description: Initialize
    languages:
      python: queue = deque([1, 2])
  - description: Enqueue
    languages:
      python: queue.append(element)
  - description: Dequeue
    languages:
      python: queue.pop()
  - description: Length
    languages:
      python: len(queue)
  - description: Empty
    languages:
      python: bool(queue)
table_type:
  columns:
  - php: PHP
  - python: Python
  - golang: Golang
  - typescript: TypeScript
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
tags:
- programming
title: Cheat sheet for programming languages
---

## Basics
{% include render_codes_table.html columns=page.table_basics.columns rows=page.table_basics.rows %}

- Strict equal in TypeScript
    - 2 variables have the same types

### Syntax for types
{% include render_codes_table.html columns=page.table_type.columns rows=page.table_type.rows %}

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
{% include render_codes_table.html columns=page.table_control_flows.columns rows=page.table_control_flows.rows %}

## Operations for common types
### Numbers
{% include render_codes_table.html columns=page.table_numbers.columns rows=page.table_numbers.rows %}

### List
{% include render_codes_table.html columns=page.table_list.columns rows=page.table_list.rows %}

### Hash
{% include render_codes_table.html columns=page.table_hash.columns rows=page.table_hash.rows %}

### Queue
{% include render_codes_table.html columns=page.table_queue.columns rows=page.table_queue.rows %}

### Heap

{% include render_codes_table.html columns=page.table_heap.columns rows=page.table_heap.rows %}

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

<table>
  <thead>
    <tr>
      <th></th>
      {% for hash in page.table_package_managers.columns %}
        {% assign column = hash | first | last %}
        <th colspan={{ column.children | size }}>{{ column.description }}</th>
      {% endfor %}
    </tr>
    <tr>
      <th></th>
      {% for hash in page.table_package_managers.columns %}
        {% assign column = hash | first | last %}
        {% for child_hash in column.children %}
          {% assign desc = child_hash | first | last %}
          <th>{{ desc | markdownify }}</th>
        {% endfor %}
      {% endfor %}
    </tr>
  </thead>

  <tbody>
    {% for row in page.table_package_managers.rows %}
    <tr>
      <td>{{ row.description }}</td>
      {% for hash in page.table_package_managers.columns -%}
        {%- assign column = hash | first | last -%}
        {%- for child_hash in column.children -%}
          {%- assign package_manager = child_hash | first | first -%}
          <td>
            {%- assign value = row.values[package_manager] -%}
            {%- if value -%}
              {%- if row.highlight_lang -%}
                {%- highlight_param bash -%}
                  {{- value -}}
                {%- endhighlight_param -%}
              {%- else -%}
                {{- value | markdownify -}}
              {%- endif -%}
            {%- endif -%}
          </td>
        {%- endfor -%}
      {%- endfor %}
    </tr>
    {% endfor %}
  </tbody>
</table>

For Python, see some articles like [this dev.to article](https://dev.to/adamghill/python-package-manager-comparison-1g98) for better comparisons.
