package com.lantoli.codemotion2012.numbers;

import org.junit.Test;

public class PruebasNumeros2 {

	@Test
	public void bitsMinValue() {
		int a = Integer.MIN_VALUE;
		System.out.println(Integer.toBinaryString(a));
		System.out.println(Integer.toBinaryString(-a));
	}
	
	@Test
	public void doubletoIntBajo() {
		double d = 1e9;		
		int i = (int) d;
		System.out.println(i);
	}

	@Test
	public void doubletoIntAlto() {
		double d = 1e12;		
		int i = (int) d;
		System.out.println(i);
	}


	@Test
	public void doubletoIntAlto2() {
		double d = 1e19;		
		int i = (int) d;
		System.out.println(i);
		System.out.println(Integer.MAX_VALUE);
	}


	@Test
	public void doubletoIntNan() {
		double d = Double.NaN;		
		int i = (int) d;
		System.out.println(i);
	}

	@Test
	public void doubletoIntInf() {
		double d = Double.POSITIVE_INFINITY;		
		int i = (int) d;
		System.out.println(i);
	}

	@Test public void varios() {
		double nan = Double.NaN;
		double inf = Double.POSITIVE_INFINITY;
		double grande = 1e20;

		int inan = (int) nan;
		int iinf = (int) inf;
		int igrande = (int) grande;
		
		String a ="";
	
	}
	
	
	@Test public void testFibformula() {
		System.out.println(fib_formula(5));  // 5
		System.out.println(fib_formula(14)); // 377
		System.out.println(fib_formula(38)); // 39088169
		
	}

	@Test public void testFibrec() {
		System.out.println(fib_rec(5));  // 5
		System.out.println(fib_rec(14)); // 377
		System.out.println(fib_rec(38)); // 39088169
		
	}
	

  	private final static double ROOT_FIVE = Math.sqrt(5);
	private final static double GOLDEN_RATIO = (1 + ROOT_FIVE) / 2;
  
	public long fib_rec(int n) {
		return (long) ( (Math.pow(GOLDEN_RATIO,n) - Math.pow(GOLDEN_RATIO,-n)) / ROOT_FIVE  + .5);
	}

	public long fib_formula(int n) {
		if (n<=1) return n;
		return fib_formula(n-1) + fib_formula(n-2);
	}

}
