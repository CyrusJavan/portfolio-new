#!/usr/bin/env bash

browserify static/js/app.jsx -t babelify -o static/js/bundle.js
