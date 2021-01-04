(docker run -p 8080:8080 --cpus 1 -m 2G primes)&
sleep 1s
hyperfine "curl http://localhost:8080/primes/100000" --runs 10
docker container kill "$(docker ps -q)"