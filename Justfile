#!/usr/bin/env just --justfile

set shell := ["bash", "-exuo", "pipefail", "-c"]
set export := true
set dotenv-load := false
set positional-arguments := true

compose := "docker compose -p things -f docker-compose.yml"

build *SERVICES:
  {{compose}} build {{SERVICES}}

run *SERVICES:
  {{compose}} up {{SERVICES}}

codegen:
  (cd services/api && go generate .)
  (cd services/app && npm run codegen)
