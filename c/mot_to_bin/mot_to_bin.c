#include <stdio.h>
#include <string.h>
#include <stdlib.h>

#define EXITCODE_SUCCESS (0)
#define EXITCODE_FAILED  (1)

#define MAX_RECORD_SIZE        (256)
#define MAX_RECORD_LINE_NUM    (512)
#define MAX_RECORD_BINARY_SIZE (128)
#define MAX_OUTPUT_BINARY_SIZE (2048)

#define SREC_INTIALIZE_BYTE (0xFF)

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
    srec_type                             type;
    unsigned long                         length;
    unsigned long                         address;
    unsigned char                         data[MAX_RECORD_BINARY_SIZE];
    unsigned char                         checksum;
} srec;

typedef struct s_records {
    srec          records[MAX_RECORD_LINE_NUM];
    unsigned long binary_size;
} srecs;

int cut_datafield(char*,char*, int);
srec_type get_srec_type(char *line);
int notify_totallength(char*);
long notify_dataaddr(char*, srec_type);
int calc_datalength(int, int);
void charactor_to_binary(char*, unsigned char*, int);
void write_srec_bin_obj(unsigned char*, unsigned char*, long, int);
void init_srec_bin_obj(unsigned char*, int);
int convert_blen_to_clen(int);
int convert_clen_to_blen(int);
void clear_buffer();

int main (int argc, char **argv) {
    FILE *fp_src, *fp_dst;
    srecs srecords;
    char record_line[MAX_RECORD_SIZE];
    char proccessed_line[MAX_RECORD_SIZE];
    unsigned char binary_line[MAX_RECORD_BINARY_SIZE];
    unsigned char srec_bin_obj[MAX_OUTPUT_BINARY_SIZE];
    // int srec_type;
    int datasize_by_line;
    long dataaddr_by_line;

    if (argc < 2) {
        printf("Error: program needs argument of <filename>\n");
        return EXITCODE_FAILED;
    }

    fp_src = fopen(argv[1], "r");
    if (fp_src == NULL) {
        printf("Error: src file open error\n");
        return EXITCODE_FAILED;
    }
    fp_dst = fopen("output.bin", "wb");
    if (fp_dst == NULL) {
        printf("Error: dst file open error\n");
        return EXITCODE_FAILED;
    }

    init_srec_bin_obj(srec_bin_obj, (int)sizeof(srec_bin_obj));

    while(fgets(record_line, MAX_RECORD_SIZE, fp_src) != NULL) {
        srec srecord;

        srecord.type = get_srec_type(record_line);
        if (srecord.type == S1) {
            dataaddr_by_line = notify_dataaddr(record_line, srecord.type);
            datasize_by_line = cut_datafield(record_line, proccessed_line, srecord.type);

            charactor_to_binary(proccessed_line, binary_line, datasize_by_line);
            write_srec_bin_obj(srec_bin_obj, binary_line, dataaddr_by_line, datasize_by_line);

            clear_buffer(record_line);
            clear_buffer(proccessed_line);
        } else {
            printf("Warning: skiped parsing line: %s", record_line);
        }
    }
    fwrite(srec_bin_obj, sizeof(unsigned char), MAX_OUTPUT_BINARY_SIZE, fp_dst);

    fclose(fp_src);
    fclose(fp_dst);

    return EXITCODE_SUCCESS;
}

int cut_datafield(char* original_line, char* proccessed_line, int type) {
    int length;

    length = notify_totallength(original_line);
    length = convert_blen_to_clen(length);
    length = calc_datalength(length, type);

    strncpy(proccessed_line, original_line + 8, length);
    strcat(proccessed_line, "\n");

    return convert_clen_to_blen(length);
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

int notify_totallength(char* original_line) {
    int length;
    char tmp_len_str[256];

    sprintf(tmp_len_str, "%c%c", original_line[2], original_line[3]);
    length = (int)strtol(tmp_len_str, 0, 16);

    return length;
}

long notify_dataaddr(char* original_line, srec_type type) {
    long dataaddr;
    char tmp_addr_str[256];

    if (type == 1) {
        sprintf(tmp_addr_str, "%c%c%c%c",
            original_line[4], original_line[5], original_line[6], original_line[7]);
        dataaddr = (long)strtol(tmp_addr_str, 0, 16);
    } else {
        dataaddr = INVALID_DATA_ADDRESS;
    }
    return dataaddr;
}

int calc_datalength(int total_length, int type) {
    if (type == 1) {
        return total_length - 6;
    } else if (type == 2) {
        return INVALID_DATA_LENGTH;
    } else if (type == 3) {
        return INVALID_DATA_LENGTH;
    } else {
        return INVALID_DATA_LENGTH;
    }
}

void charactor_to_binary(char* src, unsigned char* dst, int size) {
    char tmp_chrs[3];
    int i;

    for (i = 0; i < convert_blen_to_clen(size); i += 2) {
        sprintf(tmp_chrs, "%c%c", src[i], src[i + 1]);
        dst[i / 2] = (unsigned char)strtol(tmp_chrs, 0, 16);
    }
}

void write_srec_bin_obj(unsigned char* obj, unsigned char* line, long addr, int size) {
    memcpy(obj + addr, line, size);
}

void init_srec_bin_obj(unsigned char* obj, int size) {
    int i;

    for (i = 0; i < size; i++) {
        obj[i] = SREC_INTIALIZE_BYTE;
    }
}

int convert_clen_to_blen(int length) {
    return length / 2;
}

int convert_blen_to_clen(int length) {
    return length * 2;
}

void clear_buffer(char* buf) {
    memset(buf, '\0', strlen(buf));
}
