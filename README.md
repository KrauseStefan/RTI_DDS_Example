# RTI_DDS_Example
Quick and dirty port of the `C` `Simple_hello` example in go.

The example requires go to be setup correctly and that `lib_wrapper.go` is addapted to the local enviroment.
https://github.com/KrauseStefan/RTI_DDS_Example/blob/master/lib_wrapper/lib_wrapper.go

The example is designed to be downloaded using go get with either (-d is to avoid installing the test executables):
* `go get -d github.com/KrauseStefan/RTI_DDS_Example/writer`
* `go get -d github.com/KrauseStefan/RTI_DDS_Example/reader`
  
  
navigate to the `reader.go` and/or `writer.go` and run wither `go run writer.go` or `go run reader.go`

  
  
  
