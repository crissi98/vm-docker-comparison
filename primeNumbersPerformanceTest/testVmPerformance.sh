(vboxmanage startvm "PopOS VM")&
sleep 10s
hyperfine "curl http://192.168.178.59:8080/primes/100000" --runs 10
vboxmanage controlvm "PopOS VM" savestate
exit 0
