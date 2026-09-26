package com.lantoli.codemotion2012.excepciones_tests_logs;

import static org.junit.Assert.*;

import org.junit.Test;

public class Cobertura_02 {

	@Test public void ejemploCoche() {
		Coche coche = new Coche();
		coche.setMarca("ford");
		coche.setColor("rojo");
		assertEquals("ford", coche.getMarca());
		assertEquals("rojo", coche.getColor());		
	}
	
	
}
