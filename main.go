package main

// https://oss.oetiker.ch/rrdtool/doc/librrd.en.html

/*
#cgo pkg-config: librrd
#include <rrd.h>
//#include <rrd_client.h>
*/
import "C"

import (
	"fmt"
)

func main() {
	fmt.Printf("RRDTool Version: %v\n", C.rrd_version())
}

// TODO: CORE FUNCTIONS
//  rrd_dump_cb_r(char *filename, int opt_header, rrd_output_callback_t cb, void *user)
//  rrd_fetch_cb_register(rrd_fetch_cb_t c)
