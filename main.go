package main

import (
	// "context"
	// "context"
	"flag"
	// "fmt"
	// "log"

	// "price-fetcher/client"
	// "fmt"
	// "log"
)

func main() {
	// client  :=  client.New("http://0.0.0.0:3000")
	// price , err :=  client.FetchPrice(context.Background(),"ETH")
	// if err!= nil{
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%+v\n",price)
	// return

	listenAddr := flag.String("listernaddr", ":3000", "listen adress the service is running")
	flag.Parse()
	svc := NewLoggingService(NewMetricServce(&priceFecther{}))

	server := NewJSONAPIServer(*listenAddr, svc)
	server.Run()
}
