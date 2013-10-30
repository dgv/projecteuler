#include <stdio.h>
int main(){
    int a, b, c, d; 
    for (a=1; a<1000; a++){   
	for (b=a; b<1000; b++){                
	    c=1000-a-b;	
	    if (a*a+b*b==c*c){
		d = a*b*c;
		printf("%d*%d*%d=%d\n",a,b,c,d);
	    }
	}
    }
}