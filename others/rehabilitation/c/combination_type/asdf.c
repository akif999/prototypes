#include <stdio.h>

typedef void (*StateFunc)(void);

void state_init(void) { printf("Init\n"); }
void state_word(void) { printf("Work\n"); }
void state_error(void) { printf("Error\n"); }

StateFunc table[] = {
    state_init,
    state_word,
    state_error
};

int main(void) {
    for (int i = 0; i < 3; i++) {
        table[i]();
    }
    return(0);
}
