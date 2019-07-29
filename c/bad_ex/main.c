#include<stdio.h>

int* f(int x, int y);

int main(void) {
    int x = 1;
    int y = 2;
    int *p;

    p = f(x, y);

    printf("%p\n", p);
    printf("value is: %d\n", *p); // null pointer dereference; program exits

    return 0;
}

int* f(int x, int y) {
    int p;

    p = x * y;

    printf("%p\n", &p);
    return &p;
}
