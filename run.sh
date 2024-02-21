# /bin/bash
if [ $# -ne 2 ]; then
    go test .
    exit 0
fi
go test $1 $2
exit 0
