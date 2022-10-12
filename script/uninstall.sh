#!/bin/bash



rm -f /usr/local/bin/xii
rm -f /usr/local/bin/xxi
if [ "$(uname)" == "Darwin" ] ; then
  rm -f ~/xii
else
  cd /home/xii
fi

echo "Uninstall success"