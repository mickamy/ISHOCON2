#!/bin/bash -eux

pt-query-digest --type slowlog /var/log/mysql/slow_query.log | tee "../digetst_$(date +%Y%m%d%H%M).log"
rm -rf /var/log/mysql/slow_query.log
