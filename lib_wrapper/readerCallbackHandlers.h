#ifndef READER_CALLBACK_HANDLERS_H
#define READER_CALLBACK_HANDLERS_H

#include <ndds/ndds_c.h>

void printRecivedData(void *listener_data, DDS_DataReader* data_reader);

#endif