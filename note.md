- 初回ベンチ

```
❯ make bench
docker exec -i ishocon2-bench-1 sh -c "./benchmark --ip app:443"
2024/01/01 23:15:48 Start GET /initialize
2024/01/01 23:15:48 期日前投票を開始します
2024/01/01 23:15:49 期日前投票が終了しました
2024/01/01 23:15:49 投票を開始します  Workload: 3
2024/01/01 23:16:35 投票が終了しました
2024/01/01 23:16:35 投票者が結果を確認しています
2024/01/01 23:16:50 投票者の感心がなくなりました
2024/01/01 23:16:50 {"score": 6980, "success": 6340, "failure": 0}

~/ghq/github.com/mickamy/ISHOCON2 imp1* 1m 2s
```

- Locale/TZ 設定

```
❯ make bench
docker exec -i ishocon2-bench-1 sh -c "./benchmark --ip app:443"
2024/01/01 23:34:15 Start GET /initialize
2024/01/01 23:34:15 期日前投票を開始します
2024/01/01 23:35:03 期日前投票が終了しました
2024/01/01 23:35:03 投票を開始します  Workload: 3
2024/01/01 23:35:50 投票が終了しました
2024/01/01 23:35:50 投票者が結果を確認しています
2024/01/01 23:36:05 投票者の感心がなくなりました
2024/01/01 23:36:05 {"score": 6098, "success": 5474, "failure": 0}

~/ghq/github.com/mickamy/ISHOCON2 imp1* 1m 51s
```

TZ 設定できてない :pieng:

- update ruby to 3.1.4

```
❯ make bench
docker exec -i ishocon2-bench-1 sh -c "./benchmark --ip app:443"
2024/01/01 23:41:21 Start GET /initialize
2024/01/01 23:41:21 期日前投票を開始します
2024/01/01 23:41:29 期日前投票が終了しました
2024/01/01 23:41:29 投票を開始します  Workload: 3
2024/01/01 23:42:15 投票が終了しました
2024/01/01 23:42:15 投票者が結果を確認しています
2024/01/01 23:42:31 投票者の感心がなくなりました
2024/01/01 23:42:31 {"score": 6948, "success": 6324, "failure": 0}

~/ghq/github.com/mickamy/ISHOCON2 imp1* 1m 11s
```

- show create table

```
| votes | CREATE TABLE `votes` (
  `id` int(32) NOT NULL AUTO_INCREMENT,
  `user_id` int(32) NOT NULL,
  `candidate_id` int(11) NOT NULL,
  `keyword` text NOT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=950393 DEFAULT CHARSET=utf8mb4 |

| users | CREATE TABLE `users` (
  `id` int(32) NOT NULL AUTO_INCREMENT,
  `name` varchar(128) NOT NULL,
  `address` varchar(256) NOT NULL,
  `mynumber` varchar(32) NOT NULL,
  `votes` int(4) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `mynumber` (`mynumber`)
) ENGINE=InnoDB AUTO_INCREMENT=4000001 DEFAULT CHARSET=utf8mb4 |

| candidates | CREATE TABLE `candidates` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(128) NOT NULL,
  `political_party` varchar(128) NOT NULL,
  `sex` varchar(32) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8mb4 |
```

- bundle update and specify ruby version 

```
❯ make bench
docker exec -i ishocon2-bench-1 sh -c "./benchmark --ip app:443"
2024/01/02 00:35:45 Start GET /initialize
2024/01/02 00:35:45 期日前投票を開始します
2024/01/02 00:35:46 期日前投票が終了しました
2024/01/02 00:35:46 投票を開始します  Workload: 3
2024/01/02 00:36:31 投票が終了しました
2024/01/02 00:36:31 投票者が結果を確認しています
2024/01/02 00:36:47 投票者の感心がなくなりました
2024/01/02 00:36:47 {"score": 7144, "success": 6472, "failure": 0}

~/ghq/github.com/mickamy/ISHOCON2 imp1* 1m 2s
```

- unicorn のワーカ増やした
- candidates をメモリに載せようとしたら壊れたので戻した
  - 元々30件だからたいしたことないか？
  - TODO: 余裕できたら再チャレンジ

```
❯ make bench
docker exec -i ishocon2-bench-1 sh -c "./benchmark --ip app:443"
2024/01/02 01:19:04 Start GET /initialize
2024/01/02 01:19:05 期日前投票を開始します
2024/01/02 01:19:13 エラーメッセージに誤りがあります at POST /vote
make: *** [bench] Error 1

~/ghq/github.com/mickamy/ISHOCON2 imp1* 9s
```

- unicorn のプロセスのおかげでスコアだけ上がった

```
❯ make bench
docker exec -i ishocon2-bench-1 sh -c "./benchmark --ip app:443"
2024/01/02 01:33:23 Start GET /initialize
2024/01/02 01:33:23 期日前投票を開始します
2024/01/02 01:33:24 期日前投票が終了しました
2024/01/02 01:33:24 投票を開始します  Workload: 3
2024/01/02 01:34:09 投票が終了しました
2024/01/02 01:34:09 投票者が結果を確認しています
2024/01/02 01:34:25 投票者の感心がなくなりました
2024/01/02 01:34:25 {"score": 17144, "success": 16072, "failure": 0}
```

- workload と unicorn worker を増やすが、悪くなった

```
❯ make bench
docker exec -i ishocon2-bench-1 sh -c "./benchmark --ip app:443 --workload 8"
2024/01/02 01:40:43 Start GET /initialize
2024/01/02 01:40:43 期日前投票を開始します
2024/01/02 01:40:48 期日前投票が終了しました
2024/01/02 01:40:48 投票を開始します  Workload: 8
2024/01/02 01:41:42 投票が終了しました
2024/01/02 01:41:42 投票者が結果を確認しています
2024/01/02 01:41:57 投票者の感心がなくなりました
2024/01/02 01:41:57 {"score": 12196, "success": 10148, "failure": 0}

~/ghq/github.com/mickamy/ISHOCON2 imp1* 1m 15s
```

- workload: 4, unicorn: 12

```
❯ make bench
docker exec -i ishocon2-bench-1 sh -c "./benchmark --ip app:443 --workload 4"
2024/01/02 01:43:27 Start GET /initialize
2024/01/02 01:43:27 期日前投票を開始します
2024/01/02 01:43:55 期日前投票が終了しました
2024/01/02 01:43:55 投票を開始します  Workload: 4
2024/01/02 01:44:41 投票が終了しました
2024/01/02 01:44:41 投票者が結果を確認しています
2024/01/02 01:44:56 投票者の感心がなくなりました
2024/01/02 01:44:56 {"score": 16942, "success": 16046, "failure": 0}

~/ghq/github.com/mickamy/ISHOCON2 imp1* 1m 30s
```

- ALTER TABLE votes ADD INDEX idx_votes_candidate_id (candidate_id); 

```
❯ make bench
docker exec -i ishocon2-bench-1 sh -c "./benchmark --ip app:443 --workload 4"
2024/01/02 01:49:33 Start GET /initialize
2024/01/02 01:49:33 期日前投票を開始します
2024/01/02 01:49:34 期日前投票が終了しました
2024/01/02 01:49:34 投票を開始します  Workload: 4
2024/01/02 01:50:20 投票が終了しました
2024/01/02 01:50:20 投票者が結果を確認しています
2024/01/02 01:50:35 投票者の感心がなくなりました
2024/01/02 01:50:35 {"score": 19292, "success": 16996, "failure": 0}

~/ghq/github.com/mickamy/ISHOCON2 imp1* 1m 2s
```

- unicorn: 24 多分意味ない

```
docker exec -i ishocon2-bench-1 sh -c "./benchmark --ip app:443 --workload 4"
2024/01/02 01:52:15 Start GET /initialize
2024/01/02 01:52:15 期日前投票を開始します
2024/01/02 01:52:15 期日前投票が終了しました
2024/01/02 01:52:15 投票を開始します  Workload: 4
2024/01/02 01:53:02 投票が終了しました
2024/01/02 01:53:02 投票者が結果を確認しています
2024/01/02 01:53:17 投票者の感心がなくなりました
2024/01/02 01:53:17 {"score": 19568, "success": 17384, "failure": 0}
```

- css 配信を nginx に任せる 

```
docker exec -i ishocon2-bench-1 sh -c "./benchmark --ip app:443 --workload 4"
2024/01/02 01:58:22 Start GET /initialize
2024/01/02 01:58:22 期日前投票を開始します
2024/01/02 01:58:23 期日前投票が終了しました
2024/01/02 01:58:23 投票を開始します  Workload: 4
2024/01/02 01:59:09 投票が終了しました
2024/01/02 01:59:09 投票者が結果を確認しています
2024/01/02 01:59:25 投票者の感心がなくなりました
2024/01/02 01:59:25 {"score": 19652, "success": 17476, "failure": 0}
```

- unicorn: 16, workload: 6 24094

```
❯ make bench
docker exec -i ishocon2-bench-1 sh -c "./benchmark --ip app:443 --workload 6"
2024/01/02 02:02:20 Start GET /initialize
2024/01/02 02:02:20 期日前投票を開始します
2024/01/02 02:02:26 期日前投票が終了しました
2024/01/02 02:02:26 投票を開始します  Workload: 6
2024/01/02 02:03:12 投票が終了しました
2024/01/02 02:03:12 投票者が結果を確認しています
2024/01/02 02:03:28 投票者の感心がなくなりました
2024/01/02 02:03:28 {"score": 24094, "success": 21822, "failure": 0}
```

- candidates をメモリに載せる 21250

```
❯ make bench
docker exec -i ishocon2-bench-1 sh -c "./benchmark --ip app:443 --workload 6"
2024/01/02 02:40:03 Start GET /initialize
2024/01/02 02:40:03 期日前投票を開始します
2024/01/02 02:40:04 期日前投票が終了しました
2024/01/02 02:40:04 投票を開始します  Workload: 6
2024/01/02 02:40:51 投票が終了しました
2024/01/02 02:40:51 投票者が結果を確認しています
2024/01/02 02:41:06 投票者の感心がなくなりました
2024/01/02 02:41:06 {"score": 21250, "success": 19650, "failure": 0}
```

- votes を bulk insert 20114

```
❯ make bench
docker exec -i ishocon2-bench-1 sh -c "./benchmark --ip app:443 --workload 6"
2024/01/02 03:49:19 Start GET /initialize
2024/01/02 03:49:19 期日前投票を開始します
2024/01/02 03:49:20 期日前投票が終了しました
2024/01/02 03:49:20 投票を開始します  Workload: 6
2024/01/02 03:50:05 投票が終了しました
2024/01/02 03:50:05 投票者が結果を確認しています
2024/01/02 03:50:22 投票者の感心がなくなりました
2024/01/02 03:50:22 {"score": 20114, "success": 19682, "failure": 0}
```

- votes に count column を追加して、投票数を保存 27926

```
docker exec -i ishocon2-bench-1 sh -c "./benchmark --ip app:443 --workload 6"
2024/01/02 04:19:42 Start GET /initialize
2024/01/02 04:19:42 期日前投票を開始します
2024/01/02 04:19:43 期日前投票が終了しました
2024/01/02 04:19:43 投票を開始します  Workload: 6
2024/01/02 04:20:29 投票が終了しました
2024/01/02 04:20:29 投票者が結果を確認しています
2024/01/02 04:20:44 投票者の感心がなくなりました
2024/01/02 04:20:44 {"score": 27926, "success": 25238, "failure": 0}
```

- 遅そうなクエリの explain

```
mysql> EXPLAIN SELECT c.id, c.name, c.political_party, c.sex, v.count
    -> FROM candidates AS c
    -> LEFT OUTER JOIN
    ->   (SELECT candidate_id, IFNULL(SUM(count), 0) AS count
    ->   FROM votes
    ->   GROUP BY candidate_id) AS v
    -> ON c.id = v.candidate_id
    -> ORDER BY v.count DESC;
+----+-------------+------------+------------+-------+------------------------+------------------------+---------+---------------+-------+----------+---------------------------------+
| id | select_type | table      | partitions | type  | possible_keys          | key                    | key_len | ref           | rows  | filtered | Extra                           |
+----+-------------+------------+------------+-------+------------------------+------------------------+---------+---------------+-------+----------+---------------------------------+
|  1 | PRIMARY     | c          | NULL       | ALL   | NULL                   | NULL                   | NULL    | NULL          |    30 |   100.00 | Using temporary; Using filesort |
|  1 | PRIMARY     | <derived2> | NULL       | ref   | <auto_key0>            | <auto_key0>            | 4       | ishocon2.c.id |   128 |   100.00 | NULL                            |
|  2 | DERIVED     | votes      | NULL       | index | idx_votes_candidate_id | idx_votes_candidate_id | 4       | NULL          | 12863 |   100.00 | NULL                            |
+----+-------------+------------+------------+-------+------------------------+------------------------+---------+---------------+-------+----------+---------------------------------+
```

- add foreign key to votes to candidates 28840

```
❯ make bench
docker exec -i ishocon2-bench-1 sh -c "./benchmark --ip app:443 --workload 6"
2024/01/02 04:26:26 Start GET /initialize
2024/01/02 04:26:26 期日前投票を開始します
2024/01/02 04:26:27 期日前投票が終了しました
2024/01/02 04:26:27 投票を開始します  Workload: 6
2024/01/02 04:27:13 投票が終了しました
2024/01/02 04:27:13 投票者が結果を確認しています
2024/01/02 04:27:28 投票者の感心がなくなりました
2024/01/02 04:27:28 {"score": 28840, "success": 26120, "failure": 0}
```

- election_results の複数回コールを止める 33168

```
❯ make bench
docker exec -i ishocon2-bench-1 sh -c "./benchmark --ip app:443 --workload 6"
2024/01/02 04:29:00 Start GET /initialize
2024/01/02 04:29:01 期日前投票を開始します
2024/01/02 04:29:02 期日前投票が終了しました
2024/01/02 04:29:02 投票を開始します  Workload: 6
2024/01/02 04:29:47 投票が終了しました
2024/01/02 04:29:47 投票者が結果を確認しています
2024/01/02 04:30:02 投票者の感心がなくなりました
2024/01/02 04:30:02 {"score": 33168, "success": 28584, "failure": 0}
```

- ALTER TABLE votes ADD INDEX idx_votes_candidate_and_count (candidate_id, count); 33718

```
❯ make bench
docker exec -i ishocon2-bench-1 sh -c "./benchmark --ip app:443 --workload 6"
2024/01/02 04:39:22 Start GET /initialize
2024/01/02 04:39:22 期日前投票を開始します
2024/01/02 04:39:23 期日前投票が終了しました
2024/01/02 04:39:23 投票を開始します  Workload: 6
2024/01/02 04:40:08 投票が終了しました
2024/01/02 04:40:08 投票者が結果を確認しています
2024/01/02 04:40:23 投票者の感心がなくなりました
2024/01/02 04:40:23 {"score": 33718, "success": 27934, "failure": 0}
```

- YJIT 有効化（なぜかローカルだと遅くなったけど速くなるはずなので CI で確認） 19963

```
❯ make bench
docker exec -i ishocon2-bench-1 sh -c "./benchmark --ip app:443 --workload 6"
2024/01/02 05:08:31 Start GET /initialize
2024/01/02 05:08:32 期日前投票を開始します
2024/01/02 05:08:34 期日前投票が終了しました
2024/01/02 05:08:34 投票を開始します  Workload: 6
2024/01/02 05:09:19 投票が終了しました
2024/01/02 05:09:19 投票者が結果を確認しています
2024/01/02 05:09:35 投票者の感心がなくなりました
2024/01/02 05:09:35 {"score": 19963, "success": 16931, "failure": 0}
```

- 静的ファイルが nginx で返せていないようなので revisit（なんだかローカルが遅い） 22143

```
❯ make bench
docker exec -i ishocon2-bench-1 sh -c "./benchmark --ip app:443 --workload 6"
2024/01/02 05:41:38 Start GET /initialize
2024/01/02 05:41:38 期日前投票を開始します
2024/01/02 05:41:40 期日前投票が終了しました
2024/01/02 05:41:40 投票を開始します  Workload: 6
2024/01/02 05:42:26 投票が終了しました
2024/01/02 05:42:26 投票者が結果を確認しています
2024/01/02 05:42:41 投票者の感心がなくなりました
2024/01/02 05:42:41 {"score": 22143, "success": 18471, "failure": 0}
```

- nginx のログを捨てる 20725

```
❯ make bench
docker exec -i ishocon2-bench-1 sh -c "./benchmark --ip app:443 --workload 6"
2024/01/02 05:50:39 Start GET /initialize
2024/01/02 05:50:39 期日前投票を開始します
2024/01/02 05:50:41 期日前投票が終了しました
2024/01/02 05:50:41 投票を開始します  Workload: 6
2024/01/02 05:51:26 投票が終了しました
2024/01/02 05:51:26 投票者が結果を確認しています
2024/01/02 05:51:42 投票者の感心がなくなりました
2024/01/02 05:51:42 {"score": 20725, "success": 16637, "failure": 0}
```

- 最新の ruby を使用（3.2.2） 31606

```
❯ make bench
docker exec -i ishocon2-bench-1 sh -c "./benchmark --ip app:443 --workload 6"
2024/01/02 06:41:24 Start GET /initialize
2024/01/02 06:41:24 期日前投票を開始します
2024/01/02 06:41:25 期日前投票が終了しました
2024/01/02 06:41:25 投票を開始します  Workload: 6
2024/01/02 06:42:11 投票が終了しました
2024/01/02 06:42:11 投票者が結果を確認しています
2024/01/02 06:42:26 投票者の感心がなくなりました
2024/01/02 06:42:26 {"score": 31606, "success": 26078, "failure": 0}
```

- keyword をテーブルにする 28084

```
❯ make bench
docker exec -i ishocon2-bench-1 sh -c "./benchmark --ip app:443 --workload 6"
2024/01/02 07:23:59 Start GET /initialize
2024/01/02 07:23:59 期日前投票を開始します
2024/01/02 07:24:00 期日前投票が終了しました
2024/01/02 07:24:00 投票を開始します  Workload: 6
2024/01/02 07:24:46 投票が終了しました
2024/01/02 07:24:46 投票者が結果を確認しています
2024/01/02 07:25:23 投票者の感心がなくなりました
2024/01/02 07:25:23 {"score": 28084, "success": 23692, "failure": 0}
```

- add desc to count index 29668

```
❯ make bench
docker exec -i ishocon2-bench-1 sh -c "./benchmark --ip app:443 --workload 6"
2024/01/02 07:42:14 Start GET /initialize
2024/01/02 07:42:14 期日前投票を開始します
2024/01/02 07:42:16 期日前投票が終了しました
2024/01/02 07:42:16 投票を開始します  Workload: 6
2024/01/02 07:43:02 投票が終了しました
2024/01/02 07:43:02 投票者が結果を確認しています
2024/01/02 07:44:00 投票者の感心がなくなりました
2024/01/02 07:44:00 {"score": 29668, "success": 25884, "failure": 0}
```

- set content max age to /, /candidates/:id and /political_parties/:name 30742

```
docker exec -i ishocon2-bench-1 sh -c "./benchmark --ip app:443 --workload 6"
2024/01/02 07:52:22 Start GET /initialize
2024/01/02 07:52:23 期日前投票を開始します
2024/01/02 07:52:24 期日前投票が終了しました
2024/01/02 07:52:24 投票を開始します  Workload: 6
2024/01/02 07:53:10 投票が終了しました
2024/01/02 07:53:10 投票者が結果を確認しています
2024/01/02 07:54:03 投票者の感心がなくなりました
2024/01/02 07:54:03 {"score": 30742, "success": 26046, "failure": 0}
```

- cache at nginx level 19152

```
❯ make bench
docker exec -i ishocon2-bench-1 sh -c "./benchmark --ip app:443 --workload 6"
2024/01/02 08:43:20 Start GET /initialize
2024/01/02 08:43:21 期日前投票を開始します
2024/01/02 08:43:36 期日前投票が終了しました
2024/01/02 08:43:36 投票を開始します  Workload: 6
2024/01/02 08:44:22 投票が終了しました
2024/01/02 08:44:22 投票者が結果を確認しています
2024/01/02 08:44:47 投票者の感心がなくなりました
2024/01/02 08:44:47 {"score": 19152, "success": 16176, "failure": 0}
```

- increase nginx worker process 20122

```
❯ make bench
docker exec -i ishocon2-bench-1 sh -c "./benchmark --ip app:443 --workload 6"
2024/01/02 08:50:34 Start GET /initialize
2024/01/02 08:50:34 期日前投票を開始します
2024/01/02 08:50:36 期日前投票が終了しました
2024/01/02 08:50:36 投票を開始します  Workload: 6
2024/01/02 08:51:21 投票が終了しました
2024/01/02 08:51:21 投票者が結果を確認しています
2024/01/02 08:51:43 投票者の感心がなくなりました
2024/01/02 08:51:43 {"score": 20122, "success": 17186, "failure": 0}
```

- workload を 8 に 18968

```
❯ make bench
docker exec -i ishocon2-bench-1 sh -c "./benchmark --ip app:443 --workload 6"
2024/01/02 09:10:56 Start GET /initialize
2024/01/02 09:10:56 期日前投票を開始します
2024/01/02 09:10:57 期日前投票が終了しました
2024/01/02 09:10:57 投票を開始します  Workload: 6
2024/01/02 09:11:43 投票が終了しました
2024/01/02 09:11:43 投票者が結果を確認しています
2024/01/02 09:12:13 投票者の感心がなくなりました
2024/01/02 09:12:13 {"score": 18968, "success": 16184, "failure": 0}
```

- more cache on nginx 20442

```
❯ make bench
docker exec -i ishocon2-bench-1 sh -c "./benchmark --ip app:443 --workload 6"
2024/01/02 09:41:10 Start GET /initialize
2024/01/02 09:41:10 期日前投票を開始します
2024/01/02 09:41:21 期日前投票が終了しました
2024/01/02 09:41:21 投票を開始します  Workload: 6
2024/01/02 09:42:07 投票が終了しました
2024/01/02 09:42:07 投票者が結果を確認しています
2024/01/02 09:42:30 投票者の感心がなくなりました
2024/01/02 09:42:30 {"score": 20442, "success": 15346, "failure": 0}
```

---

以下、他の方のブログ等を読んでの対応

---

- 前回の状態で再度ベンチ 36352

```
❯ make bench
docker exec -i ishocon2-bench-1 sh -c "./benchmark --ip app:443 --workload 6"
2024/01/02 23:23:16 Start GET /initialize
2024/01/02 23:23:17 期日前投票を開始します
2024/01/02 23:23:17 期日前投票が終了しました
2024/01/02 23:23:17 投票を開始します  Workload: 6
2024/01/02 23:24:03 投票が終了しました
2024/01/02 23:24:03 投票者が結果を確認しています
2024/01/02 23:24:31 投票者の感心がなくなりました
2024/01/02 23:24:31 {"score": 36352, "success": 28576, "failure": 0}
```

- nginx の worker process は CPU と同じがいいらしい 35176

```
ishocon@8bb736ded6ff:~$ lscpu | grep CPU
CPU 操作モード:                      32-bit
CPU:                                 8
オンラインになっている CPU のリスト: 0-7
```

```
❯ make bench
docker exec -i ishocon2-bench-1 sh -c "./benchmark --ip app:443 --workload 6"
2024/01/02 23:30:27 Start GET /initialize
2024/01/02 23:30:27 期日前投票を開始します
2024/01/02 23:30:28 期日前投票が終了しました
2024/01/02 23:30:28 投票を開始します  Workload: 6
2024/01/02 23:31:13 投票が終了しました
2024/01/02 23:31:13 投票者が結果を確認しています
2024/01/02 23:32:20 投票者の感心がなくなりました
2024/01/02 23:32:20 {"score": 35176, "success": 29888, "failure": 0}
```

- unicorn process 調整（とりあえず 8）

nginx に到達できていないっぽい

```
2024/01/02 23:38:02 Get https://app:443/css/bootstrap.min.css: dial tcp 172.18.0.2:443: connect: cannot assign requested address
2024/01/02 23:38:02 Get https://app:443/css/bootstrap.min.css: dial tcp 172.18.0.2:443: connect: cannot assign requested address
2024/01/02 23:38:02 Get https://app:443/css/bootstrap.min.css: dial tcp 172.18.0.2:443: connect: cannot assign requested address
2024/01/02 23:38:02 Get https://app:443/css/bootstrap.min.css: dial tcp 172.18.0.2:443: connect: cannot assign requested address
2024/01/02 23:38:02 Get https://app:443/css/bootstrap.min.css: dial tcp 172.18.0.2:443: connect: cannot assign requested address
2024/01/02 23:38:02 Get https://app:443/: dial tcp 172.18.0.2:443: connect: cannot assign requested address
2024/01/02 23:38:02 Get https://app:443/: dial tcp 172.18.0.2:443: connect: cannot assign requested address
2024/01/02 23:38:02 Get https://app:443/candidates/19: dial tcp 172.18.0.2:443: connect: cannot assign requested address
2024/01/02 23:38:02 Get https://app:443/css/bootstrap.min.css: dial tcp 172.18.0.2:443: connect: cannot assign requested address
2024/01/02 23:38:02 Get https://app:443/css/bootstrap.min.css: dial tcp 172.18.0.2:443: connect: cannot assign requested address
2024/01/02 23:38:02 Get https://app:443/css/bootstrap.min.css: dial tcp 172.18.0.2:443: connect: cannot assign requested address
2024/01/02 23:38:02 Get https://app:443/: dial tcp 172.18.0.2:443: connect: cannot assign requested address
2024/01/02 23:38:02 Get https://app:443/css/bootstrap.min.css: dial tcp 172.18.0.2:443: connect: cannot assign requested address
2024/01/02 23:38:32 投票者の感心がなくなりました
2024/01/02 23:38:32 {"score": -6410, "success": 31845, "failure": 451}
```

- workload を 4 に変更 43824

```
❯ make bench
docker exec -i ishocon2-bench-1 sh -c "./benchmark --ip app:443 --workload 4"
2024/01/02 23:41:41 Start GET /initialize
2024/01/02 23:41:41 期日前投票を開始します
2024/01/02 23:41:42 期日前投票が終了しました
2024/01/02 23:41:42 投票を開始します  Workload: 4
2024/01/02 23:42:27 投票が終了しました
2024/01/02 23:42:27 投票者が結果を確認しています
2024/01/02 23:43:08 投票者の感心がなくなりました
2024/01/02 23:43:08 {"score": 43824, "success": 33112, "failure": 0}
```

- http2 対応 43136

```
❯ make bench
docker exec -i ishocon2-bench-1 sh -c "./benchmark --ip app:443 --workload 4"
2024/01/02 23:51:52 Start GET /initialize
2024/01/02 23:51:52 期日前投票を開始します
2024/01/02 23:51:53 期日前投票が終了しました
2024/01/02 23:51:53 投票を開始します  Workload: 4
2024/01/02 23:52:38 投票が終了しました
2024/01/02 23:52:38 投票者が結果を確認しています
2024/01/02 23:53:05 投票者の感心がなくなりました
2024/01/02 23:53:05 {"score": 43136, "success": 32568, "failure": 0}
```

- select * の見直し 37208

```
❯ make bench
docker exec -i ishocon2-bench-1 sh -c "./benchmark --ip app:443 --workload 4"
2024/01/02 23:55:51 Start GET /initialize
2024/01/02 23:55:51 期日前投票を開始します
2024/01/02 23:55:52 期日前投票が終了しました
2024/01/02 23:55:52 投票を開始します  Workload: 4
2024/01/02 23:56:37 投票が終了しました
2024/01/02 23:56:37 投票者が結果を確認しています
2024/01/02 23:57:01 投票者の感心がなくなりました
2024/01/02 23:57:01 {"score": 37208, "success": 27504, "failure": 0}
```
