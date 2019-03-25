package main

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
	"time"

	log "github.com/golang/glog"
)

const count = 10
const contentType = "application/x-amz-json-1.1"

/*
Test for APIG -> KDS - service proxy
const url = "https://SOMEDISTRO.execute-api.us-east-1.amazonaws.com/p/i"
const body = "{ \"PartitionKey\": \"sometest\", \"StreamName\": \"ingest-IngestionKinesisDataStream-1FX7KXAIT463P\", \"Data\": \"eyAieWVhIjogMSwgIm5heSI6IDAsICJldmVudElkIjogIm5ld25ld25ld3Rlc3QiLCAidXNlcklkIjogInNvbWV0aGluZ25ldyIgfQo=\" }"
*/

// This is at test for APIG -> SQS service proxy
const url = "https://SOMEDISTRO.execute-api.us-east-1.amazonaws.com/p/s"
const body = "MessageBody=lllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllll"

func init() {
	flag.Parse()
	flag.Lookup("logtostderr").Value.Set("true")
}

func main() {

	start := timestamp()
	var err error

	for i := 0; i < count; i++ {
		if err = request(); err != nil {
			log.Error(err)
		}
	}

	diff := timestamp() - start
	fmt.Println(fmt.Sprintf("T: %d", diff))

	log.Flush()
}

func request() error {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", contentType)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Non 200 response received: %s", resp.Status)
	}

	return nil
}

func timestamp() int64 {
	return time.Now().UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}
