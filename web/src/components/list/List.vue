<template>
  <div>
    <div class="card" v-for="list in schoolList" :key="list.id">
        <p> {{list.adi}} </p>
    </div>
  </div>
</template>

<script>

import axios from 'axios';
export default {
  data(){
    return {
      schoolList : []
    }
  },
  mounted: function(){
      this.etkinlikListe()
  },
  methods: {
      etkinlikListe: function() {
          axios.get('http://192.168.1.55:8090/okul/liste')
          .then((result) => {
              this.schoolList = result.data
          });
      }
  }
}


/*
created(){
  axios.get('http://192.168.1.55:8090/okul/liste')
  .then(function (response) {
    let data = response.data;
    for(let key in data){
      this.schoolList.push({...data[key], id: key});
    }
  })
  .catch(function (error) {
  console.log(error);
})*/
</script>
