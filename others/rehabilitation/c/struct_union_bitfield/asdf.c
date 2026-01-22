#include <stdio.h>

struct Point {
    int x;
    int y;
};

int main(void) {
    struct Point p = { .x = 10, .y = 20};
    printf("x=%d y=%d\n", p.x, p.y);
    printf("sizeof(Point)=%zu\n", sizeof(struct Point));

    return (0);
}

