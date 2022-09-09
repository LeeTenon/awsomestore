package conf
message Data {
message Mysql {
string host = 1;
}
message Mongo {
string host = 1;
string connectTimeout = 2;
}
message Redis {
string network = 1;
string addr = 2;
google.protobuf.Duration read_timeout = 3;
google.protobuf.Duration write_timeout = 4;
}
Mysql mysql = 1;
Mongo mongo = 2;
Redis redis = 3;
}
type Config struct {

}
