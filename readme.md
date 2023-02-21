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
scoop cache cleaner (scc) 2.0.0, 2023-02-21

Usage:
  scc <command> [path/to/scoop/cache]
      clean up the specified scoop cache directory.
      if the path is omitted, it will use the path defined in the environment variable %SCOOP%.

Command:
  -l:  list the outdated packages.
  -b:  backup the outdated packages.
  -d:  delete the outdated packages.

all other parameters will display the above information.
```

If path is omitted, it takes the environment value of `%SCOOP%` and appends `cache` to it.

For example, if `%SCOOP%` is `C:\Scoop`, `scc -l` will list outdated packages in `C:\Scoop\cache`.

[1]: https://github.com/ScoopInstaller/Scoop
