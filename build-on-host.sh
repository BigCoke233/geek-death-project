#!/bin/sh

git pull
hugo --gc --minify
pnpm pagefind --site public
