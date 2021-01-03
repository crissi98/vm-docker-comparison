./primes&
sleep 1s
hyperfine "curl http://localhost:8080/primes/100000" --runs 10
primesProcess=$(pgrep primes)
kill "$primesProcess"