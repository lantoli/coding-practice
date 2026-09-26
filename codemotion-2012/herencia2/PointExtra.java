package com.lantoli.codemotion2012.herencia2;

import static org.junit.Assert.*;

import org.junit.Test;

public class PointExtra extends Point {
	int extra;

	public PointExtra(int x, int y, int extra) {
		super(x, y);
		this.extra = extra;
	}

	public int getExtra() {
		return extra;
	}

	public void setExtra(int extra) {
		this.extra = extra;
	}

	public boolean equals(Object obj) {
		if (!(obj instanceof PointExtra)) return false;
		if (isEqualsInDifferentClass(obj)) return obj.equals(this);
		PointExtra other = (PointExtra) obj;	
		if (extra != other.extra) return false;
		return true;	
	}
}
