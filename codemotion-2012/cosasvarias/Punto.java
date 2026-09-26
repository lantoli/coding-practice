package com.lantoli.codemotion2012.cosasvarias;

import java.util.HashSet;
import java.util.Set;

public class Punto {
	public static void main(String[] args) {

		Punto p1 = new Punto(10,20);
		Punto p2 = new Punto(10,20);
		
		System.out.println(p1.equals(p2));
		
		Set<Punto> set = new HashSet<>();
		set.add(p1);
		System.out.println(set.contains(p2));
		
	}

	private int x,y;

	
	public Punto(int x, int y) {
		super();
		this.x = x;
		this.y = y;
	}
	

	public int hashCode() {
		final int prime = 31;
		int result = 1;
		result = prime * result + x;
		result = prime * result + y;
		return result;
	}

	public boolean equals(Punto p) {
		return x == p.x && y == p.y;
	}
	
	
	
	
}
