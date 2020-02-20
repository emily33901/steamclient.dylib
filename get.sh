#!/bin/sh

curl -sqL "https://steamcdn-a.akamaihd.net/client/installer/steamcmd_osx.tar.gz" | tar zxvf -
./steamcmd.sh +quit


brew install tree

tree

