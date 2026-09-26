package com.lantoli.codemotion2012.clone;


public class Punto implements Cloneable {
	private int x;
	private int y;
	public Punto(int x, int y) {
		super();
		this.x = x;
		this.y = y;
	}

	public Punto(Punto p) {
		x = p.x;
		y = p.y;
	}
	
	public int getX() {
		return x;
	}
	public int getY() {
		return y;
	}
	@Override
	protected Object clone() throws CloneNotSupportedException {
	
		return new Punto(this);
	}

	static public class PuntoB extends Punto  {

		public PuntoB(int x, int y) {
			super(x, y);
			// TODO Auto-generated constructor stub
		}
	
	}
	
	public static void main(String[] args) throws CloneNotSupportedException 
	{
		Punto punto1 = new Punto(10,20);
		Punto punto2 = (Punto) punto1.clone();
		PuntoB puntob = new PuntoB(20,30);
		PuntoB puntob2 = (PuntoB) puntob.clone();
		String a = "";
	}
	
}
