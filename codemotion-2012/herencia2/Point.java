package com.lantoli.codemotion2012.herencia2;

public class Point {

	private int x;
	private int y;

	public Point(int x, int y) {
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

	public boolean equals(Object obj) {
		if (!(obj instanceof Point)) return false;
		if (isEqualsInDifferentClass(obj)) return obj.equals(this);
		Point other = (Point) obj;
		if (x != other.x) return false;
		if (y != other.y) return false;
		return true;
	}

	final protected boolean isEqualsInDifferentClass(Object obj) {
		try {
			return ! getClass().equals(obj.getClass().getMethod("equals", Object.class).getDeclaringClass());
		} catch (Exception e) {		
			return false;
		}
	}
	
}
