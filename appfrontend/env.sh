#!/bin/sh

# skip processing of .env file and use just one expected environment variable

CONFFILE="/usr/share/nginx/html/env-config.js"

echo "window._env_ = {" > $CONFFILE
echo "  REACT_APP_APIHOSTPORT: \"$REACT_APP_APIHOSTPORT\"" >> $CONFFILE
echo "}" >> $CONFFILE

cat $CONFFILE

# ignore the rest of file
exit

## Recreate config file
#echo env.sh script started...
#whoami
#pwd
#ls -la
#echo

#echo "" > ./env-config.js
##rm -rf ./env-config.js
##touch ./env-config.js

## Add assignment 
#echo "window._env_ = {" >> ./env-config.js

## Read each line in .env file
## Each line represents key=value pairs
#while read -r line || [[ -n "$line" ]];
#do
#  # Split env variables by character `=`
#  if printf '%s\n' "$line" | grep -q -e '='; then
#    varname=$(printf '%s\n' "$line" | sed -e 's/=.*//')
#    varvalue=$(printf '%s\n' "$line" | sed -e 's/^[^=]*=//')
#  fi
#
#  # Read value of current variable if exists as Environment variable
#  value=$(printf '%s\n' "${!varname}")
#  # Otherwise use value from .env file
#  [[ -z $value ]] && value=${varvalue}
#  
#  # Append configuration property to JS file
#  echo "  $varname: \"$value\"," >> ./env-config.js
#done < .env

#echo "}" >> ./env-config.js

#cat ./env-config.js
#echo
#echo env.sh script finished!
