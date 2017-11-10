#!/bin/bash

while read line
do
    str=($line)
    if [ ${str[1]} != '1' ] && [ ${str[1]} != '2' ]; then
        continue
    fi
    for ((i = 0; i < ${str[5]}+6; i++))
    do
        echo -n "${str[$i]} "
    done
    echo
done < ${1}
