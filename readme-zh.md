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
scoop cache cleaner (scc) 2.1.3, 2023-03-20

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
scoop cache cleaner (scc) 2.1.3, 2023-03-20

List obsolete packages in: F:\Scoop\cache

     Name                             Version               Extension     Size

   1 act                              0.2.42                   64.zip     5.39 MB
   2 anki                             2.1.58                    dl.7z   136.44 MB
   3 antidupl.net                     2.3.10                    dl.7z     3.45 MB
   4 audacity                         3.2.4                   x64.zip    19.97 MB
   5 baidunetdisk                     7.25.0.4                  dl.7z   236.99 MB
   6 beekeeper                        0.11.18               eeper.exe    40.31 MB
   7 beekeeper                        0.11.19                            40.31 MB
   8 beyondcompare                    4.4.5.27371             x64.msi    17.22 MB
   9 bookxnotepro                     2.0.0.1103            30204.zip    62.93 MB
  10 bookxnotepro                     2.0.0.1104            30219.zip    62.95 MB
  11 broot                            1.20.2                    2.zip    22.43 MB
  12 Buzz                             0.7.1                    tar.gz   273.65 MB
  13 ccleaner                         6.06.10144            up606.zip    42.04 MB
  14 chezmoi                          2.31.0                amd64.zip     9.74 MB
  15 chezmoi                          2.31.1                              9.75 MB
  16 Clash-for-Windows                0.20.17                  win.7z    68.45 MB
  17 Clash-for-Windows                0.20.18                            68.54 MB
  18 containerd                       1.6.19                   tar.gz    31.84 MB
  19 containerd                       1.6.9                              31.00 MB
  20 copytranslator                   11.0.0                  win.zip    81.73 MB
  21 cpu-z                            2.04                     en.zip     3.28 MB
  22 croc                             9.6.3                 64bit.zip     2.66 MB
  23 curl                             7.88.1                   tar.xz     6.05 MB
  24 curl                             7.88.1_1                            6.05 MB
  25 d2                               0.2.4                    tar.gz    15.78 MB
  26 d2                               0.2.5                              17.36 MB
  27 dbmate                           1.16.2                bmate.exe     7.73 MB
  28 dbmate                           2.0.0                              13.62 MB
  29 dbmate                           2.0.1                              13.61 MB
  30 dbvis                            14.0                      0.zip   139.01 MB
  31 dingtalk                         7.0.10.2189102            dl.7z   393.43 MB
  32 diskusage                        1.1.0                 amd64.zip     2.74 MB
  33 diskusage                        1.1.1                               2.74 MB
  34 duckdb                           0.7.0                 amd64.zip     7.85 MB
  35 dumo                             2.25.2.122             dumo.zip     2.81 MB
  36 edrawmax-cn                      12.0.7                setup.exe   267.07 MB
  37 EverythingToolbar                1.0.2                     2.msi     3.09 MB
  38 feishu                           5.30.10                   10.7z   204.34 MB
  39 feishu                           5.31.6                     6.7z   205.54 MB
  40 feishu                           5.32.4                     4.7z   208.14 MB
  41 ffmpeg                           5.1.2                  build.7z    45.23 MB
  42 firefox-zh-cn                    110.0                 irefox.7z    55.29 MB
  43 firefox-zh-cn                    110.0.1                            55.26 MB
  44 flow-launcher                    1.13.0                table.zip   178.40 MB
  45 foxit-reader                     12.1.0.1525           setup.exe   189.84 MB
  46 geogebra                         6.0.760.0                 0.zip    95.79 MB
  47 gimp                             2.10.32-1                 1.exe   252.99 MB
  48 git                              2.39.2.windows.1          dl.7z    46.14 MB
  49 github                           3.1.8                    dl.7z_   132.47 MB
  50 glaryutilities                   5.201.0.230           table.zip    24.57 MB
  51 go                               1.20.1                amd64.zip   108.08 MB
  52 googlechrome                     109.0.5414.129            dl.7z    88.81 MB
  53 goreleaser                       1.16.0                   64.zip    15.56 MB
  54 gradle                           8.0.1                   all.zip   159.90 MB
  55 hydrus-network                   517                    only.zip   220.74 MB
  56 hydrus-network                   518                               223.68 MB
  57 hydrus-network                   519                               223.70 MB
  58 intellij-idea-ultimate-portable  2022.3.2-223.8617.56    win.zip     1.02 GB
  59 intellij-idea-ultimate-portable  2022.3.2-223.8617.56  table.ps1     1.13 KB
  60 intellij-idea-ultimate-portable  2022.3.3-223.8836.35    win.zip     1.03 GB
  61 intellij-idea-ultimate-portable  2022.3.3-223.8836.35  table.ps1     1.13 KB
  62 jenkins-lts                      2.375.3               nkins.jar    89.87 MB
  63 lapce                            0.2.5                 table.zip    17.31 MB
  64 lapce                            0.2.6                              17.52 MB
  65 liquibase                        4.19.0                    0.zip    65.76 MB
  66 liquibase                        4.19.1                    1.zip    65.97 MB
  67 logseq                           0.8.18                    dl.7z   168.56 MB
  68 lunacy                           8.7.2                     2.exe   120.13 MB
  69 maven                            3.9.0                   bin.zip     8.70 MB
  70 mdbook                           0.4.27                 msvc.zip     3.79 MB
  71 micronaut                        3.8.5                     5.zip    25.27 MB
  72 micronaut                        3.8.6                     6.zip    25.27 MB
  73 microsoft-teams                  1.6.00.376            ull.nupkg   131.05 MB
  74 mindmaster                       10.1.0                setup.exe   108.57 MB
  75 naiveproxy                       110.0.5481.100-1        x64.zip     3.10 MB
  76 nodejs                           19.7.0                   x64.7z    17.71 MB
  77 nodejs                           19.8.0                             17.73 MB
  78 notepadplusplus                  8.4.9                   x64.zip     5.39 MB
  79 obsidian                         1.1.15                    dl.7z    69.37 MB
  80 obsidian                         1.1.16                             69.38 MB
  81 onefetch                         2.15.1                   tar.gz     5.23 MB
  82 opera                            95.0.4635.46              dl.7z    93.03 MB
  83 opera                            96.0.4693.20                       93.81 MB
  84 opera                            96.0.4693.31                       94.08 MB
  85 opera                            96.0.4693.50                       94.08 MB
  86 paint.net                        5.0.1                   x64.zip    86.64 MB
  87 pandoc                           3.1                      64.zip    25.06 MB
  88 plantuml                         1.2023.1              ntuml.jar    10.56 MB
  89 plantuml                         1.2023.2                           10.60 MB
  90 postman                          10.11.1               ull.nupkg   158.85 MB
  91 postman                          10.12.0                           163.36 MB
  92 powertoys                        0.67.1                  x64.exe   167.66 MB
  93 powertoys                        0.68.0                            169.23 MB
  94 questdb                          7.0.0                    tar.gz     6.79 MB
  95 schemacrawler                    16.19.7                   7.msi   226.85 MB
  96 scoop-cache-cleaner              2.1.1                     1.tgz     1.65 MB
  97 shadowsocks-rust                 1.15.2                 msvc.zip     9.95 MB
  98 sourcetree                       3.4.12                ull.nupkg    24.21 MB
  99 springboot                       3.0.2                   bin.zip     4.48 MB
 100 springboot                       3.0.3                               4.49 MB
 101 StrokesPlus.net                  0.5.7.2                   2.zip    23.50 MB
 102 syncback                         10.2.88.0             Setup.exe    44.75 MB
 103 syncthing                        1.23.1                    1.zip     9.54 MB
 104 SysGauge                         9.1.12                    dl.7z     5.44 MB
 105 SysGauge                         9.2.18                              5.45 MB
 106 tabby                            1.0.188                   dl.7z    92.14 MB
 107 tabby                            1.0.189                            92.54 MB
 108 tabby                            1.0.191                            89.99 MB
 109 teamviewer                       15.38.3               table.zip    50.28 MB
 110 teamviewer                       15.39.5                            50.36 MB
 111 telegram                         4.6.3                     3.zip    47.78 MB
 112 telegram                         4.6.5                     5.zip    47.79 MB
 113 tencentmeeting                   3.14.10.401               dl.7z   183.46 MB
 114 tencentmeeting                   3.14.10.401           setup.exe   183.46 MB
 115 thunderbird                      102.8.0                   dl.7z    53.86 MB
 116 todesk                           4.6.1.3                   dl.7z    76.21 MB
 117 tradingview                      1.0.9                 View.msix    89.57 MB
 118 trid                             2.24-23.02.18         ddefs.zip     1.79 MB
 119 trid                             2.24-23.02.18           w32.zip    47.98 KB
 120 trid                             2.24-23.02.22         ddefs.zip     1.80 MB
 121 trid                             2.24-23.02.22           w32.zip    47.98 KB
 122 trid                             2.24-23.02.28         ddefs.zip     1.80 MB
 123 trid                             2.24-23.02.28           w32.zip    47.98 KB
 124 trid                             2.24-23.03.04         ddefs.zip     1.80 MB
 125 trid                             2.24-23.03.04           w32.zip    47.98 KB
 126 trid                             2.24-23.03.08         ddefs.zip     1.80 MB
 127 trid                             2.24-23.03.08           w32.zip    47.98 KB
 128 trid                             2.24-23.03.11         ddefs.zip     1.80 MB
 129 trid                             2.24-23.03.11           w32.zip    47.98 KB
 130 trid                             2.24-23.03.14         ddefs.zip     1.80 MB
 131 trid                             2.24-23.03.14           w32.zip    47.98 KB
 132 trilium                          0.58.8                    8.zip   100.83 MB
 133 trilium                          0.59.1                    1.zip   101.55 MB
 134 ugrep                            3.10.0                   ug.exe     1.84 MB
 135 ugrep                            3.10.1                              1.84 MB
 136 usql                             0.13.10               amd64.zip    24.19 MB
 137 usql                             0.13.11                            24.20 MB
 138 usql                             0.13.12                            24.43 MB
 139 ventoy                           1.0.88                ndows.zip    15.86 MB
 140 video-compare                    20230117                 64.zip    54.18 MB
 141 video-compare                    20230223                           54.18 MB
 142 video-compare                    20230306                           54.18 MB
 143 vivaldi                          5.7.2921.53               dl.7z    87.80 MB
 144 vivaldi                          5.7.2921.60                        87.81 MB
 145 vscode                           1.75.1                    dl.7z   125.22 MB
 146 vscode                           1.76.0                            120.76 MB
 147 vscode                           1.76.1                            120.76 MB
 148 watchexec                        1.20.5                 msvc.zip     1.92 MB
 149 watchexec                        1.21.1                              2.00 MB
 150 whatsapp                         2.2305.7                  dl.7z   152.37 MB
 151 whatsapp                         2.2306.9                          152.50 MB
 152 wireshark                        4.0.3                     dl.7z    75.08 MB
 153 yanknote                         3.49.0                    0.zip   129.01 MB
 154 yanknote                         3.50.1                    1.zip   129.00 MB
 155 zotero                           6.0.20                    dl.7z    49.08 MB
 156 zotero                           6.0.22                             49.08 MB

-----------------------
File found            : 495
Software found        : 313
Obsolete Package found: 156
Obsolete Package Size : 12.09 GB
-----------------------
```

使用 `scc -d` 将直接删除以上过期的安装包。

[1]: https://github.com/ScoopInstaller/Scoop
