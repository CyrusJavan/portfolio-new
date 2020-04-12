{{ define "BlogArticle" }}
<div class="container">
    <div class="row justify-content-start px-2">
        <div class="col-4-md mt-4">
            <div>
                <h1>{{ .Title }}</h1>
            </div>
        </div>
    </div>

    <div class="row justify-content-start mt-3 px-2">
        <div class="col-4-md">
            <span class="text-muted">{{ .Date }}</span>
        </div>
    </div>

    <div class="row justify-content-start px-2">
        <div class="col-4-md mt-1">
            <span> <div class="author-image d-inline-block"><img class="w-100 h-100 border border-dark rounded-circle" src="{{ .AuthorImage }}" alt="author image"></div> <strong>{{ .Author }}</strong></span>
        </div>
    </div>

    <div class="row justify-content-start mt-3 px-2">
        <div class="col-4-md mt-4">
            {{ .Content }}
        </div>
    </div>

    <div class="row justify-content-start mt-3 px-2">
        <div class="col-4-md mt-1 mb-5">
            {{ template "_PostSnippetTags" .Tags }}
        </div>
    </div>
</div>
{{end}}