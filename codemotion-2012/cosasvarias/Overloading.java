package com.lantoli.codemotion2012.cosasvarias;

import org.junit.Test;

public class Overloading {

	public String funcion(String a) {
		return null;
	}


	public Object funcion(Object a) {
		return null;
	}
	
	@Test public void prueba1() {
		Overloading a = new Overloading();
		Object o1 = "hola";
		String s1 = "hola";
		String o2 = (String) a.funcion(o1);
		String s2 = a.funcion(s1); 
	}
}
