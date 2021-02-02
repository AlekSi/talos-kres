// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package golangci

const config = `
# options for analysis running
run:
  timeout: 10m
  issues-exit-code: 1
  tests: true
  build-tags: []
  skip-dirs: []
  skip-dirs-use-default: true
  skip-files: []
  modules-download-mode: readonly

# output configuration options
output:
  format: colored-line-number
  print-issued-lines: true
  print-linter-name: true
  uniq-by-line: true
  path-prefix: ""


# all available settings of specific linters
linters-settings:
  dogsled:
    max-blank-identifiers: 2
  dupl:
    threshold: 150
  errcheck:
    check-type-assertions: true
    check-blank: true
  exhaustive:
    default-signifies-exhaustive: false
  funlen:
    lines: 60
    statements: 40
  gci:
    local-prefixes: %[1]s
  gocognit:
    min-complexity: 30
  nestif:
    min-complexity: 5
  goconst:
    min-len: 3
    min-occurrences: 3
  gocritic:
    disabled-checks: []
  gocyclo:
    min-complexity: 20
  godot:
    check-all: false
  godox:
    keywords: # default keywords are TODO, BUG, and FIXME, these can be overwritten by this setting
      - NOTE
      - OPTIMIZE # marks code that should be optimized before merging
      - HACK # marks hack-arounds that should be removed before merging
  gofmt:
    simplify: true
  goimports:
    local-prefixes: %[1]s
  golint:
    min-confidence: 0.8
  gomnd:
    settings: {}
  gomodguard: {}
  govet:
    check-shadowing: true
    enable-all: true
  depguard:
    list-type: blacklist
    include-go-root: false
  lll:
    line-length: 200
    tab-width: 4
  maligned:
    suggest-new: true
  misspell:
    locale: US
    ignore-words: []
  nakedret:
    max-func-lines: 30
  prealloc:
    simple: true
    range-loops: true # Report preallocation suggestions on range loops, true by default
    for-loops: false # Report preallocation suggestions on for loops, false by default
  nolintlint:
    allow-unused: false
    allow-leading-space: false
    allow-no-explanation: []
    require-explanation: false
    require-specific: true
  rowserrcheck: {}
  testpackage:
  unparam:
    check-exported: false
  unused:
    check-exported: false
  whitespace:
    multi-if: false   # Enforces newlines (or comments) after every multi-line if statement
    multi-func: false # Enforces newlines (or comments) after every multi-line function signature
  wsl:
    strict-append: true
    allow-assign-and-call: true
    allow-multiline-assign: true
    allow-cuddle-declarations: false
    allow-trailing-comment: false
    force-case-trailing-whitespace: 0
    force-err-cuddling: false
    allow-separated-leading-comment: false
  gofumpt:
    extra-rules: false

linters:
  enable-all: true
  disable:
    - gas
    - typecheck
    - gochecknoglobals
    - gochecknoinits
    - funlen
    - godox
    - gomnd
    - goerr113
    - nestif
    - wrapcheck
    - paralleltest
    - exhaustivestruct
    - forbidigo
  disable-all: false
  fast: false


issues:
  exclude: []
  exclude-rules: []
  exclude-use-default: false
  exclude-case-sensitive: false
  max-issues-per-linter: 10
  max-same-issues: 3

  new: false

severity:
  default-severity: error

  case-sensitive: false
`
