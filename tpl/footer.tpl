{{ define "Footer"}}
<footer class="footer mt-auto">
    <div class="container">
        <div class="row justify-content-center mt-2">
            <div class="col-6-md">
                <div class="d-flex justify-content-around">
                    <a class="mx-3" href="https://www.linkedin.com/in/cyrusjavan"><i class="fa fa-linkedin social-icon"></i></a>
                    <a class="mx-3" href="https://www.github.com/CyrusJavan"><i class="fa fa-github social-icon"></i></a>
                    <a class="mx-3" href="mailto:javan.cyrus+website@gmail.com"><i class="fa fa-envelope social-icon"></i></a>
                </div>
            </div>
        </div>
        <div class="row justify-content-center">
            <div class="col-6-md">
                <div class="d-flex justify-content-center">
                    <span class="text-muted"> &copy; {{.Year}} Cyrus Javan</span>
                </div>
            </div>
        </div>
    </div>
</footer>
{{end}}