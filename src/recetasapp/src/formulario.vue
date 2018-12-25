<template>  
  <div class="container">
                <div id="app" class="col-sm-8 col-sm-offset-2">
                        <h1>Agregar Receta</h1>

                        <form >

                            <input type="text"  v-model="recipe_title">                     
                                
                            <input type="text" v-model="ingredient">
                                
                            <button v-on:click="addIngredient" class="btn btn-normal" >Agregar ingrediente</button> 
                            
                            <input type="text" v-model="step">
                                
                            <button v-on:click="addStep" class="btn btn-normal">Agregar Paso</button>
                            
                            <button v-on:click="addRecipe" class="btn btn-success">
                                    Agregar Receta
                            </button>
                           
                            
                        </form>
                        
                </div>
        </div>
        
        
</template>

<script>


export default {
  name: 'app',

  data () {
    return {
      recipe_title: '',
      ingredient: '',
      step: '',
      ingredients: [],
      steps:[],
      tracks: []
    }
  },

  methods: {
    addRecipe: function(e){
            e.preventDefault();
           
           console.log(this.ingredients);
           
             
            axios.post('http://localhost:8000/new', {  

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
            
            /*axios.post('http://localhost:8000/new', {
                "recipe_id": this.recipe_id,
                "recipe_title": this.recipe_title
            })*/
            

            /*this.recetas.push({
                titulo: this.recetas.titulo,
                ingredientes:['asfasf'],
                pasos:['afasdfsf']
            })*/
            
        },        
        addIngredient: function (e) {
            e.preventDefault();
            //console.log(this.ingredients)
            this.ingredients.push(this.ingredient);

        },

        addStep: function (e) {
            e.preventDefault();
            console.log(this.steps)
            this.steps.push(this.step);

          }
  }
}
</script>

<style lang="scss">
  

  .results {
    margin-top: 50px;
  }
</style>