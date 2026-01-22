#include <stdio.h>

void f(int x) {
    x = 100;
}

void ff(int *x) {
    // int big[100000000000];
    *x = 100;
}

int main(void) {
    int a = 10;
    f(a);
    printf("%d\n", a);
    ff(&a);
    printf("%d\n", a);

    return (0);
}
