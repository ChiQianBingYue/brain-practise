#include <stdio.h>

int main(int argc, char **argv)
{
    // int num = 0xAAAAAAAAA;
    // int x = 8;
    // x << 4;
    unsigned x = 0x80000000;
    printf("%u, %d\n", x, x);

    return 0;
}
