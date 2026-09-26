package com.lantoli.codemotion2012.mutabilidad;

import java.util.HashMap;
import java.util.Map;

public class Punto2 {

	private int x;
	private int y;
	public Punto2(int x, int y) {
		this.x = x;
		this.y = y;
	}
	public int getX() {
		return x;
	}
	public int getY() {
		return y;
	}

	public void setX(int x) {
		this.x = x;
	}

	
	@Override
	public int hashCode() {
		final int prime = 31;
		int result = 1;
		result = prime * result + x;
		result = prime * result + y;
		return result;
	}
	@Override
	public boolean equals(Object obj) {
		if (this == obj)
			return true;
		if (obj == null)
			return false;
		if (getClass() != obj.getClass())
			return false;
		Punto2 other = (Punto2) obj;
		if (x != other.x)
			return false;
		if (y != other.y)
			return false;
		return true;
	}
	@Override
	public String toString() {
		return "Punto2 [x=" + x + ", y=" + y + "]";
	}
	
	public static void main(String[] args)  {
		Map<Punto2,String> mapa = new HashMap<>();
		Punto2 punto = new Punto2(10,20);
		mapa.put(punto, "mi punto");
		System.out.println(mapa.get(punto)); // "mi punto"
		punto.setX(30);
		System.out.println(mapa.get(punto)); // null		
	}

}
