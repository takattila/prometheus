#!/bin/bash

function init() {
  JOB_ID=$(
    if [[ "${GITHUB_WORKFLOW}" ]]; then
      echo "${GITHUB_WORKFLOW}"
    else
      echo 1
    fi
  )

  CACHE_DIR=~/.test_cache

  GOPATH=$(
    if [[ -z "${GITHUB_WORKFLOW}" ]]; then
      echo "${GOPATH}"
    else
      echo "${HOME}/go"
    fi
  )
  mkdir -p "${CACHE_DIR}"
}

function installTools() {
  if [[ ! -e "${GOPATH}/bin/benchstat" ]]; then
    printf "== Installing benchstat...\n"
    go get -u golang.org/x/perf/cmd/benchstat
  fi
}

function runBenchmark() {
  local file=$1
  go test -bench=. >"${file}"
  cat "${file}"
}

function benchmark() {
  local benchStat="${GOPATH}/bin/benchstat"

  if [[ ! -e "${CACHE_DIR}"/old-"${JOB_ID}"-bench.out ]]; then
    printf "== Running benchmark tests...\n\n"
    runBenchmark "${CACHE_DIR}"/old-"${JOB_ID}"-bench.out
  else
    rm -f "${CACHE_DIR}"/new-"${JOB_ID}"-bench.out
    printf "== Running benchmark tests and comparing with old test results...\n\n"
    runBenchmark "${CACHE_DIR}"/new-"${JOB_ID}"-bench.out

    printf "\n== Comparison:\n\n"
    "${benchStat}" "${CACHE_DIR}"/old-"${JOB_ID}"-bench.out "${CACHE_DIR}"/new-"${JOB_ID}"-bench.out
    cat "${CACHE_DIR}"/new-"${JOB_ID}"-bench.out >"${CACHE_DIR}"/old-"${JOB_ID}"-bench.out
  fi
}

function main() {
  init
  installTools
  benchmark
}

main
