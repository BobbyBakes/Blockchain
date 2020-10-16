#!/usr/bin/env bash


# Run test coverage on each subdirectories and merge the coverage profile.

echo "Testing with race detection and producing coverage report"
set +x
echo "mode: atomic" > coverage.txt

FAIL=0

# Standard go tooling behavior is to ignore dirs with leading underscores.
for dir in ./ $(find . -maxdepth 10 -not -path '.' -not -path './.git*' \
                     -not -path '*/_*' -not -path './cmd*' -not -path './release*' \
                     -not -path './vendor*' -type d)
do
    if ls $dir/*.go &> /dev/null; then
        GORACE="halt_on_error=1" go test -race -covermode=atomic -coverprofile=$dir/profile.tmp $dir
        if [ $? != 0 ]; then
            FAIL=1
        fi
        if [ -f $dir/profile.tmp ]; then
            cat $dir/profile.tmp | tail -n +2 >> coverage.txt
            rm $dir/profile.tmp
        fi
    fi
done

exit $FAIL
