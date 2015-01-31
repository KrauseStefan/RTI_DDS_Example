package lib_wrapper

/*

#cgo CFLAGS: -DRTI_UNIX -DRTI_LINUX -DRTI_64BIT -m64 -I . -I/home/stefan/RTI/ndds.5.1.0/include/ -I/home/stefan/RTI/ndds.5.1.0/include/ndds
#cgo LDFLAGS: -m64 -Wl,--no-as-needed -ldl -lnsl -lm -lpthread -lrt -lnddscz -lnddscorez -L/home/stefan/RTI/ndds.5.1.0/lib/x64Linux2.6gcc4.4.5 -L .

#include <stdio.h>
#include "readerCallbackHandlers.h"
#include <ndds/ndds_c.h>


*/
import "C"

import (
	"fmt"
	"unsafe"
)

type DomainParticipant struct {
	init     bool
	realInst *C.DDS_DomainParticipant
}

type Topic struct {
	init     bool
	realInst *C.DDS_Topic
}

type StringDataWriter struct {
	init     bool
	realInst *C.DDS_StringDataWriter
}

type StringDataReader struct {
	init     bool
	realInst *C.DDS_DataReader
}

func (p *DomainParticipant) Error() string {
	if p.realInst == nil && p.init {
		return "Unable to create domain participant."
	}
	if p.realInst != nil {
		return "DomainParticipant is valid"
	}
	return "DomainParticipant is initialized incorectly"
}

func Create_participant() (participant DomainParticipant) {
	participantPtr := C.DDS_DomainParticipantFactory_create_participant(C.DDS_DomainParticipantFactory_get_instance(), 0, &C.DDS_PARTICIPANT_QOS_DEFAULT, nil, C.DDS_STATUS_MASK_NONE)
	dp := DomainParticipant{realInst: participantPtr, init: true}

	return dp
}

func (p *DomainParticipant) CreateTopic(topicName string) Topic {
	topicNameCharPtr := C.CString(topicName)
	defer C.free(unsafe.Pointer(topicNameCharPtr))
	topic := C.DDS_DomainParticipant_create_topic(p.realInst,
		topicNameCharPtr,
		C.DDS_StringTypeSupport_get_type_name(), /* Type name */
		&C.DDS_TOPIC_QOS_DEFAULT,                /* Topic QoS */
		nil, /* Listener  */
		C.DDS_STATUS_MASK_NONE)

	return Topic{init: true, realInst: topic}
}

func (p *DomainParticipant) CreateStringDatawriter(topic Topic) StringDataWriter {

	dataWriter := C.DDS_DomainParticipant_create_datawriter(
		p.realInst,
		topic.realInst,
		&C.DDS_DATAWRITER_QOS_DEFAULT, /* QoS */
		nil, /* Listener */
		C.DDS_STATUS_MASK_NONE)

	stringDataWriter := C.DDS_StringDataWriter_narrow(dataWriter)

	return StringDataWriter{init: true, realInst: stringDataWriter}
}

func (p *DomainParticipant) CreateDatareader(topic Topic) StringDataReader {

	//listener := C.DDS_DataReaderListener_INITIALIZER
	on_data_available := C.DDS_DataReaderListener_DataAvailableCallback(C.printRecivedData)
	listener := C.struct_DDS_DataReaderListener{on_data_available: on_data_available}

	var dataReader *C.DDS_DataReader = C.DDS_DomainParticipant_create_datareader(
		p.realInst,
		topic.realInst._as_TopicDescription,
		&C.DDS_DATAREADER_QOS_DEFAULT, /* QoS */
		&listener,                     /* Listener */
		C.DDS_DATA_AVAILABLE_STATUS)

	stringDataReader := C.DDS_StringDataReader_narrow(dataReader)

	return StringDataReader{init: true, realInst: stringDataReader}
}

//func (sw *StringDataReader) takeNextSample(message string) bool {
//	messageCharPtr := C.CString(message)

//	retcode := C.DDS_StringDataReader_take_next_sample(
//		sw.realInst,
//		messageCharPtr,
//		&C.DDS_HANDLE_NIL)

//	C.free(unsafe.Pointer(messageCharPtr))

//	return retcode == C.DDS_RETCODE_OK
//}

func (sw *StringDataWriter) Write(message string) bool {
	messageCharPtr := C.CString(message)

	retcode := C.DDS_StringDataWriter_write(
		sw.realInst,
		messageCharPtr,
		&C.DDS_HANDLE_NIL)

	C.free(unsafe.Pointer(messageCharPtr))

	return retcode == C.DDS_RETCODE_OK
}

////export DataCallback
func DataCallback(listener_data unsafe.Pointer, data_reader *C.DDS_DataReader) { //listener_data *C.DDS_DataReader) {
	fmt.Println("data recived")
}

//export MyFunction
func MyFunction(msg *C.char) { //listener_data *C.DDS_DataReader) {
	message := C.GoString(msg)
	fmt.Println(message)
}
