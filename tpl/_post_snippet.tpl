{{ define "_PostSnippet" }}
<div class="row justify-content-start mt-2 px-4">
    <div class="col-4-md">
        <span class="text-muted">{{ .Date }}</span>
        <h3><a href="/blog/{{ .Slug }}">{{ .Title }}</a></h3>
        {{ template "_PostSnippetTags" .Tags }}
        <p>
            <div class="mr-3">
                <span>{{ .Snippet }}</span>
            </div>
        </p>
        <!-- <p>
            <span> <div class="author-image d-inline-block"><img class="w-100 h-100 border border-dark rounded-circle" src="{{ .AuthorImage }}" alt="author image"></div> <strong>{{ .Author }}</strong></span>
        </p> -->
    </div>
</div>
{{end}}