# flamingo
####一个 Golang API 框架 ！按照 Resful 的请求规则你可以自己写，这是个JWT验证的登录验证框架。


未必要重新创建数据库，你也可以直接连接之前已经存在的数据库，除非有必要。

DB_MIGRATE=true（.env.example 文件里有） 则表示你需要重新迁移数据库

#
###直接用了gin框架来搭建。
###数据库部分我用的 gorm 。
#

####参考：<br/>
https://github.com/Gourouting/singo <br/>
https://github.com/Wangjiaxing123/JwtDemo