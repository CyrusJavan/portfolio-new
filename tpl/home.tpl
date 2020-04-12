{{ define "Home" }}
<div class="container">
    <div class="row justify-content-start mt-3 pl-2">
        <div class="col-4-md">
            <h2>Blog Posts</h2>
        </div>
    </div>
    {{ range .Articles }}
        {{ template "_PostSnippet" . }}
    {{ else }}
        <div class="row justify-content-center mt-1">
            <div class="col-4-md">
                <h3>No Posts, Coming Soon.</h3>
            </div>
        </div>
    {{ end }}
    <div class="row justify-content-start mt-3 pl-2">
        <div class="col-4-md">
            <h2>Talks</h2>
        </div>
    </div>
    <div class="row justify-content-center mt-1">
        <div class="col-4-md">
            <h3>No Talks, Coming Soon.</h3>
        </div>
    </div>
</div>
{{end}}