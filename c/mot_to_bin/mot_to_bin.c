#include <stdio.h>
#include <string.h>
#include <stdlib.h>

#define EXITCODE_SUCCESS (0)
#define EXITCODE_FAILED  (1)

#define RECORD_MAXSIZE        (256)
#define RECORD_LINE_MAXNUM    (512)
#define RECORD_BINARY_MAXSIZE (128)
#define OUTPUT_BINARY_MAXSIZE (2048)

#define LENGTH_FIELD_BYTE_SIZE     (1)
#define S1_ADDRESS_FIELD_BYTE_SIZE (2)

#define SREC_INITIALIZE_BYTE (0xFF)

#define INVALID_DATA_LENGTH  (0xFF)
#define INVALID_DATA_ADDRESS (0xFFFFFFFF)

typedef enum s_record_type {
    S1 = 1,
    S2,
    S3,
    S4,
    S5,
    S6,
    S7,
    S8,
    S9,
    INVALID_TYPE
} srec_type;

typedef struct s_record {
    srec_type     type;
    unsigned long length;
    unsigned long address;
    unsigned char data[RECORD_BINARY_MAXSIZE];
    unsigned char checksum;
} srec;

typedef struct s_records {
    srec          records[RECORD_LINE_MAXNUM];
    unsigned long number_of_data_records;
} srecs;

srec_type get_srec_type(char *line);
unsigned long get_length(char *line);
unsigned long get_address(char *line, srec_type);
unsigned long get_data(srec *rec, char *line);
unsigned long get_datalength(unsigned long length, srec_type type);
void string_to_bytes(char *str, unsigned char *bytes, unsigned long str_len);
void make_binary_file(srecs recs, FILE *fp);
void filling_bytes(unsigned char *obj, unsigned long size, unsigned char filler);
unsigned long get_record_binary_size(srecs recs);
unsigned long get_characterlength(unsigned long byte_length);
void dump_bytes(unsigned char *bytes, unsigned long size);

int main (int argc, char **argv) {
    FILE *fp_src, *fp_dst;
    srecs srecords;
    char record_line[RECORD_MAXSIZE];
    int rec_idx;

    if (argc < 2) {
        printf("Error: program needs argument of <filename>\n");
        return EXITCODE_FAILED;
    }

    fp_src = fopen(argv[1], "r");
    if (fp_src == NULL) {
        printf("Error: src file open error\n");
        return EXITCODE_FAILED;
    }

    rec_idx = 0;
    while(fgets(record_line, RECORD_MAXSIZE, fp_src) != NULL) {
        srec srecord;

        srecord.type = get_srec_type(record_line);
        if (srecord.type == S1) {
            srecord.length   = get_length(record_line);
            srecord.address  = get_address(record_line, srecord.type);
            get_data(&srecord, record_line);

            srecords.records[rec_idx] = srecord;
            rec_idx++;
        } else {
            printf("Warning: skiped parsing line (not S1 record): %s", record_line);
        }
    }
    srecords.number_of_data_records = rec_idx;

    fp_dst = fopen("output.bin", "wb");
    if (fp_dst == NULL) {
        printf("Error: dst file open error\n");
        return EXITCODE_FAILED;
    }
    make_binary_file(srecords, fp_dst);

    fclose(fp_src);
    fclose(fp_dst);

    return EXITCODE_SUCCESS;
}

srec_type get_srec_type(char *line) {
    srec_type type;

    if (line[0] == 'S') {
        type = (srec_type)line[1] - 48;
    } else {
        type = INVALID_TYPE;
    }

    return type;
}

unsigned long get_length(char* line) {
    unsigned long len;
    char len_str[3];

    sprintf(len_str, "%c%c", line[2], line[3]);
    len = (unsigned long)strtol(len_str, 0, 16);

    return len;
}

unsigned long get_address(char* line, srec_type type) {
    unsigned long addr;
    char tmp_addr_str[9];

    if (type == S1) {
        sprintf(tmp_addr_str, "%c%c%c%c", line[4], line[5], line[6], line[7]);
        addr = (unsigned long)strtol(tmp_addr_str, 0, 16);
    } else {
        addr = INVALID_DATA_ADDRESS;
    }

    return addr;
}

unsigned long get_data(srec *rec, char *line) {
    char          str[RECORD_MAXSIZE];
    unsigned long  byte_datalen;

    byte_datalen = get_datalength(rec->length, rec->type);

    if (rec->type == S1) {
        strncpy(str, line + 8, get_characterlength(byte_datalen));
        strcat(str, "\n");
    } else {
        // not supported S2, S3 type
    }

    string_to_bytes(str, rec->data, get_characterlength(byte_datalen));

    return byte_datalen;
}

unsigned long get_datalength(unsigned long length, srec_type type) {
    if (type == S1) {
        return length - (LENGTH_FIELD_BYTE_SIZE + S1_ADDRESS_FIELD_BYTE_SIZE);
    } else if (type == S2) {
        return INVALID_DATA_LENGTH;
    } else if (type == S3) {
        return INVALID_DATA_LENGTH;
    } else {
        return INVALID_DATA_LENGTH;
    }
}

void string_to_bytes(char *str, unsigned char *bytes, unsigned long str_len) {
    char tmp_chrs[3];
    int i;

    for (i = 0; i < str_len; i += 2) {
        sprintf(tmp_chrs, "%c%c", str[i], str[i + 1]);
        bytes[i / 2] = (unsigned char)strtol(tmp_chrs, 0, 16);
    }
}

void make_binary_file(srecs recs, FILE *fp) {
    unsigned char bytes[OUTPUT_BINARY_MAXSIZE];
    int i;

    filling_bytes(bytes, OUTPUT_BINARY_MAXSIZE, SREC_INITIALIZE_BYTE);

    for (i = 0; i < recs.number_of_data_records; i++) {
        memcpy(bytes + recs.records[i].address, recs.records[i].data,
               get_datalength(recs.records[i].length, recs.records[i].type));
    }
    fwrite(bytes, sizeof(unsigned char), get_record_binary_size(recs), fp);
}

void filling_bytes(unsigned char *obj, unsigned long size, unsigned char filler) {
    int i;

    for (i = 0; i < size; i++) {
        obj[i] = filler;
    }
}

unsigned long get_record_binary_size(srecs recs) {
    unsigned long start;
    unsigned long end;
    unsigned long end_record_datasize;

    start               = recs.records[0].address;
    end                 = recs.records[recs.number_of_data_records -1].address;
    end_record_datasize = get_datalength(recs.records[recs.number_of_data_records -1].length,
                          recs.records[recs.number_of_data_records -1].type);
    return (end - start) + end_record_datasize;
}

unsigned long get_characterlength(unsigned long byte_length) {
    return byte_length * 2;
}

void dump_bytes(unsigned char *bytes, unsigned long size) {
    int i;

    for (i = 0; i < size; i++) {
        printf("%02X", bytes[i]);
        if (i % 16 == 0) {
            printf("\n");
        }
    }
    printf("\n");
}
