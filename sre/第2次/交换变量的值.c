//��ʹ�õ���������,��������������ֵ 
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
