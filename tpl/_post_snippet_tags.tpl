{{ define "_PostSnippetTags" }}
    {{ range . }}
        <a class="px-1 mr-1 bg-light text-info" href="#">{{ . }}</a>
    {{end}}
{{end}}