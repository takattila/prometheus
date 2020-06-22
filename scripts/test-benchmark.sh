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

function runBenchmark() {
    local file=$1
     go test -bench=. > "${file}"
     cat "${file}"
}

function benchmark() {
  if [[ ! -e old-"${JOB_ID}"-bench.out ]]; then
    printf "== Running benchmark tests...\n\n"
    runBenchmark old-"${JOB_ID}"-bench.out
  else
    rm -f new-"${JOB_ID}"-bench.out
    printf "== Running benchmark test and comparing with old test results...\n\n"
    runBenchmark new-"${JOB_ID}"-bench.out

    printf "\n== Comparison:\n\n"
    "$GOPATH/bin/benchcmp" old-"${JOB_ID}"-bench.out new-"${JOB_ID}"-bench.out >benchcmp.out
    cat benchcmp.out

    "$GOPATH/bin/benchviz" >"${JOB_ID}".svg <benchcmp.out
    if [[ -z "${GITHUB_WORKFLOW}" ]]; then
      xdg-open "${JOB_ID}".svg
    fi

    cat new-"${JOB_ID}"-bench.out >old-"${JOB_ID}"-bench.out
  fi
}

function main() {
  installTools
  benchmark
}

main
