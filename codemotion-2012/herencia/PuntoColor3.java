package com.lantoli.codemotion2012.herencia;

import static org.junit.Assert.*;

import org.junit.Test;

public class PuntoColor3 extends Punto2 {
	String color;
	public PuntoColor3(int x, int y, String color) {
	
		super(x, y);
		this.color = color;
		// TODO Auto-generated constructor stub
	}
	@Override
	public int hashCode() {
		final int prime = 31;
		int result = super.hashCode();
		result = prime * result + ((color == null) ? 0 : color.hashCode());
		return result;
	}
	@Override
	public boolean equals(Object obj) {
		 if (!(obj instanceof Punto2)) return false;
		   if (!(obj instanceof PuntoColor3)) return obj.equals(this);

		   return super.equals(obj) && ((PuntoColor3)obj).color.equals(color);	
	}
	

}
