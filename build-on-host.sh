#!/bin/sh

git pull
pnpm build:uno:prod
hugo --gc --minify
pnpm pagefind --site public
