# Scoop Cache Cleaner

[English](readme.md)

## 一、 用途

`Scoop Cache Cleaner (scc)` 是清理 [SCOOP][1] 所安装的应用程序安装文件缓存的命令行工具。

由于 [SCOOP][1] 只运行于 Windows 平台，所以，虽然本程序未使用任何与平台相关的 API，但在非 Windows 平台上运行也没有意义。

## 二、 为何不使用 `scoop cache rm *`

`scoop cache rm *` 为清空整个缓存。而我希望缓存中保留所有安装程序的最新版。这样，通过 [SCOOP][1] 反复安装应用程序时就不必再下载了。这是个相对特殊的操作，主要原因是为了`玩`：我有数台虚拟机，要来回做试验。而且，每一两个月都会重新安装一遍 Windows，这些应该程序肯定也要一键安装，当然不希望每次都下载半天。

如果不清理缓存目录，这些安装程序是很占用存储空间的。

## 三、 使用方法

```text {.line-numbers}
Copyright (c) 1999-2023 Not a dream Co., Ltd.
scoop cache cleaner (scc) 1.0.0, 2023-01-25

Usage:
  scc [path/to/scoop/cache]
      指定待清理的目录。
  scc -e
      通过 SCOOP 环境变量确定要清理的目录。

  all other parameters will display the above information.
```

`%SCOOP%` 环境变量是在安装 [SCOOP][1] 设置的。例如, 若 `%SCOOP%` 值为 `C:\Scoop`，则 `scc -e` 将清理 `C:\Scoop\cache`。

所以，若已按默认值设置环境变量或缓存目录（`%SCOOP%\cache`），则使用 `scc -e` 最为方便。

[1]: https://github.com/ScoopInstaller/Scoop
