package com.lantoli.codemotion2012.herencia;


import static org.junit.Assert.*;

import java.util.Date;

import org.junit.Test;

public class Timestamp {

	@Test public void testTimestamp() {
		java.sql.Timestamp ts = new java.sql.Timestamp(100);
		Date date = new Date(100);
		
		assertEquals(date,ts);
		assertEquals(ts,date); // FALSE
	}
}
