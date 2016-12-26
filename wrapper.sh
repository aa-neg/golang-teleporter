echo "hello!"

OUTPUT=$(tp $@)

echo "$?"
echo "$OUTPUT"
# if [ 300 -gt 100 ]
# then
#     echo "Hey thats a large number."
#     pwd
# fi