#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main(void) {
    {
        int arr[3] = {10, 20, 30};
        int *p = arr;

        printf("p[1]=%d\n", p[1]);
        printf("*(p + 1)=%d\n", *(p + 1));
    }

    {
        int a[10];
        int *p = a;

        printf("sizeof(a)=%llu\n", sizeof(a));
        printf("sizeof(p)=%llu\n", sizeof(p));
    }

    {
        char s[] = "abc";
        char *p = s;

        printf("%llu\n", sizeof(s));
        printf("%s\n", s);
        while (*p) {
          putchar(*p);
          p++;
        }
        printf("\n");
    }

    {
        char *buf = malloc(100);
        if (!buf) {
            return (1);
        }

        strcpy(buf, "hello");
        printf("%s\n", buf);
    }
}
