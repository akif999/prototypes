#include <stdio.h>

int main (void) {
    int a = 10;
    int *p = &a;

    printf("a     = %d\n", a);
    printf("&a    = %p\n", (void*)&a);
    printf("p     = %p\n", (void*)p);
    printf("p     = %d\n", *p);

    *p = 20;
    printf("a after *p=20 Å® %d\n", a);

    return (0);
}
