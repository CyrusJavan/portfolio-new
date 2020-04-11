<!doctype html>
<html lang="en">

<head>
  <!-- Global site tag (gtag.js) - Google Analytics -->
  <script async src="https://www.googletagmanager.com/gtag/js?id=UA-163289496-1"></script>
  <script>
    window.dataLayer = window.dataLayer || [];
    function gtag(){dataLayer.push(arguments);}
    gtag('js', new Date());

    gtag('config', 'UA-163289496-1');
  </script>
  <title>Cyrus Javan | Software Engineer</title>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
  <meta name="description" content="Cyrus Javan is a Software Engineer, Jiu Jitsu Practioner and Aquarium Enthusiast currently residing in Silicon Valley.">
  <link rel="canonical" href="https://cyrusjavan.com/" />

  <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css" integrity="sha384-MCw98/SFnGE8fJT3GXwEOngsV7Zt27NXFoaoApmYm81iuXoPkFOJwJ8ERdknLPMO" crossorigin="anonymous">
  <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css">
  <link rel="stylesheet" href="/static/css/main.css">

  <link rel="apple-touch-icon" sizes="180x180" href="/static/favicon/apple-touch-icon.png">
  <link rel="icon" type="image/png" sizes="32x32" href="/static/favicon/favicon-32x32.png">
  <link rel="icon" type="image/png" sizes="16x16" href="/static/favicon/favicon-16x16.png">
  <link rel="manifest" href="/static/favicon/site.webmanifest">
  <link rel="mask-icon" href="/static/favicon/safari-pinned-tab.svg" color="#5bbad5">
  <meta name="msapplication-TileColor" content="#da532c">
  <meta name="theme-color" content="#ffffff">
</head>

<body class="d-flex flex-column h-100">
  <div class="page-content flex-fill">
    {{ template "Header" . }}
    {{ if eq .Page "Home"}}
      {{ template "Home" . }}
    {{ else if eq .Page "Blog" }}
      {{ template "Blog" .Articles }}
    {{ else if eq .Page "Talks" }}
      {{ template "Talks" . }}
    {{ else if eq .Page "About" }}
      {{ template "About" . }}
    {{ else if eq .Page "BlogArticle"}}
      {{ template "BlogArticle" .Article }}
    {{end}}
  </div>
  {{ template "Footer" . }}
  {{ template "BodyJS"}}
</body>

</html>