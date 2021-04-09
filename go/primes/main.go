package main
import "fmt"
import "./sieve"
func main() {
     fmt.Println("Hello from package main")
     primes:=sieve.Sieve()
     cur_prime:= <-primes
     for cur_prime<100 {
     	 fmt.Println(cur_prime)
	 cur_prime=<-primes
     }
}
