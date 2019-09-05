package main

import (
	"log"

	"github.com/aoripov/goav/avcodec"
	"github.com/aoripov/goav/avdevice"
	"github.com/aoripov/goav/avfilter"
	"github.com/aoripov/goav/avformat"
	"github.com/aoripov/goav/avutil"
	"github.com/aoripov/goav/swresample"
	"github.com/aoripov/goav/swscale"
)

func main() {

	// Register all formats and codecs
	avformat.AvRegisterAll()
	avcodec.AvcodecRegisterAll()

	log.Printf("AvFilter Version:\t%v", avfilter.AvfilterVersion())
	log.Printf("AvDevice Version:\t%v", avdevice.AvdeviceVersion())
	log.Printf("SWScale Version:\t%v", swscale.SwscaleVersion())
	log.Printf("AvUtil Version:\t%v", avutil.AvutilVersion())
	log.Printf("AvCodec Version:\t%v", avcodec.AvcodecVersion())
	log.Printf("Resample Version:\t%v", swresample.SwresampleLicense())

}
