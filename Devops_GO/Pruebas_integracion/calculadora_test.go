package main

import "testing"
import "fmt"
            
func TestSuma(t *testing.T) {
            suma := Suma(7, 7)
            if suma != 14 {
                        t.Error("Se esperaba 14 y se obtuvo", suma)
            }
            
            fmt.Println("Prueba de suma exitosa")
}
 
func TestResta(t *testing.T) {
            resta := Resta(7, 7)
            if resta != 0 {
                        t.Error("Se esperaba 0 y se obtuvo", resta)
            }
            
            fmt.Println("Prueba de resta exitosa")
}
 
func TestMultiplicacion(t *testing.T) {
            multiplicacion := Multiplicacion(7, 7)
            if multiplicacion != 49 {
                        t.Error("Se esperaba 49 y se obtuvo", multiplicacion)
            }
            
            fmt.Println("Prueba de multiplicacion exitosa")
}
 
func TestDivision(t *testing.T) {
            division := Division(7, 7)
            if division != 1 {
                        t.Error("Se esperaba 1 y se obtuvo", division)
            }
            
            fmt.Println("Prueba de division exitosa")
}
 
func TestPromedio(t *testing.T) {
            promedio := Promedio(7, 7, 7, 7, 7)
            if promedio != 7 {
                        t.Error("Se esperaba 7 y se obtuvo", promedio)
            }
            
            fmt.Println("Prueba del promedio exitosa")
}
