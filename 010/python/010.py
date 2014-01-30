import time
import math
start = time.time()

def primo(n):
    if n%2==0: return False    
    return not any(n%i==0 for i in range(3, int(math.sqrt(n))+1,2))  
  
sum = 0

for i in range(2, 2000000):
    if primo(i): 
        sum += i

print 'sum:',sum+2
end = time.time()
timer = end-start
print "time:",timer
      
