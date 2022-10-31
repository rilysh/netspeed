## netspeed
A utility to check your average Internet speed

## Install
### From source
```sh
git clone https://github.com/rilysh/netspeed
cd netspeed
# enable build.sh permission attribute as executable
chmod +x build.sh
./build.sh
```
#### Info
Using GCCGo (which is a front-end of GCC), you can reduce the binary size approxmately 80%. The standard Go compiler is doesn't yet support manual optimzations, where in GCCGo, you can pass all kind of gcc parameters.
