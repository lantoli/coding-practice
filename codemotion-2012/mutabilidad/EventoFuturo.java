package com.lantoli.codemotion2012.mutabilidad;

import java.util.Date;

public class EventoFuturo {

	private Date fechaEvento = new Date();
	
	EventoFuturo() {
	}

	public Date getFechaEvento() {
		return fechaEvento;
	}

	public void setFechaEvento(Date fechaEvento) {
		if (new Date().after(fechaEvento)) throw new IllegalArgumentException("la fecha tiene que ser en el futuro");
		this.fechaEvento = fechaEvento;
	}
	
	
	public static void main(String[] args) throws Exception {
		Date pasado = new Date();
		System.out.println("pasado: " + pasado);
		Thread.sleep(1000);
		EventoFuturo evt = new EventoFuturo();
		evt.setFechaEvento(pasado);
		
		evt.getFechaEvento().setDate(pasado.getDate());
		System.out.println(evt.getFechaEvento());
	}
	
}
