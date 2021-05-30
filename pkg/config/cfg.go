package cfg

import mcon "github.com/carlosd-ss/star-wars/pkg/mongo"

var (
	Connect    = mcon.Connect()
	Db         = "starwars"
	Collection = "planet"
	Port       = ":8080"
)
