# Nmap Clean Output

You have a bunch of scan data from nmap in `-oG` format and you want to get a list in the format of `IP:PORT`, and your bash pipeline is getting so out of hand that you start questioning reality...

## Usage
Convert this:
```
v@netrunner:~/project/scan$ grep -E "Host: ([0-9\.]{7,15}).+Ports:.+([0-9]{1,5})/open+" *.og
nmap-10.2.3.4.og:Host: 10.2.3.4 ()  Ports: 10000/open/tcp//http//COOL SERVICE HERE12/
nmap-10.1.1.50.og:Host: 10.1.1.50 ()  Ports: 80/open/tcp//http//COOL SERVICE HERE httpd/
nmap-10.1.1.51.og:Host: 10.1.1.51 ()  Ports: 80/open/tcp//http//COOL SERVICE HERE httpd/, 443/open/tcp//ssl|https?///
nmap-10.1.1.53.og:Host: 10.1.1.53 ()  Ports: 80/open/tcp//http//COOL SERVICE HERE httpd/, 443/open/tcp//ssl|https?///
nmap-10.1.1.54.og:Host: 10.1.1.54 ()  Ports: 80/open/tcp//http//COOL SERVICE HERE httpd/, 443/open/tcp//ssl|https?///
nmap-10.1.1.55.og:Host: 10.1.1.55 ()  Ports: 80/open/tcp//http//COOL SERVICE HERE httpd/, 443/open/tcp//ssl|https?///
nmap-10.1.1.56.og:Host: 10.1.1.56 ()  Ports: 443/open/tcp//ssl|https?///
nmap-10.1.1.57.og:Host: 10.1.1.57 ()  Ports: 80/open/tcp//http//COOL SERVICE HERE httpd/, 443/open/tcp//ssl|https?///
nmap-10.1.1.102.og:Host: 10.1.1.102 ()  Ports: 80/open/tcp//http//Microsoft HTTPAPI httpd 2.0 (SSDP|UPnP)/, 443/open/tcp//ssl|http//Microsoft HTTPAPI httpd 2.0 (SSDP|UPnP)/
nmap-10.1.1.112.og:Host: 10.1.1.112 (www.localdomain.local)   Ports: 80/open/tcp//http//Microsoft IIS httpd 10.0/
nmap-10.1.1.119.og:Host: 10.1.1.119 ()  Ports: 80/open/tcp//http//Microsoft HTTPAPI httpd 2.0 (SSDP|UPnP)/, 443/open/tcp//ssl|http//Microsoft HTTPAPI httpd 2.0 (SSDP|UPnP)/
nmap-10.1.1.212.og:Host: 10.1.1.212 ()  Ports: 21/open/tcp//ftp?///, 443/open/tcp//ssl|http//COOL SERVICE HERE2/, 990/open/tcp//ssl|nagios-nsca//Nagios NSCA/
nmap-10.1.1.32.og:Host: 10.1.1.32 ()    Ports: 80/open/tcp//http//COOL SERVICE HERE56/, 443/open/tcp//ssl|https?///
nmap-10.1.1.94.og:Host: 10.1.1.94 ()    Ports: 80/open/tcp//http//Microsoft HTTPAPI httpd 2.0 (SSDP|UPnP)/, 443/open/tcp//ssl|https///
nmap-10.1.1.95.og:Host: 10.1.1.95 ()    Ports: 443/open/tcp//ssl|https///
nmap-10.1.1.7.og:Host: 10.1.1.7 (localdomain.local)        Ports: 80/open/tcp//http//COOL SERVICE HERE httpd/, 443/open/tcp//ssl|https?///, 4414/open/tcp//tcpwrapped///, 5254/open/tcp//tcpwrapped///, 7483/open/tcp//tcpwrapped///, 8286/open/tcp//tcpwrapped///, 10298/open/tcp//tcpwrapped///, 11304/open/tcp//tcpwrapped///, 17770/open/tcp//tcpwrapped///, 19260/open/tcp//tcpwrapped///, 20336/open/tcp//tcpwrapped///, 21542/open/tcp//tcpwrapped///, 22109/open/tcp//tcpwrapped///, 23862/open/tcp//tcpwrapped///, 24475/open/tcp//tcpwrapped///, 26707/open/tcp//tcpwrapped///, 27858/open/tcp//tcpwrapped///, 30255/open/tcp//tcpwrapped///, 33096/open/tcp//tcpwrapped///, 36895/open/tcp//tcpwrapped///, 39185/open/tcp//tcpwrapped///, 40966/open/tcp//tcpwrapped///, 43232/open/tcp//tcpwrapped///, 43277/open/tcp//tcpwrapped///, 48640/open/tcp//tcpwrapped///, 54673/open/tcp//tcpwrapped///, 63471/open/tcp//tcpwrapped///, 65255/open/tcp//tcpwrapped///    Ignored State: filtered (46)
nmap-10.1.1.6.og:Host: 10.1.1.6 (home.localdomain.local)  Ports: 443/open/tcp//ssl|https///, 1677/open/tcp//tcpwrapped///, 4272/open/tcp//tcpwrapped///, 4721/open/tcp//tcpwrapped///, 5101/open/tcp//tcpwrapped///, 6084/open/tcp//tcpwrapped///, 8850/open/tcp//tcpwrapped///, 12321/open/tcp//tcpwrapped///, 15098/open/tcp//tcpwrapped///, 17129/open/tcp//tcpwrapped///, 17245/open/tcp//tcpwrapped///, 25398/open/tcp//tcpwrapped///, 25998/open/tcp//tcpwrapped///, 29568/open/tcp//tcpwrapped///, 31150/open/tcp//tcpwrapped///, 31455/open/tcp//tcpwrapped///, 31776/open/tcp//tcpwrapped///, 41429/open/tcp//tcpwrapped///, 42539/open/tcp//tcpwrapped///, 52248/open/tcp//tcpwrapped///, 55303/open/tcp//tcpwrapped///, 55605/open/tcp//tcpwrapped///, 57730/open/tcp//tcpwrapped///, 58894/open/tcp//tcpwrapped///, 59285/open/tcp//tcpwrapped///, 64125/open/tcp//tcpwrapped/// Ignored State: filtered (42)
```

Into this:
```
v@netrunner:~/project/scan$ grep -E "Host: ([0-9\.]{7,15}).+Ports:.+([0-9]{1,5})/open+" *.og | nmapclean
[...]
10.1.1.57:80
10.1.1.57:443
10.1.1.102:80
10.1.1.102:443
10.1.1.112:80
10.1.1.119:80
10.1.1.119:443
10.1.1.212:21
10.1.1.212:443
10.1.1.212:990
10.1.1.32:80
10.1.1.32:443
10.1.1.94:80
10.1.1.94:443
10.1.1.95:443
[...]
```

If you don't want `tcpwrapped` ports, you can do something like:
```
$ grep -E "/open/" *.og | grep -wv "tcpwrapped" | nmapclean
``
