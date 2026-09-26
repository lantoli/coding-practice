package com.lantoli.codemotion2012.herencia;

import static org.junit.Assert.*;
import static org.junit.Assert.assertEquals;

import org.junit.Ignore;
import org.junit.Test;

public class PuntoTest {

	@Ignore
	@Test public void prueba() {
		Punto punto1 = new Punto(10,20);
		Punto punto2 = new Punto(10,20);
		PuntoColor puntoColor = new PuntoColor(10,20,"azul");
		
		assertEquals(punto1,punto2);
		assertEquals(punto1,puntoColor);
	}


	@Ignore
	@Test public void prueba2() {
		Punto2 punto = new Punto2(10,20);
		PuntoColor2 puntoColor = new PuntoColor2(10,20,"azul");

		assertTrue(punto.equals(puntoColor));
		assertFalse(puntoColor.equals(punto)); 		
	}



	@Test public void prueba3() {
		PuntoColor3 pcolor1 = new PuntoColor3(10,20,"azul");
		Punto2 punto = new Punto2(10,20);
		PuntoColor3 pcolor2 = new PuntoColor3(10,20,"verde");
		
		assertTrue(pcolor1.equals(punto));
		assertTrue(punto.equals(pcolor2));
		assertFalse(pcolor1.equals(pcolor2)); 		
	}
}
