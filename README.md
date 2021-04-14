# pkgmsg
Common updates checker that notifies via `notify-send` command.

```
Usage: pkgmsg [options]
Options:
  -h, --help		print this help
  -v, --version		print version
  -t, --time NUMBER	time interval in seconds to check package manager for new updates (default 3600)
  -x, --xbps		use XBPS as desired package manager
  -p, --pacman		use Pacman as desired package manager
  -a, --apk		use Apk as desired package manager
  -g, --apt-get		use apt-get desired package manager
```