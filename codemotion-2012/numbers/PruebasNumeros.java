package com.lantoli.codemotion2012.numbers;

import static org.junit.Assert.*;

import org.junit.Test;

public class PruebasNumeros {

	@Test public void questions() {
		
		int int1 = 0;  
		assertEquals(-int1, Math.abs(int1));

		int int2 = Integer.MIN_VALUE;  
		assertEquals(-int2, Math.abs(int2));
		
		double dbl1 = Double.NaN;
		assertFalse(dbl1 == dbl1);
		
	}
	
	@Test(expected=ArithmeticException.class) public void ex1() {
		int int1 = 5;
		int int2 = 0;
		int int3 = int1 / int2;	
	}

	@Test(expected=ArithmeticException.class) public void ex2() {
		int int1 = 5;
		int int2 = 0;
		int int3 = int1 % int2;	
	}

	@Test public void ex3() {
		int int1 = Integer.MAX_VALUE - 2;
		int int2 = int1 + 3;
		System.out.println(int2);  // -2147483648
	}

	@Test public void ex3b() {
		int int1 = 1_000_000_000;
		int int2 = int1 * 2_000_000_000;
		System.out.println(int2);  // 1321730048
	}

	
	@Test public void ex4() {
		double dbl1 = Math.sqrt(-1);
		System.out.println(dbl1); // NaN
	}

	@Test public void ex5() {
		double dbl1 = 5;
		double dbl2 = 0;
		double dbl3 = dbl1 % dbl2;	
		System.out.println(dbl3); // NaN
		
		double dbl4 = dbl1 / dbl2;	
		System.out.println(dbl4); // Inifinity

	}

	@Test public void ex6() {
		double dbl1 = Double.MAX_VALUE;
		double dbl2 = dbl1 + 1;
		System.out.println(dbl1); // 1.7976931348623157E308
		System.out.println(dbl2); // 1.7976931348623157E308
	}
		
	@Test public void ex7() {
		double a = 1.00 - 9 * .10;
		System.out.println(1.00 - 9 * .10); // 0.09999999999999998
	}


	@Test public void ex8() {
		long l = 1_000_000_000_000L;
		int i = (int) l;  // -727379968
		System.out.println(i);
	}
	
	@Test public void boxing() {
		Integer a1 = new Integer(10);
		Integer a2 = new Integer(10);
		Integer b = 20;
		int c = 20;
		
		System.out.println(a1 < b);
		System.out.println(a1 < c);
		System.out.println(c > a2);
		System.out.println(a1 <= a2);  
		System.out.println(a1 == a2);
		System.out.println(a1 != a2);
		
	}
}
