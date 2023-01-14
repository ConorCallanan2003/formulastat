<template>
    <div :class=this.backgroundColor>
    <div class="columns is-mobile is-gapless is-multiline">
      <div class="column is-full mb-2">
        <h1 class="subtitle has-text-weight-bold is-size-4 has-text-info-light"> {{this.title}}</h1>
      </div>
      <div v-for="(item, index) in listItems" v-bind:key="item"  class="column is-half">
        <h1 class="ml-2 is-size-5 subtitle has-text-primary-light">
            {{ (index+1) + '. ' + item }}
        </h1>
      </div>    
    </div>
  </div>
    
</template>
  
<script>
export default {
name: 'StandingsVue',
props: {
    listItems: Array,
    backgroundColor: String,
    title: String
},
}
</script>
  