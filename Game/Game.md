#### 分布式游戏服务器 用户连接信息如何管理





#### 非对称加密 应用于保存用户连接信息

#### 弱交互的卡牌等游戏 

client登陆  server用privatekey 对(uid + 时间戳)加密得到pass-key , 返回给client
之后都用http通讯 客户端保存pass-key到内存  服务器不需要保存 因为每次都可以利用私钥解密 
确认多次连接的 用户身份 