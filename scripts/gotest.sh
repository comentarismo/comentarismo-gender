#!/usr/bin/env bash

LEARNGENDER=true godep go test -v $(go list ./gender | grep -v /vendor/);
godep go test -v $(go list ./lang | grep -v /vendor/);
godep go test -v $(go list ./server | grep -v /vendor/);
