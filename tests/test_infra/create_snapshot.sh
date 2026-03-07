#!/bin/bash
export CONTEXT="dev"
/opt/homebrew/bin/hcloud --context ${CONTEXT} server create-image --type snapshot fw-test-sys-int --description opnsense-$(date '+%Y%m%d-%H%M') --label test=opnsens
