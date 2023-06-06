# XIU2/SNIProxy

[![Go Version](https://img.shields.io/github/go-mod/go-version/XIU2/SNIProxy.svg?style=flat-square&label=Go&color=00ADD8&logo=go)](https://github.com/XIU2/SNIProxy/)
[![Release Version](https://img.shields.io/github/v/release/XIU2/SNIProxy.svg?style=flat-square&label=Release&color=00ADD8&logo=github)](https://github.com/XIU2/SNIProxy/releases/latest)
[![GitHub license](https://img.shields.io/github/license/XIU2/SNIProxy.svg?style=flat-square&label=License&color=00ADD8&logo=github)](https://github.com/XIU2/SNIProxy/)
[![GitHub Star](https://img.shields.io/github/stars/XIU2/SNIProxy.svg?style=flat-square&label=Star&color=00ADD8&logo=github)](https://github.com/XIU2/SNIProxy/)
[![GitHub Fork](https://img.shields.io/github/forks/XIU2/SNIProxy.svg?style=flat-square&label=Fork&color=00ADD8&logo=github)](https://github.com/XIU2/SNIProxy/)

🧷 自用的一个功能很简单的 SNI Proxy 顺便分享出来给有同样需求的人，用得上的话可以点个⭐支持下~

****

## \# 软件特性

1. **支持** 全平台、全系统（Go 语言特性）
2. **支持** Socks5 前置代理（比如可以套 WARP+）
3. **支持** 允许所有域名 或 仅允许指定域名（允许域名自身及其所有子域名）

****

## \# 使用方法

<details>
<summary><code><strong>「 点击查看 Windows 系统下的使用示例 」</strong></code></summary>

****

### 下载

下载已编译好的可执行文件并解压：

1. [Github Releases](https://github.com/XIU2/SNIProxy/releases)  
2. [蓝奏云](https://pan.lanzouf.com/b077bn2ri)(密码:xiu2)

### 配置

找到配置文件 `config.yaml` 右键菜单 - 打开方式 - 记事本。

根据配置文件中的注释（# 开头的）修改配置内容并保存。

### 运行

可双击运行 `sniproxy.exe` 文件。

或者也可以在 CMD 命令行中进入软件所在目录并运行 `sniproxy.exe`：

```cmd
:: 进入解压后的 sniproxy 程序所在目录（记得修改下面示例路径）
cd /d C:\xxx\sniproxy

:: 运行（不带参数）
sniproxy.exe

:: 运行（带参数示例）
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

# 运行（不带参数）
./sniproxy

# 运行（带参数示例）
./sniproxy -c "config.yaml"

# 后台运行（带参数示例）
nohup ./sniproxy -c "config.yaml" > "sni.log" 2>&1 &
```

</details>

****

<details>
<summary><code><strong>「 点击查看 Mac 系统下的使用示例 」</strong></code></summary>

****

下载已编译好的可执行文件并解压：

1. [Github Releases](https://github.com/XIU2/SNIProxy/releases)  
2. [蓝奏云](https://pan.lanzouf.com/b077bn2ri)(密码:xiu2)

```yaml
# 进入 sniproxy 压缩包所在目录（记得修改下面示例路径）
cd /xxx/xxx

# 解压（不需要删除旧文件，会直接覆盖，自行根据需求替换 文件名）
tar -zxf sniproxy_linux_amd64.tar.gz

# 赋予执行权限
chmod a+x sniproxy

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

## Credit

Source from [FastGitORG/F-Proxy-Agent](https://github.com/FastGitORG/F-Proxy-Agent)(GPL-3.0) and [TachibanaSuzume/SNIProxyGo](https://github.com/TachibanaSuzume/SNIProxyGo)(GPL-3.0)

****

## License

The GPL-3.0 License.