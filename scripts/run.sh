# !/bin/sh

APP_BASE=http://localhost:8081

function getAllEmployees() {
  CMD="curl -sSX GET $APP_BASE/employees | jq ."
  echo 'CMD='$CMD
  /bin/sh -c "$CMD"
}

function getEmployeeById() {
  ID=$1
  CMD="curl -sSX GET $APP_BASE/employees/$ID | jq ."
  echo 'CMD='$CMD
  /bin/sh -c "$CMD"
}

function saveEmployee() {
  NAME=$1
  CMD="curl -sSX POST $APP_BASE/employees -H 'content-type: application/json' -d '{\"name\":\"$NAME\"}' "
  echo 'CMD='$CMD
  /bin/sh -c "$CMD"
}

#saveEmployee Sachin-A
#saveEmployee Sachin-B
#saveEmployee Sachin-C

getAllEmployees
#getEmployeeById 2

