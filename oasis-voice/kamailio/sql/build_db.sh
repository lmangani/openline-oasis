#!/bin/bash
FILES="standard-create.sql permissions-create.sql carriers.sql"
if [ "x$1" == "xlocal-kube" ]; then
  while [ -z "$pod" ]; do
    pod=$(kubectl get pods -n openline-development|grep oasis-postgres|grep Running| cut -f1 -d ' ')
    if [ -z "$pod" ]; then
      echo "database not ready waiting"
      sleep 1
    fi
  done
  echo $FILES |xargs cat|kubectl exec -n openline-development -it $pod -- psql -U $SQL_USER $SQL_DATABASE
else
  echo $FILES |xargs cat| PGPASSWORD=$SQL_PASSWORD  psql -h $SQL_HOST $SQL_USER $SQL_DATABASE
fi