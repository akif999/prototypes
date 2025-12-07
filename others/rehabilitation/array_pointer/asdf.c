#include <stdio.h>

int main(void) {
    int arr[3] = {10, 20, 30};
    int *p = arr;

    printf("arr      = %p\n", (void*)arr);
    printf("&arr[0]  = %p\n", (void*)&arr[0]);

    printf("p[1]       = %d\n", p[1]);
    printf("*(arr + 2) = %d\n", *(arr + 2));
}
