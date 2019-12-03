#!/bin/sh
set -e


readonly DAY="$1"

cp -r Template "Day$DAY"

cd "Day$DAY"

perl -pi -e "s/var DAY = -1/var DAY = $DAY/g" Template.go
mv Template.go "Day${DAY}.go"

