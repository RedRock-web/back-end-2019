//产生不同的随机数 
#include <stdio.h>
#include <stdlib.h>
#include <time.h>

void test();
int main(void)
{
    test();
    
    return 0;
}
void test()
{
    static int sum;
    static int randNum;
    
    while(10 > sum)
    {
    	
        srand((unsigned)time(NULL));
        randNum = rand() % 10;
        
        sum += randNum;
        printf("rand num is %d,sum is %d\n",randNum,sum);
    }
    printf("the final sum is %d\n",sum);
}
