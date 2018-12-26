<template>  
  <div class="container">
                <div id="app" class="col-sm-8 col-sm-offset-2">
                        <h1>RECETAS</h1>
                        <div >
                          <b-btn @click="clean" v-b-modal.modalPrevent>Crear nueva receta</b-btn>
                          <nav class="navbar navbar-light bg-light">
                          <a class="navbar-brand">Buscar receta</a>
                          <form class="form-inline">
                          <input class="form-control mr-sm-2" type="search" placeholder="Titulo receta" aria-label="Search" v-model="buscar">
                          <button @click="searchRecipe" class="btn btn-outline-success my-2 my-sm-0" type="submit">Buscar</button>
                          <button @click="closeSearch" type="button" class="close" aria-label="Close">
                          <span aria-hidden="true">&times;</span>
                          </button>
                          </form>
                          </nav>
                          <!-- Main UI -->
                          
                          <!-- Modal Component -->
                          <b-modal id="modalPrevent"
                                   ref="modal"
                                   title="Agrega una receta"
                                   @ok="handleOk"
                                   >
                          <form @submit.stop.prevent="addRecipe">
                          <b-container fluid>
                          <p class="font-weight-bold" >Título:</p>
                          <b-form-row class="mb-4"><input type="text" placeholder="Escriba el titulo" v-model="recipe_title"></b-form-row>                                             
                          <p class="font-weight-bold" >Ingredientes:</p>
                          <b-form-row align-h="end" v-for="ingredient in this.ingredients">
                            <b-col ><ul><li>{{ingredient}}</li></ul></b-col></b-form-row>
                          <p>Escriba ingrediente por ingrediente</p>
                          <b-form-row class="mb-3"><input type="text" placeholder="Escriba ingrediente" v-model="ingredient">                                
                            <button @click="addIngredient" class="btn btn-primary" >Agregar ingrediente</button></b-form-row>                           
                          <p class="font-weight-bold" >Pasos:</p>
                          <b-form-row   v-for="step in this.steps">
                            <b-col sm="6"><ul><li>{{step}}</li></ul></b-col></b-form-row>
                          <p>Escriba paso por paso</p>
                          <b-form-row><input type="text" placeholder="Escriba un paso" v-model="step">                                
                            <button @click="addStep" class="btn btn-primary">Agregar Paso</button> </b-form-row> 
                          </b-container>
                          </form>
                          </b-modal>
                        </div> 

                        <div >
                          
                          <!-- Main UI -->
                          
                          <!-- Modal2 Component -->
                          <b-modal id="modalPrevent2"
                                   ref="modal"
                                   title="Editar una receta"
                                   @ok="editRecipe"
                                   >
                          <form>
                          <b-container fluid>
                          <p class="font-weight-bold" >Título:</p>
                          <b-form-row class="mb-4">
                            <b-col sm="6"><input type="text" v-model="editRecipeData.recipe_title"></b-col></b-form-row>                                             

                          <p class="font-weight-bold" >Ingredientes:</p>
                          <b-form-row align-h="end" v-for="(ingredient,i) in editRecipeData.ingredients">
                            <b-col ><ul><li>{{ingredient}}</li></ul></b-col>
                            <b-col ><button type="button" @click="deleteIngredient(i)" class="btn btn-danger" >Eliminar</button></b-col>
                          </b-form-row>

                          <p class="font-weight-bold">Pasos:</p>
                          <b-form-row  class="mb-4" v-for="(step,s) in editRecipeData.steps">
                            <b-col sm="6"><ul><li>{{step}}</li></ul></b-col>
                            <b-col sm="6"><button  type="button" @click="deleteStep(s)" class="btn btn-danger" >Eliminar</button></b-col>
                          </b-form-row>

                          <p class="font-weight-bold">Agregar nuevo ingrediente</p>
                          <b-form-row align-h="between" class="mb-3">
                            <b-col ><input type="text" placeholder="Escriba ingrediente" v-model="editRecipeData.ingredient"></b-col>                                
                            <b-col ><button @click="addIngredientEdit" class="btn btn-primary" >Agregar ingrediente</button></b-col>
                          </b-form-row>                           
                            <p class="font-weight-bold">Agregar nuevo paso</p>
                          <b-form-row align-h="between" class="mb-4">
                            <b-col ><input type="text" placeholder="Escriba un paso" v-model="editRecipeData.step"></b-col>                                
                            <b-col ><button @click="addStepEdit" class="btn btn-primary">Agregar Paso</button> </b-col>
                          </b-form-row> 
                          </b-container>
                          </form>
                          </b-modal>
                        </div>


                        <div v-if="showInfo" class="card">
	                        <div class="card-header" >
	                                  <p class="font-weight-bold">Detalles Receta</p>
	                        </div>
	                        <div class="card-body">
	      	                <p >Titulo: {{recipe_title}}</p>
	                        <p >Ingredientes:</p>
                        
                          <ul>
                          <li v-for="i in ingredients" > {{ i }}</li></ul>
	                        <p >Pasos:</p>
                          <ul><li v-for="p in steps" > {{ p }}</li></ul>
                          <button @click="cerrarDetalle" class="btn btn-primary">Cerrar detalle</button>
	                        </div>
                          
		                    </div>
                                       
                        <table class="table">
                            <thead>
                                <tr>
                                    <th>Título Receta</th>
                                    <th>Ver Detalle</th>
                                    <th>Modificar</th>
                                    <th>Eliminar</th>                                    
                                </tr>
                            </thead>
                            <tbody>
                               <tr v-for="receta in recetas" >
                                   <td>{{receta.recipe_title}}</td>
                                   
                                   <td>
                                        <button @click="showRecipe(receta)" class="btn btn-primary">
                                                Ver
                                        </button>
                                   </td>
                                   <td>
                                       <button class="btn btn-success" @click="cargar(receta)" v-b-modal.modalPrevent2>
                                           Modificar
                                       </button> 
                                   </td>
                                   <td>
                                        <button v-on:click="deleteRecipe(receta)" class="btn btn-danger">
                                                Eliminar
                                        </button>
                                   </td>                                   
                               </tr> 
                            </tbody>
                        </table>
                    </div>
        </div>
</template>

<script>

//this.prototype.$axios = axios

export default {
  name: 'app',

  data () {
    return {
      
      recipe_title: '',
        ingredient: '',
        step:'',
        showInfo: false, 
        buscar: '',
        recipe:{},
        ingredients: [],
        steps:[],        
      recetas:[],
      
      editRecipeData: {
        recipe_title : '',
        ingredient: '',
        step: '',
        ingredients: [],
        steps: [],
      },
    }
  },

  created: function(){
        this.getRecipes();
        
     },

  methods: {
    search () {
      this.tracks = tracks
    },
    getRecipes: function() {
           
      
      this.$axios.get('http://localhost:8000/show')
      .then(response => {
       
        this.recetas = [],
        this.recetas = response.data;
        
        
             });
        },
    clearName (evt) {
      evt.preventDefault();
    },
 
    handleOk (evt) {
      
      evt.preventDefault()
                    
      var vm = true;
      if (!this.recipe_title) {
        alert('Por favor agregar titulo')
      }else if(this.recetas != null){
        this.recetas.forEach(element => {
               if(element.recipe_title == this.recipe_title){                
                  alert('YA EXISTE UNA RECETA CON ESE TITULO')
                  vm = false;
               }
             });
      }
      
      if(!this.ingredients.length){
        alert('Por favor agregar ingredientes')
      } else if(!this.steps.length){
        alert('Por favor agregar pasos')
      }else if(vm){
        
        this.addRecipe(evt)
      }
    },

    handleOkEditar (evt){

        evt.preventDefault();

      if(!this.editRecipeData.ingredients.length){
        alert('Por favor agregar ingredientes')
      } else if(!this.editRecipeData.steps.length){
        alert('Por favor agregar pasos')
      }else{
        
        

         this.editRecipe(evt);
      }


     // this.editRecipe(evt);
      
    },
    
    addRecipe: function(e){
            e.preventDefault();
           
          
            this.$axios.post('http://localhost:8000/new', {  

                recipe_title: this.recipe_title,
                ingredients: this.ingredients,
                steps: this.steps                                
              })

              .then(function (response) {
                console.log(response);
              })
              .catch(function (error) {
                
                console.log(error);
              });

            
            this.recetas.push({recipe_title:this.recipe_title,ingredients:this.ingredients,steps:this.steps});     
                       
            this.recipe_title = '',
            this.ingredient = '',
            this.step = '',
            this.ingredients = [],
            this.steps = []       
            alert('RECETA AGREGADA CON ÉXITO');
            this.recipe_title = '',
            
            this.$refs.modal.hide();
        },        
        addIngredient: function (e) {
            e.preventDefault();
            var vm = true;
         if(this.ingredients != null){
        this.ingredients.forEach(element => {
               if(element == this.ingredient){                
                  alert('YA EXISTE ESE INGREDIENTE EN ESA RECETA')
                  vm = false;
               }
             });
      }
          if(vm){
            
            this.ingredients.push(this.ingredient);}
            
            this.ingredient = '';

        },

        addStep: function (e) {
            e.preventDefault();
            
            var vm = true;
         if(this.steps != null){
        this.steps.forEach(element => {
               if(element == this.step){                
                  alert('YA EXISTE ESE PASO EN ESA RECETA')
                  vm = false;
               }
             });
      }
          if(vm){


            this.steps.push(this.step);}
            this.step = '';

          },

        showRecipe: function(receta){
          this.showInfo = true;

          this.$axios.get('http://localhost:8000/show/'+receta.recipe_title)
      .then(response => {
        
        this.recipe_title = response.data.recipe_title;
        this.ingredients = response.data.ingredients;
          this.steps = response.data.steps;
       
        
             });

                    
        },
        clean: function(){
          console.log("INGREDIENTES:"+this.ingredients)
          this.recipe_title = '',
            this.ingredient = '',
            this.step = '',
            this.ingredients = [],
            this.steps = []

        },

        cargar: function(receta){

          this.$axios.get('http://localhost:8000/show/'+receta.recipe_title)
      .then(response => {
        
        this.editRecipeData.recipe_title  = response.data.recipe_title;
        this.editRecipeData.ingredients = response.data.ingredients;
        this.editRecipeData.steps = response.data.steps;
        this.editRecipeData.ingredient = '';
        this.editRecipeData.step = '';
        
        
             });

        this.recipe = receta;
        
        },
        cerrarDetalle: function(){

            this.showInfo = false;
            
            
        },
        editRecipe: function(e){
          
            e.preventDefault();

            this.showInfo = false;

             if(!this.editRecipeData.ingredients.length){
        alert('Por favor agregar ingredientes')
      } else if(!this.editRecipeData.steps.length){
        alert('Por favor agregar pasos')
      }else{
           
          
            this.$axios.post('http://localhost:8000/edit/'+this.recipe.recipe_title, {  

                recipe_title: this.editRecipeData.recipe_title,
                ingredients: this.editRecipeData.ingredients,
                steps: this.editRecipeData.steps                                
              })

              .then(function (response) {
                console.log(response);
              })
              .catch(function (error) {
                
                console.log(error);
              });

            
            var index = this.recetas.indexOf(this.recipe);
            

            this.recetas[index].recipe_title = this.editRecipeData.recipe_title;
            this.recetas[index].ingredients = this.editRecipeData.ingredients;
            this.recetas[index].steps = this.editRecipeData.steps;

           
            this.$refs.modal.hide();

      }

        },

        deleteIngredient: function(i){

                   
          this.editRecipeData.ingredients.pop(i);
          
          

        },

        addIngredientEdit: function(e){

            e.preventDefault();
          var vm = true;
         if(this.editRecipeData.ingredients != null){
        this.editRecipeData.ingredients.forEach(element => {
               if(element == this.editRecipeData.ingredient){                
                  alert('YA EXISTE ESE INGREDIENTE EN ESA RECETA')
                  vm = false;
               }
             });
      }
          if(vm){
         this.editRecipeData.ingredients.push(this.editRecipeData.ingredient);}
         


         this.editRecipeData.ingredient = '';

        },

        deleteStep: function(i){

          
          
          this.editRecipeData.steps.pop(i);
          

        },

        addStepEdit: function(e){
          e.preventDefault();
           var vm = true;
         if(this.editRecipeData.steps != null){
        this.editRecipeData.steps.forEach(element => {
               if(element == this.editRecipeData.step){                
                  alert('YA EXISTE ESE PASO EN ESA RECETA')
                  vm = false;
               }
             });
      }
          if(vm){

         this.editRecipeData.steps.push(this.editRecipeData.step);}
         this.editRecipeData.step = '';
        },

        deleteRecipe: function(receta){

          var index = this.recetas.indexOf(receta);
          this.recetas.splice(index,1);

          console.log(index)
          console.log(this.recetas)

          this.$axios.post('http://localhost:8000/delete/'+receta.recipe_title)

              .then(function (response) {
                console.log(response);
              })
              .catch(function (error) {
                
                console.log(error);
              });

            //this.getRecipes();

        },

        searchRecipe: function(e){

          e.preventDefault();
          var receta = {};
          var vm = false;
          this.recetas.forEach(element => {
               if(element.recipe_title == this.buscar){                
                  
                  receta = element;
                  vm = true;
               }
             });

            if(receta){

              this.recetas = [];
              this.recetas.push(receta);
              

            }
            
            if(!vm){

              alert('NO EXISTE NINGUNA RECETA CON ESE TÍTULO');
              this.getRecipes();

            }

        },

        closeSearch: function(){

            this.getRecipes();
            this.buscar = '';

        }

  },

  computed: {
    filteredList() {
      return this.recetas.filter(receta => {
        return receta.recipe_title.toLowerCase().includes(this.buscar.toLowerCase())
      })
    }
  }
  }

</script>

<style lang="scss">
  
  .results {
    margin-top: 50px;
  }
</style>