#!/bin/bash

if [ "$#" != "1" ]; then
	echo "Must supply the path to an initialized domain"
	exit 1
fi

set -o nounset
set -o errexit

gowhich() {
	WHICH=$(which which)
	echo -n "$(PATH="${GOPATH//://bin:}/bin" $WHICH "$1")"
}

DOMAIN="$1"
FAKE_PASS=BogusPass
TPM="/dev/tpm0"
PCRS="17,18" # PCR registers of TPM 
AIKBLOB="${HOME}/aikblob"

# Make sure we have sudo privileges before using them to try to start linux_host
# below.
sudo test true

sudo "$(gowhich linux_host)" -config_path ${DOMAIN}/tao.config \
	-host_type stacked -host_channel_type tpm \
	-tpm_device $TPM -tpm_pcrs $PCRS -tpm_aik_path $AIKBLOB &
HOSTPID=$!

echo "Waiting for linux_host to start"
sleep 5

DSPID=$("$(gowhich tao_launch)" -sock ${DOMAIN}/linux_tao_host/admin_socket \
	"$(gowhich demo_server)" -config=${DOMAIN}/tao.config)
"$(gowhich tao_launch)" -sock ${DOMAIN}/linux_tao_host/admin_socket \
	"$(gowhich demo_client)" -config=${DOMAIN}/tao.config > /dev/null


echo "Waiting for the tests to finish"
sleep 5

echo "\n\nClient output:"
cat /tmp/demo_client.INFO

echo "\n\nServer output:"
cat /tmp/demo_server.INFO

echo "Cleaning up remaining programs"
kill $DSPID
sudo kill $HOSTPID
sudo rm -f ${DOMAIN}/linux_tao_host/admin_socket
