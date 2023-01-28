# Scoop Cache Cleaner

[中文版](readme-zh.md)

## What is it

`Scoop Cache Cleaner (scc)` is a command line tool used to clean up the [SCOOP][1] cache directory. Although no platform-specified API is used, it only applicable to Windows because [SCOOP][1] only runs on Windows.

## Why not `scoop cache rm *`

Running `scoop cache rm *` will empty the cache directory without leaving any file. However, I still want to keep all latest setup files after running `scoop update *`.

If you do not run `spoon cache rm *` or manually clean the directory, it will occupy a lot of disk space.

## Usage

```text {.line-numbers}
Copyright (c) 1999-2023 Not a dream Co., Ltd.
scoop cache cleaner (scc) 1.0.0, 2023-01-25

Usage:
  scc [path/to/scoop/cache]
      clean up the specified scoop cache directory.
  scc -e
      clean up scoop cache directory defined in the environment.

  all other parameters will display the above information.
```

For parameter `-e`, it takes the environment value of `%SCOOP%` and appends `cache` to it.

For example, if `%SCOOP%` is `C:\Scoop`, `scc -e` will clean `C:\Scoop\cache`.

[1]: https://github.com/ScoopInstaller/Scoop
