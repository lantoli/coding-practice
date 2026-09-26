package com.lantoli.codemotion2012.herencia;

import static org.junit.Assert.*;

import org.junit.Test;

public class PuntoColor2 extends Punto2 {
	String color;
	public PuntoColor2(int x, int y, String color) {
	
		super(x, y);
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
		if (this == obj)
			return true;
		if (!super.equals(obj))
			return false;
		if (!(obj instanceof PuntoColor2))
			return false;
		PuntoColor2 other = (PuntoColor2) obj;
		if (color == null) {
			if (other.color != null)
				return false;
		} else if (!color.equals(other.color))
			return false;
		return true;
	}
	

}
