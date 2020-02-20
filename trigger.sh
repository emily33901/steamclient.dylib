#!/bin/sh

git add .
git commit -m "`cat lasttime.txt`"
git tag `cat lasttime.txt`
git push