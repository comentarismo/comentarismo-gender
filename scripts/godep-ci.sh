#!/usr/bin/env bash

export GOPATH=/var/lib/jenkins/jobs/comentarismo-gender/workspace;
export PATH=$PATH:/var/lib/jenkins/jobs/comentarismo-gender/workspace/bin;

go get github.com/tools/godep;
go get github.com/stretchr/testify;
go get github.com/smartystreets/goconvey;
go get github.com/drewolson/testflight;
go get github.com/tsenart/vegeta;
cd src/comentarismo-gender;
godep restore;