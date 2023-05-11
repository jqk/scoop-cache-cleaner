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
scoop cache cleaner (scc) 2.1.3, 2023-05-11

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
scoop cache cleaner (scc) 2.1.3, 2023-05-11

List obsolete packages in: F:\Scoop\cache

     Name                             Version               Extension     Size

   1 act                              0.2.43                86_64.zip     5.39 MB
   2 act                              0.2.44                86_64.zip     5.42 MB
   3 AgentRansack                     3367                  _3367.exe    46.97 MB
   4 AgentRansack                     3385                  _3385.exe    46.73 MB
   5 AgentRansack                     3386                  _3386.exe    46.74 MB
   6 alldup                           4.4.56                table.zip    20.10 MB
   7 alldup                           4.5.40                table.zip    27.83 MB
   8 alldup                           4.5.42                table.zip    27.84 MB
   9 alldup                           4.5.44                table.zip    27.93 MB
  10 anki                             2.1.60                exe_dl.7z   136.49 MB
  11 anki                             2.1.61                exe_dl.7z   136.83 MB
  12 audacity                         3.2.5                 5-x64.zip    19.97 MB
  13 audacity                         3.3.0                 0-x64.zip    20.97 MB
  14 audacity                         3.3.1                 1-x64.zip    21.70 MB
  15 baidunetdisk                     7.26.0.10             exe_dl.7z   239.37 MB
  16 baidunetdisk                     7.27.1.5              exe_dl.7z   273.05 MB
  17 beekeeper                        0.12.0                eeper.exe    40.95 MB
  18 beekeeper                        0.12.1                eeper.exe    40.95 MB
  19 beekeeper-studio                 3.8.9                 exe_dl.7z    63.51 MB
  20 beekeeper-studio                 3.9.5                 exe_dl.7z    67.22 MB
  21 beekeeper-studio                 3.9.8                 exe_dl.7z    67.21 MB
  22 bookxnotepro                     2.0.0.1105            30306.zip    62.93 MB
  23 broot                            1.21.1                .21.1.zip    22.51 MB
  24 broot                            1.21.2                .21.2.zip    22.43 MB
  25 Buzz                             0.7.2                 ws.tar.gz   274.21 MB
  26 ccleaner                         6.1.0                 up610.zip    43.01 MB
  27 charles                          4.6.3                 win64.msi    56.83 MB
  28 chezmoi                          2.33.0                amd64.zip    10.23 MB
  29 chezmoi                          2.33.1                amd64.zip    10.25 MB
  30 chezmoi                          2.33.3                amd64.zip    10.25 MB
  31 Clash-for-Windows                0.20.20               20-win.7z    73.46 MB
  32 Clash-for-Windows                0.20.21               21-win.7z    73.47 MB
  33 Clash-for-Windows                0.20.22               22-win.7z    73.47 MB
  34 containerd                       1.7.0                 64.tar.gz    32.61 MB
  35 curl                             8.0.1_4               gw.tar.xz     6.05 MB
  36 curl                             8.0.1_5               gw.tar.xz     6.05 MB
  37 curl                             8.0.1_6               gw.tar.xz     6.06 MB
  38 curl                             8.0.1_7               gw.tar.xz     6.16 MB
  39 curl                             8.0.1_8               gw.tar.xz     6.16 MB
  40 d2                               0.2.6                 64.tar.gz    17.39 MB
  41 d2                               0.3.0                 64.tar.gz    17.58 MB
  42 d2                               0.4.0                 64.tar.gz    17.64 MB
  43 d2                               0.4.1                 64.tar.gz    18.12 MB
  44 dbvis                            14.0.1                4_0_1.zip   139.07 MB
  45 dbvis                            23.1                  _23_1.zip   168.44 MB
  46 diskgenius                       5.4.6.1441            1_x64.zip    46.32 MB
  47 diskusage                        1.1.2                 amd64.zip     2.81 MB
  48 diskusage                        1.1.3                 amd64.zip     2.81 MB
  49 dismplusplus                     10.1.1002.1           002.1.zip     3.60 MB
  50 ditto                            3.24.238.0            238_0.zip     4.67 MB
  51 dngrep                           3.2.279.0             9.x64.msi     6.78 MB
  52 dotpeek                          2022.3.3              tpeek.exe    85.98 MB
  53 dotpeek                          2023.1                tpeek.exe    90.22 MB
  54 draw.io                          20.8.16               exe_dl.7z    96.56 MB
  55 draw.io                          21.1.2                exe_dl.7z    97.00 MB
  56 draw.io                          21.2.1                exe_dl.7z    99.20 MB
  57 edrawmax-cn                      12.0.8                setup.exe   269.04 MB
  58 everything                       1.4.1.1022            2.x64.zip     1.72 MB
  59 EverythingToolbar                1.0.3                 1.0.3.msi     3.09 MB
  60 EverythingToolbar                1.0.4                 1.0.4.msi     3.21 MB
  61 fastcopy                         4.2.2                 aller.exe     3.57 MB
  62 fastcopy                         5.0.2                 aller.exe     4.13 MB
  63 fastcopy                         5.0.3                 aller.exe     4.13 MB
  64 fastcopy                         5.0.4                 aller.exe     4.13 MB
  65 fastcopy                         5.0.5                 aller.exe     4.13 MB
  66 feishu                           6.0.5                 -6.0.5.7z   208.44 MB
  67 feishu                           6.1.5                 -6.1.5.7z   210.90 MB
  68 feishu                           6.2.5                 -6.2.5.7z   212.55 MB
  69 firefox                          111.0.1               exe_dl.7z    55.54 MB
  70 firefox                          112.0                 exe_dl.7z    55.78 MB
  71 firefox                          112.0.1               exe_dl.7z    55.77 MB
  72 firefox                          112.0.2               exe_dl.7z    55.78 MB
  73 flow-launcher                    1.14.0                table.zip   178.54 MB
  74 flow-launcher                    1.14.1                table.zip   178.54 MB
  75 foobar2000                       1.6.16                exe_dl.7z     4.63 MB
  76 foxit-reader                     12.1.1.15289          setup.exe   189.83 MB
  77 geogebra                         6.0.766.0             766-0.zip    95.81 MB
  78 geogebra                         6.0.772.0             772-0.zip    95.81 MB
  79 geogebra                         6.0.774.0             774-0.zip    95.81 MB
  80 geogebra                         6.0.775.0             775-0.zip    95.81 MB
  81 git                              2.40.0.windows.1      exe_dl.7z    46.82 MB
  82 github                           3.2.0                 32_dl.7z_   132.48 MB
  83 github                           3.2.1                 32_dl.7z_   132.79 MB
  84 github                           3.2.2                 32_dl.7z_   133.13 MB
  85 glaryutilities                   5.203.0.232           table.zip    24.68 MB
  86 gnucash                          5.0                   setup.exe   144.04 MB
  87 go                               1.20.2                amd64.zip   108.69 MB
  88 go                               1.20.3                amd64.zip   108.73 MB
  89 googlechrome                     109.0.5414.75         exe_dl.7z    88.82 MB
  90 googlechrome                     112.0.5615.121        exe_dl.7z    89.08 MB
  91 googlechrome                     112.0.5615.138        exe_dl.7z    89.16 MB
  92 googlechrome                     112.0.5615.50         exe_dl.7z    89.14 MB
  93 googlechrome                     112.0.5615.87         exe_dl.7z    89.15 MB
  94 googlechrome                     113.0.5672.64         exe_dl.7z    90.52 MB
  95 goreleaser                       1.16.2                86_64.zip    15.82 MB
  96 goreleaser                       1.17.0                86_64.zip    16.00 MB
  97 goreleaser                       1.17.1                86_64.zip    16.00 MB
  98 goreleaser                       1.17.2                86_64.zip    16.00 MB
  99 goreleaser                       1.18.1                86_64.zip    16.29 MB
 100 graalvm-jdk11                    22.3.1                2.3.1.zip   237.54 MB
 101 graalvm-jdk17                    22.3.1                2.3.1.zip   245.00 MB
 102 gradle                           8.0.2                 2-all.zip   159.90 MB
 103 gradle                           8.1                   1-all.zip   160.51 MB
 104 graphviz                         8.0.1                 exe_dl.7z     4.84 MB
 105 graphviz                         8.0.2                 exe_dl.7z     4.83 MB
 106 graphviz                         8.0.3                 exe_dl.7z     4.83 MB
 107 graphviz                         8.0.4                 exe_dl.7z     4.83 MB
 108 gsudo                            2.0.4                 2.0.4.zip     9.61 MB
 109 gsudo                            2.0.6                 p.x64.msi     2.21 MB
 110 gsudo                            2.0.7                 p.x64.msi     2.21 MB
 111 gsudo                            2.0.8                 p.x64.msi     2.21 MB
 112 HeidiSQL                         12.4                  table.zip    18.08 MB
 113 hydrus-network                   521                   .only.zip   223.73 MB
 114 hydrus-network                   522                   .only.zip   223.84 MB
 115 hydrus-network                   523                   .only.zip   223.90 MB
 116 hydrus-network                   524                   .only.zip   223.93 MB
 117 hydrus-network                   525                   .only.zip   228.43 MB
 118 imhex                            1.27.1                86_64.zip    23.05 MB
 119 intellij-idea-ultimate-portable  2022.3.3-223.8836.41  3.win.zip     1.03 GB
 120 intellij-idea-ultimate-portable  2022.3.3-223.8836.41  table.ps1     1.13 KB
 121 intellij-idea-ultimate-portable  2023.1-231.8109.175   1.win.zip     1.03 GB
 122 intellij-idea-ultimate-portable  2023.1-231.8109.175   table.ps1     1.13 KB
 123 jenkins-lts                      2.387.1               nkins.jar    93.78 MB
 124 jenkins-lts                      2.387.2               nkins.jar    93.80 MB
 125 joplin                           2.10.17               exe_dl.7z   196.67 MB
 126 joplin                           2.10.18               exe_dl.7z   195.33 MB
 127 kopia                            0.12.1                s-x64.zip    11.23 MB
 128 kopiaui                          0.12.1                1-win.zip    93.56 MB
 129 kopiaui                          0.12.1                exe_dl.7z    68.61 MB
 130 kopiaui                          0.13.0                0-win.zip   104.02 MB
 131 ktlint                           0.48.2                tlint.jar    62.13 MB
 132 lazygit                          0.37.0                86_64.zip     5.35 MB
 133 lazygit                          0.38.0                .download     4.02 MB
 134 lazygit                          0.38.1                86_64.zip     5.39 MB
 135 liquibase                        4.20.0                .20.0.zip    66.00 MB
 136 liquibase                        4.21.0                .21.0.zip    66.12 MB
 137 listen1desktop                   2.27.0                exe_dl.7z   171.41 MB
 138 logseq                           0.8.9                 exe_dl.7z   175.29 MB
 139 logseq                           0.9.0                 exe_dl.7z   168.89 MB
 140 logseq                           0.9.1                 exe_dl.7z   168.89 MB
 141 logseq                           0.9.2                 exe_dl.7z   169.07 MB
 142 logseq                           0.9.3                 exe_dl.7z   169.26 MB
 143 lunacy                           9.0.8                 9.0.8.exe   122.21 MB
 144 lx-music                         2.2.0                 -green.7z    64.74 MB
 145 manictime                        2023.1.1.0            n-x64.zip    78.92 MB
 146 micronaut                        3.8.7                 3.8.7.zip    25.27 MB
 147 micronaut                        3.8.8                 3.8.8.zip    25.28 MB
 148 micronaut                        3.8.9                 3.8.9.zip    25.29 MB
 149 micronaut                        3.9.0                 3.9.0.zip    25.45 MB
 150 microsoft-teams                  1.6.00.11166          ull.nupkg   134.58 MB
 151 microsoft-teams                  1.6.00.4472           ull.nupkg   134.18 MB
 152 miktex                           22.7                  7-x64.exe   132.71 MB
 153 mindmaster                       10.5.2                setup.exe   129.04 MB
 154 minikube                         1.29.0                ikube.exe    77.25 MB
 155 minikube                         1.30.0                ikube.exe    79.93 MB
 156 mobaxterm                        23.0                  v23.0.zip    28.30 MB
 157 mvndaemon                        0.9.0                 amd64.zip    18.64 MB
 158 naiveproxy                       111.0.5563.64-1       n-x64.zip     3.08 MB
 159 naiveproxy                       112.0.5615.49-1       n-x64.zip     3.08 MB
 160 naiveproxy                       113.0.5672.62-1       n-x64.zip     3.08 MB
 161 neovim                           0.8.3                 win64.zip    42.29 MB
 162 nginx                            1.23.3                .23.3.zip     1.67 MB
 163 nodejs                           19.8.1                in-x64.7z    17.73 MB
 164 nodejs                           19.9.0                in-x64.7z    17.69 MB
 165 nodejs                           20.0.0                in-x64.7z    17.66 MB
 166 notepadplusplus                  8.5.1                 e.x64.zip     5.43 MB
 167 obs-studio                       29.0.2                l-x64.zip   139.52 MB
 168 obsidian                         1.1.9                 exe_dl.7z    69.32 MB
 169 obsidian                         1.2.7                 exe_dl.7z    71.45 MB
 170 ocenaudio                        3.11.22               _v3.11.22    59.06 MB
 171 ocenaudio                        3.11.23               _v3.11.23    59.07 MB
 172 ocenaudio                        3.11.24               _v3.11.24    59.23 MB
 173 okular                           22.12.1-1265          exe_dl.7z   118.04 MB
 174 okular                           22.12.3               x86_64.7z   117.73 MB
 175 okular                           22.12.3-1357          x86_64.7z   117.77 MB
 176 okular                           22.12.3-1359          x86_64.7z   117.77 MB
 177 okular                           22.12.3-1360          x86_64.7z   117.78 MB
 178 okular                           23.04.0-1363          x86_64.7z   117.76 MB
 179 okular                           23.04.0-1364          x86_64.7z   117.76 MB
 180 okular                           23.04.0-1365          x86_64.7z   117.76 MB
 181 okular                           23.04.0-1367          x86_64.7z   117.73 MB
 182 okular                           23.04.0-1368          x86_64.7z   117.72 MB
 183 okular                           23.04.0-1369          x86_64.7z   117.73 MB
 184 okular                           23.04.0-1370          x86_64.7z   117.73 MB
 185 okular                           23.04.0-1372          x86_64.7z   117.72 MB
 186 okular                           23.04.0-1374          x86_64.7z   117.72 MB
 187 okular                           23.04.0-1375          x86_64.7z   117.73 MB
 188 okular                           23.04.0-1376          x86_64.7z   117.72 MB
 189 okular                           23.04.0-1378          x86_64.7z   117.73 MB
 190 okular                           23.04.0-1379          x86_64.7z   117.72 MB
 191 okular                           23.04.0-1380          x86_64.7z   117.73 MB
 192 onefetch                         2.16.0                in.tar.gz     5.27 MB
 193 onefetch                         2.17.0                in.tar.gz     5.28 MB
 194 openai-translator                0.0.37                en-US.msi     9.94 MB
 195 openai-translator                0.0.38                en-US.msi    10.26 MB
 196 openai-translator                0.0.39                en-US.msi    10.29 MB
 197 openai-translator                0.0.40                en-US.msi    10.29 MB
 198 openai-translator                0.0.41                en-US.msi    10.29 MB
 199 openai-translator                0.0.43                en-US.msi    10.29 MB
 200 opera                            97.0.4719.28          exe_dl.7z    95.48 MB
 201 opera                            97.0.4719.43          exe_dl.7z    95.48 MB
 202 opera                            97.0.4719.63          exe_dl.7z    95.49 MB
 203 opera                            97.0.4719.83          exe_dl.7z    95.50 MB
 204 paint.net                        5.0.2                 e.x64.zip    86.75 MB
 205 PandaOCRPro                      5.46                  _5.46.zip    11.94 MB
 206 plantuml                         1.2023.5              ntuml.jar    10.61 MB
 207 postman                          10.12.13              exe_dl.7z   165.11 MB
 208 postman                          10.13.0               exe_dl.7z   159.51 MB
 209 potplayer                        230207                exe_dl.7z    32.80 MB
 210 powertoys                        0.68.1                1-x64.exe   169.21 MB
 211 powertoys                        0.69.0                0-x64.exe   206.96 MB
 212 python                           3.11.2                setup.exe    24.15 MB
 213 questdb                          7.0.1                 in.tar.gz     6.80 MB
 214 questdb                          7.1                   in.tar.gz     6.77 MB
 215 quicker                          1.36.5.0              6.5.0.msi    19.02 MB
 216 rufus                            3.22                  rufus.exe     1.35 MB
 217 rustup                           1.25.2                -init.exe     9.53 MB
 218 shotcut                          22.12.21              21221.zip   135.89 MB
 219 slack                            4.31.155              ull.nupkg   109.61 MB
 220 snappy-driver-installer-origin   1.12.10.750           0.750.zip     6.17 MB
 221 snappy-driver-installer-origin   1.12.11.751           1.751.zip     6.18 MB
 222 snappy-driver-installer-origin   1.12.12.753           2.753.zip     6.18 MB
 223 springboot                       3.0.5                 5-bin.zip     4.50 MB
 224 switchhosts                      4.1.1.6077            exe_dl.7z   112.62 MB
 225 syncback                         10.2.112.0            Setup.exe    44.77 MB
 226 syncback                         10.2.116.0            Setup.exe    44.74 MB
 227 syncthing                        1.23.2                .23.2.zip     9.73 MB
 228 syncthing                        1.23.3                .23.3.zip     9.67 MB
 229 SysGauge                         9.3.12                exe_dl.7z     5.46 MB
 230 SysGauge                         9.4.16                exe_dl.7z     5.46 MB
 231 tabby                            1.0.196               exe_dl.7z    89.97 MB
 232 tailscale                        1.38.2                amd64.msi    18.47 MB
 233 tailscale                        1.38.3                amd64.msi    18.48 MB
 234 tailscale                        1.38.4                amd64.msi    18.68 MB
 235 teamviewer                       15.40.8               table.zip    52.16 MB
 236 teamviewer                       15.41.7               table.zip    52.98 MB
 237 teamviewer                       15.41.8               table.zip    52.98 MB
 238 telegram                         4.7.1                 4.7.1.zip    48.01 MB
 239 telegram                         4.8.0                 4.8.0.zip    48.33 MB
 240 thunderbird                      102.10.0              exe_dl.7z    53.90 MB
 241 thunderbird                      102.10.1              exe_dl.7z    53.89 MB
 242 thunderbird                      102.11.0              exe_dl.7z    53.93 MB
 243 thunderbird                      102.9.0               exe_dl.7z    53.87 MB
 244 todesk                           4.6.1.4               exe_dl.7z    77.91 MB
 245 todesk                           4.6.2.1               exe_dl.7z    85.96 MB
 246 todesk                           4.6.2.2               exe_dl.7z    86.02 MB
 247 tradingview                      2.2.0                 View.msix   102.91 MB
 248 tradingview                      2.3.1                 View.msix   102.47 MB
 249 tradingview                      2.3.3                 View.msix   102.47 MB
 250 trid                             2.24-23.03.28         ddefs.zip     1.82 MB
 251 trid                             2.24-23.03.28         d_w32.zip    47.98 KB
 252 trid                             2.24-23.04.01         ddefs.zip     1.82 MB
 253 trid                             2.24-23.04.01         d_w32.zip    47.98 KB
 254 trid                             2.24-23.04.03         ddefs.zip     1.83 MB
 255 trid                             2.24-23.04.03         d_w32.zip    47.98 KB
 256 trid                             2.24-23.04.05         ddefs.zip     1.83 MB
 257 trid                             2.24-23.04.05         d_w32.zip    47.98 KB
 258 trid                             2.24-23.04.08         ddefs.zip     1.83 MB
 259 trid                             2.24-23.04.08         d_w32.zip    47.98 KB
 260 trid                             2.24-23.04.10         ddefs.zip     1.83 MB
 261 trid                             2.24-23.04.10         d_w32.zip    47.98 KB
 262 trid                             2.24-23.04.12         ddefs.zip     1.84 MB
 263 trid                             2.24-23.04.12         d_w32.zip    47.98 KB
 264 trid                             2.24-23.04.14         ddefs.zip     1.84 MB
 265 trid                             2.24-23.04.14         d_w32.zip    47.98 KB
 266 trid                             2.24-23.04.17         ddefs.zip     1.84 MB
 267 trid                             2.24-23.04.17         d_w32.zip    47.98 KB
 268 trid                             2.24-23.04.19         ddefs.zip     1.84 MB
 269 trid                             2.24-23.04.19         d_w32.zip    47.98 KB
 270 trid                             2.24-23.04.23         ddefs.zip     1.85 MB
 271 trid                             2.24-23.04.23         d_w32.zip    47.98 KB
 272 trid                             2.24-23.04.25         ddefs.zip     1.85 MB
 273 trid                             2.24-23.04.25         d_w32.zip    47.98 KB
 274 trid                             2.24-23.04.28         ddefs.zip     1.85 MB
 275 trid                             2.24-23.04.28         d_w32.zip    47.98 KB
 276 trid                             2.24-23.04.30         ddefs.zip     1.85 MB
 277 trid                             2.24-23.04.30         d_w32.zip    47.98 KB
 278 trid                             2.24-23.05.02         ddefs.zip     1.85 MB
 279 trid                             2.24-23.05.02         d_w32.zip    47.98 KB
 280 trid                             2.24-23.05.04         ddefs.zip     1.85 MB
 281 trid                             2.24-23.05.04         d_w32.zip    47.98 KB
 282 trid                             2.24-23.05.07         ddefs.zip     1.86 MB
 283 trid                             2.24-23.05.07         d_w32.zip    47.98 KB
 284 trilium                          0.59.3                .59.3.zip   101.55 MB
 285 typst                            0.1                   -msvc.zip     8.49 MB
 286 typst                            0.1.0                 -msvc.zip     8.49 MB
 287 typst                            0.2.0                 -msvc.zip     9.61 MB
 288 typst                            0.3.0                 -msvc.zip    10.27 MB
 289 typst                            23-03-21-2            -msvc.zip     8.03 MB
 290 ugrep                            3.11.0                xe_ug.exe     1.84 MB
 291 ugrep                            3.11.1                xe_ug.exe     1.84 MB
 292 umi-ocr                          1.3.3                 guages.7z    50.49 MB
 293 umi-ocr                          1.3.3                 v1.3.3.7z    82.41 MB
 294 usql                             0.13.8                amd64.zip    23.59 MB
 295 usql                             0.14.0                amd64.zip    23.99 MB
 296 usql                             0.14.1                amd64.zip    24.01 MB
 297 usql                             0.14.2                amd64.zip    24.83 MB
 298 usql                             0.14.3                amd64.zip    24.86 MB
 299 v2ray                            5.3.0                 ws-64.zip    11.16 MB
 300 v2rayn                           5.39                  2rayN.zip     8.56 MB
 301 v2rayn                           6.21                  2rayN.zip    20.63 MB
 302 vcredist2022                     14.34.31931.0         t.x86.exe    13.19 MB
 303 vcredist2022                     14.34.31931.0         t.x64.exe    24.29 MB
 304 ventoy                           1.0.90                ndows.zip    15.87 MB
 305 VisualParadigm                   17.0.20230201         lFree.zip   738.46 MB
 306 vivaldi                          5.7.2921.65           exe_dl.7z    87.81 MB
 307 vivaldi                          5.7.2921.68           exe_dl.7z    87.80 MB
 308 vivaldi                          6.0.2979.11           exe_dl.7z    93.28 MB
 309 vivaldi                          6.0.2979.15           exe_dl.7z    93.31 MB
 310 vscode                           1.76.2                ble_dl.7z   120.76 MB
 311 vscode                           1.77.0                ble_dl.7z   119.64 MB
 312 vscode                           1.77.1                ble_dl.7z   119.62 MB
 313 vscode                           1.77.2                ble_dl.7z   119.62 MB
 314 vscode                           1.77.3                ble_dl.7z   119.62 MB
 315 vscode                           1.78.0                ble_dl.7z   123.13 MB
 316 whatsapp                         2.2310.5              pkg_dl.7z   153.09 MB
 317 whatsapp                         2.2314.11             pkg_dl.7z   151.73 MB
 318 whatsapp                         2.2317.10             pkg_dl.7z   152.10 MB
 319 whatsapp                         2.2317.11             pkg_dl.7z   152.27 MB
 320 whatsapp                         2.2318.10             pkg_dl.7z   152.40 MB
 321 wireshark                        4.0.4                 exe_dl.7z    75.10 MB
 322 wpsoffice                        11.2.0.11513          11513.exe   214.18 MB
 323 wpsoffice                        11.2.0.11516          11516.exe   214.16 MB
 324 wpsoffice                        11.2.0.11536          11536.exe   214.03 MB
 325 ximalaya                         3.3.6                 win_dl.7z    66.51 MB
 326 yanknote                         3.50.2                .50.2.zip   128.21 MB
 327 yanknote                         3.51.0                .51.0.zip   128.21 MB
 328 zotero                           6.0.23                exe_dl.7z    49.10 MB
 329 zotero                           6.0.25                exe_dl.7z    49.18 MB
 330 zoxide                           0.9.0                 -msvc.zip   522.12 KB

-----------------------
File found            : 689
Software found        : 334
Obsolete Package found: 330
Obsolete Package Size : 23.46 GB
-----------------------
```

使用 `scc -d` 将直接删除以上过期的安装包。

[1]: https://github.com/ScoopInstaller/Scoop
