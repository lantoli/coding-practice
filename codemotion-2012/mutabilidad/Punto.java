package com.lantoli.codemotion2012.mutabilidad;


public class Punto  {
	private int x,y;
	public Punto(int x, int y) {
		super();
		this.x = x;
		this.y = y;
	}

	public int getX() {
		return x;
	}
	public int getY() {
		return y;
	}
	
	public static void main(String[] args) throws CloneNotSupportedException 
	{
		Punto punto1 = new Punto(10,20);
		
		MiPunto punto = new MiPunto(10,20);
		setPunto(punto);
		punto.setX(30);
	}
	
	
	public static class MiPunto extends Punto {
		private int x,y;
		 MiPunto(int x, int y) {
				super(x,y);
				this.x = x;
				this.y = y;
		 }
		public int getX() {
			return x;
		}
		public void setX(int x) {
			this.x = x;
		}
		public int getY() {
			return y;
		}
		public void setY(int y) {
			this.y = y;
		}
		 
	}

	static void setPunto(Punto p)  {
	
	}
}
