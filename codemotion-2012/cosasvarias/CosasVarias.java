package com.lantoli.codemotion2012.cosasvarias;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

import org.junit.Test;

public class CosasVarias {

	@Test public void test1() {
		String[] listaStr = { "hola ", "que ", "tal ", "estas" };	
		String res = "";
		for (String str : listaStr) {
			res += str;
		}
		System.out.println(res);
	}

	@Test public void test2a() {
		long timeIni = System.nanoTime();
	
		for (int i=0; i< 1_000_000; i++) {
			String str = "a" + "b" + "c" + "d" + "e" + "f";
 		}
		
		long timeEnd = System.nanoTime();
		System.out.println("tiempo 2a (ns): " + (timeEnd-timeIni) );
	}
	
	@Test public void test2b() {
		long timeIni = System.nanoTime();
	
		for (int i=0; i< 1_000_000; i++) {
			String str = new StringBuilder().append("a").append("b").append("c").append("d").append("e").append("f").toString();
 		}
		
		long timeEnd = System.nanoTime();
		System.out.println("tiempo 2b (ns): " + (timeEnd-timeIni) );
	}
	
	@Test public void test3a() {
		List<String> datos = new ArrayList<>();
		datos.add("esto");
		datos.add("es");
		datos.add("una");
		datos.add("prueba");
		System.out.println(datos);	
		
	}
	
	@Test public void test3b() {
		String[] elms = { "esto", "es", "una", "prueba" };
		List<String> datos = Arrays.asList(elms);
		System.out.println(datos);	
		
	}
	
	
	@Test public void test4a() {
		Map<String,Integer> datos = new HashMap<>(); // Java 7
		datos.put("dato1", 10);
		datos.put("dato2", 20);
		datos.put("dato3", 30);
		System.out.println(datos);	
		
	}

	@Test public void test4b() {
		Map<String,Integer> datos = new HashMap<String,Integer>() {{
			put("dato1", 10);
			put("dato2", 20);
			put("dato3", 30);
		}};
		System.out.println(datos);	
		
	}

	@Test public void test5() {
		new Utilidades().ejecuta();
		new BloqueInicializacion();
	}
	
	public class Padre {
	
	};
	

	public class Hijo extends Padre {
		
	};
	
	
	public class Utilidades {
		public void ejecuta() {
			Padre p = new Padre();
			Hijo h = new Hijo();
			Padre p2 = new Hijo();
			llama(p);
			llama(h);
			llama(p2);
			
			// llama el padre
			// llama el hijo
			// llama el padre

		}
		
		public void llama(Padre obj) {
			System.out.println("llama el padre");
			
			
		}

		public void llama(Hijo obj) {
			System.out.println("llama el hijo");
		}

	}
	
	public static class BloqueInicializacion {
		int a = 10;
		{ 
			a = 15; 
		};
		
		public BloqueInicializacion() {
			a = 10;
		}
		
	}


	
}
