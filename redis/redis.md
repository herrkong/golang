
### redis 用作消息队列
rpush  blpop 


#### redis缓存

##### 缓存淘汰机制
LRU 最近最少使用 双向链表


##### redis缓存模式

Cache Aside模式

查询数据 先查询缓存 如果缓存不存在 会再查询数据库 获取到数据之后再更新到缓存

存在的并发风险： 查询缓存没有查到 去查询数据库 写入缓存 另外一个写操作线程修改数据库 更新这条记录 缓存中记录的是脏数据

（高并发情况下延时双删策略 写库前后都将删除缓存key）


#### redis缓存穿透
数据在数据库和缓存中都不存在 做了两次无用的查询 造成数据库压力激增 挂掉

##### 缓存穿透的解决方案
1 设置缓存空值或缺省值 查询为空时将空结果也进行缓存 过期时间设为较短 下次查询直接从缓存中取值 避免查表
2 bloomfilter布隆过滤器 存所有可能的查询 bitmap 先查key是否存在 再去查缓存


#### redis缓存击穿 
某个访问非常频繁的缓存过期了 导致查数据库

#####
对于访问非常频繁的缓存不设置过期时间



#### redis 缓存雪崩

redis中有大量缓存同时过期或者redis挂了  大量请求无法在redis中处理 发送到mysql 导致压力暴增 挂了

##### 大量缓存同时到期 导致过量请求无法处理的解决方案

1 数据预热 在大并发到来之前 主动加载rediskey 避免直接查数据库
2 设置不同的过期时间 让缓存失效的时间点分布均匀
3 双层缓存策略 原始缓存 拷贝一份 拷贝缓存 原始缓存过期时间短 拷贝缓存 过期时间长
4 服务降级 不同数据采取不同的方案  非核心数据直接返回预定义信息


#### 高性能并发指标

QPS : query per second 每秒查询

TPS : transaction per second 每秒事务

RT : Response-Time 响应时间

并发数 

吞吐量： 单个request对 cpu消耗越高  外部接口 io速度越慢 系统吞吐能力越低 


#### mysql和redis 数据一致性问题
对于大多数redis缓存都是设置为只读的 先更新数据库再删除更新redis来保证数据一致性


数据库 并发访问的薄弱环节

读数据流程：  先查redis 查不到 再查mysql 并把查询结果写redis

写数据流程 可能产生 数据不一致的情况


先更新mysql  在删除更新redis key 之前 mysql挂了 没有成功写入mysql  

先删除redis 在更新mysql之前 再来查询 缓存为空  则去数据库查并写入redis 此时redis的值为脏数据



##### 1 延时双删策略

在写库前后都进行redis.del()操作 并设置超时时间

1） 删除缓存
2）写数据库
3） 休眠一段时间
4） 删除缓存


给缓存设置过期时间 时间过期了 后面的读请求 自然会把数据库中的值填入缓存



#### redis 集群模式

##### 主从复制模式

master服务器下多个slave服务器   主服务器可以进行读写操作 从服务器只可读 只同步主服务器的数据

主服务器挂了需要人工介入 将某个从服务器切换为主服务器

##### Sentinel（哨兵）模式

哨兵作为一个独立的进程 监控主服务器和从服务器  发送命令等待服务器响应 检测到master挂了 自动切换slave和master
并通过发布订阅模式 通知其他从服务器修改配置文件 切换主机

还可以做多个哨兵 互相监控

自动容灾处理 自动切换主从


##### Cluster 集群模式 redis官方

分布式存储 每台redis服务器上都存储着不同的数据

redis节点之间都是互联互通的 
客户端直接连接其中的节点
集群中如果超半数的节点挂了 集群不可用了






