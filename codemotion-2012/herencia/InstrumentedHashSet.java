package com.lantoli.codemotion2012.herencia;

import java.util.Arrays;
import java.util.Collection;
import java.util.HashSet;

	public class InstrumentedHashSet<E> extends HashSet<E> {
			private static final long serialVersionUID = 1L;
			private int addCount = 0;
		
			@Override public boolean add(E e) { 
				addCount++; 
				return super.add(e); 
			} 
			
			@Override public boolean addAll(Collection<? extends E> c) 
			{ 
				addCount += c.size(); 
				return super.addAll(c); 
			}
			
			private int addCount() {
				// TODO Auto-generated method stub
				return addCount;
			}
			
			public static void main(String[] args) {
				InstrumentedHashSet<String> s = new InstrumentedHashSet<>(); 
				s.addAll(Arrays.asList("Snap", "Crackle", "Pop"));
				System.out.println(s.addCount());
			}

}
