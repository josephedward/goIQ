vars=$("/Applications/Google Chrome.app/Contents/MacOS/Google Chrome" --remote-debugging-port=9222) &
echo "First  word of var: '${vars[0]}'" &&
echo "Second word of var: '${vars[1]}'" &&
echo "Third  word of var: '${vars[2]}'" &&
echo "Number of words in var: '${#vars[@]}'" 