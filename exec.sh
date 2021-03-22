#!/bin/bash

readonly local DIR="${1}"
readonly local QUALITY="70"

if [ -z ${DIR} ]; then
	echo "引数がありません"
	exit
fi

ls -1 ${DIR}/*.jpg | while read SRC; do
	DIST=$(echo "${SRC}" | sed -e 's/\.jpg/.webp/g')
	./trans_image ${SRC} ${DIST} ${QUALITY}
done


