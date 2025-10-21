#!/usr/bin/env bash

#------------------------------------------------------------------------------
# @file
# Builds a Hugo site hosted on a Cloudflare Worker.
#
# The Cloudflare Worker automatically installs Node.js dependencies.
#------------------------------------------------------------------------------

main() {

  DART_SASS_VERSION=1.93.2
  GO_VERSION=1.25.1
  HUGO_VERSION=0.151.0

  export TZ=Asia/Shanghai

  # ===== Install All Dependencies ===== #

  echo "🏗️ Installing all dependencies..."

  # Install Dart Sass
  echo "Installing Dart Sass ${DART_SASS_VERSION}..."
  curl -sLJO "https://github.com/sass/dart-sass/releases/download/${DART_SASS_VERSION}/dart-sass-${DART_SASS_VERSION}-linux-x64.tar.gz"
  tar -C "${HOME}/.local" -xf "dart-sass-${DART_SASS_VERSION}-linux-x64.tar.gz"
  rm "dart-sass-${DART_SASS_VERSION}-linux-x64.tar.gz"
  export PATH="${HOME}/.local/dart-sass:${PATH}"

  # Install Go
  echo "Installing Go ${GO_VERSION}..."
  curl -sLJO "https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz"
  tar -C "${HOME}/.local" -xf "go${GO_VERSION}.linux-amd64.tar.gz"
  rm "go${GO_VERSION}.linux-amd64.tar.gz"
  export PATH="${HOME}/.local/go/bin:${PATH}"

  # Install Hugo
  echo "Installing Hugo ${HUGO_VERSION}..."
  curl -sLJO "https://github.com/gohugoio/hugo/releases/download/v${HUGO_VERSION}/hugo_extended_${HUGO_VERSION}_linux-amd64.tar.gz"
  mkdir "${HOME}/.local/hugo"
  tar -C "${HOME}/.local/hugo" -xf "hugo_extended_${HUGO_VERSION}_linux-amd64.tar.gz"
  rm "hugo_extended_${HUGO_VERSION}_linux-amd64.tar.gz"
  export PATH="${HOME}/.local/hugo:${PATH}"

  # Verify installations
  echo "Verifying installations..."
  echo Dart Sass: "$(sass --version)"
  echo Go: "$(go version)"
  echo Hugo: "$(hugo version)"

  # ===== Execute Pre-Hugo Actions ===== #

  echo "🏗️ Running Custom Go Scripts"

  # compile scripts
  ./scripts/build_scripts.sh
  # execute them
  ./scripts/extract_highlights
  ./scripts/get_external_link_rank
  ./scripts/get_internal_link_graph

  echo "🏗️ Building UnoCSS..."
  npm run build:uno:prod

  echo "🏗️ Configuring Git..."
  git config core.quotepath false
  if [ "$(git rev-parse --is-shallow-repository)" = "true" ]; then
    git fetch --unshallow
  fi

  # ===== Building Hugo ===== #

  echo "🏗️ Building Hugo..."
  hugo --gc --minify

  # ===== Execute Post-Hugo Actions ===== #

  echo "🏗️ Indexing Pagefind..."
  npx pagefind --site public

  # ===== END ===== #
  echo "✅ Build completed successfully!"
}

set -euo pipefail
main "$@"
