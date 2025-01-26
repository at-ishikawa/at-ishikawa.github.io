---
# https://mmistakes.github.io/minimal-mistakes/docs/layouts/#splash-page-layout
layout: home
title: at-ishikawa's technical notebooks
author_profile: true
excerpt: Notes for software development and operations
header:
  overlay_image: /assets/images/home/overlay.jpg
  overlay_filter: rgba(0, 128, 128, 0.5)
  caption: "Image credit: [**xresch on pixabay**](https://pixabay.com/photos/circle-tech-technology-abstract-5090539/)"
# feature_row:
#   - image_path: /assets/images/home/feature_row_orders.jpg
#     image_caption: "Image credit: [**markusspiske on pixabay**](https://pixabay.com/photos/hacker-cyber-code-angrfiff-3655668/)"
#     alt: "No image"
#     title: "Recent posts"
#     url: "/year-archive/"
#     btn_label: "See recent posts"
#     btn_class: "btn--inverse"
#   - image_path: /assets/images/home/feature_row_post.jpg
#     image_caption: "Image credit: [**joffi on pixabay**](https://pixabay.com/photos/hacking-hacker-computer-internet-1685092/)"
#     alt: "No image"
#     title: "Posts by tags"
#     url: "/tags/"
#     btn_label: "See tags"
#     btn_class: "btn--inverse"
---

My GitHub repositories I use for the development and operation:

- [Fish Completion Interceptor](https://github.com/at-ishikawa/fish-completion-interceptor/): Enable to show suggestions like kubectl for completion
- [GitHub Project Prometheus Exporter](https://github.com/at-ishikawa/github_project_prometheus_exporter): Export prometheus metrics for a GitHub project


## Search articles

I've written many articles and here are the links to search them;
- [Search articles by time](/year-archive/)
- [Search articles by tags](/tags/)

## Recent articles

{{< recent_posts.inline >}}
<ul>
  {{- range $i, $v := where .Site.RegularPages "Section" "posts" -}}{{ if lt $i 10 -}}
    <li>
      <a href="{{ .RelPermalink }}">{{.Date.Format "2006/01/02"}} -- {{ .LinkTitle }}</a>
    </li>
  {{- end -}}{{- end -}}
</ul>
{{< / recent_posts.inline >}}
