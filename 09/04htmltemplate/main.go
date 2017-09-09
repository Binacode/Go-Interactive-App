package main

import (
    "html/template"
    "log"
    "net/http"
)

// Persona contiene los datos de una Persona
type Persona struct{
    Nombre string
    Edad uint8
}

func renderizar(w http.ResponseWriter, r *http.Request){
    p := &Persona{"Jonathan", 27}
    t, err := template.ParseFiles("./views/index.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    t.Execute(w, p)
    // t.Execute(w, nil)  //puede ser de esta otra manera
}

func main()  {
    mux := http.NewServeMux()
    mux.HandleFunc("/", renderizar)

    log.Println("Ejecutando server en http://localhost:8080")
    log.Println(http.ListenAndServe(":8080", mux))
    log.Println("Ejecución terminada.")
}
