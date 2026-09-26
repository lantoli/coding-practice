package com.lantoli.codemotion2012.excepciones_tests_logs;

import java.sql.Connection;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.sql.Statement;

import org.junit.Test;

public class CerrarRecursos {

	@Test public void comprobarCerrarAntiguo() {
		 String query = "";
		 Connection con = getConnection();

 try  {
	 Statement stmt = con.createStatement();
     ResultSet rs = stmt.executeQuery(query);

     while (rs.next()) {
   	  // HACER COSAS CON EL RESULTADO
     }
      
     stmt.close();

 } catch (SQLException e) {
	 // GESTIONAR EXCEPCION
 }
	}

	
	@Test public void comprobarCerrarAntiguo1() {
		 String query = "";
		 Connection con = getConnection();
		 
	Statement stmt = null;
		 
	try  {
		stmt = con.createStatement();
		ResultSet rs = stmt.executeQuery(query);

		while (rs.next()) {
			// HACER COSAS CON EL RESULTADO
		}
     
 
} catch (SQLException e) {
	 // GESTIONAR EXCEPCION
} finally {
	 try {
		if (stmt != null) {
			stmt.close();
		}
	} catch (SQLException e) {
		// TODO Auto-generated catch block
		e.printStackTrace();
	}

}
	}	
	
	@Test public void comprobarCerrarAntiguo2() {
		 String query = "";
		 Connection con = getConnection();

		 try (Statement stmt = con.createStatement()) {
		      ResultSet rs = stmt.executeQuery(query);

		      while (rs.next()) {
		    	  // HACER COSAS CON EL RESULTADO
		      }

		    } catch (SQLException e) {
		    }
	}

	private Connection getConnection() {
		return null;
	}
	
	void miFunciondeServicio() throws ServiceException {
		try {
			///
		} catch (Exception e) {
	//		throw crearExcepcionServicio(e);
		}
	}
	
	static class ServiceException extends Exception {
		
	}
	
}
