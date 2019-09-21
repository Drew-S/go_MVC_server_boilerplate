for i in $(find wwwroot/src/scss -iname '*.scss'); do
    sass $i $(echo $i | sed 's/\/src\/scss\//\/build\/css\//' | sed 's/scss$/css/')
done