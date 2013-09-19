/*
 * http://en.wikipedia.org/wiki/Prime_number
 * http://www.s-anand.net/euler.html
 */

var prime_list = [2, 3, 5, 7, 11, 13, 17, 19, 23];
var lastn = prime_list[prime_list.length-1];

function _refresh(x) {
	while (lastn <= x){
		lastn = lastn + 1;
		if (_isprime(lastn) > -1){
			prime_list[prime_list.length] = lastn;
		} 
	}
}

function _isprime(n){
	var isprime = -1;
	if ((n>=2) && 1 || 0) {
		isprime = n;
	}
	for (var i=0; i <= prime_list.length; i++){
		if (prime_list[i] * prime_list[i] > n){		// ... up to sqrt(n)
			break;
		}
		if ((n % prime_list[i]) !==0){
			isprime = 0
			break;
		}
	}
	return isprime;
}

function prime(x){
	while (prime_list.length <= x){
		lastn = lastn + 1;
		if (_isprime(lastn) > -1){
			prime_list[prime_list.length] = lastn;
		}
	}
	return prime_list[x];
}

function isprime(x){
	_refresh(x);
	if (prime_list.indexOf(x) == -1){
		return 0;
	} else { 
		return 1; 
	}
}

function num_factors(n){
     var div = 1;
     var x = 0;
     while (n > 1){
         var c = 1;
         while ((n % prime(x)) == 0){
             c = c + 1;
             n = Math.floor(n / prime(x));
		 }
		 x = x + 1;
		 div = div * c;
	 }
     return div
}