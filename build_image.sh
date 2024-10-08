#!/bin/bash

now=`date -d today +"%Y-%m-%d-%H-%M-%S"`
img=mqtt_brige:$now

docker build -t $img .

rs="last build ok: \n$img\n"

echo -e "$rs"
echo -e "$rs" >> last_build_mqtt_brige.txt
