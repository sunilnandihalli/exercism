package sieve

func Sieve() <-chan int {
     gen := func () (ch chan int) {
     	 ch = make(chan int)
	 go func () {
	    for i:=2;true;i++ {
	    	ch<-i
	    }
	 }();
	 return
     }

     filter := func(ch chan int, x int) chan int {
     	ret := make(chan int)
	go func() {
	   for nx:= range ch {
	       if nx%x != 0 {
	       	  ret <- nx
	       }
	    
	   }
	}()
	return ret
     }
     ret:=make(chan int)
     cur_chan := gen()
     go func() {
     	for {
	    nxt_prime:= <- cur_chan
	    ret<-nxt_prime
	    cur_chan = filter(cur_chan,nxt_prime)
	}
     }()
     return ret
     
}