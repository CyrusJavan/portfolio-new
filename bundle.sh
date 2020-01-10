#!/usr/bin/env bash

watchify static/js/app.js -t babelify -o static/js/bundle.js
