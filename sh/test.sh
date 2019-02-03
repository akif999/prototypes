#!/bin/bash

function is_include() {
    args=$@
    for a in $args
    do
        echo $a
    done
}

# check number of arguments
if [ $# -gt 7 ]; then
    echo "1 arguments required, but has $#"
    exit 1
fi

# processing each line
while read line
do
    str=($line)
    # if line is not record, pass it.
    if [ ${str[1]} != '1' ] && [ ${str[1]} != '2' ]; then
        continue
    fi
    # if hex(iD) field is not equal argument, pass it.
    is_include=is_include $@
    if [   ]; then
        continue
    fi
    for ((i = 0; i < ${str[5]}+6; i++))
    do
            echo -n "${str[$i]} "
    done
    echo
done < ${1}
exit 0
