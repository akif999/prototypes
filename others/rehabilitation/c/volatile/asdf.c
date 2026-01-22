#include <stdio.h>

volatile int flag = 0;

int main(void) {
    while (flag == 0) {
        // nothing to do
    }
    printf("flag changed\n");
}
