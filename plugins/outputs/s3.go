package s3

import (
	"fmt"
	"io"
	"os"
	"time"
	"strconv"

	"gopkg.in/Clever/pathio.v3"
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/outputs"
	"github.com/influxdata/telegraf/plugins/serializers"
)

type S3 struct {
	Bucket string
	AccessKey string
	SecretKey string

	writers []io.Writer
	closers []io.Closer

	serializer serializers.Serializer
}

var sampleConfig = `
  ## Bucket to write to
  bucket = "bucketname"
  ## AWS Access Key
  access_key = "123456789"
  ## AWS Secret Key
  secret_key = "123456789"

  ## Data format to output.
  ## Each data format has its own unique set of configuration options, read
  ## more about them here:
  ## https://github.com/influxdata/telegraf/blob/master/docs/DATA_FORMATS_OUTPUT.md
  data_format = "influx"
`

func (f *S3) SetSerializer(serializer serializers.Serializer) {
	f.serializer = serializer
}

func (s *S3) Connect() error {
	return nil
}

func (s *S3) Close() error {
	return nil
}

func (s *S3) SampleConfig() string {
	return sampleConfig
}

func (s *S3) Description() string {
	return "Send telegraf metrics to AWS S3"
}

func (s *S3) Write(metrics []telegraf.Metric) error {
	if len(metrics) == 0 {
		return nil
	}
	var output []byte
	var writeErr error = nil
	for _, metric := range metrics {
		var b []byte
		var err error
		b, err = s.serializer.Serialize(metric)		
		if err != nil {
			return fmt.Errorf("failed to serialize message: %s", err)
		}
		output = append(output, b...)
	}
	filename := "metrics-" + strconv.FormatInt(time.Now().Unix(), 10)
	path := "s3://" + s.Bucket + "/" + filename
	/*Set environment variables if AWS keys are provided*/
	if s.AccessKey != "" && s.SecretKey != "" {
		os.Setenv("AWS_ACCESS_KEY_ID", s.AccessKey)
		os.Setenv("AWS_SECRET_ACCESS_KEY", s.SecretKey)		
	}		

	err := pathio.Write(path, output) 

	if err != nil {
		writeErr = fmt.Errorf("E! failed to write message: %s, %s", output, err)
	}	
	return writeErr
}

func init() {
	outputs.Add("s3", func() telegraf.Output {
		return &S3{}
	})
}
