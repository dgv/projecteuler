package euler;

import java.math.BigDecimal;
import java.math.RoundingMode;

public class Q016 {

	public static void main(String[] args) {
		BigDecimal n = (new BigDecimal("2")).pow(1000);
		int soma = 0;
		BigDecimal d = new BigDecimal("10");

		while (n.compareTo(BigDecimal.ZERO) > 0) {
			soma = soma + n.remainder(d).intValue();
			n = n.divide(d).setScale(0, RoundingMode.FLOOR);
		}
		System.out.print(soma);
	}
}
