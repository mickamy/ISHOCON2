初回ベンチ（DB index などは ruby から引き継ぎ） 357969

```
❯ make bench
docker exec -i ishocon2-bench-1 sh -c "./benchmark --ip app:443 --workload 7"
2024/01/09 09:42:23 Start GET /initialize
2024/01/09 09:42:23 期日前投票を開始します
2024/01/09 09:42:24 期日前投票が終了しました
2024/01/09 09:42:24 投票を開始します  Workload: 7
2024/01/09 09:43:09 投票が終了しました
2024/01/09 09:43:09 投票者が結果を確認しています
2024/01/09 09:43:24 投票者の感心がなくなりました
2024/01/09 09:43:24 {"score": 357969, "success": 234209, "failure": 0}
```

alp の結果。go が速すぎて笑う

```
ishocon@2ee24875c3fc:~/scripts$ run_pt-query.sh
+ rm -rf sloq-query.log.analyzed
+ pt-query-digest --type slowlog /var/log/mysql/slow-query.log
^C# Caught SIGINT.
^C# Exiting on SIGINT.
ishocon@2ee24875c3fc:~/scripts$ exit

~/ghq/github.com/mickamy/ISHOCON2 try-with-go-lang* 4m 18s
❯ docker exec -it ishocon2-app-1 bash
ishocon@420e599cf577:~$ htop
ishocon@420e599cf577:~$ cd scripts/
ishocon@420e599cf577:~/scripts$ ls
modify_db_schema.sql  run_alp.sh  run_pt-query.sh
ishocon@420e599cf577:~/scripts$ ./run_alp.sh
+ sudo cat /var/log/nginx/access.log
+ alp ltsv --sort sum -o count,2xx,3xx,4xx,5xx,method,uri,avg,sum
+--------+--------+-----+-----+-----+--------+--------------------------------------------------------------------------------------+-------+---------+
| COUNT  |  2XX   | 3XX | 4XX | 5XX | METHOD |                                         URI                                          |  AVG  |   SUM   |
+--------+--------+-----+-----+-----+--------+--------------------------------------------------------------------------------------+-------+---------+
| 1000   | 1000   | 0   | 0   | 0   | GET    | /candidates/20                                                                       | 0.000 | 0.001   |
| 976    | 976    | 0   | 0   | 0   | GET    | /candidates/23                                                                       | 0.000 | 0.001   |
| 6970   | 6970   | 0   | 0   | 0   | GET    | /political_parties/%E5%9B%BD%E6%B0%91%E5%85%83%E6%B0%97%E5%85%9A                     | 0.000 | 0.001   |
| 920    | 920    | 0   | 0   | 0   | GET    | /candidates/24                                                                       | 0.000 | 0.003   |
| 1      | 1      | 0   | 0   | 0   | GET    | /initialize                                                                          | 0.004 | 0.004   |
| 961    | 961    | 0   | 0   | 0   | GET    | /candidates/13                                                                       | 0.000 | 0.004   |
| 903    | 903    | 0   | 0   | 0   | GET    | /candidates/29                                                                       | 0.000 | 0.004   |
| 937    | 937    | 0   | 0   | 0   | GET    | /candidates/8                                                                        | 0.000 | 0.005   |
| 929    | 929    | 0   | 0   | 0   | GET    | /candidates/10                                                                       | 0.000 | 0.005   |
| 933    | 933    | 0   | 0   | 0   | GET    | /candidates/14                                                                       | 0.000 | 0.006   |
| 930    | 930    | 0   | 0   | 0   | GET    | /candidates/30                                                                       | 0.000 | 0.006   |
| 965    | 965    | 0   | 0   | 0   | GET    | /candidates/25                                                                       | 0.000 | 0.007   |
| 913    | 913    | 0   | 0   | 0   | GET    | /candidates/17                                                                       | 0.000 | 0.007   |
| 926    | 926    | 0   | 0   | 0   | GET    | /candidates/12                                                                       | 0.000 | 0.007   |
| 943    | 943    | 0   | 0   | 0   | GET    | /candidates/9                                                                        | 0.000 | 0.009   |
| 947    | 947    | 0   | 0   | 0   | GET    | /candidates/5                                                                        | 0.000 | 0.010   |
| 897    | 897    | 0   | 0   | 0   | GET    | /candidates/27                                                                       | 0.000 | 0.010   |
| 905    | 905    | 0   | 0   | 0   | GET    | /candidates/22                                                                       | 0.000 | 0.011   |
| 959    | 959    | 0   | 0   | 0   | GET    | /candidates/26                                                                       | 0.000 | 0.012   |
| 940    | 940    | 0   | 0   | 0   | GET    | /candidates/6                                                                        | 0.000 | 0.014   |
| 894    | 894    | 0   | 0   | 0   | GET    | /candidates/15                                                                       | 0.000 | 0.014   |
| 933    | 933    | 0   | 0   | 0   | GET    | /candidates/7                                                                        | 0.000 | 0.015   |
| 931    | 931    | 0   | 0   | 0   | GET    | /candidates/2                                                                        | 0.000 | 0.015   |
| 928    | 928    | 0   | 0   | 0   | GET    | /candidates/1                                                                        | 0.000 | 0.016   |
| 876    | 876    | 0   | 0   | 0   | GET    | /candidates/18                                                                       | 0.000 | 0.017   |
| 956    | 956    | 0   | 0   | 0   | GET    | /candidates/11                                                                       | 0.000 | 0.017   |
| 897    | 897    | 0   | 0   | 0   | GET    | /candidates/19                                                                       | 0.000 | 0.017   |
| 994    | 994    | 0   | 0   | 0   | GET    | /candidates/3                                                                        | 0.000 | 0.019   |
| 921    | 921    | 0   | 0   | 0   | GET    | /candidates/16                                                                       | 0.000 | 0.022   |
| 917    | 917    | 0   | 0   | 0   | GET    | /candidates/28                                                                       | 0.000 | 0.023   |
| 925    | 925    | 0   | 0   | 0   | GET    | /candidates/21                                                                       | 0.000 | 0.030   |
| 6874   | 6874   | 0   | 0   | 0   | GET    | /political_parties/%E5%9B%BD%E6%B0%91%E5%B9%B3%E5%92%8C%E5%85%9A                     | 0.000 | 0.058   |
| 910    | 910    | 0   | 0   | 0   | GET    | /candidates/4                                                                        | 0.000 | 0.073   |
| 7155   | 7155   | 0   | 0   | 0   | GET    | /political_parties/%E5%A4%A2%E5%AE%9F%E7%8F%BE%E5%85%9A                              | 0.000 | 0.122   |
| 7066   | 7066   | 0   | 0   | 0   | GET    | /political_parties/%E5%9B%BD%E6%B0%9110%E4%BA%BA%E5%A4%A7%E6%B4%BB%E8%BA%8D%E5%85%9A | 0.000 | 0.228   |
| 60344  | 60344  | 0   | 0   | 0   | GET    | /css/bootstrap.min.css                                                               | 0.000 | 0.342   |
| 4317   | 4317   | 0   | 0   | 0   | GET    | /                                                                                    | 0.016 | 70.004  |
| 102955 | 102955 | 0   | 0   | 0   | POST   | /vote                                                                                | 0.003 | 314.879 |
+--------+--------+-----+-----+-----+--------+--------------------------------------------------------------------------------------+-------+---------+
```

candidates のキャッシュをしようとする 358429

```
❯ make bench
docker exec -i ishocon2-bench-1 sh -c "./benchmark --ip app:443 --workload 7"
2024/01/09 10:48:10 Start GET /initialize
2024/01/09 10:48:10 期日前投票を開始します
2024/01/09 10:48:10 期日前投票が終了しました
2024/01/09 10:48:10 投票を開始します  Workload: 7
2024/01/09 10:48:55 投票が終了しました
2024/01/09 10:48:55 投票者が結果を確認しています
2024/01/09 10:49:10 投票者の感心がなくなりました
2024/01/09 10:49:10 {"score": 358429, "success": 236589, "failure": 0}
```

GIN をリリースモードで起動 322438

```
❯ make bench
docker exec -i ishocon2-bench-1 sh -c "./benchmark --ip app:443 --workload 7"
2024/01/09 10:51:49 Start GET /initialize
2024/01/09 10:51:49 期日前投票を開始します
2024/01/09 10:51:50 期日前投票が終了しました
2024/01/09 10:51:50 投票を開始します  Workload: 7
2024/01/09 10:52:35 投票が終了しました
2024/01/09 10:52:35 投票者が結果を確認しています
2024/01/09 10:52:50 投票者の感心がなくなりました
2024/01/09 10:52:50 {"score": 322438, "success": 216894, "failure": 0}
```

go: fix setHTMLTemplate is not thread safe warning 371622

```
❯ make bench
docker exec -i ishocon2-bench-1 sh -c "./benchmark --ip app:443 --workload 7"
2024/01/10 09:01:50 Start GET /initialize
2024/01/10 09:01:50 期日前投票を開始します
2024/01/10 09:01:51 期日前投票が終了しました
2024/01/10 09:01:51 投票を開始します  Workload: 7
2024/01/10 09:02:36 投票が終了しました
2024/01/10 09:02:36 投票者が結果を確認しています
2024/01/10 09:02:51 投票者の感心がなくなりました
2024/01/10 09:02:51 {"score": 371622, "success": 246886, "failure": 0}
```

go: run in release mode of gin 381180

```
❯ make bench
docker exec -i ishocon2-bench-1 sh -c "./benchmark --ip app:443 --workload 7"
2024/01/10 09:19:10 Start GET /initialize
2024/01/10 09:19:10 期日前投票を開始します
2024/01/10 09:19:11 期日前投票が終了しました
2024/01/10 09:19:11 投票を開始します  Workload: 7
2024/01/10 09:19:56 投票が終了しました
2024/01/10 09:19:56 投票者が結果を確認しています
2024/01/10 09:20:11 投票者の感心がなくなりました
2024/01/10 09:20:11 {"score": 381180, "success": 256388, "failure": 0}
```

middleware: disable logs 344721

```
❯ make bench
docker exec -i ishocon2-bench-1 sh -c "./benchmark --ip app:443 --workload 7"
2024/01/10 10:59:18 期日前投票を開始します
2024/01/10 10:59:19 期日前投票が終了しました
2024/01/10 10:59:19 投票を開始します  Workload: 7
2024/01/10 11:04:35 投票が終了しました
2024/01/10 11:04:35 投票者が結果を確認しています
2024/01/10 11:04:50 投票者の感心がなくなりました
2024/01/10 11:04:50 {"score": 344721, "success": 233409, "failure": 0}
```

content_keyword の導入 and increase connection pools of DB to 12 361385

```
❯ make bench
docker exec -i ishocon2-bench-1 sh -c "./benchmark --ip app:443 --workload 7"
2024/01/10 11:31:22 Start GET /initialize
2024/01/10 11:31:22 期日前投票を開始します
2024/01/10 11:31:22 期日前投票が終了しました
2024/01/10 11:31:22 投票を開始します  Workload: 7
2024/01/10 11:32:08 投票が終了しました
2024/01/10 11:32:08 投票者が結果を確認しています
2024/01/10 11:32:23 投票者の感心がなくなりました
2024/01/10 11:32:23 {"score": 361385, "success": 240017, "failure": 0}
```

disable gin logging 422739

```
❯ make bench
docker exec -i ishocon2-bench-1 sh -c "./benchmark --ip app:443 --workload 7"
2024/01/10 11:35:30 Start GET /initialize
2024/01/10 11:35:30 期日前投票を開始します
2024/01/10 11:35:30 期日前投票が終了しました
2024/01/10 11:35:30 投票を開始します  Workload: 7
2024/01/10 11:36:15 投票が終了しました
2024/01/10 11:36:15 投票者が結果を確認しています
2024/01/10 11:36:30 投票者の感心がなくなりました
2024/01/10 11:36:30 {"score": 422739, "success": 293819, "failure": 0}
```

disable static 410681

```
❯ make bench
docker exec -i ishocon2-bench-1 sh -c "./benchmark --ip app:443 --workload 7"
2024/01/10 12:11:46 Start GET /initialize
2024/01/10 12:11:46 期日前投票を開始します
2024/01/10 12:11:46 期日前投票が終了しました
2024/01/10 12:11:46 投票を開始します  Workload: 7
2024/01/10 12:12:31 投票が終了しました
2024/01/10 12:12:31 投票者が結果を確認しています
2024/01/10 12:12:46 投票者の感心がなくなりました
2024/01/10 12:12:46 {"score": 410681, "success": 280665, "failure": 0}
```
