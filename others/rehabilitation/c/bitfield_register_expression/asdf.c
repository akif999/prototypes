#include <stdio.h>
#include <stdint.h>

struct Reg {
    uint32_t enable : 1;
    uint32_t mode   : 3;
    uint32_t        : 4;
    uint32_t value  : 8;
};

int main(void) {
    struct Reg r = {0};
    r.enable = 1;
    r.mode = 3;
    r.value = 0xAA;

    printf("sizeof(Reg)=%zu\n", sizeof(struct Reg));
}
