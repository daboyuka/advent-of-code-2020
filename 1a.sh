#!/bin/bash

cat >x
<x sed 's/\(.*\)/2020 \1 - p/' | dc >y
( comm -12 <(sort x) <(sort y); echo "* p" ) | dc
