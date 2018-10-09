// La casa tiene verdes jardines
  // Y en cada esquina crecen edelweis
// En la terraza siempre hay dos sillas
  // Para los voces que el viento le trae
    // Pero el viento no pasa la noche
// Y ella por los días se va 
  // Porque las plantas ya no están mustias
    // Pero el campo siempre será darkgreen
      // Vuela con ojos vivaces
// Sueña despierta cómo podría todo ser
  // Pues la casa tiene vallas
    // Pero en su cabeza nunca las haya
      // Tiene un todo por descubrir
        // Aunque el viento el todo ya lo encontró
          // En su jardín 

package main
  import (
    "net/http"
      )
        func goyo(w http.ResponseWriter, r *http.Request) {
          w.Write([]byte("Boby?"))
            w.Write([]byte("Boby, is that you?"))
              }
                func main() {
http.HandleFunc("/sotano", goyo)
  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
      } 
        }  
