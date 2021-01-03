(vboxmanage startvm "PopOS VM")&
curl http://localhost:8080/hello
res=$?
while test $res != 0
do
  curl http://192.168.178.59:8080/hello --max-time 0,001
  res=$?
done
exit 0
