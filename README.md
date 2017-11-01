Introduction:
=============

This is **Go** version of [C library](https://sourceforge.net/p/rootsh/code/ci/master/tree/)

Rootsh is a wrapper for a shell which will make a copy of everything printed
on your terminal. Its main purpose is to give ordinary users a shell with
root privileges while keeping an eye on what they type. This is accomplished
by allowing them to execute rootsh via the sudo command. Unlike a simple
"sudo -s" which is the usual way doing this, "sudo rootsh" will send their
terminal keystrokes and output to a logfile and eventually to a remote
syslog server, where they are out of reach and safe from manipulation.

Usage:
======

Local machine:
```bash
$ sudo ./rootsh
[root@arch rootsh]# ls
Gopkg.lock  Gopkg.toml	log.txt  README.md  rootsh  rootsh.go  vendor
[root@arch rootsh]# cal
    November 2017
Su Mo Tu We Th Fr Sa
          1  2  3  4
 5  6  7  8  9 10 11
12 13 14 15 16 17 18
19 20 21 22 23 24 25
26 27 28 29 30

[root@arch rootsh]# date
Wed Nov  1 13:43:45 +04 2017
[root@arch rootsh]# exit
exit
```

Local log:
```bash
$ cat log.txt
[root@arch rootsh]# ls
Gopkg.lock  Gopkg.toml	log.txt  README.md  rootsh  rootsh.go  vendor
[root@arch rootsh]# cal
    November 2017
Su Mo Tu We Th Fr Sa
          1  2  3  4
 5  6  7  8  9 10 11
12 13 14 15 16 17 18
19 20 21 22 23 24 25
26 27 28 29 30

[root@arch rootsh]# date
Wed Nov  1 13:43:45 +04 2017
[root@arch rootsh]# exit
exit
```

Remote syslog:
```bash
Nov  1 13:43:21 arch rootsh[25105]: #033]0;root@arch:/home/.../rootsh#007[root@arch rootsh]# ls#015
Nov  1 13:43:21 arch rootsh[25105]: Gopkg.lock  Gopkg.toml#011log.txt  README.md  rootsh  rootsh.go  vendor#015
Nov  1 13:43:27 arch rootsh[25105]: #033]0;root@arch:/home/.../rootsh#007[root@arch rootsh]# cal#015
Nov  1 13:43:27 arch rootsh[25105]:     November 2017   #015
Nov  1 13:43:27 arch rootsh[25105]: Su Mo Tu We Th Fr Sa#015
Nov  1 13:43:27 arch rootsh[25105]:          #033[7m 1#033[27m  2  3  4 #015
Nov  1 13:43:27 arch rootsh[25105]:  5  6  7  8  9 10 11 #015
Nov  1 13:43:27 arch rootsh[25105]: 12 13 14 15 16 17 18 #015
Nov  1 13:43:27 arch rootsh[25105]: 19 20 21 22 23 24 25 #015
Nov  1 13:43:27 arch rootsh[25105]: 26 27 28 29 30       #015
Nov  1 13:43:27 arch rootsh[25105]:                      #015
Nov  1 13:43:45 arch rootsh[25105]: #033]0;root@arch:/home/.../rootsh#007[root@arch rootsh]# date#015
Nov  1 13:43:45 arch rootsh[25105]: Wed Nov  1 13:43:45 +04 2017#015
Nov  1 13:43:47 arch rootsh[25105]: #033]0;root@arch:/home/.../rootsh#007[root@arch rootsh]# exit#015
Nov  1 13:43:47 arch rootsh[25105]: exit#015
```
