package cassandra

import (
	"github.com/boniface/gonosqltests/domain"
	"time"
)

func Create(number int64) domain.Results {

	results := domain.Results{
		Duration: 10,
		End:      time.Now(),
		Objects:  number,
		Start:    time.Now(),
	}
	return results
}
