package main

import (
	"database/sql"
	"fmt"
	"log"	
	"encoding/json"	
	//"strconv"	
	"net/http"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

var recipes []Recipe
var results []string
 
type Recipe struct { 
	Recipe_title string `json:"recipe_title,omitempty"` 
	Ingredients []string `json:"ingredients,omitempty"` 
	Steps []string `json:"steps,omitempty"`
}


func Conection() (db *sql.DB){
	db, err := sql.Open("postgres", "postgresql://root@localhost:26257/?sslmode=disable")
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	return db

}

func Show(w http.ResponseWriter,req *http.Request){

	db := Conection()
	var recetas []Recipe
	var pasos []string
	var ingredientes []string
	var recipeID int64
	var title string

	

	rows, err := db.Query("select recipe_id, recipe_title FROM recipes;")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		
		
		if err := rows.Scan(&recipeID, &title); err != nil {
			log.Fatal(err)
		}

		ingredientes = SearchIngredients(title)
		pasos = SearchSteps(title)

		
		fmt.Printf("Name : %s\n", title)
		recetas = append(recetas, Recipe{Recipe_title: title, Ingredients:ingredientes, Steps:pasos})
	}

	

	json.NewEncoder(w).Encode(recetas)
}

func SearchIngredients(titulo string) ([]string){

	db := Conection();
	var ingredientes []string

	rowsi, erri := db.Query("select ingredient_name from ingredients INNER JOIN (select ingredient_id from recipesingredients INNER JOIN  (SELECT recipe_id FROM recipes WHERE recipe_title = '"+titulo+"') as recipe ON recipesingredients.recipe_id = recipe.recipe_id) as ingredientes ON ingredients.ingredient_id = ingredientes.ingredient_id;")
		if erri != nil {
			log.Fatal(erri)
		}

		for rowsi.Next(){
			var ingrediente string
			if err := rowsi.Scan(&ingrediente); err != nil {
				log.Fatal(err)
			}
			ingredientes = append(ingredientes, ingrediente)
		}

	return ingredientes
}

func SearchSteps(titulo string) ( []string){

	 db := Conection();
	 var pasos []string
	

		rowsp, errp := db.Query("select step_description from steps INNER JOIN (select step_id from recipessteps INNER JOIN  (SELECT recipe_id FROM recipes WHERE recipe_title = '"+titulo+"') as recipe ON recipessteps.recipe_id = recipe.recipe_id) as pasos ON steps.step_id = pasos.step_id;")
		if errp != nil {
			log.Fatal(errp)
		}

		for rowsp.Next(){
			var paso string
			if err := rowsp.Scan(&paso); err != nil {
				log.Fatal(err)
			}
			pasos = append(pasos, paso)
		}		
		
		return pasos

}
func New(w http.ResponseWriter,req *http.Request){

	db := Conection()

	decoder := json.NewDecoder(req.Body)
    var recipe Recipe
    erro := decoder.Decode(&recipe)
    if erro != nil {
        panic(erro)
    }
   
	title := recipe.Recipe_title
	
	
	filas,_:= db.Query("SELECT recipe_id FROM recipes WHERE recipe_title = '"+title+"';")
	if !filas.Next() {
		fmt.Printf("NO HAY NINGUNA RECETA IGUAL\n")
		if _, err := db.Exec("INSERT INTO recipes (recipe_title) VALUES ( '"+title+"');"); err != nil {
			log.Fatal(err)
		}
	}else{

		log.Fatal("RECETA CON ESE TITULO YA EXISTE")

	}
	

	

	//var ids []int64

	for i := 0; i < len(recipe.Ingredients); i++ {
	fmt.Printf("ENTRE \n")
	name := recipe.Ingredients[i]
	fmt.Printf(recipe.Ingredients[i]+"\n")

	res,_:= db.Query("SELECT ingredient_id FROM ingredients WHERE ingredient_name = '"+name+"';")
	if !res.Next() {
		fmt.Printf("NO HAY NINGUN INGREDIENTE IGUAL\n")
		if _, err := db.Exec("INSERT INTO ingredients (ingredient_name) VALUES ( '"+name+"');"); err != nil {
			log.Fatal(err)
		}
	}else {fmt.Printf("ENCONTRE UNA IGUAL\n")}

	rows, err := db.Query("SELECT (SELECT recipe_id FROM recipes WHERE recipe_title = '"+title+"') AS recipe, (SELECT ingredient_id FROM ingredients WHERE ingredient_name = '"+name+"') AS ingredient ;")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		
		var recipeID string
		var ingredientID string
		if err := rows.Scan(&recipeID, &ingredientID); err != nil {
			log.Fatal(err)
		}
		//fmt.Printf(recipeID)
		if _, err := db.Exec("INSERT INTO recipesingredients (recipe_id,ingredient_id) VALUES ( '"+recipeID+"','"+ingredientID+"');"); err != nil {
			log.Fatal(err)
		}

	}

	}

	for i := 0; i < len(recipe.Steps); i++ {
		fmt.Printf("ENTRE R \n")
		desc := recipe.Steps[i]
		fmt.Printf(recipe.Steps[i]+"\n")
	
		res,_:= db.Query("SELECT step_id FROM steps WHERE step_description = '"+desc+"';")
		if !res.Next() {
			fmt.Printf("NO HAY NINGUN PASO IGUAL\n")
			if _, err := db.Exec("INSERT INTO steps (step_description) VALUES ( '"+desc+"');"); err != nil {
				log.Fatal(err)
			}
		}else {fmt.Printf("ENCONTRE UN PASO IGUAL\n")}
	
		rows, err := db.Query("SELECT (SELECT recipe_id FROM recipes WHERE recipe_title = '"+title+"') AS recipe, (SELECT step_id FROM steps WHERE step_description = '"+desc+"') AS step ;")
		if err != nil {
			log.Fatal(err)
		}
	
		for rows.Next() {
			
			var recipeID string
			var stepID string
			if err := rows.Scan(&recipeID, &stepID); err != nil {
				log.Fatal(err)
			}
		
			if _, err := db.Exec("INSERT INTO recipessteps (recipe_id,step_id) VALUES ( '"+recipeID+"','"+stepID+"');"); err != nil {
				log.Fatal(err)
			}
	
		}
	
		}
	

	

	rows, err := db.Query("select recipe_title FROM recipes;")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		//var recipeID int
		var title string
		if err := rows.Scan(&title); err != nil {
			log.Fatal(err)
		}
		recipes = append(recipes, Recipe{Recipe_title: title})
	}



}

func EditRecipe(w http.ResponseWriter,req *http.Request){

	db := Conection()	
	params:=mux.Vars(req)
	tituloOriginal,_:=params["title"]
	fmt.Printf("titulo original original:"+tituloOriginal)	
	var tituloEditado string
	var ingredientesEdit []string
	var pasosEdit []string

	decoder := json.NewDecoder(req.Body)
    var recipe Recipe
    erro := decoder.Decode(&recipe)
    if erro != nil {
        panic(erro)
    }
   
	tituloEditado = recipe.Recipe_title
	ingredientesEdit = recipe.Ingredients
	pasosEdit = recipe.Steps

	fmt.Printf(tituloEditado)
	fmt.Printf(ingredientesEdit[0])
	fmt.Printf(pasosEdit[0])

	if _, erro := db.Exec("update recipes set recipe_title ='"+tituloEditado+"' where recipe_title = '"+tituloOriginal+"';"); erro != nil {
		log.Fatal(erro)
		fmt.Printf("ERRA")
	}
	
	
	if _, err := db.Exec("delete from recipesingredients where recipe_id = (select recipe_id from recipes where recipe_title = '"+tituloEditado+"');"); err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec("delete from recipessteps where recipe_id = (select recipe_id from recipes where recipe_title = '"+tituloEditado+"');"); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(ingredientesEdit); i++ {

		fmt.Printf("ENTRE \n")
		name := ingredientesEdit[i]
		fmt.Printf(ingredientesEdit[i]+"\n")
	
		res,_:= db.Query("SELECT ingredient_id FROM ingredients WHERE ingredient_name = '"+name+"';")
		if !res.Next() {
			fmt.Printf("NO HAY NINGUN INGREDIENTE IGUAL\n")
			if _, err := db.Exec("INSERT INTO ingredients (ingredient_name) VALUES ( '"+name+"');"); err != nil {
				log.Fatal(err)
			}
		}else {fmt.Printf("ENCONTRE UNA IGUAL\n")}
	
		rows, err := db.Query("SELECT (SELECT recipe_id FROM recipes WHERE recipe_title = '"+tituloEditado+"') AS recipe, (SELECT ingredient_id FROM ingredients WHERE ingredient_name = '"+name+"') AS ingredient ;")
		if err != nil {
			log.Fatal(err)
		}

		
		for rows.Next() {
			
			var recipeID string
			var ingredientID string
			if err := rows.Scan(&recipeID, &ingredientID); err != nil {
				log.Fatal(err)
			}
			//fmt.Printf(recipeID)
			if _, err := db.Exec("INSERT INTO recipesingredients (recipe_id,ingredient_id) VALUES ( '"+recipeID+"','"+ingredientID+"');"); err != nil {
				log.Fatal(err)
			}
	
		}
	
		}



		for i := 0; i < len(pasosEdit); i++ {
			fmt.Printf("ENTRE R \n")
			desc := pasosEdit[i]
			fmt.Printf(pasosEdit[i]+"\n")
		
			res,_:= db.Query("SELECT step_id FROM steps WHERE step_description = '"+desc+"';")
			if !res.Next() {
				fmt.Printf("NO HAY NINGUN PASO IGUAL\n")
				if _, err := db.Exec("INSERT INTO steps (step_description) VALUES ( '"+desc+"');"); err != nil {
					log.Fatal(err)
				}
			}else {fmt.Printf("ENCONTRE UN PASO IGUAL\n")}
		
			rows, err := db.Query("SELECT (SELECT recipe_id FROM recipes WHERE recipe_title = '"+tituloEditado+"') AS recipe, (SELECT step_id FROM steps WHERE step_description = '"+desc+"') AS step ;")
			if err != nil {
				log.Fatal(err)
			}
		
			for rows.Next() {
				
				var recipeID string
				var stepID string
				if err := rows.Scan(&recipeID, &stepID); err != nil {
					log.Fatal(err)
				}
				//fmt.Printf(recipeID)
				if _, err := db.Exec("INSERT INTO recipessteps (recipe_id,step_id) VALUES ( '"+recipeID+"','"+stepID+"');"); err != nil {
					log.Fatal(err)
				}
		
			}
		
			}
	

}



func ShowRecipe(w http.ResponseWriter,req *http.Request){

	db := Conection()
	params:=mux.Vars(req)
	titulo,_:=params["title"]
	fmt.Printf(titulo)
	var recipeID int64
	var ingredientesShow []string
	var pasosShow []string

	filas,_:= db.Query("SELECT recipe_id FROM recipes WHERE recipe_title = '"+titulo+"';")
	if !filas.Next() {		
		
			if err := filas.Scan(&recipeID); err != nil {
				log.Fatal(err)
			}
	}
	

	ingredientesShow = SearchIngredients(titulo)
	pasosShow = SearchSteps(titulo)

		fmt.Printf(ingredientesShow[0])

		
		fmt.Printf(pasosShow[0])

		json.NewEncoder(w).Encode(Recipe{Recipe_title: titulo, Ingredients:ingredientesShow, Steps:pasosShow})
		


}

func Delete(w http.ResponseWriter,req *http.Request){
	db := Conection()

	params:=mux.Vars(req)
	titulo,_:=params["title"]

	if _, err := db.Exec("delete from recipesingredients where recipe_id = (select recipe_id from recipes where recipe_title = '"+titulo+"');"); err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec("delete from recipessteps where recipe_id = (select recipe_id from recipes where recipe_title = '"+titulo+"');"); err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec("delete from recipes where recipe_title = '"+titulo+"';"); err != nil {
		log.Fatal(err)
	}
	

}

func main() {
	

	router := mux.NewRouter()
	router.HandleFunc("/show", Show).Methods("GET")
	router.HandleFunc("/show/{title}", ShowRecipe).Methods("GET")
	router.HandleFunc("/new", New).Methods("POST")
	router.HandleFunc("/delete/{title}", Delete).Methods("POST")
	router.HandleFunc("/edit/{title}", EditRecipe).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"}), handlers.AllowedOrigins([]string{"*"}))(router)))
	
}