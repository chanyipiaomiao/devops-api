#!/bin/bash

git clone https://github.com/xluohome/phonedata

[[ -e ../data/phone.dat ]] && mv ../data/phone.dat ../data/phone_`date +%F_%H%M%S`.dat

cp phonedata/phone.dat ../data/phone.dat

rm -rf phonedata