    这是个正在练习以及不断丰富的基于REST形式的api服务
使用Gin框架搭建的服务：
    Gin与node的koa框架非常相似，它在go的http包的基础上做了一层封装，感觉像express->koa的感觉。

数据库方面使用gorm操作：
    创建表，定义表关系，增删改查都非常方便。
    
身份验证jwt-go:
    基于是web服务，用jwt能够为更过的客户端提供服务
# adminvideo

models包 链接数据库 自动迁移

createtable包 定义表

operating包  对每个表的操作

router包  定义路由

v1 v2 v3 包  实现路由

public 存放静态文件

middleware包 定义中间件

seeds包 生成测试数据
