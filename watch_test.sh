echo "test watching..."
fswatch -0 src/ | xargs -0 -n 1 ./test.sh
