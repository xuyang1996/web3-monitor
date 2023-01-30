https://www.jianshu.com/p/32335bd372dc

```
docker run -itd --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root mysql

docker run -itd --name phpmyadmin --link mysql -e PMA_HOST=mysql -p 9999:80 phpmyadmin/phpmyadmin
```

