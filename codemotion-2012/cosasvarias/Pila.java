package com.lantoli.codemotion2012.cosasvarias;

public class Pila {

	static class NodoPila implements Nodo {
		
	}
	
	
	public void optimizacion() {
		Object instance = null;
		
		
		
	}
	
	public class MiClase {
		
	}
	private static MiClase instancia;
	
	public static MiClase getInstancia() {
		if (instancia == null) {
			instancia = inicializaInstancia();
		}
		return instancia;
	}

	private static synchronized MiClase inicializaInstancia() {
		if (instancia != null) {
			return instancia;
		}
		return null;
	}
}
