for i in $(find wwwroot/src/scss -iname '*.scss'); do
    sass $i $(echo $i | sed 's/^src/build/' | sed 's/scss$/css/') &>/dev/null
done