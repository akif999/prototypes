#include <stdio.h>
#include <stdint.h>

union U {
    uint32_t raw;
    uint8_t bytes[4];
};

int main(void) {
    union U u;
    u.raw = 0x11223344;

    for (int i = 0; i < 4;i++) {
        printf("%02X ", u.bytes[i]);
    }
    printf("\n");
}
