a = 1
b = 1
c = 0
d = 0

while (c < 4000000):
	c = a + b; 
	a = c + b; 
	b = a + c; 

	if (c < 4000000):
		d += c

print d