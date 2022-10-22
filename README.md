# golangcoding

grpcurl --plaintext localhost:9092 list
Currency
grpc.reflection.v1alpha.ServerReflection


grpcurl --plaintext localhost:9092 list Currenc
y
Currency.GetRate


grpcurl --plaintext localhost:9092 describe Currency.GetRate
Currency.GetRate is a method:
rpc GetRate ( .RateRequest ) returns ( .RateResponse );


grpcurl --plaintext -d '{"Base":"USD","Destination":"Yen"}' localhost:9092 Currency.GetRate
{
  "Rate": 0.5
}