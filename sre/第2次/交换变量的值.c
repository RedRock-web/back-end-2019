//不使用第三个变量,交换两个变量的值 
#include <stdio.h>

int main(void)
{
    int a = 0, b = 1;
    a ^= b;
    b ^= a;
    a ^= b;
    
    printf("a = %d, b = %d", a, b);
    
    return 0;
}
