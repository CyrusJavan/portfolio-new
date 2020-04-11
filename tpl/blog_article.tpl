{{ define "BlogArticle" }}
<div class="container">
    <div class="row justify-content-start mt-3 pl-2">
        <div class="col-4-md">
            <div class="mt-4">
                <h1>{{ .Title }}</h1>
            </div>
            <span class="text-muted">{{ .Date }}</span>
            <p>
                <span> <div class="author-image d-inline-block"><img class="w-100 h-100 border border-dark rounded-circle" src="{{ .AuthorImage }}" alt="author image"></div> <strong>{{ .Author }}</strong></span>
            </p>
            <p>
                {{ .Content }}
            </p>
            <p>
                {{ template "_PostSnippetTags" .Tags }}
            </p>
        </div>
    </div>
</div>
{{end}}