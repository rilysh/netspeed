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
Using GCCGo (which is a front-end of GCC), you can reduce the binary size by approximately 80%. The standard Go compiler doesn't yet support manual optimizations, whereas, in GCCGo, you can pass all kinds of GCC parameters.
