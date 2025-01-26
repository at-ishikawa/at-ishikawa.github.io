---
title: Articles on each year
---

{{< yearly_archive.inline >}}
  {{ $yearlyArticles := newScratch }}
  {{- range where .Site.RegularPages "Section" "posts" -}}
    {{ $year := .Date.Format "2006" }}
    {{ $article := dict "date" (.Date.Format "2006/01/02") "title" .LinkTitle "permalink" .RelPermalink }}

    {{ $existingYear := $yearlyArticles.Get $year }}
    {{ if $existingYear }}
      {{ $yearlyArticles.Set $year (append $existingYear (slice $article)) }}
    {{ else }}
      {{ $yearlyArticles.Set $year (slice $article) }}
    {{ end }}
  {{- end -}}

  {{ $sortedYears := slice }}
  {{ range $year, $articles := $yearlyArticles.Values }}
    {{ $sortedYears = $sortedYears | append $year }}
  {{ end }}
  {{ $sortedYears = sort $sortedYears "value" "desc" }}

  {{- range $year := $sortedYears -}}
    {{ $articles := sort ($yearlyArticles.Get $year) "date" "desc" }}
    <h2>{{ $year }}</h2>
    <ul>
      {{- range $articles -}}
        <li>
          <a href="{{ .permalink }}">{{ .date }} -- {{ .title }}</a>
        </li>
      {{- end -}}
    </ul>
  {{- end -}}

{{< / yearly_archive.inline >}}
