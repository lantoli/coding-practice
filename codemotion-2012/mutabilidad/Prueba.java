package com.lantoli.codemotion2012.mutabilidad;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

public class Prueba {

	
	
	
	public static void main(String[] args) {
		// TODO Auto-generated method stub
		List<String> a = Collections.emptyList();
		//a.add("hola");
		
		List<String> miLista = new ArrayList<>();
		miLista.add("hola");
		List<String> miLista2 = Collections.unmodifiableList(miLista);
		miLista2.add("vale");
	}

}
