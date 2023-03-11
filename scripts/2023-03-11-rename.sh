#! /bin/bash

## Rename files from _docs to _posts

for filename in $(find _docs -name '*.md'); do
    committed_date=$(git log -n 1 --oneline --date="format:%Y-%m-%d" --pretty="format:%cd" $filename)
    title=$(grep '^title: ' $filename | cut -d ':' -f 2 | sed 's/ /-/g' )
    if [ -z "$title" ]; then
        echo "No title:" $filename
        exit 1
    fi
    date=$(git grep "date: " $filename)
    if [ -z "$date" ]; then
        sed -i'' -e "0,/---/{s/---/---\ndate: ${committed_date}/}" $filename
    fi

    destination=${committed_date}${title}.md
    echo "from: $filename, to: _posts/$destination"
    mv $filename _posts/$destination
done
