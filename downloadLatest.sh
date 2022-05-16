#!/bin/sh
# Copyright 2022 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


set -e

# Determines the operating system.
OS="$(uname)"
if [ "${OS}" = "Darwin" ] ; then
  OSEXT="Darwin"
else
  OSEXT="Linux"
fi

# Determine the latest apigeecli version by version number ignoring alpha, beta, and rc versions.
if [ "${APIGEECLI_VERSION}" = "" ] ; then
  APIGEECLI_VERSION="$(curl -sL https://github.com/apigee/apigeecli/releases/latest | \
                  grep -i release | grep -o 'v[0-9].[0-9][0-9][0-9]' | tail -1)"
  APIGEECLI_VERSION="${APIGEECLI_VERSION##*/}"
fi

LOCAL_ARCH=$(uname -m)
if [ "${TARGET_ARCH}" ]; then
    LOCAL_ARCH=${TARGET_ARCH}
fi

case "${LOCAL_ARCH}" in
  x86_64)
    APIGEECLI_ARCH=x86_64
    ;;
  armv8*)
    APIGEECLI_ARCH=arm64
    ;;
  aarch64*)
    APIGEECLI_ARCH=arm64
    ;;
  *)
    echo "This system's architecture, ${LOCAL_ARCH}, isn't supported"
    exit 1
    ;;
esac

if [ "${APIGEECLI_VERSION}" = "" ] ; then
  printf "Unable to get latest apigeecli version. Set APIGEECLI_VERSION env var and re-run. For example: export APIGEECLI_VERSION=v1.104"
  exit 1;
fi

NAME="apigeecli_$APIGEECLI_VERSION"
URL="https://github.com/apigee/apigeecli/releases/download/${APIGEECLI_VERSION}/apigeecli_${APIGEECLI_VERSION}_${OSEXT}_${APIGEECLI_ARCH}.zip"

download_cli() {
  printf "\nDownloading %s from %s ...\n" "$NAME" "$URL"
  if ! curl -o /dev/null -sIf "$URL"; then
    printf "\n%s is not found, please specify a valid APIGEECLI_VERSION and TARGET_ARCH\n" "$URL"
    exit 1
  fi
  curl -fsLO "$URL"
  filename="apigeecli_${APIGEECLI_VERSION}_${OSEXT}_${APIGEECLI_ARCH}.zip"
  unzip "${filename}"
  rm "${filename}"
}


download_cli

printf ""
printf "\napigeecli %s Download Complete!\n" "$APIGEECLI_VERSION"
printf "\n"
printf "apigeecli has been successfully downloaded into the %s folder on your system.\n" "$NAME"
printf "\n"

while true; do
    read -p "Do you want to move the apigeecli binary to /usr/bin?" yn
    case $yn in
        [Yy]* ) break;;
        [Nn]* ) exit 0;;
        * ) echo "Please enter yes or no.";;
    esac
done

sudo mv -f apigeecli_${APIGEECLI_VERSION}_*/apigeecli /usr/bin
