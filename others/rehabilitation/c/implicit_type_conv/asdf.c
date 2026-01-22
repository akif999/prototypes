#include <stdio.h>

int main (void) {
    unsigned int u = 1;
    int i = -1;

//    if (i < u) {
    if (i < (int)(u)) {
        printf("i < u\n");
    }
}
