# pttoolkit
Oracle Performance Tuning Toolkit

# Author
My name is Stefano Trallori. You can find me on [Linkedin](https://www.linkedin.com/in/stefanotrallori/)

# General information
Welcome to the pttoolkit. This small project is an attempt to create a small swiss-army-knife-like tool for Oracle performance tuning on linux and linux-like systems.
It's supposed to include all most commonly used tools to diagnose performance and tackle performance tuning problems.

# Directory structure
```
   |-contrib
   |---dgiles
   |-----swingbench
   |-------data
   |-------install
   |---kclosson
   |-----SLOB
   |-------data
   |-------install
   |---mpagano
   |-----sqld360
   |-------data
   |-------install
   |---tpoder
   |-----dstat
   |-------data
   |-------install
   |-----psnapper
   |-------data
   |-------install
   |-----tpt-oracle
   |-------data
   |-------install
   |-oracle
   |---client
   |-----Darwin
   |-------19c
   |-----linux
   |-------19c
   |---network
   |-----admin
   |-----log
   |-----trc
   |---sqlcl
   |-----bin
   |-----lib
   |-------ext
```

Main points are
- Under oracle/ are all binary files (installation files as well as extracted ones) which come from the oracle website.
This would be - in most cases - the oracle client in different flavors / versions.

- Under contrib/ you can find the third party tools. The structure is "contrib/$contributor/$project". Under $project there are two directories: "data" and "install".
"install" contains the install script. "data" is the actual directory where the extracted binaries will live.
As far as possible, the binaries are downloaded "on the fly" using git clone from the official repository.


# Credits
Thanks to 
- Tanel Poder (https://tanelpoder.com)
- Dominic Giles (https://www.dominicgiles.com/index.html)
- Kevin Closson (https://kevinclosson.net)
- Mauro Pagano (https://mauro-pagano.com)

Without you, none of this would have been possible.
