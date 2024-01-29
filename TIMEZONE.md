# Go-MySQL-Driver TimeZone メモ

基本的には以下でよい

- parseTime は true
- loc に指定するタイムゾーンと MySQL のタイムゾーンを一致させる

これらがずれると、UTC として解釈されるべき値が、JST として解釈されてしまうといった事態が発生する。このような状態に一度なってしまうと後から修正するのは困難。

参考：[go-sql-driver/mysqlと日時データ型とタイムゾーン](https://zenn.dev/utsushiiro/articles/e8d5343cc374a9)

## 実例

DB コンテナのタイムゾーンは UTC

```sql
mysql> show variables like '%time_zone%';
+------------------+--------+
| Variable_name    | Value  |
+------------------+--------+
| system_time_zone | UTC    |
| time_zone        | SYSTEM |
+------------------+--------+
2 rows in set (0.01 sec)

mysql> SELECT @@GLOBAL.time_zone, @@SESSION.time_zone;
+--------------------+---------------------+
| @@GLOBAL.time_zone | @@SESSION.time_zone |
+--------------------+---------------------+
| SYSTEM             | SYSTEM              |
+--------------------+---------------------+
1 row in set (0.00 sec)

mysql> 
```

クライアントのタイムゾーンが Asia/Tokyo の場合、loc パラメータが Local だと JST 時刻が採用され SQL となる。

```go
const DB_USER = "root"
const DB_PASS = "pass"
const DB_HOST = "localhost"
const DB_PORT = "3306"
const DB_NAME = "go-gorm-example"
const DB_LOCA = "Local" // <=== Localになっているため、ローカルのタイムゾーンになる

func DB() *gorm.DB {
	dsn := DB_USER + ":" + DB_PASS + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?charset=utf8mb4&parseTime=True&loc="+DB_LOCA
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
```

出力される SQL は Asia/Tokyo

```sh
INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`name`,`name_kana`,`used_at`,`age`,`role`,`company_id`) VALUES ('2024-01-29 22:25:55.81','2024-01-29 22:25:55.81',NULL,'名前',NULL,'0000-00-00 00:00:00',NULL,NULL,NULL)
```

loca のパラメータを UTC にする。

```go
const DB_LOCA = "UTC" // <=== UTCにする
```

出力される SQL は UTC になる。

```sh
INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`name`,`name_kana`,`used_at`,`age`,`role`,`company_id`) VALUES ('2024-01-29 13:25:55.81','2024-01-29 13:25:55.81',NULL,'名前',NULL,'0000-00-00 00:00:00',NULL,NULL,NULL)
```

ただ、この場合、返ってきた時刻も UTC になるため、意識して JST に戻す必要がある。

loc　は Go-MySQL-Driver の仕様、下記参照。

参考：[Go-MySQL-Driver#Parameters loc](https://github.com/go-sql-driver/mysql?tab=readme-ov-file#loc)

## NOW()の扱い

NOW() または CURRENT_TIMESTAMP() は @@SESSION.time_zone の時刻となる。

```sql
mysql> SELECT @@GLOBAL.time_zone, @@SESSION.time_zone;
+--------------------+---------------------+
| @@GLOBAL.time_zone | @@SESSION.time_zone |
+--------------------+---------------------+
| SYSTEM             | SYSTEM              | <= コンテナのタイムゾーンは UTC
+--------------------+---------------------+
1 row in set (0.00 sec)

mysql> SELECT NOW();
+---------------------+
| NOW()               |
+---------------------+
| 2024-01-29 13:35:00 |  <= UTC 時刻
+---------------------+
1 row in set (0.00 sec)

mysql> 
```

JST として扱う場合は、セッションタイムゾーンを Asia/Tokyo に変更する。

```sql
mysql> SET SESSION time_zone = 'Asia/Tokyo';
Query OK, 0 rows affected (0.00 sec)

mysql> SELECT @@GLOBAL.time_zone, @@SESSION.time_zone;
+--------------------+---------------------+
| @@GLOBAL.time_zone | @@SESSION.time_zone |
+--------------------+---------------------+
| SYSTEM             | Asia/Tokyo          | <= JST に変更
+--------------------+---------------------+
1 row in set (0.00 sec)

mysql> SELECT NOW();
+---------------------+
| NOW()               |
+---------------------+
| 2024-01-29 22:37:06 |  <= JST 時刻
+---------------------+
1 row in set (0.00 sec)

mysql>
```

クラウドのマネージド SQL サーバーはタイムゾーンは UTC として扱い、クライアント側が意識して SQL を発効する必要がある。
