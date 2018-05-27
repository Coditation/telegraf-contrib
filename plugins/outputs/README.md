# S3 Output Plugin

This plugin writes telegraf metrics to AWS S3

### Configuration
```
[[outputs.s3]]
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
```
