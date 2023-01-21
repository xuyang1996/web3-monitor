set -u
set -e

function echo_error {
    [ $# -eq 0 ] && return
    input="$@"
    DATE_N=`date "+%Y-%m-%d %H:%M:%S"`
    echo "${DATE_N} [ERROR] $input"
}
function echo_info {
    [ $# -eq 0 ] && return
    input="$@"
    DATE_N=`date "+%Y-%m-%d %H:%M:%S"`
    echo "${DATE_N} [INFO] $input"
}


[[ $# -ne 1 ]] && echo_error "Usage: ./$0 api|model|docker" && exit 1
WOKR_DIR=""
type=`uname` 
[[ "${type}" == "Darwin" ]] && WORK_DIR=$(dirname $(realpath $0))
[[ "${type}" == "Linux" ]] && WORK_DIR=$(dirname $(readlink -f $0))
cd ${WORK_DIR}
echo_info "WORK_DIR: ${WORK_DIR}"


function api {
    goctl api go -api ./api/service.api -dir ./ --style=go_zero
    go mod tidy
}

function model {
    goctl model mysql ddl --src=./ddl/*.sql --dir=./internal/model --style=go_zero
}

function docker {
    rm -f Dockerfile
    goctl docker -go main.go
}


case $1 in
api)
    api
;;
model)
    model
;;
docker)
    docker
;;
*)
    echo_error "Usage: $0 api|model|docker" && exit 1
;;
esac