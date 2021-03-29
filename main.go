package main

import (
	"context"
	"log"
	"net"
	"strings"
	"time"

	"arzbrooz.com/currencyPriceProto"
	"github.com/go-redis/redis/v8"

	"google.golang.org/grpc"
)

type server struct {
	rdb *redis.Client
}

func (s *server) GetCurrencyPrice(req *currencyPriceProto.RequestCurrencyPrice, stream currencyPriceProto.CurrencyPrice_GetCurrencyPriceServer) error {
	for {
		time.Sleep(10 * time.Second)
		currenciesResponse := []*currencyPriceProto.CurrencyInfo{}
		val, _ := s.rdb.Do(context.TODO(), "keys", "MARKET_*").Result()
		for _, value := range val.([]interface{}) {
			currencyPrice, _ := s.rdb.Do(context.TODO(), "get", value).Result()
			currencyMarketPrice := &currencyPriceProto.CurrencyInfo{
				Market: strings.Split(value.(string), "_")[1],
				Price:  currencyPrice.(string),
			}
			currenciesResponse = append(currenciesResponse, currencyMarketPrice)
		}
		stream.Send(&currencyPriceProto.ResponseCurrencyPrice{
			CurrencyPriceList: currenciesResponse,
		})
	}
}

func main() {
	log.Println("starting grpc service in port 7070")
	lis, _ := net.Listen("tcp", "0.0.0.0:7070")
	rdb := redis.NewClient(&redis.Options{Addr: "0.0.0.0:6379", Password: "", DB: 0})
	defer rdb.Close()
	grpcServer := grpc.NewServer()
	currencyPriceProto.RegisterCurrencyPriceServer(grpcServer, &server{rdb: rdb})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("error in serving grpc: %v", err)
	}

}
