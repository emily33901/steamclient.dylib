#!/bin/sh

git add .
git commit -m "`cat lasttime.txt`"
git tag `cat lasttime.txt | sed "s/:/./g"`
git push origin `cat lasttime.txt | sed "s/:/./g"`