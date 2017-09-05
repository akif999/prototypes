#include <stdio.h>
#include <string.h>
#include <stdlib.h>

#define PG_SECTION_SIZE 368
#define BUF_SIZE        256

int cut_datafield(char*,char*, int);
int notify_srectype(char*);
int notify_totallength(char*);
long notify_dataaddr(char*, int);
int calc_datalength(int, int);
void charactor_to_binary(char*, unsigned char*, int);
void write_srec_bin_obj(unsigned char*, unsigned char*, long, int);
void init_srec_bin_obj(unsigned char*, int);
int convert_blen_to_clen(int);
int convert_clen_to_blen(int);
void clear_buffer();

int main (void) {
    FILE *fp_src, *fp_dst;
    char original_line[BUF_SIZE];
    char proccessed_line[BUF_SIZE];
    unsigned char binary_line[BUF_SIZE];
    unsigned char srec_bin_obj[PG_SECTION_SIZE];
    int srec_type;
    int datasize_by_line;
    long dataaddr_by_line;

    init_srec_bin_obj(srec_bin_obj, (int)sizeof(srec_bin_obj));

    fp_src = fopen("s1rec_sample", "r");
    if (fp_src == NULL) {
        printf("Error:src file open error\n");
        return 1;
    }
    fp_dst = fopen("s1rec_sample.bin", "wb");
    if (fp_dst == NULL) {
        printf("Error:dst file open error\n");
        return 1;
    }

    while(fgets(original_line, 256, fp_src) != NULL) {
        srec_type = notify_srectype(original_line);
        if (srec_type == 1) {
            dataaddr_by_line = notify_dataaddr(original_line, srec_type);
            datasize_by_line = cut_datafield(original_line, proccessed_line, srec_type);

            charactor_to_binary(proccessed_line, binary_line, datasize_by_line);
            write_srec_bin_obj(srec_bin_obj, binary_line, dataaddr_by_line, datasize_by_line);

            clear_buffer(original_line);
            clear_buffer(proccessed_line);
        }
    }
    fwrite(srec_bin_obj, sizeof(unsigned char), PG_SECTION_SIZE, fp_dst);

    fclose(fp_src);
    fclose(fp_dst);

    return 0;
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

int notify_srectype(char* original_line) {
    int type;

    type = (int)original_line[1] - 48;

    return type;
}

int notify_totallength(char* original_line) {
    int length;
    char tmp_len_str[256];

    sprintf(tmp_len_str, "%c%c", original_line[2], original_line[3]);
    length = (int)strtol(tmp_len_str, 0, 16);

    return length;
}

long notify_dataaddr(char* original_line, int type) {
    long dataaddr;
    char tmp_addr_str[256];

    if (type == 1) {
        sprintf(tmp_addr_str, "%c%c%c%c",
            original_line[4], original_line[5], original_line[6], original_line[7]);
        dataaddr = (long)strtol(tmp_addr_str, 0, 16);
    } else {
        dataaddr = 0xFFFFFFFF;
    }
    return dataaddr;
}

int calc_datalength(int total_length, int type) {
    if (type == 1) {
        return total_length - 6;
    } else if (type == 2) {
        return 0xFF;
    } else if (type == 3) {
        return 0xFF;
    } else {
        return 0xFF;
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
        obj[i] = 0xFF;
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
