GO_ENV=test
ROOT_PATH=`cat .env | grep ROOT_PATH| sed -e 's/ROOT_PATH=\(.*\)/\1/'`
ROOT_PATH=`echo $ROOT_PATH` GO_ENV=`echo $GO_ENV` go test test -v . | sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/''
