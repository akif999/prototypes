#include <stdio.h>

struct A {
    char c;
    int  i;
};

int main(void) {
    struct A a[2];
    printf("sizeof(A)=%zu\n", sizeof(struct A));
    printf("total=%zu\n", sizeof(a));
}
