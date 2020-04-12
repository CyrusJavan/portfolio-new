{{ define "Editor" }}
<div class="container">
    <div class="row justify-content-center mt-3">
        <div class="col">
            <div class="form-group">
                <label for="articleTitle">Article Title</label>
                <input type="text" class="form-control" id="articleTitle"
                value="{{ .Title }}">
              </div>
        </div>
    </div>
    <div class="row justify-content-center mt-3">
        <div class="col">
            <form method="post">
                <textarea id="mytextarea" name="mytextarea">
                  {{ if . }}
                    {{ .Content }}
                  {{ else }}
                    No article
                  {{end}}
                </textarea>
            </form>
            <div class="mt-4 mb-4">
                <button id="editor-submit" type="button" class="btn btn-primary">Save</button>
            </div>
        </div>
    </div>
</div>
{{end}}