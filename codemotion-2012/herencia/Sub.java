package com.lantoli.codemotion2012.herencia;

public final class Sub extends Super { 
	
	private final Integer num;  
	
	Sub(int n) { 
		num = new Integer(n); 
	} 
	
	// Overriding method invoked by superclass constructor 
	
	@Override public void overrideMe() {
		int res = num + 1;
		System.out.println("sumando 1: " + res);
	} 
	
	public static void main(String[] args) { 
		Sub sub = new Sub(10); 
	} 

}