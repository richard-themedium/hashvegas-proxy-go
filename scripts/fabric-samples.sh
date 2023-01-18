#!/bin/bash

# Get fabric sample project
curl -LO https://bit.ly/2ysbOFE &&
  mv ./2ysbOFE fabric-installer.sh &&
  chmod 0755 ./fabric-installer.sh

# Install essential binaries
./fabric-installer.sh 2.4.7 1.5.5

