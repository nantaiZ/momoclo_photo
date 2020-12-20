# !/bin/bash

# MySQLサーバーが起動するまでrealizeを実行せずにループで待機する
until mysqladmin ping -h db -P 3306 --silent; do
    echo 'wating for mysqld to be connectable...'
    sleep 2
done

echo "api is starting...!"
exec realize start --run
