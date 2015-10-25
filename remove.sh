for file in `find *.orig -type f -print .`
do
	echo "Deleting file $file"
	git rm $file -f       
done
