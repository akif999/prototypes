#include <stdio.h>

// #define SQR(x) x*x
#define SQR(x) ((x)*(x))

int main(void) {
    printf("%d\n", SQR(1 + 2));
}
