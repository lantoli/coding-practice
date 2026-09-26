package com.lantoli.codemotion2012.herencia2;

import static org.junit.Assert.*;
import static org.junit.Assert.assertEquals;

import org.junit.Ignore;
import org.junit.Test;

public class PointEqualTest {

	@Test public void normalPointEqualsExpections() {
		Point p1 = new Point(10,20);
		Point p2 = new Point(10,20);
		Point dif = new Point(10,30);
		assertEquals(p1,p2);
		assertEquals(p2,p1);
		assertEquals(p1,p1);
		assertFalse(p1.equals(dif));
		assertFalse(dif.equals(p1));
	}

	@Test public void normalPointExtraEqualsExpections() {
		PointExtra p1 = new PointExtra(10,20,30);
		PointExtra p2 = new PointExtra(10,20,30);
		PointExtra dif = new PointExtra(10,20,40);
		assertEquals(p1,p2);
		assertEquals(p2,p1);
		assertEquals(p1,p1);
		assertFalse(p1.equals(dif));
		assertFalse(dif.equals(p1));
	}

	@Test public void symmetry() {
		Point p = new Point(10,20);
		PointExtra pe = new PointExtra(10,20,30);
		
		assertFalse (p.equals(pe));
		assertFalse (pe.equals(p));
	}

	@Test public void transitivity() {
		PointExtra pe1 = new PointExtra(10,20,30);
		Point p = new Point(10,20);
		PointExtra pe2 = new PointExtra(10,20,40);
		
		assertFalse (pe1.equals(p));
		assertFalse (p.equals(pe2));
		assertFalse (pe1.equals(pe2));

	}
	
	@Test public void liskov() {
		Point p = new Point(10,20);
		PointNotSignficant pns = new PointNotSignficant(10,20);
		assertEquals(p,pns);
		assertEquals(pns,p);
		
	}
}
