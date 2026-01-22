#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main(void) {

    // NG
    int a = 1;
    int b = 1;
    // if (a = b) {}
    if (a == b) {}
    const char src[10] = "123456789";
    char buf[10];
    // strcpy(buf, src);
    memcpy(buf, src, sizeof(buf));
    printf("%s\n", buf);

    int ptr[10];
    printf("%llu\n", sizeof(ptr) / sizeof(ptr[0]));

    int *p = malloc(10);
    free(p);
    // p[0] = 100;
    p[0] = (int)(NULL);
    printf("%d", p[0]);

    return (0);
}
