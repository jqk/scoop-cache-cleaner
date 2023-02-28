# Scoop Cache Cleaner

[中文版](readme-zh.md)

## What is it

`Scoop Cache Cleaner (scc)` is a command line tool used to clean up the [SCOOP][1] cache directory. Although no platform-specified API is used, it only applicable to Windows because [SCOOP][1] only runs on Windows.

## Why not `scoop cache rm *`

Running `scoop cache rm *` will empty the cache directory without leaving any file. However, I still want to keep all latest setup files after running `scoop update *`.

If you do not run `spoon cache rm *` or manually clean the directory, it will occupy lots of disk space.

## Usage

```text {.line-numbers}
$ scc

Copyright (c) 1999-2023 Not a dream Co., Ltd.
scoop cache cleaner (scc) 2.1.2, 2023-02-28

Usage:
  scc <command> [path/to/scoop/cache]
      clean up the specified scoop cache directory.
      if the path is omitted, it will use the path defined in the environment variable %SCOOP%.

Command:
  -l:  list the obsolete packages.
  -b:  backup the obsolete packages.
  -d:  delete the obsolete packages.

all other parameters will display the above information.
```

If path is omitted, it takes the environment value of `%SCOOP%` and appends `cache` to it.

For example, if `%SCOOP%` is `C:\Scoop`, `scc -l` will list obsolete packages in `C:\Scoop\cache`.

## Install

### Download from GitHub

Go to <https://github.com/jqk/scoop-cache-cleaner/releases>, download the latest package and unzip it.

There is only scc.exe in the compressed file, which can be executed in any directory.

### Using Scoop

If [Scoop][1] is already installed, run:

```powershell {.line-numbers}
scoop bucket add jqk https://github.com/jqk/scoopbucket
scoop install scoop-cache-cleaner
```

## Example

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

Using `scc -d` will directly delete the above installation packages.

[1]: https://github.com/ScoopInstaller/Scoop
