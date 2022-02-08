# 说明文档. 
### Server(服务端):  ### 

#### 目录结构: 
- data: 持久化数据 xml文件目录. 
- handler: http接口请求处理方法目录（登陆、注册、获取数据、登出）
- inits:服务初始化方法
- model:数据模型
- public:公共定义常量、变量（配置定义、错误定义）
- routes:路由
- utils:通用方法
#### 使用第三方库：
1. gin. 
2. jwt-go.  
3. crypto. 
##### go版本：1.17
运行说明：
启动前先go mod tidy 安装引用库

### Client(客户端)：
采用 react+TypeScript 开发
开启前先 npm install下载node_modules

#### 使用的第三方库：
1. react-router-dom
2. bootstrap
3. react-cookies


![](https://github.com/zadminmax/jwt_server/blob/main/C539E5D1-FC2E-4C78-9F65-409D411C42D0.jpeg?raw=true)
