var SECONDS = 1000;
$(document).ready(function () {
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
