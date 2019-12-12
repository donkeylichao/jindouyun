### jindouyun 教程 ###

#### 安装 ####
1. 安装go环境(已安装环境忽略)

- [go下载地址](https://studygolang.com/dl)
- 选择Apple macOS软件下载点击安装(其他系统自行选择合适版本)
- 安装完成运行如下命令查看是否安装成功
```
go version
```
2. 下载代码
```
git clone git@github.com:donkeylichao/jindouyun.git
```
- 运行如下命令编译软件
```
go build main.go
```
- 编译后当前目录生成main可执行文件，通过如下命令运行程序，然后根据提示操作。
```
./main
```
> <font color=red>配置文件必须和可执行文件在同一目录下，必须使用当前目录下的方式执行，否则无法正确加载配置文件</font>

3. 配置文件配置方式
```
{
  "address": "", //金斗云接口地址
  "app_id": "", //金斗云appid
  "app_key": "", //金斗云appkey
  "user":"", //保险公司账号
  "pass":"", //保险公司账号密码
  "proxy_id":"" //代理地址id
}
```
> 当前只支持代理操作和人保账号操作，太平洋和平安没开发完成。

