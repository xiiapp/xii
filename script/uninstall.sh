#!/bin/bash




if [ "$(uname)" == "Darwin" ] ; then
  echo "Enter your password to uninstall "
  rm -rf ~/xii
  sudo rm -f /usr/local/bin/xii
  sudo rm -f /usr/local/bin/xxi
else
  rm -rf /home/xii
  sudo rm -f /usr/local/bin/xii
  sudo rm -f /usr/local/bin/xxi
fi



echo "Uninstall success"