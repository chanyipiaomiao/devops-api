#!/bin/bash

git clone https://gitee.com/lionsoul/ip2region.git

[[ -e ../data/ip2region.db ]] && mv ../data/ip2region.db ../data/ip2region_`date +%F_%H%M%S`.db

cp ip2region/data/ip2region.db ../data/ip2region.db

rm -rf ip2region