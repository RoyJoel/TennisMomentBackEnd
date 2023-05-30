# Details

Date : 2023-05-15 21:51:57

Directory /Users/jasonzhang/TennisMoment/TennisMoment_backend

Total : 56 files,  5108 codes, 1209 comments, 825 blanks, all 7142 lines

[Summary](results.md) / Details / [Diff Summary](diff.md) / [Diff Details](diff-details.md)

## Files
| filename | language | code | comment | blank | total |
| :--- | :--- | ---: | ---: | ---: | ---: |
| [Makefile](/Makefile) | Makefile | 8 | 0 | 3 | 11 |
| [README.md](/README.md) | Markdown | 2 | 0 | 1 | 3 |
| [app.properties](/app.properties) | Properties | 10 | 0 | 1 | 11 |
| [cmd/grpc/main.go](/cmd/grpc/main.go) | Go | 5 | 0 | 3 | 8 |
| [cmd/web/main.go](/cmd/web/main.go) | Go | 5 | 0 | 3 | 8 |
| [go.mod](/go.mod) | Go Module File | 60 | 0 | 4 | 64 |
| [go.sum](/go.sum) | Go Checksum File | 1,069 | 0 | 1 | 1,070 |
| [main 21.54.44.go](/main%2021.54.44.go) | Go | 10 | 0 | 3 | 13 |
| [package/cache/TennisMomentCache.go](/package/cache/TennisMomentCache.go) | Go | 44 | 5 | 10 | 59 |
| [package/config/mysql.go](/package/config/mysql.go) | Go | 40 | 0 | 5 | 45 |
| [package/config/redis.go](/package/config/redis.go) | Go | 14 | 0 | 5 | 19 |
| [package/config/web.go](/package/config/web.go) | Go | 21 | 0 | 5 | 26 |
| [package/dao/impl/TMImpl.go](/package/dao/impl/TMImpl.go) | Go | 772 | 5 | 126 | 903 |
| [package/middleware/middleware.go](/package/middleware/middleware.go) | Go | 42 | 28 | 12 | 82 |
| [package/model/Address.go](/package/model/Address.go) | Go | 68 | 32 | 16 | 116 |
| [package/model/Bill.go](/package/model/Bill.go) | Go | 49 | 2 | 11 | 62 |
| [package/model/Cart.go](/package/model/Cart.go) | Go | 23 | 1 | 7 | 31 |
| [package/model/Club.go](/package/model/Club.go) | Go | 59 | 2 | 12 | 73 |
| [package/model/ClubEvent.go](/package/model/ClubEvent.go) | Go | 23 | 1 | 7 | 31 |
| [package/model/ClubManager.go](/package/model/ClubManager.go) | Go | 23 | 1 | 7 | 31 |
| [package/model/ClubMember.go](/package/model/ClubMember.go) | Go | 23 | 1 | 7 | 31 |
| [package/model/Commodity.go](/package/model/Commodity.go) | Go | 59 | 2 | 11 | 72 |
| [package/model/Event.go](/package/model/Event.go) | Go | 64 | 2 | 12 | 78 |
| [package/model/EventGame.go](/package/model/EventGame.go) | Go | 23 | 1 | 7 | 31 |
| [package/model/EventPlayer.go](/package/model/EventPlayer.go) | Go | 23 | 1 | 7 | 31 |
| [package/model/Game.go](/package/model/Game.go) | Go | 129 | 2 | 13 | 144 |
| [package/model/Order.go](/package/model/Order.go) | Go | 63 | 32 | 15 | 110 |
| [package/model/Player.go](/package/model/Player.go) | Go | 91 | 2 | 11 | 104 |
| [package/model/PlayerStats.go](/package/model/PlayerStats.go) | Go | 23 | 1 | 7 | 31 |
| [package/model/Relationship.go](/package/model/Relationship.go) | Go | 23 | 1 | 7 | 31 |
| [package/model/Schedule.go](/package/model/Schedule.go) | Go | 49 | 2 | 11 | 62 |
| [package/model/Stats.go](/package/model/Stats.go) | Go | 75 | 1 | 7 | 83 |
| [package/model/User.go](/package/model/User.go) | Go | 67 | 1 | 7 | 75 |
| [package/model/option.go](/package/model/option.go) | Go | 45 | 2 | 11 | 58 |
| [package/rpc/TennisMoment.go](/package/rpc/TennisMoment.go) | Go | 20 | 3 | 4 | 27 |
| [package/rpc/client/GameClient.go](/package/rpc/client/GameClient.go) | Go | 33 | 3 | 6 | 42 |
| [package/rpc/client/PlayerClient.go](/package/rpc/client/PlayerClient.go) | Go | 33 | 3 | 6 | 42 |
| [package/rpc/client/StatsClient.go](/package/rpc/client/StatsClient.go) | Go | 1 | 35 | 6 | 42 |
| [package/rpc/impl/TennisMoment.go](/package/rpc/impl/TennisMoment.go) | Go | 104 | 11 | 12 | 127 |
| [package/utils/IntMatrix.go](/package/utils/IntMatrix.go) | Go | 170 | 9 | 32 | 211 |
| [package/utils/md5.go](/package/utils/md5.go) | Go | 28 | 0 | 4 | 32 |
| [package/utils/time.go](/package/utils/time.go) | Go | 19 | 4 | 6 | 29 |
| [package/web/auth/auth.go](/package/web/auth/auth.go) | Go | 83 | 3 | 15 | 101 |
| [package/web/controller/TennisMomentControllerImpl.go](/package/web/controller/TennisMomentControllerImpl.go) | Go | 340 | 17 | 91 | 448 |
| [package/web/interceptor/interceptor.go](/package/web/interceptor/interceptor.go) | Go | 23 | 11 | 4 | 38 |
| [package/web/router.go](/package/web/router.go) | Go | 76 | 18 | 16 | 110 |
| [proto/Game.pb.go](/proto/Game.pb.go) | Go | 283 | 303 | 69 | 655 |
| [proto/Player.pb.go](/proto/Player.pb.go) | Go | 560 | 47 | 71 | 678 |
| [proto/Stats.pb.go](/proto/Stats.pb.go) | Go | 1 | 590 | 68 | 659 |
| [tennismoment_app.sql](/tennismoment_app.sql) | SQL | 195 | 1 | 18 | 214 |
| [test/TennisMoment_test.go](/test/TennisMoment_test.go) | Go | 1 | 0 | 1 | 2 |
| [test/app.properties](/test/app.properties) | Properties | 10 | 0 | 1 | 11 |
| [test/auth_test.go](/test/auth_test.go) | Go | 8 | 0 | 4 | 12 |
| [test/role_test.go](/test/role_test.go) | Go | 1 | 13 | 4 | 18 |
| [test/time_test.go](/test/time_test.go) | Go | 14 | 0 | 6 | 20 |
| [test/user_test.go](/test/user_test.go) | Go | 1 | 10 | 6 | 17 |

[Summary](results.md) / Details / [Diff Summary](diff.md) / [Diff Details](diff-details.md)