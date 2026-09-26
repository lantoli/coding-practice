package com.lantoli.codemotion2012.excepciones_tests_logs;

import static org.junit.Assert.*;

import java.io.FileInputStream;
import java.io.FileNotFoundException;

import org.junit.Test;

public class ExcepcionesEnTests {

	@Test public void pruebaExcepciones() {
		String fichero = "hola.txt";
		try {
			FileInputStream fin = new FileInputStream(fichero);
			// HACER COSAS CON EL FICHERO
		} catch (FileNotFoundException e) {
			e.printStackTrace();
			fail(e.getMessage());
		} 
	}
	
	@Test public void pruebaExcepciones2() throws FileNotFoundException {
		String fichero = "hola.txt";

		FileInputStream fin = new FileInputStream(fichero);
		// HACER COSAS CON EL FICHERO
	}



}
