./hello&
curl http://localhost:8080/hello
res=$?
while test $res != 0
do
  curl http://localhost:8080/hello
  res=$?
done
exit 0