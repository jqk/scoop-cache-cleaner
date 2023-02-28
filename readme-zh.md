# Scoop Cache Cleaner

[English](readme.md)

## 一、 用途

`Scoop Cache Cleaner (scc)` 是清理 [SCOOP][1] 所安装的应用程序安装文件缓存的`命令行工具`。

由于 [SCOOP][1] 只运行于 Windows 平台，所以，虽然本程序未使用任何与平台相关的 API，但在非 Windows 平台上运行也没有意义。

## 二、 为何不使用 `scoop cache rm *`

`scoop cache rm *` 将清空整个缓存目录。而我希望缓存中保留所有安装程序的最新版。这样，通过 [SCOOP][1] 反复安装应用程序时就不必再下载了。这是个相对特殊的操作，主要原因是为了`玩`：我有数台虚拟机，要来回做试验。而且，每隔两三个月都会重新安装一遍 Windows，这些应该程序肯定也要一键安装，当然不希望每次都下载半天。

如果不清理缓存目录，这些安装程序是很占用存储空间的。

## 三、 使用方法

```text {.line-numbers}
$ scc

Copyright (c) 1999-2023 Not a dream Co., Ltd.
scoop cache cleaner (scc) 2.1.2, 2023-02-28

Usage:
  scc <command> [path/to/scoop/cache]
      清理指定文件夹中的安装包。如果不指定路径，则使用环境变量 %SCOOP% 中定义的路径。

Command:
  -l:  显示过期的安装包。
  -b:  备份过期的安装包。
  -d:  删除过期的安装包。

如果命令行参数非以上格式，则显示本屏信息。
```

`%SCOOP%` 环境变量是在安装 [SCOOP][1] 设置的。例如, 若 `%SCOOP%` 值为 `C:\Scoop`，则 `scc -l` 将显示 `C:\Scoop\cache` 中的过期安装包。

所以，若已按默认值设置环境变量或缓存目录（`%SCOOP%\cache`），则使用 `scc -d` 最为方便，直接删除过期安装文件。

## 四、 安装方法

### 4.1 直接下载

从 <https://github.com/jqk/scoop-cache-cleaner/releases> 下载并解压最新的压缩包 `scc-windows-vX.X.X.tgz`。压缩包内只有一个文件 `scc.exe`，可放在任何目录下执行。

### 4.2 通过 Scoop

如果已经安装了 [Scoop][1]，则可通过以下命令安装本程序。

```powershell {.line-numbers}
scoop bucket add jqk https://github.com/jqk/scoopbucket
scoop install scoop-cache-cleaner
```

## 五、 示例

```text {.line-numbers}
$ scc -l

Copyright (c) 1999-2023 Not a dream Co., Ltd.
scoop cache cleaner (scc) 2.1.2, 2023-02-28

List obsolete packages in: F:\Scoop\cache

     Name               Version            Size

   1 anki               2.1.58           136.44 MB
   2 beekeeper          0.11.18           40.31 MB
   3 bookxnotepro       2.0.0.1103        62.93 MB
   4 Clash-for-Windows  0.20.17           68.45 MB
   5 curl               7.88.1             6.05 MB
   6 curl               7.88.1_1           6.05 MB
   7 dbmate             1.16.2             7.73 MB
   8 dbvis              14.0             139.01 MB
   9 duckdb             0.7.0              7.85 MB
  10 dumo               2.25.2.122         2.81 MB
  11 EverythingToolbar  1.0.2              3.09 MB
  12 feishu             5.30.10          204.34 MB
  13 foxit-reader       12.1.0.1525      189.84 MB
  14 googlechrome       109.0.5414.129    88.81 MB
  15 hydrus-network     517              220.74 MB
  16 obsidian           1.1.15            69.37 MB
  17 onefetch           2.15.1             5.23 MB
  18 opera              95.0.4635.46      93.03 MB
  19 paint.net          5.0.1             86.64 MB
  20 plantuml           1.2023.1          10.56 MB
  21 postman            10.11.1          158.85 MB
  22 questdb            7.0.0              6.79 MB
  23 springboot         3.0.2              4.48 MB
  24 SysGauge           9.1.12             5.44 MB
  25 tabby              1.0.188           92.14 MB
  26 teamviewer         15.38.3           50.28 MB
  27 telegram           4.6.3             47.78 MB
  28 trid               2.24-23.02.18      1.79 MB
  29 trid               2.24-23.02.18     47.98 KB
  30 vivaldi            5.7.2921.53       87.80 MB
  31 whatsapp           2.2305.7         152.37 MB
  32 yanknote           3.49.0           129.01 MB
  33 zotero             6.0.20            49.08 MB

-----------------------
File found            : 366
Software found        : 307
Obsolete Package found: 33
Obsolete Package Size : 2.18 GB
-----------------------
```

使用 `scc -d` 将直接删除以上过期的安装包。

[1]: https://github.com/ScoopInstaller/Scoop
