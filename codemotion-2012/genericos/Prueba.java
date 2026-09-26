package com.lantoli.codemotion2012.genericos;

import java.util.ArrayList;
import java.util.List;

public class Prueba {

	/**
	 * @param args
	 */
	public static void main(String[] args) {
		// TODO Auto-generated method stub

	}

	// List-based generic reduction 
	static <E> E reduce(List<E> list, Function<E> f, E initVal) { 
		List<E> snapshot; 
		synchronized(list) { 
			snapshot = new ArrayList<E>(list); 
		} 
		E result = initVal; 
		for (E e : snapshot){ 
			result = f.apply(result, e); 
		}
		return result; }
}
