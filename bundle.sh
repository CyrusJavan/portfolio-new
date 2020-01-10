#!/usr/bin/env bash

watchify static/js/app.jsx -t babelify -o static/js/bundle.js
