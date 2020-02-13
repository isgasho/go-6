#!/bin/sh

licenses="
dialog/LICENSE.txt
glcaps/LICENSE.txt
humanize/LICENSE.txt
operator/LICENSE.txt
ximage/LICENSE.txt
ximage/xcolor/LICENSE.txt
"

dest="LICENSE.txt"

printf "tawesoft.co.uk/go\n" > "$dest"
for license in $licenses; do
    printf "\n--------------------------------------------------------------------------------\n\n" >> "$dest"
    cat "$license" >> "$dest"
done
