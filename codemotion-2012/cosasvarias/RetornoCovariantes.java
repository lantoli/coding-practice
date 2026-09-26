package com.lantoli.codemotion2012.cosasvarias;

import org.junit.Test;

public class RetornoCovariantes {

	public class A {
		protected Object resultado() throws Exception {
			return null;
		}
	
		public Object fun(String a)  {
			System.out.println("fun string A");
			return null;
		}
	
		
	}
	
	public class B extends A {

		@Override
		public String resultado()  throws MiExc {
			return "hola";
		} 
		public Object fun(Object a) {
			System.out.println("fun object A");
			return null;
		}

	}
	
	class MiExc extends Exception {
	}
	
	@Test public void prueba1() {
		B b = new B();
		b.fun("hola");
	
	}
	
	
}
