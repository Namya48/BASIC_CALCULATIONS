#include <stdio.h>
int main() {    

    int num1, num2, sum;
    
    printf("Enter first integers:\n");
    scanf("%d", &num1);

    printf("Enter Your Second Integer:\n");
    scanf("%d", &num2);
    
    sum=num1+num2;
    
    printf("%d + %d = %d", num1, num2, sum);
    return 0;
}
