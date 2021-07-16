# dlevel
Get any level of subdomain from ```1....N``` 
## Install

```bash
go get -u github.com/MPaandeey/dlevel
```
## Usage Example
```
📄 files.txt
hackerone.com
info.hackerone.com
managed.hackerone.com
forwarding.hackerone.com
a.ns.hackerone.com
b.ns.hackerone.com

​▶ ​​cat files.txt | dlevel 2
ns.hackerone.com
info.hackerone.com
managed.hackerone.com
forwarding.hackerone.com

▶ cat files.txt | dlevel 3
a.ns.hackerone.com
b.ns.hackerone.com

```
## Credits
Inspired by [0xdln1](https://github.com/0xdln1/getlevels) and [tomnomnom](https://github.com/tomnomnom/anew) code for help to remove duplicates .