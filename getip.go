package main

import (
    "fmt"
    "net"
    "flag"
    "os"
)

func main(){
    var input string

	fullDetails:=  flag.Bool("full",false,"Get full details (IP and domain, if not set only print IP)")
	
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "[GetIP] : a binary to getIp from CLI v:0.1\n[Usage] : echo 'hostname.tld' | getip\n")
		flag.VisitAll(func(f *flag.Flag) {
			fmt.Fprintf(os.Stderr, "\t-%v: %v\n", f.Name,f.Usage) // f.Name, f.Value
		})
	}

	flag.Parse()



    for {
        _, err := fmt.Scanln(&input)
        if err != nil {
            break
            
        }else{
            ips, _ := net.LookupIP(input)
            for _, ip := range ips {
                if ipv4 := ip.To4(); ipv4 != nil {
                    if(*fullDetails){
                        fmt.Println(ipv4,input)
                    }else{
                        fmt.Println(ipv4)
                    }
                    
                }
            }
        }
    }
}

