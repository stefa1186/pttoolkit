#!/bin/bash
if [ -d instantclient_19_3 ]; then
  rm -rf instantclient_19_3 
fi
for file in $(ls -1 instantclient*19.3.0.0.0*.zip)
do
 unzip $file
 if [ ${?} -ne 0 ]; then
   exit 1
 fi
done
