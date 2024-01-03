#!/bin/bash -eux

sudo mkdir -p /var/nginx/cache_1 /var/nginx/cache_2 /var/nginx/cache_3 /var/nginx/cache_4 /var/nginx/cache_5
sudo nginx -t
sudo service nginx start
sudo service mysql start || true # なぜか失敗する(調査中)
sudo chown -R mysql:mysql /var/lib/mysql /var/run/mysqld
sudo service mysql start  # 正しく起動
sudo mysql -u root -pishocon -e 'CREATE DATABASE IF NOT EXISTS ishocon2;' && \
sudo mysql -u root -pishocon -e "CREATE USER IF NOT EXISTS ishocon IDENTIFIED BY 'ishocon';" && \
sudo mysql -u root -pishocon -e 'GRANT ALL ON *.* TO ishocon;' && \
cd ~/data && tar -jxvf ishocon2.dump.tar.bz2 && sudo mysql -u root -pishocon ishocon2 < ~/data/ishocon2.dump
sudo mysql -u root -pishocon -e '
ALTER TABLE votes ADD COLUMN count int(4) NOT NULL;
ALTER TABLE votes ADD INDEX (candidate_id, count DESC);
ALTER TABLE votes DROP COLUMN keyword;
ALTER TABLE users ADD INDEX (name, address, mynumber);
';

echo 'setup completed.'

exec "$@"
