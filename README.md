# XIU2/SNIProxy

[![Go Version](https://img.shields.io/github/go-mod/go-version/XIU2/SNIProxy.svg?style=flat-square&label=Go&color=00ADD8&logo=go)](https://github.com/XIU2/SNIProxy/)
[![Release Version](https://img.shields.io/github/v/release/XIU2/SNIProxy.svg?style=flat-square&label=Release&color=00ADD8&logo=github)](https://github.com/XIU2/SNIProxy/releases/latest)
[![GitHub license](https://img.shields.io/github/license/XIU2/SNIProxy.svg?style=flat-square&label=License&color=00ADD8&logo=github)](https://github.com/XIU2/SNIProxy/)
[![GitHub Star](https://img.shields.io/github/stars/XIU2/SNIProxy.svg?style=flat-square&label=Star&color=00ADD8&logo=github)](https://github.com/XIU2/SNIProxy/)
[![GitHub Fork](https://img.shields.io/github/forks/XIU2/SNIProxy.svg?style=flat-square&label=Fork&color=00ADD8&logo=github)](https://github.com/XIU2/SNIProxy/)

🧷 自用的一个功能很简单的 SNI Proxy 顺便分享出来给有同样需求的人，用得上的话可以**点个⭐支持下~**

SNI Proxy 是一个无需加解密的反向代理工具，根据传入的 SNI(域名) 来自动转发流量至该域名的源站。

> _分享我其他开源项目：[**TrackersList.com** - 全网热门 BT Tracker 列表！有效提高 BT 下载速度~](https://github.com/XIU2/TrackersListCollection) <img src="https://img.shields.io/github/stars/XIU2/TrackersListCollection.svg?style=flat-square&label=Star&color=4285dd&logo=github" height="16px" />_  
> _[**CloudflareSpeedTest** - 🌩「自选优选 IP」测试 Cloudflare CDN 延迟和速度，获取最快 IP~](https://github.com/XIU2/CloudflareSpeedTest) <img src="https://img.shields.io/github/stars/XIU2/CloudflareSpeedTest.svg?style=flat-square&label=Star&color=4285dd&logo=github" height="16px" />_  
> _[**UserScript** - 🐵 Github 高速下载、知乎增强、自动无缝翻页、护眼模式 等十几个**油猴脚本**~](https://github.com/XIU2/UserScript) <img src="https://img.shields.io/github/stars/XIU2/UserScript.svg?style=flat-square&label=Star&color=4285dd&logo=github" height="16px" />_


****

## \# 软件介绍

1. **支持** 全平台、全系统（Go 语言特性）
2. **支持** Socks5 前置代理（比如可以套 WARP+）
3. **支持** 允许所有域名  仅允许指定域名（包含域名自身及其所有子域名）

> 注意！SNI Proxy 仅为我个人自写自用，**可靠性、稳定性**等方面**不如专业的商业软件（如 Nginx、HAProxy）**，因此在正式的**生产环境下不建议使用本软件**，如造成损失，根据 GPL-3.0 本项目无需承担责任（溜了溜了~

****

SNI Proxy 的工作流程大概如下：

1. 解析传入连接中的 TLS/SSL 握手消息，以获取客户端发送的 **SNI 域名**信息。
2. 检查域名是否在允许列表中（或开启了 `allow_all_hosts`），如果不在将中断连接，反之继续。
3. 使用系统 DNS 解析 SNI 域名获得 IP 地址（即该域名的源站服务器 IP 地址）。
4. 将流量转发给该域名的源站 **IP:443**，在客户端和源站服务器之间进行数据传输（即 TCP 中转/端口转发）。

```javascript
// 本地通过 Hosts、DNS 等方法将 example.com 域名指向 SNI Proxy 服务器的 IP，然后：
访问 example.com <=> SNIProxy <=> 源站(example.com)

// 如果 SNI Proxy 开启了前置代理，那么就是这样：
访问 example.com <=> SNIProxy <=> Socks5 <=> 源站(example.com)
```

> SNIProxy 本质也算是一种**端口转发（中转）**，但不同于端口转发只能指定一个**固定的目标 IP**，SNIProxy 可以通过 DNS 解析传入的域名来获得**灵活的目标 IP**（传入不同的域名走不同目标 IP，可**同时存在**且**互不干扰**）。

****

## \# 使用方法

<details>
<summary><code><strong>「 点击查看 Windows 系统下的使用示例 」</strong></code></summary>

****

### 下载

下载已编译好的可执行文件并解压：

1. [Github Releases](https://github.com/XIU2/SNIProxy/releases)  
2. [蓝奏云](https://pan.lanzouj.com/b077bn2ri)(密码:xiu2)

### 配置

找到配置文件 `config.yaml` 右键菜单 - 打开方式 - 记事本。

根据下面的 [配置文件说明](https://github.com/XIU2/SNIProxy#-配置文件说明-configyaml) 来自定义配置内容并保存。

### 运行

双击运行 `sniproxy.exe` 文件。

或者在 CMD 命令行中进入软件所在目录并运行 `sniproxy.exe`：

```yaml
# 进入解压后的 sniproxy 程序所在目录（记得修改下面示例路径）
cd /d C:\xxx\sniproxy

# 运行（不带参数）
sniproxy.exe

# 运行（带参数示例）
sniproxy.exe -c "config.yaml"
```
</details>

****

<details>
<summary><code><strong>「 点击查看 Linux 系统下的使用示例 」</strong></code></summary>

****

以下命令仅为示例，版本号和文件名请前往 [**Releases**](https://github.com/XIU2/SNIProxy/releases) 查看。

```yaml
# 如果是第一次使用，则建议创建新文件夹（后续更新时，跳过该步骤）
mkdir sniproxy

# 进入文件夹（后续更新，只需要从这里重复下面的下载、解压命令即可）
cd sniproxy

# 下载 sniproxy 压缩包（自行根据需求替换 URL 中 [版本号] 和 [文件名]）
wget -N https://github.com/XIU2/SNIProxy/releases/download/v1.0.0/sniproxy_linux_amd64.tar.gz
# 如果你是在国内服务器上下载，那么请使用下面这几个镜像加速：
# wget -N https://download.fastgit.org/XIU2/SNIProxy/releases/download/v1.0.0/sniproxy_linux_amd64.tar.gz
# wget -N https://ghproxy.com/https://github.com/XIU2/SNIProxy/releases/download/v1.0.0/sniproxy_linux_amd64.tar.gz
# 如果下载失败的话，尝试删除 -N 参数（如果是为了更新，则记得提前删除旧压缩包 rm sniproxy_linux_amd64.tar.gz ）

# 解压（不需要删除旧文件，会直接覆盖，自行根据需求替换 文件名）
tar -zxf sniproxy_linux_amd64.tar.gz

# 赋予执行权限
chmod +x sniproxy

# 编辑配置文件（根据下面的 配置文件说明 来自定义配置内容并保存(按下 Ctrl+X 然后再按 2 下回车)
nano config.yaml

# 运行（不带参数）
./sniproxy

# 运行（带参数示例）
./sniproxy -c "config.yaml"

# 后台运行（带参数示例）
nohup ./sniproxy -c "config.yaml" > "sni.log" 2>&1 &
```

> 另外，强烈建议顺便提高一下 [系统文件句柄数上限](https://github.com/XIU2/SNIProxy#-提高系统文件句柄数上限-避免报错-too-many-open-files)，避免遇到报错 **too many open files**  

> 另外，如果你希望 **开机启动、后台运行、方便管理** 等，那么可以将其 [注册为系统服务](https://github.com/XIU2/SNIProxy#-linux-配置为系统服务-systemd---以支持开机启动后台运行等)。

</details>

****

<details>
<summary><code><strong>「 点击查看 Mac 系统下的使用示例 」</strong></code></summary>

****

下载已编译好的可执行文件并解压：

1. [Github Releases](https://github.com/XIU2/SNIProxy/releases)  
2. [蓝奏云](https://pan.lanzouj.com/b077bn2ri)(密码:xiu2)

```yaml
# 进入 sniproxy 压缩包所在目录（记得修改下面示例路径）
cd /xxx/xxx

# 解压（不需要删除旧文件，会直接覆盖，自行根据需求替换 文件名）
tar -zxf sniproxy_linux_amd64.tar.gz

# 赋予执行权限
chmod a+x sniproxy

# 编辑配置文件（根据下面的 配置文件说明 来自定义配置内容并保存(按下 Contrl+X 然后再按 2 下回车)
nano config.yaml

# 运行（不带参数）
./sniproxy

# 运行（带参数示例）
./sniproxy -c "config.yaml"
```

</details>

****

```css
home@xiu:~# ./sniproxy -h

SNIProxy vX.X.X
https://github.com/XIU2/SNIProxy

参数：
    -c config.yaml
        配置文件 (默认 config.yaml)
    -l sni.log
        日志文件 (默认 无)
    -d
        调试模式 (默认 关)
    -v
        程序版本
    -h
        帮助说明
```

****

## \# 其他说明

#### \# 配置文件说明 (config.yaml)

<details>
<summary><code><strong>「 点击展开 查看内容 」</strong></code></summary>

****

> **注意：** 配置文件是 YAML 格式，即按照缩进（即每行前面的空格数量）来确定层级关系的，因此不懂的话请按照默认配置文件内示例的格式为准，其中 ` # ` 的是注释（会被程序忽略），不需要的配置可以注释掉。

目前配置文件中的配置项没几个，分别为：

```yaml
# 监听端口（注意需要引号），常见示例如下：
# ":443"            省略 IP 只写端口，代表监听本机所有 IPv4+IPv6 地址的 443 端口
# "0.0.0.0:443"     代表监听本机所有 IPv4 地址的 443 端口
# "127.0.0.1:443"   代表监听本机本地 IPv4 地址的 443 端口（只有本机可访问）
# "[::]:443"        代表监听本机所有 IPv6 地址的 443 端口
# "[::1]:443"       代表监听本机本地 IPv6 地址的 443 端口（只有本机可访问）
# 上面示例中的 IP 地址也可以换成例如你的外网 IP，这样的话就只能从该外网 IP 访问了
listen_addr: ":443"

# 可选：启用 Socks5 前置代理
# （启用前：访客 <=> SNIProxy <=> 目标网站
# （启用后：访客 <=> SNIProxy <=> Socks5 <=> 目标网站
# （比如可以套 WARP+，那样就变成：访客 <=> SNIProxy <=> WARP+ <=> 目标网站
enable_socks5: true
# 可选：配置 Socks5 代理地址
socks_addr: 127.0.0.1:40000

# 可选：允许所有域名（开启后会忽略下面的 rules 列表）
allow_all_hosts: true

# 可选：仅允许指定域名（和上面的 allow_all_hosts 二选一）
# 指定域名后，则代表允许 域名自身 及其 所有子域名 访问服务（以下方两个为例，√ 代表允许，× 代表阻止）
rules:
  - example.com #    example.com  √ 、a.example.com  √ 、a.a.example.com  √
  - b.example2.com # example2.com × 、b.example2.com √ 、c.b.example2.com √
```

****

一些示例：

1. 允许所有域名访问

```yaml
listen_addr: ":443"
allow_all_hosts: true
```

> 注意，开启 allow_all_hosts 时，可能会被他人扫描到而滥用，请悉知！  
> 建议做一些限制，例如只使用 IPv6（`"[::]:443"`）或防火墙限制 443 端口的可访问 IP。

2. 仅允许指定域名

```yaml
listen_addr: ":443"
rules:
  - example.com
  - b.example2.com
```

3. 允许所有域名访问 + 启用前置代理

```yaml
listen_addr: ":443"
enable_socks5: true
socks_addr: 127.0.0.1:40000
allow_all_hosts: true
```

4. 仅允许指定域名 + 启用前置代理

```yaml
listen_addr: ":443"
enable_socks5: true
socks_addr: 127.0.0.1:40000
rules:
  - example.com
  - b.example2.com
```

</details>

****

#### \# Linux 配置为系统服务 (systemd - 以支持开机启动、后台运行等)

<details>
<summary><code><strong>「 点击展开 查看内容 」</strong></code></summary>

****

新建一个空的名叫 **sniproxy** 的系统服务配置文件：

```yaml
nano /etc/systemd/system/sniproxy.service
```

修改以下内容后（`ExecStart=` 后面的路径、参数）后粘贴进文件内：

```ini
[Unit]
Description=SNI Proxy
After=network.target

[Service]
ExecStart=/home/sniproxy/sniproxy -c /home/sniproxy/config.yaml -l /home/sniproxy/sni.log

[Install]
WantedBy=multi-user.target
```

设置 **sniproxy** 开机启动并立即启动：

```yaml
# 设置开机启动
systemctl enable sniproxy

# 立即启动
systemctl start sniproxy
```

其他可能会用到的命令：

```yaml
# 停止
systemctl stop sniproxy

# 查看运行状态
systemctl status sniproxy

# 查看完整日志
cat /home/sniproxy/sni.log

# 实时监听日志（会实时显示最新日志内容）
tail -f /home/sniproxy/sni.log
```
</details>

****

#### \# 提高系统文件句柄数上限 (避免报错 too many open files)

<details>
<summary><code><strong>「 点击展开 查看内容 」</strong></code></summary>

****

Linux 系统下，一些人可能会遇到报错（日志如下）：
```
接受连接请求时出错: accept tcp [::]:443: accept4: too many open files
```

这是因为系统的文件句柄数耗尽了（默认 1024），提高系统文件句柄数上限可有效缓解该问题（不能完全解决，因为理论上，当打开文件、连接等等足够多时，迟早会耗尽，一般来说不管是做代理还是做网站，这个操作都是必须的）。

- **临时提高**（重启后恢复为 1024）
```shell
ulimit -n 65535
```

- **永久提高**（重启后依然为 65535，当然打开文件后手动删除就恢复了）
```shell
echo "* soft nofile 65535
* hard nofile 65535
root soft nofile 65535
root hard nofile 65535" >> /etc/security/limits.conf
```

执行以上命令后，需要重启 SNI Proxy 来使其生效，如果还不行请尝试重启系统。

</details>


****

## 问题反馈

如果你遇到什么问题，可以先去 [**Issues**](https://github.com/XIU2/SNIProxy/issues)、[Discussions](https://github.com/XIU2/SNIProxy/discussions) 里看看是否有别人问过了（记得去看下  [**Closed**](https://github.com/XIU2/SNIProxy/issues?q=is%3Aissue+is%3Aclosed) 的）。  
如果没找到类似问题，请新开个 [**Issues**](https://github.com/XIU2/SNIProxy/issues/new) 来告诉我！

> **注意**！_与 `反馈问题、功能建议` 无关的，请前往项目内部 论坛 讨论（上面的 `💬 Discussions`_  

****

## 赞赏支持

![微信赞赏](https://github.com/XIU2/XIU2/blob/master/img/zs-01.png)![支付宝赞赏](https://github.com/XIU2/XIU2/blob/master/img/zs-02.png)

****


## 手动编译

<details>
<summary><code><strong>「 点击展开 查看内容 」</strong></code></summary>

****

为了方便，我是在编译的时候将版本号写入代码中的 version 变量，因此你手动编译时，需要像下面这样在 `go build` 命令后面加上 `-ldflags` 参数来指定版本号：

```bash
go build -ldflags "-s -w -X main.version=v1.0.0"
# 在 SNIProxy 目录中通过命令行（例如 CMD、Bat 脚本）运行该命令，即可编译一个可在和当前设备同样系统、位数、架构的环境下运行的二进制程序（Go 会自动检测你的系统位数、架构）且版本号为 v1.0.0
```

如果想要在 Windows 64位系统下编译**其他系统、架构、位数**，那么需要指定 **GOOS** 和 **GOARCH** 变量。

例如在 Windows 系统下编译一个适用于 **Linux 系统 amd 架构 64 位**的二进制程序：

```bat
SET GOOS=linux
SET GOARCH=amd64
go build -ldflags "-s -w -X main.version=v1.0.0"
```

例如在 Linux 系统下编译一个适用于 **Windows 系统 amd 架构 32 位**的二进制程序：

```bash
GOOS=windows
GOARCH=386
go build -ldflags "-s -w -X main.version=v1.0.0"
```

> 可以运行 `go tool dist list` 来查看当前 Go 版本支持编译哪些组合。

****

当然，为了方便批量编译，我会专门指定一个变量为版本号，后续编译直接调用该版本号变量即可。  
同时，批量编译的话，还需要分开放到不同文件夹才行（或者文件名不同），需要加上 `-o` 参数指定。

```bat
:: Windows 系统下是这样：
SET version=v1.0.0
SET GOOS=linux
SET GOARCH=amd64
go build -o Releases\sniproxy_linux_amd64\sniproxy -ldflags "-s -w -X main.version=%version%"
```

```bash
# Linux 系统下是这样：
version=v1.0.0
GOOS=windows
GOARCH=386
go build -o Releases/sniproxy_windows_386/sniproxy.exe -ldflags "-s -w -X main.version=${version}"
```

</details>

****

## Credit

The source code has been adapted from [FastGitORG/F-Proxy-Agent](https://github.com/FastGitORG/F-Proxy-Agent) and [TachibanaSuzume/SNIProxyGo](https://github.com/TachibanaSuzume/SNIProxyGo) .  

Thank them for their help!

****

## License

The GPL-3.0 License.