#!/bin/bash

JOB_ID=$(
  if [[ "${GITHUB_WORKFLOW}" ]]; then
    echo "${GITHUB_WORKFLOW}"
  else
    echo 1
  fi
)

function installTools() {
  if [[ "$(
    command -v benchcmp >/dev/null
    echo $?
  )" == "1" ]]; then
    printf "== Installing benchcmp...\n\n"
    go install golang.org/x/tools/cmd/benchcmp
  fi
  if [[ "$(
    command -v benchviz >/dev/null
    echo $?
  )" == "1" ]]; then
    printf "== Installing benchviz...\n\n"
    go install github.com/ajstarks/svgo/benchviz
  fi
}

function tailFile() {
  local file=$1
  tail -n 1 -f "${file}" -s 0.1 2>/dev/null | while read -r line; do
    if [[ "${line}" == "PASS" ]]; then
      break
    fi
    echo "${line}"
  done
}

function nohupFile() {
  local file=$1
  (nohup go test -bench=. >"${file}" &) 2>/dev/null
  sleep 0.2
}

function benchmark() {
  if [[ ! -e old-"${JOB_ID}"-bench.out ]]; then
    printf "== Running benchmark tests...\n\n"
    nohupFile old-"${JOB_ID}"-bench.out
    tailFile old-"${JOB_ID}"-bench.out
  else
    rm -f new-"${JOB_ID}"-bench.out
    printf "== Running benchmark test and comparing with old test results...\n\n"
    nohupFile new-"${JOB_ID}"-bench.out
    tailFile new-"${JOB_ID}"-bench.out

    printf "\n== Comparison:\n\n"
    "$GOPATH/bin/benchcmp" old-"${JOB_ID}"-bench.out new-"${JOB_ID}"-bench.out >benchcmp.out
    cat benchcmp.out
    cat new-"${JOB_ID}"-bench.out >old-"${JOB_ID}"-bench.out
    "$GOPATH/bin/benchviz" >"${JOB_ID}".svg <benchcmp.out
    if [[ -z "${GITHUB_WORKFLOW}" ]]; then
      xdg-open "${JOB_ID}".svg
    fi
  fi
}

function main() {
  installTools
  benchmark
}

main
