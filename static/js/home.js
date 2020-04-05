var SECONDS = 1000;
$(document).ready(function () {
    trackPageViews();
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