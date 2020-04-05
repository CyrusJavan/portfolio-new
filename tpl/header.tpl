{{ define "Header" }}
<header>
    <nav class="navbar navbar-expand-md navbar-light bg-white border-bottom border-secondary">
        <div class="container">
            <a class="navbar-brand" href="/">Cyrus Javan</a>
            <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNavAltMarkup" aria-controls="navbarNavAltMarkup" aria-expanded="false" aria-label="Toggle navigation">
                <span class="navbar-toggler-icon"></span>
              </button>
              <div class="collapse navbar-collapse" id="navbarNavAltMarkup">
                <div class="navbar-nav">
                  <a class="nav-item nav-link{{ if eq .Page "Blog" }} active{{end}}" href="/blog">Blog</a>
                  <a class="nav-item nav-link{{ if eq .Page "Talks" }} active{{end}}" href="/talks">Talks</a>
                  <a class="nav-item nav-link{{ if eq .Page "About" }} active{{end}}" href="/about">About</a>
                </div>
              </div>
        </div>
    </nav>
</header>
{{end}}