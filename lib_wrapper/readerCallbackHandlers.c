
#include "readerCallbackHandlers.h"
#include "_cgo_export.h"

void printRecivedData(void *listener_data, DDS_DataReader* data_reader){

	DDS_StringDataReader *string_reader = NULL;
	char				  sample[250]; 
	struct DDS_SampleInfo info;

	string_reader = DDS_StringDataReader_narrow(data_reader);

	DDS_StringDataReader_take_next_sample(
						string_reader,
						sample,
						&info);

	MyFunction(sample);
}