package com.lantoli.codemotion2012.mutabilidad;


import java.util.List;


public class Alumno {

	private String nombre;
	private String primerApellido;
	private String segundoApellido;
	private List<Materia> materiasMatriculadas;

	public List<Materia> getMateriasMatriculadas() {
		return materiasMatriculadas;
	}
	public void setMateriasMatriculadas(List<Materia> materiasMatriculadas) {
		this.materiasMatriculadas = materiasMatriculadas;
	}
}