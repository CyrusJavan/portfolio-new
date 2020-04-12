var SECONDS = 1000;
$(document).ready(function () {
    //trackPageViews();

    $('#editor-submit').click(function () {
        var snippetLen = 400;
        $.post($(location).attr('pathname'), {
            content: tinymce.get('mytextarea').getContent(),

            // Is this proper? Maybe the snippet should just be generated server side from the
            // content itself.
            snippet: tinymce.get('mytextarea').getContent({format: 'text'}).substring(0, snippetLen) + "...",
            title: $('#articleTitle').val()
        });
    });
});

function trackPageView(type) {
    $.post('/api/track', {
        type: type,
        page: location.pathname,
    });
}

function trackPageViews() {
    trackPageView("000s");

    setTimeout(function () {
        trackPageView("010s");
    }, 10 * SECONDS);

    setTimeout(function () {
        trackPageView("060s");
    }, 60 * SECONDS);

    setTimeout(function () {
        trackPageView("300s");
    }, 300 * SECONDS);
}