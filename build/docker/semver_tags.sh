#!/bin/sh

# This script creates a docker tag list based on the image name ($1), the branch ($2) and the semantic version ($3).
# The result (using thornode, mainnet, 1.2.3 as an example), should generate the following tags...
# mainnet
# mainnet-1
# mainnet-1.2
# mainnet-1.2.3

echo " -t $1:$2 -t $1:$2-$(echo "$3" | awk -F '.' '{print $1}') -t $1:$2-$(echo "$3" | awk -F '.' '{print $1"."$2}') -t $1:$2-$3 "
