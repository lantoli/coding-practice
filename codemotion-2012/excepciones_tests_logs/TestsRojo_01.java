package com.lantoli.codemotion2012.excepciones_tests_logs;

import static org.junit.Assert.*;

import org.junit.Test;

public class TestsRojo_01 {


	@Test 
	public void testEjemploSencillo() {
		assertTrue(5 < 8);
	}

	@Test public void testEjemplo() {
		assertTrue(5 < 8);
	}
	
	
	@Test public void testEjemplo2() {
		assertTrue(8 < 5);
	}
	
}
