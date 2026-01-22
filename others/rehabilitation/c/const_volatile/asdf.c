#include <stdio.h>

int main(void) {
    int a = 10;
    int b = 20;

    const int *p1 = &a;
    int * const p2 = &b;
}
