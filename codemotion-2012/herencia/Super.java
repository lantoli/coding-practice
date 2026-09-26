package com.lantoli.codemotion2012.herencia;

public class Super { 
	
	// Broken - constructor invokes an overridable method 
	public Super() { 
		overrideMe(); 
	} 
	
	public void overrideMe() { } 
}
